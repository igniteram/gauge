// Copyright 2015 ThoughtWorks, Inc.

// This file is part of Gauge.

// Gauge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge.  If not, see <http://www.gnu.org/licenses/>.

package execution

import (
	"time"

	"strings"

	"github.com/getgauge/gauge/execution/result"
	m "github.com/getgauge/gauge/gauge_messages"
)

func mergeDataTableSpecResults(sResult *result.SuiteResult) *result.SuiteResult {
	suiteRes := result.NewSuiteResult(sResult.Tags, time.Now())
	suiteRes.IsFailed = sResult.IsFailed
	suiteRes.ExecutionTime = sResult.ExecutionTime
	suiteRes.PostSuite = sResult.PostSuite
	suiteRes.PreSuite = sResult.PreSuite
	suiteRes.UnhandledErrors = sResult.UnhandledErrors
	suiteRes.Timestamp = sResult.Timestamp
	suiteRes.ProjectName = sResult.ProjectName
	suiteRes.Environment = sResult.Environment
	suiteRes.Tags = sResult.Tags
	combinedResults := make(map[string][]*result.SpecResult)
	for _, res := range sResult.SpecResults {
		fileName := res.ProtoSpec.GetFileName()
		combinedResults[fileName] = append(combinedResults[fileName], res)
	}
	for _, res := range combinedResults {
		mergedRes := res[0]
		if len(res) > 1 {
			mergedRes = mergeResults(res)
		}
		if mergedRes.GetFailed() {
			suiteRes.SpecsFailedCount++
		} else if mergedRes.Skipped {
			suiteRes.SpecsSkippedCount++
		}
		suiteRes.SpecResults = append(suiteRes.SpecResults, mergedRes)
	}
	return suiteRes
}

func mergeResults(results []*result.SpecResult) *result.SpecResult {
	specResult := &result.SpecResult{ProtoSpec: &m.ProtoSpec{IsTableDriven: true}}
	var scnResults []*m.ProtoItem
	table := &m.ProtoTable{}
	dataTableScnResults := make(map[string][]*m.ProtoTableDrivenScenario)
	max := results[0].ExecutionTime
	for _, res := range results {
		specResult.ExecutionTime += res.ExecutionTime
		if res.ExecutionTime > max {
			max = res.ExecutionTime
		}
		if res.GetFailed() {
			specResult.IsFailed = true
		}
		for _, item := range res.ProtoSpec.Items {
			switch item.ItemType {
			case m.ProtoItem_Scenario:
				scnResults = append(scnResults, item)
				modifySpecStats(item.Scenario, specResult)
			case m.ProtoItem_TableDrivenScenario:
				scnResults = append(scnResults, item)
				heading := item.TableDrivenScenario.Scenario.ScenarioHeading
				item.TableDrivenScenario.TableRowIndex = int32(len(table.Rows) - 1)
				dataTableScnResults[heading] = append(dataTableScnResults[heading], item.TableDrivenScenario)
			case m.ProtoItem_Table:
				table.Headers = item.Table.Headers
				table.Rows = append(table.Rows, item.Table.Rows...)
			}
		}
		addHookFailure(table, res.GetPreHook(), specResult.AddPreHook)
		addHookFailure(table, res.GetPostHook(), specResult.AddPostHook)
	}
	if InParallel {
		specResult.ExecutionTime = max
	}
	aggregateDataTableScnStats(dataTableScnResults, specResult)
	specResult.ProtoSpec.FileName = results[0].ProtoSpec.FileName
	specResult.ProtoSpec.Tags = results[0].ProtoSpec.Tags
	specResult.ProtoSpec.SpecHeading = results[0].ProtoSpec.SpecHeading
	specResult.ProtoSpec.Items = getItems(table, scnResults, results)
	return specResult
}

func addHookFailure(table *m.ProtoTable, f []*m.ProtoHookFailure, add func(...*m.ProtoHookFailure)) {
	for _, h := range f {
		h.TableRowIndex = int32(len(table.Rows) - 1)
	}
	add(f...)
}

func getItems(table *m.ProtoTable, scnResults []*m.ProtoItem, results []*result.SpecResult) (items []*m.ProtoItem) {
	index := 0
	for _, item := range results[0].ProtoSpec.Items {
		switch item.ItemType {
		case m.ProtoItem_Scenario, m.ProtoItem_TableDrivenScenario:
			items = append(items, scnResults[index])
			index++
		case m.ProtoItem_Table:
			items = append(items, &m.ProtoItem{ItemType: m.ProtoItem_Table, Table: table})
		default:
			items = append(items, item)
		}
	}
	items = append(items, scnResults[index:]...)
	return
}

func aggregateDataTableScnStats(results map[string][]*m.ProtoTableDrivenScenario, specResult *result.SpecResult) {
	for _, dResult := range results {
		isFailed := 0
		isSkipped := 0
		for _, res := range dResult {
			if res.Scenario.ExecutionStatus == m.ExecutionStatus_FAILED {
				isFailed = 1
			} else if res.Scenario.ExecutionStatus == m.ExecutionStatus_SKIPPED &&
				!strings.Contains(res.Scenario.SkipErrors[0], "--table-rows") {
				isSkipped = 1
			}
		}
		specResult.ScenarioFailedCount += isFailed
		specResult.ScenarioSkippedCount += isSkipped
		specResult.ScenarioCount++
	}
}

func modifySpecStats(scn *m.ProtoScenario, specRes *result.SpecResult) {
	switch scn.ExecutionStatus {
	case m.ExecutionStatus_SKIPPED:
		specRes.ScenarioSkippedCount++
		return
	case m.ExecutionStatus_FAILED:
		specRes.ScenarioFailedCount++
	}
	specRes.ScenarioCount++
}
