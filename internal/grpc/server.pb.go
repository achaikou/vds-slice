// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: internal/grpc/server.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CoordinateSystem int32

const (
	CoordinateSystem_ANNOTATION CoordinateSystem = 0
	CoordinateSystem_INDEX      CoordinateSystem = 1
	CoordinateSystem_WORLD      CoordinateSystem = 2
)

// Enum value maps for CoordinateSystem.
var (
	CoordinateSystem_name = map[int32]string{
		0: "ANNOTATION",
		1: "INDEX",
		2: "WORLD",
	}
	CoordinateSystem_value = map[string]int32{
		"ANNOTATION": 0,
		"INDEX":      1,
		"WORLD":      2,
	}
)

func (x CoordinateSystem) Enum() *CoordinateSystem {
	p := new(CoordinateSystem)
	*p = x
	return p
}

func (x CoordinateSystem) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoordinateSystem) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_grpc_server_proto_enumTypes[0].Descriptor()
}

func (CoordinateSystem) Type() protoreflect.EnumType {
	return &file_internal_grpc_server_proto_enumTypes[0]
}

func (x CoordinateSystem) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoordinateSystem.Descriptor instead.
func (CoordinateSystem) EnumDescriptor() ([]byte, []int) {
	return file_internal_grpc_server_proto_rawDescGZIP(), []int{0}
}

type Interpolation int32

const (
	Interpolation_NEAREST    Interpolation = 0
	Interpolation_LINEAR     Interpolation = 1
	Interpolation_CUBIC      Interpolation = 2
	Interpolation_ANGULAR    Interpolation = 3
	Interpolation_TRIANGULAR Interpolation = 4
)

// Enum value maps for Interpolation.
var (
	Interpolation_name = map[int32]string{
		0: "NEAREST",
		1: "LINEAR",
		2: "CUBIC",
		3: "ANGULAR",
		4: "TRIANGULAR",
	}
	Interpolation_value = map[string]int32{
		"NEAREST":    0,
		"LINEAR":     1,
		"CUBIC":      2,
		"ANGULAR":    3,
		"TRIANGULAR": 4,
	}
)

func (x Interpolation) Enum() *Interpolation {
	p := new(Interpolation)
	*p = x
	return p
}

func (x Interpolation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Interpolation) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_grpc_server_proto_enumTypes[1].Descriptor()
}

func (Interpolation) Type() protoreflect.EnumType {
	return &file_internal_grpc_server_proto_enumTypes[1]
}

func (x Interpolation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Interpolation.Descriptor instead.
func (Interpolation) EnumDescriptor() ([]byte, []int) {
	return file_internal_grpc_server_proto_rawDescGZIP(), []int{1}
}

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid          string `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Part         int32  `protobuf:"varint,2,opt,name=part,proto3" json:"part,omitempty"`
	Parts        int32  `protobuf:"varint,3,opt,name=parts,proto3" json:"parts,omitempty"`
	StartSampleN int32  `protobuf:"varint,4,opt,name=startSampleN,proto3" json:"startSampleN,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_internal_grpc_server_proto_rawDescGZIP(), []int{0}
}

func (x *Info) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *Info) GetPart() int32 {
	if x != nil {
		return x.Part
	}
	return 0
}

func (x *Info) GetParts() int32 {
	if x != nil {
		return x.Parts
	}
	return 0
}

func (x *Info) GetStartSampleN() int32 {
	if x != nil {
		return x.StartSampleN
	}
	return 0
}

type Connection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url        string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Credential string `protobuf:"bytes,2,opt,name=credential,proto3" json:"credential,omitempty"`
}

func (x *Connection) Reset() {
	*x = Connection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connection) ProtoMessage() {}

func (x *Connection) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connection.ProtoReflect.Descriptor instead.
func (*Connection) Descriptor() ([]byte, []int) {
	return file_internal_grpc_server_proto_rawDescGZIP(), []int{1}
}

func (x *Connection) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Connection) GetCredential() string {
	if x != nil {
		return x.Credential
	}
	return ""
}

type Coordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Coordinate) Reset() {
	*x = Coordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinate) ProtoMessage() {}

func (x *Coordinate) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinate.ProtoReflect.Descriptor instead.
func (*Coordinate) Descriptor() ([]byte, []int) {
	return file_internal_grpc_server_proto_rawDescGZIP(), []int{2}
}

func (x *Coordinate) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Coordinate) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type FenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info             *Info         `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Connection       *Connection   `protobuf:"bytes,2,opt,name=connection,proto3" json:"connection,omitempty"`
	CoordinateSystem int32         `protobuf:"varint,3,opt,name=coordinateSystem,proto3" json:"coordinateSystem,omitempty"` // TODO change to enum
	Interpolation    int32         `protobuf:"varint,4,opt,name=interpolation,proto3" json:"interpolation,omitempty"`       //TODO change to enum
	Coordinates      []*Coordinate `protobuf:"bytes,5,rep,name=coordinates,proto3" json:"coordinates,omitempty"`
}

func (x *FenceRequest) Reset() {
	*x = FenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FenceRequest) ProtoMessage() {}

func (x *FenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FenceRequest.ProtoReflect.Descriptor instead.
func (*FenceRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_server_proto_rawDescGZIP(), []int{3}
}

func (x *FenceRequest) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *FenceRequest) GetConnection() *Connection {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *FenceRequest) GetCoordinateSystem() int32 {
	if x != nil {
		return x.CoordinateSystem
	}
	return 0
}

func (x *FenceRequest) GetInterpolation() int32 {
	if x != nil {
		return x.Interpolation
	}
	return 0
}

func (x *FenceRequest) GetCoordinates() []*Coordinate {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

type FenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info  *Info  `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Fence []byte `protobuf:"bytes,2,opt,name=fence,proto3" json:"fence,omitempty"`
}

func (x *FenceResponse) Reset() {
	*x = FenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FenceResponse) ProtoMessage() {}

func (x *FenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FenceResponse.ProtoReflect.Descriptor instead.
func (*FenceResponse) Descriptor() ([]byte, []int) {
	return file_internal_grpc_server_proto_rawDescGZIP(), []int{4}
}

func (x *FenceResponse) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *FenceResponse) GetFence() []byte {
	if x != nil {
		return x.Fence
	}
	return nil
}

var File_internal_grpc_server_proto protoreflect.FileDescriptor

var file_internal_grpc_server_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x6e,
	0x65, 0x73, 0x65, 0x69, 0x73, 0x6d, 0x69, 0x63, 0x22, 0x66, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4e,
	0x22, 0x3e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x22, 0x28, 0x0a, 0x0a, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x0c,
	0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x22, 0xf8, 0x01, 0x0a, 0x0c, 0x46,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x6e, 0x65, 0x73,
	0x65, 0x69, 0x73, 0x6d, 0x69, 0x63, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x36, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x6e, 0x65, 0x73, 0x65, 0x69, 0x73, 0x6d,
	0x69, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x6f,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x0b, 0x63,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6f, 0x6e, 0x65, 0x73, 0x65, 0x69, 0x73, 0x6d, 0x69, 0x63, 0x2e, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x73, 0x22, 0x4b, 0x0a, 0x0d, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x6e, 0x65, 0x73, 0x65, 0x69, 0x73, 0x6d, 0x69,
	0x63, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x66, 0x65, 0x6e,
	0x63, 0x65, 0x2a, 0x38, 0x0a, 0x10, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x10, 0x02, 0x2a, 0x50, 0x0a, 0x0d,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a,
	0x07, 0x4e, 0x45, 0x41, 0x52, 0x45, 0x53, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x49,
	0x4e, 0x45, 0x41, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x55, 0x42, 0x49, 0x43, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x4e, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x10, 0x03, 0x12, 0x0e,
	0x0a, 0x0a, 0x54, 0x52, 0x49, 0x41, 0x4e, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x10, 0x04, 0x32, 0x4f,
	0x0a, 0x0a, 0x4f, 0x6e, 0x65, 0x73, 0x65, 0x69, 0x73, 0x6d, 0x69, 0x63, 0x12, 0x41, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x6f, 0x6e, 0x65, 0x73, 0x65,
	0x69, 0x73, 0x6d, 0x69, 0x63, 0x2e, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x6e, 0x65, 0x73, 0x65, 0x69, 0x73, 0x6d, 0x69, 0x63, 0x2e,
	0x46, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x71,
	0x75, 0x69, 0x6e, 0x6f, 0x72, 0x2f, 0x76, 0x64, 0x73, 0x2d, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_grpc_server_proto_rawDescOnce sync.Once
	file_internal_grpc_server_proto_rawDescData = file_internal_grpc_server_proto_rawDesc
)

func file_internal_grpc_server_proto_rawDescGZIP() []byte {
	file_internal_grpc_server_proto_rawDescOnce.Do(func() {
		file_internal_grpc_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_grpc_server_proto_rawDescData)
	})
	return file_internal_grpc_server_proto_rawDescData
}

var file_internal_grpc_server_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_internal_grpc_server_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_grpc_server_proto_goTypes = []interface{}{
	(CoordinateSystem)(0), // 0: oneseismic.CoordinateSystem
	(Interpolation)(0),    // 1: oneseismic.Interpolation
	(*Info)(nil),          // 2: oneseismic.Info
	(*Connection)(nil),    // 3: oneseismic.Connection
	(*Coordinate)(nil),    // 4: oneseismic.Coordinate
	(*FenceRequest)(nil),  // 5: oneseismic.FenceRequest
	(*FenceResponse)(nil), // 6: oneseismic.FenceResponse
}
var file_internal_grpc_server_proto_depIdxs = []int32{
	2, // 0: oneseismic.FenceRequest.info:type_name -> oneseismic.Info
	3, // 1: oneseismic.FenceRequest.connection:type_name -> oneseismic.Connection
	4, // 2: oneseismic.FenceRequest.coordinates:type_name -> oneseismic.Coordinate
	2, // 3: oneseismic.FenceResponse.info:type_name -> oneseismic.Info
	5, // 4: oneseismic.Oneseismic.GetFence:input_type -> oneseismic.FenceRequest
	6, // 5: oneseismic.Oneseismic.GetFence:output_type -> oneseismic.FenceResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_grpc_server_proto_init() }
func file_internal_grpc_server_proto_init() {
	if File_internal_grpc_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_grpc_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_grpc_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_grpc_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coordinate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_grpc_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FenceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_grpc_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FenceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_grpc_server_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_grpc_server_proto_goTypes,
		DependencyIndexes: file_internal_grpc_server_proto_depIdxs,
		EnumInfos:         file_internal_grpc_server_proto_enumTypes,
		MessageInfos:      file_internal_grpc_server_proto_msgTypes,
	}.Build()
	File_internal_grpc_server_proto = out.File
	file_internal_grpc_server_proto_rawDesc = nil
	file_internal_grpc_server_proto_goTypes = nil
	file_internal_grpc_server_proto_depIdxs = nil
}
