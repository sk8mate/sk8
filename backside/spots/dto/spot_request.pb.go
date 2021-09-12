// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: proto/spot_request.proto

package dto

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

type SpotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address     string       `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Coordinates *Coordinates `protobuf:"bytes,3,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
	Lighting    bool         `protobuf:"varint,4,opt,name=lighting,proto3" json:"lighting,omitempty"`
	Friendly    bool         `protobuf:"varint,5,opt,name=friendly,proto3" json:"friendly,omitempty"`
	Verified    bool         `protobuf:"varint,6,opt,name=verified,proto3" json:"verified,omitempty"`
}

func (x *SpotRequest) Reset() {
	*x = SpotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_spot_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpotRequest) ProtoMessage() {}

func (x *SpotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spot_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpotRequest.ProtoReflect.Descriptor instead.
func (*SpotRequest) Descriptor() ([]byte, []int) {
	return file_proto_spot_request_proto_rawDescGZIP(), []int{0}
}

func (x *SpotRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SpotRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SpotRequest) GetCoordinates() *Coordinates {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *SpotRequest) GetLighting() bool {
	if x != nil {
		return x.Lighting
	}
	return false
}

func (x *SpotRequest) GetFriendly() bool {
	if x != nil {
		return x.Friendly
	}
	return false
}

func (x *SpotRequest) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

type Coordinates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat  float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Long float64 `protobuf:"fixed64,2,opt,name=long,proto3" json:"long,omitempty"`
}

func (x *Coordinates) Reset() {
	*x = Coordinates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_spot_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinates) ProtoMessage() {}

func (x *Coordinates) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spot_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinates.ProtoReflect.Descriptor instead.
func (*Coordinates) Descriptor() ([]byte, []int) {
	return file_proto_spot_request_proto_rawDescGZIP(), []int{1}
}

func (x *Coordinates) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Coordinates) GetLong() float64 {
	if x != nil {
		return x.Long
	}
	return 0
}

var File_proto_spot_request_proto protoreflect.FileDescriptor

var file_proto_spot_request_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e,
	0x22, 0xc4, 0x01, 0x0a, 0x0b, 0x53, 0x70, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x33,
	0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x33, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2f, 0x64, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_spot_request_proto_rawDescOnce sync.Once
	file_proto_spot_request_proto_rawDescData = file_proto_spot_request_proto_rawDesc
)

func file_proto_spot_request_proto_rawDescGZIP() []byte {
	file_proto_spot_request_proto_rawDescOnce.Do(func() {
		file_proto_spot_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_spot_request_proto_rawDescData)
	})
	return file_proto_spot_request_proto_rawDescData
}

var file_proto_spot_request_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_spot_request_proto_goTypes = []interface{}{
	(*SpotRequest)(nil), // 0: main.SpotRequest
	(*Coordinates)(nil), // 1: main.Coordinates
}
var file_proto_spot_request_proto_depIdxs = []int32{
	1, // 0: main.SpotRequest.coordinates:type_name -> main.Coordinates
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_spot_request_proto_init() }
func file_proto_spot_request_proto_init() {
	if File_proto_spot_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_spot_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpotRequest); i {
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
		file_proto_spot_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coordinates); i {
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
			RawDescriptor: file_proto_spot_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_spot_request_proto_goTypes,
		DependencyIndexes: file_proto_spot_request_proto_depIdxs,
		MessageInfos:      file_proto_spot_request_proto_msgTypes,
	}.Build()
	File_proto_spot_request_proto = out.File
	file_proto_spot_request_proto_rawDesc = nil
	file_proto_spot_request_proto_goTypes = nil
	file_proto_spot_request_proto_depIdxs = nil
}
