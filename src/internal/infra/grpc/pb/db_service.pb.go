// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: proto/db_service.proto

package pb

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

type DbStatusEnum int32

const (
	DbStatusEnum_UNINITIALIZATED DbStatusEnum = 0
	DbStatusEnum_INDEX_CREATED   DbStatusEnum = 1
	DbStatusEnum_DB_CREATED      DbStatusEnum = 2
)

// Enum value maps for DbStatusEnum.
var (
	DbStatusEnum_name = map[int32]string{
		0: "UNINITIALIZATED",
		1: "INDEX_CREATED",
		2: "DB_CREATED",
	}
	DbStatusEnum_value = map[string]int32{
		"UNINITIALIZATED": 0,
		"INDEX_CREATED":   1,
		"DB_CREATED":      2,
	}
)

func (x DbStatusEnum) Enum() *DbStatusEnum {
	p := new(DbStatusEnum)
	*p = x
	return p
}

func (x DbStatusEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DbStatusEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_db_service_proto_enumTypes[0].Descriptor()
}

func (DbStatusEnum) Type() protoreflect.EnumType {
	return &file_proto_db_service_proto_enumTypes[0]
}

func (x DbStatusEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DbStatusEnum.Descriptor instead.
func (DbStatusEnum) EnumDescriptor() ([]byte, []int) {
	return file_proto_db_service_proto_rawDescGZIP(), []int{0}
}

type DropDbRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DropDbRequest) Reset() {
	*x = DropDbRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropDbRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropDbRequest) ProtoMessage() {}

func (x *DropDbRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropDbRequest.ProtoReflect.Descriptor instead.
func (*DropDbRequest) Descriptor() ([]byte, []int) {
	return file_proto_db_service_proto_rawDescGZIP(), []int{0}
}

type DropDbResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DropDbResponse) Reset() {
	*x = DropDbResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropDbResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropDbResponse) ProtoMessage() {}

func (x *DropDbResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropDbResponse.ProtoReflect.Descriptor instead.
func (*DropDbResponse) Descriptor() ([]byte, []int) {
	return file_proto_db_service_proto_rawDescGZIP(), []int{1}
}

type DbStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DbStatusRequest) Reset() {
	*x = DbStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DbStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbStatusRequest) ProtoMessage() {}

func (x *DbStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbStatusRequest.ProtoReflect.Descriptor instead.
func (*DbStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_db_service_proto_rawDescGZIP(), []int{2}
}

type DbStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     DbStatusEnum `protobuf:"varint,1,opt,name=status,proto3,enum=pb.DbStatusEnum" json:"status,omitempty"`
	IndexCount int32        `protobuf:"varint,2,opt,name=indexCount,proto3" json:"indexCount,omitempty"`
}

func (x *DbStatusResponse) Reset() {
	*x = DbStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DbStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbStatusResponse) ProtoMessage() {}

func (x *DbStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbStatusResponse.ProtoReflect.Descriptor instead.
func (*DbStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_db_service_proto_rawDescGZIP(), []int{3}
}

func (x *DbStatusResponse) GetStatus() DbStatusEnum {
	if x != nil {
		return x.Status
	}
	return DbStatusEnum_UNINITIALIZATED
}

func (x *DbStatusResponse) GetIndexCount() int32 {
	if x != nil {
		return x.IndexCount
	}
	return 0
}

type PopulateIndexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxItems int32 `protobuf:"varint,1,opt,name=maxItems,proto3" json:"maxItems,omitempty"`
}

func (x *PopulateIndexRequest) Reset() {
	*x = PopulateIndexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopulateIndexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopulateIndexRequest) ProtoMessage() {}

func (x *PopulateIndexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopulateIndexRequest.ProtoReflect.Descriptor instead.
func (*PopulateIndexRequest) Descriptor() ([]byte, []int) {
	return file_proto_db_service_proto_rawDescGZIP(), []int{4}
}

func (x *PopulateIndexRequest) GetMaxItems() int32 {
	if x != nil {
		return x.MaxItems
	}
	return 0
}

type PopulateIndexResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IndexCount int32 `protobuf:"varint,1,opt,name=indexCount,proto3" json:"indexCount,omitempty"`
}

func (x *PopulateIndexResponse) Reset() {
	*x = PopulateIndexResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopulateIndexResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopulateIndexResponse) ProtoMessage() {}

func (x *PopulateIndexResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopulateIndexResponse.ProtoReflect.Descriptor instead.
func (*PopulateIndexResponse) Descriptor() ([]byte, []int) {
	return file_proto_db_service_proto_rawDescGZIP(), []int{5}
}

func (x *PopulateIndexResponse) GetIndexCount() int32 {
	if x != nil {
		return x.IndexCount
	}
	return 0
}

type PopulateDbRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Concurrency int32 `protobuf:"varint,1,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	Timeout     int32 `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *PopulateDbRequest) Reset() {
	*x = PopulateDbRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopulateDbRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopulateDbRequest) ProtoMessage() {}

func (x *PopulateDbRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopulateDbRequest.ProtoReflect.Descriptor instead.
func (*PopulateDbRequest) Descriptor() ([]byte, []int) {
	return file_proto_db_service_proto_rawDescGZIP(), []int{6}
}

func (x *PopulateDbRequest) GetConcurrency() int32 {
	if x != nil {
		return x.Concurrency
	}
	return 0
}

func (x *PopulateDbRequest) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type PartialTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dbread   int32 `protobuf:"varint,1,opt,name=dbread,proto3" json:"dbread,omitempty"`
	Getapi   int32 `protobuf:"varint,2,opt,name=getapi,proto3" json:"getapi,omitempty"`
	Getpage  int32 `protobuf:"varint,3,opt,name=getpage,proto3" json:"getpage,omitempty"`
	Scanpage int32 `protobuf:"varint,4,opt,name=scanpage,proto3" json:"scanpage,omitempty"`
	Dbwrite  int32 `protobuf:"varint,5,opt,name=dbwrite,proto3" json:"dbwrite,omitempty"`
}

func (x *PartialTime) Reset() {
	*x = PartialTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartialTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartialTime) ProtoMessage() {}

func (x *PartialTime) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartialTime.ProtoReflect.Descriptor instead.
func (*PartialTime) Descriptor() ([]byte, []int) {
	return file_proto_db_service_proto_rawDescGZIP(), []int{7}
}

func (x *PartialTime) GetDbread() int32 {
	if x != nil {
		return x.Dbread
	}
	return 0
}

func (x *PartialTime) GetGetapi() int32 {
	if x != nil {
		return x.Getapi
	}
	return 0
}

func (x *PartialTime) GetGetpage() int32 {
	if x != nil {
		return x.Getpage
	}
	return 0
}

func (x *PartialTime) GetScanpage() int32 {
	if x != nil {
		return x.Scanpage
	}
	return 0
}

func (x *PartialTime) GetDbwrite() int32 {
	if x != nil {
		return x.Dbwrite
	}
	return 0
}

type DeputadoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error        bool         `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	ErrorString  string       `protobuf:"bytes,2,opt,name=errorString,proto3" json:"errorString,omitempty"`
	Id           int32        `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Name         string       `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Milliseconds int32        `protobuf:"varint,5,opt,name=milliseconds,proto3" json:"milliseconds,omitempty"`
	PartialTime  *PartialTime `protobuf:"bytes,6,opt,name=partialTime,proto3" json:"partialTime,omitempty"`
}

func (x *DeputadoResponse) Reset() {
	*x = DeputadoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeputadoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeputadoResponse) ProtoMessage() {}

func (x *DeputadoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeputadoResponse.ProtoReflect.Descriptor instead.
func (*DeputadoResponse) Descriptor() ([]byte, []int) {
	return file_proto_db_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeputadoResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *DeputadoResponse) GetErrorString() string {
	if x != nil {
		return x.ErrorString
	}
	return ""
}

func (x *DeputadoResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeputadoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeputadoResponse) GetMilliseconds() int32 {
	if x != nil {
		return x.Milliseconds
	}
	return 0
}

func (x *DeputadoResponse) GetPartialTime() *PartialTime {
	if x != nil {
		return x.PartialTime
	}
	return nil
}

type CancelPopulateDbRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelPopulateDbRequest) Reset() {
	*x = CancelPopulateDbRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelPopulateDbRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelPopulateDbRequest) ProtoMessage() {}

func (x *CancelPopulateDbRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelPopulateDbRequest.ProtoReflect.Descriptor instead.
func (*CancelPopulateDbRequest) Descriptor() ([]byte, []int) {
	return file_proto_db_service_proto_rawDescGZIP(), []int{9}
}

type CancelPopulateDbResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelPopulateDbResponse) Reset() {
	*x = CancelPopulateDbResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_db_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelPopulateDbResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelPopulateDbResponse) ProtoMessage() {}

func (x *CancelPopulateDbResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_db_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelPopulateDbResponse.ProtoReflect.Descriptor instead.
func (*CancelPopulateDbResponse) Descriptor() ([]byte, []int) {
	return file_proto_db_service_proto_rawDescGZIP(), []int{10}
}

var File_proto_db_service_proto protoreflect.FileDescriptor

var file_proto_db_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x0f, 0x0a, 0x0d,
	0x44, 0x72, 0x6f, 0x70, 0x44, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a,
	0x0e, 0x44, 0x72, 0x6f, 0x70, 0x44, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x11, 0x0a, 0x0f, 0x44, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x5c, 0x0a, 0x10, 0x44, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x32, 0x0a, 0x14, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x37, 0x0a, 0x15, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4f, 0x0a,
	0x11, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x8d,
	0x01, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x62, 0x72, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x64, 0x62, 0x72, 0x65, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x74, 0x61, 0x70, 0x69,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x74, 0x61, 0x70, 0x69, 0x12, 0x18,
	0x0a, 0x07, 0x67, 0x65, 0x74, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x67, 0x65, 0x74, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x61, 0x6e,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x63, 0x61, 0x6e,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x62, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x62, 0x77, 0x72, 0x69, 0x74, 0x65, 0x22, 0xc5,
	0x01, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x75, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x12, 0x31, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x1a, 0x0a, 0x18, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x6f, 0x70, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x44, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x46, 0x0a,
	0x0c, 0x44, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x13, 0x0a,
	0x0f, 0x55, 0x4e, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x42, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0xcf, 0x02, 0x0a, 0x09, 0x44, 0x62, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x44, 0x72, 0x6f, 0x70, 0x44, 0x62, 0x12, 0x11, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x44, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x44, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08, 0x44, 0x62, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x0d, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x50, 0x6f, 0x70, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x44, 0x62, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x70, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x44, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70,
	0x62, 0x2e, 0x44, 0x65, 0x70, 0x75, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4f, 0x0a, 0x10, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x62, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x73, 0x72, 0x63, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_db_service_proto_rawDescOnce sync.Once
	file_proto_db_service_proto_rawDescData = file_proto_db_service_proto_rawDesc
)

func file_proto_db_service_proto_rawDescGZIP() []byte {
	file_proto_db_service_proto_rawDescOnce.Do(func() {
		file_proto_db_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_db_service_proto_rawDescData)
	})
	return file_proto_db_service_proto_rawDescData
}

var file_proto_db_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_db_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_db_service_proto_goTypes = []interface{}{
	(DbStatusEnum)(0),                // 0: pb.DbStatusEnum
	(*DropDbRequest)(nil),            // 1: pb.DropDbRequest
	(*DropDbResponse)(nil),           // 2: pb.DropDbResponse
	(*DbStatusRequest)(nil),          // 3: pb.DbStatusRequest
	(*DbStatusResponse)(nil),         // 4: pb.DbStatusResponse
	(*PopulateIndexRequest)(nil),     // 5: pb.PopulateIndexRequest
	(*PopulateIndexResponse)(nil),    // 6: pb.PopulateIndexResponse
	(*PopulateDbRequest)(nil),        // 7: pb.PopulateDbRequest
	(*PartialTime)(nil),              // 8: pb.PartialTime
	(*DeputadoResponse)(nil),         // 9: pb.DeputadoResponse
	(*CancelPopulateDbRequest)(nil),  // 10: pb.CancelPopulateDbRequest
	(*CancelPopulateDbResponse)(nil), // 11: pb.CancelPopulateDbResponse
}
var file_proto_db_service_proto_depIdxs = []int32{
	0,  // 0: pb.DbStatusResponse.status:type_name -> pb.DbStatusEnum
	8,  // 1: pb.DeputadoResponse.partialTime:type_name -> pb.PartialTime
	1,  // 2: pb.DbService.DropDb:input_type -> pb.DropDbRequest
	3,  // 3: pb.DbService.DbStatus:input_type -> pb.DbStatusRequest
	5,  // 4: pb.DbService.PopulateIndex:input_type -> pb.PopulateIndexRequest
	7,  // 5: pb.DbService.PopulateDb:input_type -> pb.PopulateDbRequest
	10, // 6: pb.DbService.CancelPopulateDb:input_type -> pb.CancelPopulateDbRequest
	2,  // 7: pb.DbService.DropDb:output_type -> pb.DropDbResponse
	4,  // 8: pb.DbService.DbStatus:output_type -> pb.DbStatusResponse
	6,  // 9: pb.DbService.PopulateIndex:output_type -> pb.PopulateIndexResponse
	9,  // 10: pb.DbService.PopulateDb:output_type -> pb.DeputadoResponse
	11, // 11: pb.DbService.CancelPopulateDb:output_type -> pb.CancelPopulateDbResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_db_service_proto_init() }
func file_proto_db_service_proto_init() {
	if File_proto_db_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_db_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropDbRequest); i {
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
		file_proto_db_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropDbResponse); i {
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
		file_proto_db_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DbStatusRequest); i {
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
		file_proto_db_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DbStatusResponse); i {
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
		file_proto_db_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopulateIndexRequest); i {
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
		file_proto_db_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopulateIndexResponse); i {
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
		file_proto_db_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopulateDbRequest); i {
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
		file_proto_db_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartialTime); i {
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
		file_proto_db_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeputadoResponse); i {
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
		file_proto_db_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelPopulateDbRequest); i {
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
		file_proto_db_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelPopulateDbResponse); i {
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
			RawDescriptor: file_proto_db_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_db_service_proto_goTypes,
		DependencyIndexes: file_proto_db_service_proto_depIdxs,
		EnumInfos:         file_proto_db_service_proto_enumTypes,
		MessageInfos:      file_proto_db_service_proto_msgTypes,
	}.Build()
	File_proto_db_service_proto = out.File
	file_proto_db_service_proto_rawDesc = nil
	file_proto_db_service_proto_goTypes = nil
	file_proto_db_service_proto_depIdxs = nil
}
