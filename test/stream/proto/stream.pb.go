// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: test/stream/proto/stream.proto

package proto

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

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_stream_proto_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_test_stream_proto_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_test_stream_proto_stream_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Point) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location *Point `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_stream_proto_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_test_stream_proto_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_test_stream_proto_stream_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Response) GetLocation() *Point {
	if x != nil {
		return x.Location
	}
	return nil
}

var File_test_stream_proto_stream_proto protoreflect.FileDescriptor

var file_test_stream_proto_stream_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x23, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x42, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x57, 0x0a, 0x0a, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x54, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x06, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x53, 0x65, 0x6e,
	0x64, 0x4e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x06, 0x2e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x1d, 0x5a, 0x1b, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_stream_proto_stream_proto_rawDescOnce sync.Once
	file_test_stream_proto_stream_proto_rawDescData = file_test_stream_proto_stream_proto_rawDesc
)

func file_test_stream_proto_stream_proto_rawDescGZIP() []byte {
	file_test_stream_proto_stream_proto_rawDescOnce.Do(func() {
		file_test_stream_proto_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_stream_proto_stream_proto_rawDescData)
	})
	return file_test_stream_proto_stream_proto_rawDescData
}

var file_test_stream_proto_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_stream_proto_stream_proto_goTypes = []interface{}{
	(*Point)(nil),    // 0: Point
	(*Response)(nil), // 1: Response
}
var file_test_stream_proto_stream_proto_depIdxs = []int32{
	0, // 0: Response.location:type_name -> Point
	0, // 1: StreamTest.GetData:input_type -> Point
	0, // 2: StreamTest.SendNGetData:input_type -> Point
	1, // 3: StreamTest.GetData:output_type -> Response
	1, // 4: StreamTest.SendNGetData:output_type -> Response
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_test_stream_proto_stream_proto_init() }
func file_test_stream_proto_stream_proto_init() {
	if File_test_stream_proto_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_stream_proto_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_test_stream_proto_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_test_stream_proto_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_stream_proto_stream_proto_goTypes,
		DependencyIndexes: file_test_stream_proto_stream_proto_depIdxs,
		MessageInfos:      file_test_stream_proto_stream_proto_msgTypes,
	}.Build()
	File_test_stream_proto_stream_proto = out.File
	file_test_stream_proto_stream_proto_rawDesc = nil
	file_test_stream_proto_stream_proto_goTypes = nil
	file_test_stream_proto_stream_proto_depIdxs = nil
}
