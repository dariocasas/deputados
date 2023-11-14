// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: proto/fotos_service.proto

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

type GetFotosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFotosRequest) Reset() {
	*x = GetFotosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fotos_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFotosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFotosRequest) ProtoMessage() {}

func (x *GetFotosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fotos_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFotosRequest.ProtoReflect.Descriptor instead.
func (*GetFotosRequest) Descriptor() ([]byte, []int) {
	return file_proto_fotos_service_proto_rawDescGZIP(), []int{0}
}

type Redes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Red []string `protobuf:"bytes,1,rep,name=red,proto3" json:"red,omitempty"`
}

func (x *Redes) Reset() {
	*x = Redes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fotos_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Redes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Redes) ProtoMessage() {}

func (x *Redes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fotos_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Redes.ProtoReflect.Descriptor instead.
func (*Redes) Descriptor() ([]byte, []int) {
	return file_proto_fotos_service_proto_rawDescGZIP(), []int{1}
}

func (x *Redes) GetRed() []string {
	if x != nil {
		return x.Red
	}
	return nil
}

type FotoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FotoUrl  string `protobuf:"bytes,2,opt,name=fotoUrl,proto3" json:"fotoUrl,omitempty"`
	Nome     string `protobuf:"bytes,3,opt,name=nome,proto3" json:"nome,omitempty"`
	Partido  string `protobuf:"bytes,4,opt,name=partido,proto3" json:"partido,omitempty"`
	Uf       string `protobuf:"bytes,5,opt,name=uf,proto3" json:"uf,omitempty"`
	Url      string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	Redes    *Redes `protobuf:"bytes,7,opt,name=redes,proto3" json:"redes,omitempty"`
	FotoUrl2 string `protobuf:"bytes,8,opt,name=fotoUrl2,proto3" json:"fotoUrl2,omitempty"`
}

func (x *FotoResponse) Reset() {
	*x = FotoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fotos_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FotoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FotoResponse) ProtoMessage() {}

func (x *FotoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fotos_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FotoResponse.ProtoReflect.Descriptor instead.
func (*FotoResponse) Descriptor() ([]byte, []int) {
	return file_proto_fotos_service_proto_rawDescGZIP(), []int{2}
}

func (x *FotoResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FotoResponse) GetFotoUrl() string {
	if x != nil {
		return x.FotoUrl
	}
	return ""
}

func (x *FotoResponse) GetNome() string {
	if x != nil {
		return x.Nome
	}
	return ""
}

func (x *FotoResponse) GetPartido() string {
	if x != nil {
		return x.Partido
	}
	return ""
}

func (x *FotoResponse) GetUf() string {
	if x != nil {
		return x.Uf
	}
	return ""
}

func (x *FotoResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *FotoResponse) GetRedes() *Redes {
	if x != nil {
		return x.Redes
	}
	return nil
}

func (x *FotoResponse) GetFotoUrl2() string {
	if x != nil {
		return x.FotoUrl2
	}
	return ""
}

var File_proto_fotos_service_proto protoreflect.FileDescriptor

var file_proto_fotos_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x19, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x64, 0x22, 0xc5, 0x01,
	0x0a, 0x0c, 0x46, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x66, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x66, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x64, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x75, 0x66, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x75, 0x66, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x05, 0x72, 0x65, 0x64, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x64,
	0x65, 0x73, 0x52, 0x05, 0x72, 0x65, 0x64, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x74,
	0x6f, 0x55, 0x72, 0x6c, 0x32, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x74,
	0x6f, 0x55, 0x72, 0x6c, 0x32, 0x32, 0x45, 0x0a, 0x0c, 0x46, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x74, 0x6f,
	0x73, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x74, 0x6f, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x6f, 0x74, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x1c, 0x5a, 0x1a,
	0x73, 0x72, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_fotos_service_proto_rawDescOnce sync.Once
	file_proto_fotos_service_proto_rawDescData = file_proto_fotos_service_proto_rawDesc
)

func file_proto_fotos_service_proto_rawDescGZIP() []byte {
	file_proto_fotos_service_proto_rawDescOnce.Do(func() {
		file_proto_fotos_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_fotos_service_proto_rawDescData)
	})
	return file_proto_fotos_service_proto_rawDescData
}

var file_proto_fotos_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_fotos_service_proto_goTypes = []interface{}{
	(*GetFotosRequest)(nil), // 0: pb.GetFotosRequest
	(*Redes)(nil),           // 1: pb.Redes
	(*FotoResponse)(nil),    // 2: pb.FotoResponse
}
var file_proto_fotos_service_proto_depIdxs = []int32{
	1, // 0: pb.FotoResponse.redes:type_name -> pb.Redes
	0, // 1: pb.FotosService.GetFotos:input_type -> pb.GetFotosRequest
	2, // 2: pb.FotosService.GetFotos:output_type -> pb.FotoResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_fotos_service_proto_init() }
func file_proto_fotos_service_proto_init() {
	if File_proto_fotos_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_fotos_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFotosRequest); i {
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
		file_proto_fotos_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Redes); i {
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
		file_proto_fotos_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FotoResponse); i {
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
			RawDescriptor: file_proto_fotos_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_fotos_service_proto_goTypes,
		DependencyIndexes: file_proto_fotos_service_proto_depIdxs,
		MessageInfos:      file_proto_fotos_service_proto_msgTypes,
	}.Build()
	File_proto_fotos_service_proto = out.File
	file_proto_fotos_service_proto_rawDesc = nil
	file_proto_fotos_service_proto_goTypes = nil
	file_proto_fotos_service_proto_depIdxs = nil
}