// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: to_do/event/v1/to_do_created.proto

package eventv1

import (
	v1 "github.com/hampgoodwin/todo/gen/proto/go/to_do/model/v1"
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

type ToDoCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToDo *v1.ToDo `protobuf:"bytes,1,opt,name=to_do,json=toDo,proto3" json:"to_do,omitempty"`
}

func (x *ToDoCreated) Reset() {
	*x = ToDoCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_to_do_event_v1_to_do_created_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDoCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDoCreated) ProtoMessage() {}

func (x *ToDoCreated) ProtoReflect() protoreflect.Message {
	mi := &file_to_do_event_v1_to_do_created_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDoCreated.ProtoReflect.Descriptor instead.
func (*ToDoCreated) Descriptor() ([]byte, []int) {
	return file_to_do_event_v1_to_do_created_proto_rawDescGZIP(), []int{0}
}

func (x *ToDoCreated) GetToDo() *v1.ToDo {
	if x != nil {
		return x.ToDo
	}
	return nil
}

var File_to_do_event_v1_to_do_created_proto protoreflect.FileDescriptor

var file_to_do_event_v1_to_do_created_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x38, 0x0a, 0x0b, 0x54, 0x6f, 0x44, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x29, 0x0a, 0x05, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x6f, 0x44, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x44, 0x6f, 0x42, 0xbd, 0x01, 0x0a, 0x12, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x42, 0x10, 0x54, 0x6f, 0x44, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x61, 0x6d, 0x70, 0x67, 0x6f, 0x6f, 0x64, 0x77, 0x69, 0x6e, 0x2f, 0x74, 0x6f,
	0x64, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f,
	0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x45, 0x58, 0xaa, 0x02, 0x0d, 0x54,
	0x6f, 0x44, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x54,
	0x6f, 0x44, 0x6f, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x54,
	0x6f, 0x44, 0x6f, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x54, 0x6f, 0x44, 0x6f, 0x3a,
	0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_to_do_event_v1_to_do_created_proto_rawDescOnce sync.Once
	file_to_do_event_v1_to_do_created_proto_rawDescData = file_to_do_event_v1_to_do_created_proto_rawDesc
)

func file_to_do_event_v1_to_do_created_proto_rawDescGZIP() []byte {
	file_to_do_event_v1_to_do_created_proto_rawDescOnce.Do(func() {
		file_to_do_event_v1_to_do_created_proto_rawDescData = protoimpl.X.CompressGZIP(file_to_do_event_v1_to_do_created_proto_rawDescData)
	})
	return file_to_do_event_v1_to_do_created_proto_rawDescData
}

var file_to_do_event_v1_to_do_created_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_to_do_event_v1_to_do_created_proto_goTypes = []interface{}{
	(*ToDoCreated)(nil), // 0: to_do.event.v1.ToDoCreated
	(*v1.ToDo)(nil),     // 1: to_do.model.v1.ToDo
}
var file_to_do_event_v1_to_do_created_proto_depIdxs = []int32{
	1, // 0: to_do.event.v1.ToDoCreated.to_do:type_name -> to_do.model.v1.ToDo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_to_do_event_v1_to_do_created_proto_init() }
func file_to_do_event_v1_to_do_created_proto_init() {
	if File_to_do_event_v1_to_do_created_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_to_do_event_v1_to_do_created_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDoCreated); i {
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
			RawDescriptor: file_to_do_event_v1_to_do_created_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_to_do_event_v1_to_do_created_proto_goTypes,
		DependencyIndexes: file_to_do_event_v1_to_do_created_proto_depIdxs,
		MessageInfos:      file_to_do_event_v1_to_do_created_proto_msgTypes,
	}.Build()
	File_to_do_event_v1_to_do_created_proto = out.File
	file_to_do_event_v1_to_do_created_proto_rawDesc = nil
	file_to_do_event_v1_to_do_created_proto_goTypes = nil
	file_to_do_event_v1_to_do_created_proto_depIdxs = nil
}