// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: calculator.proto

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

type AddTwoNumbRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNum      float64 `protobuf:"fixed64,1,opt,name=first_num,json=firstNum,proto3" json:"first_num,omitempty"`
	SecondNum     float64 `protobuf:"fixed64,2,opt,name=second_num,json=secondNum,proto3" json:"second_num,omitempty"`
	OperationType int64   `protobuf:"varint,3,opt,name=operation_type,json=operationType,proto3" json:"operation_type,omitempty"`
}

func (x *AddTwoNumbRequest) Reset() {
	*x = AddTwoNumbRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTwoNumbRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTwoNumbRequest) ProtoMessage() {}

func (x *AddTwoNumbRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTwoNumbRequest.ProtoReflect.Descriptor instead.
func (*AddTwoNumbRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *AddTwoNumbRequest) GetFirstNum() float64 {
	if x != nil {
		return x.FirstNum
	}
	return 0
}

func (x *AddTwoNumbRequest) GetSecondNum() float64 {
	if x != nil {
		return x.SecondNum
	}
	return 0
}

func (x *AddTwoNumbRequest) GetOperationType() int64 {
	if x != nil {
		return x.OperationType
	}
	return 0
}

type AdditionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	ErrResult string  `protobuf:"bytes,2,opt,name=err_result,json=errResult,proto3" json:"err_result,omitempty"`
}

func (x *AdditionResponse) Reset() {
	*x = AdditionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdditionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdditionResponse) ProtoMessage() {}

func (x *AdditionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdditionResponse.ProtoReflect.Descriptor instead.
func (*AdditionResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *AdditionResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *AdditionResponse) GetErrResult() string {
	if x != nil {
		return x.ErrResult
	}
	return ""
}

var File_calculator_proto protoreflect.FileDescriptor

var file_calculator_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x76,
	0x0a, 0x11, 0x41, 0x64, 0x64, 0x54, 0x77, 0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x12,
	0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x49, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0x5f, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x54, 0x77, 0x6f, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x77, 0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x64, 0x65, 0x6b, 0x75, 0x6e, 0x6c, 0x65, 0x44, 0x61, 0x6c, 0x6c, 0x79, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculator_proto_rawDescOnce sync.Once
	file_calculator_proto_rawDescData = file_calculator_proto_rawDesc
)

func file_calculator_proto_rawDescGZIP() []byte {
	file_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_proto_rawDescData)
	})
	return file_calculator_proto_rawDescData
}

var file_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_calculator_proto_goTypes = []interface{}{
	(*AddTwoNumbRequest)(nil), // 0: calculator.AddTwoNumbRequest
	(*AdditionResponse)(nil),  // 1: calculator.AdditionResponse
}
var file_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.AdditionService.AddTwoNumbers:input_type -> calculator.AddTwoNumbRequest
	1, // 1: calculator.AdditionService.AddTwoNumbers:output_type -> calculator.AdditionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_proto_init() }
func file_calculator_proto_init() {
	if File_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTwoNumbRequest); i {
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
		file_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdditionResponse); i {
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
			RawDescriptor: file_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_proto_msgTypes,
	}.Build()
	File_calculator_proto = out.File
	file_calculator_proto_rawDesc = nil
	file_calculator_proto_goTypes = nil
	file_calculator_proto_depIdxs = nil
}
