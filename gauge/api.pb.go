// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	GetProjectRootRequest
	GetProjectRootResponse
	GetAllStepsRequest
	GetAllStepsResponse
	GetAllSpecsRequest
	GetAllSpecsResponse
	GetStepValueRequest
	GetStepValueResponse
	ErrorResponse
	APIMessage
*/
package main

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type APIMessage_APIMessageType int32

const (
	APIMessage_GetProjectRootRequest  APIMessage_APIMessageType = 1
	APIMessage_GetProjectRootResponse APIMessage_APIMessageType = 2
	APIMessage_GetAllStepsRequest     APIMessage_APIMessageType = 3
	APIMessage_GetAllStepResponse     APIMessage_APIMessageType = 4
	APIMessage_GetAllSpecsRequest     APIMessage_APIMessageType = 5
	APIMessage_GetAllSpecsResponse    APIMessage_APIMessageType = 6
	APIMessage_GetStepValueRequest    APIMessage_APIMessageType = 7
	APIMessage_GetStepValueResponse   APIMessage_APIMessageType = 8
	APIMessage_ErrorResponse          APIMessage_APIMessageType = 9
)

var APIMessage_APIMessageType_name = map[int32]string{
	1: "GetProjectRootRequest",
	2: "GetProjectRootResponse",
	3: "GetAllStepsRequest",
	4: "GetAllStepResponse",
	5: "GetAllSpecsRequest",
	6: "GetAllSpecsResponse",
	7: "GetStepValueRequest",
	8: "GetStepValueResponse",
	9: "ErrorResponse",
}
var APIMessage_APIMessageType_value = map[string]int32{
	"GetProjectRootRequest":  1,
	"GetProjectRootResponse": 2,
	"GetAllStepsRequest":     3,
	"GetAllStepResponse":     4,
	"GetAllSpecsRequest":     5,
	"GetAllSpecsResponse":    6,
	"GetStepValueRequest":    7,
	"GetStepValueResponse":   8,
	"ErrorResponse":          9,
}

func (x APIMessage_APIMessageType) Enum() *APIMessage_APIMessageType {
	p := new(APIMessage_APIMessageType)
	*p = x
	return p
}
func (x APIMessage_APIMessageType) String() string {
	return proto.EnumName(APIMessage_APIMessageType_name, int32(x))
}
func (x *APIMessage_APIMessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(APIMessage_APIMessageType_value, data, "APIMessage_APIMessageType")
	if err != nil {
		return err
	}
	*x = APIMessage_APIMessageType(value)
	return nil
}

type GetProjectRootRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetProjectRootRequest) Reset()         { *m = GetProjectRootRequest{} }
func (m *GetProjectRootRequest) String() string { return proto.CompactTextString(m) }
func (*GetProjectRootRequest) ProtoMessage()    {}

type GetProjectRootResponse struct {
	ProjectRoot      *string `protobuf:"bytes,1,req,name=projectRoot" json:"projectRoot,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetProjectRootResponse) Reset()         { *m = GetProjectRootResponse{} }
func (m *GetProjectRootResponse) String() string { return proto.CompactTextString(m) }
func (*GetProjectRootResponse) ProtoMessage()    {}

func (m *GetProjectRootResponse) GetProjectRoot() string {
	if m != nil && m.ProjectRoot != nil {
		return *m.ProjectRoot
	}
	return ""
}

type GetAllStepsRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetAllStepsRequest) Reset()         { *m = GetAllStepsRequest{} }
func (m *GetAllStepsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllStepsRequest) ProtoMessage()    {}

type GetAllStepsResponse struct {
	Steps            []string `protobuf:"bytes,1,rep,name=steps" json:"steps,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GetAllStepsResponse) Reset()         { *m = GetAllStepsResponse{} }
func (m *GetAllStepsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllStepsResponse) ProtoMessage()    {}

func (m *GetAllStepsResponse) GetSteps() []string {
	if m != nil {
		return m.Steps
	}
	return nil
}

type GetAllSpecsRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetAllSpecsRequest) Reset()         { *m = GetAllSpecsRequest{} }
func (m *GetAllSpecsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllSpecsRequest) ProtoMessage()    {}

type GetAllSpecsResponse struct {
	Specs            []*ProtoSpec `protobuf:"bytes,1,rep,name=specs" json:"specs,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *GetAllSpecsResponse) Reset()         { *m = GetAllSpecsResponse{} }
func (m *GetAllSpecsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllSpecsResponse) ProtoMessage()    {}

func (m *GetAllSpecsResponse) GetSpecs() []*ProtoSpec {
	if m != nil {
		return m.Specs
	}
	return nil
}

type GetStepValueRequest struct {
	StepText         *string `protobuf:"bytes,1,req,name=stepText" json:"stepText,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetStepValueRequest) Reset()         { *m = GetStepValueRequest{} }
func (m *GetStepValueRequest) String() string { return proto.CompactTextString(m) }
func (*GetStepValueRequest) ProtoMessage()    {}

func (m *GetStepValueRequest) GetStepText() string {
	if m != nil && m.StepText != nil {
		return *m.StepText
	}
	return ""
}

type GetStepValueResponse struct {
	StepValue        *string  `protobuf:"bytes,1,req,name=stepValue" json:"stepValue,omitempty"`
	Parameters       []string `protobuf:"bytes,2,rep,name=parameters" json:"parameters,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GetStepValueResponse) Reset()         { *m = GetStepValueResponse{} }
func (m *GetStepValueResponse) String() string { return proto.CompactTextString(m) }
func (*GetStepValueResponse) ProtoMessage()    {}

func (m *GetStepValueResponse) GetStepValue() string {
	if m != nil && m.StepValue != nil {
		return *m.StepValue
	}
	return ""
}

func (m *GetStepValueResponse) GetParameters() []string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type ErrorResponse struct {
	Error            *string `protobuf:"bytes,1,req,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ErrorResponse) Reset()         { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()    {}

func (m *ErrorResponse) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

type APIMessage struct {
	MessageType         *APIMessage_APIMessageType `protobuf:"varint,1,req,name=messageType,enum=main.APIMessage_APIMessageType" json:"messageType,omitempty"`
	MessageId           *int64                     `protobuf:"varint,2,req,name=messageId" json:"messageId,omitempty"`
	ProjectRootRequest  *GetProjectRootRequest     `protobuf:"bytes,3,opt,name=projectRootRequest" json:"projectRootRequest,omitempty"`
	ProjectRootResponse *GetProjectRootResponse    `protobuf:"bytes,4,opt,name=projectRootResponse" json:"projectRootResponse,omitempty"`
	AllStepsRequest     *GetAllStepsRequest        `protobuf:"bytes,5,opt,name=allStepsRequest" json:"allStepsRequest,omitempty"`
	AllStepsResponse    *GetAllStepsResponse       `protobuf:"bytes,6,opt,name=allStepsResponse" json:"allStepsResponse,omitempty"`
	AllSpecsRequest     *GetAllSpecsRequest        `protobuf:"bytes,7,opt,name=allSpecsRequest" json:"allSpecsRequest,omitempty"`
	AllSpecsResponse    *GetAllSpecsResponse       `protobuf:"bytes,8,opt,name=allSpecsResponse" json:"allSpecsResponse,omitempty"`
	StepValueRequest    *GetStepValueRequest       `protobuf:"bytes,9,opt,name=stepValueRequest" json:"stepValueRequest,omitempty"`
	StepValueResponse   *GetStepValueResponse      `protobuf:"bytes,10,opt,name=stepValueResponse" json:"stepValueResponse,omitempty"`
	Error               *ErrorResponse             `protobuf:"bytes,11,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized    []byte                     `json:"-"`
}

func (m *APIMessage) Reset()         { *m = APIMessage{} }
func (m *APIMessage) String() string { return proto.CompactTextString(m) }
func (*APIMessage) ProtoMessage()    {}

func (m *APIMessage) GetMessageType() APIMessage_APIMessageType {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return APIMessage_GetProjectRootRequest
}

func (m *APIMessage) GetMessageId() int64 {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return 0
}

func (m *APIMessage) GetProjectRootRequest() *GetProjectRootRequest {
	if m != nil {
		return m.ProjectRootRequest
	}
	return nil
}

func (m *APIMessage) GetProjectRootResponse() *GetProjectRootResponse {
	if m != nil {
		return m.ProjectRootResponse
	}
	return nil
}

func (m *APIMessage) GetAllStepsRequest() *GetAllStepsRequest {
	if m != nil {
		return m.AllStepsRequest
	}
	return nil
}

func (m *APIMessage) GetAllStepsResponse() *GetAllStepsResponse {
	if m != nil {
		return m.AllStepsResponse
	}
	return nil
}

func (m *APIMessage) GetAllSpecsRequest() *GetAllSpecsRequest {
	if m != nil {
		return m.AllSpecsRequest
	}
	return nil
}

func (m *APIMessage) GetAllSpecsResponse() *GetAllSpecsResponse {
	if m != nil {
		return m.AllSpecsResponse
	}
	return nil
}

func (m *APIMessage) GetStepValueRequest() *GetStepValueRequest {
	if m != nil {
		return m.StepValueRequest
	}
	return nil
}

func (m *APIMessage) GetStepValueResponse() *GetStepValueResponse {
	if m != nil {
		return m.StepValueResponse
	}
	return nil
}

func (m *APIMessage) GetError() *ErrorResponse {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterEnum("main.APIMessage_APIMessageType", APIMessage_APIMessageType_name, APIMessage_APIMessageType_value)
}