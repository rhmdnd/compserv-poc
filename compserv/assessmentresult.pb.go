// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.14.0
// source: compserv/assessmentresult.proto

package compserv

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

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Control string `protobuf:"bytes,2,opt,name=control,proto3" json:"control,omitempty"`
	Rule    string `protobuf:"bytes,3,opt,name=rule,proto3" json:"rule,omitempty"`
	Outcome string `protobuf:"bytes,4,opt,name=outcome,proto3" json:"outcome,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compserv_assessmentresult_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_compserv_assessmentresult_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_compserv_assessmentresult_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Result) GetControl() string {
	if x != nil {
		return x.Control
	}
	return ""
}

func (x *Result) GetRule() string {
	if x != nil {
		return x.Rule
	}
	return ""
}

func (x *Result) GetOutcome() string {
	if x != nil {
		return x.Outcome
	}
	return ""
}

var File_compserv_assessmentresult_proto protoreflect.FileDescriptor

var file_compserv_assessmentresult_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x70, 0x73, 0x65, 0x72, 0x76, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x73,
	0x73, 0x6d, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6a, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72,
	0x75, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x32, 0x3d, 0x0a,
	0x10, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x29, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x68, 0x6d, 0x64, 0x6e,
	0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x73, 0x65, 0x72, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_compserv_assessmentresult_proto_rawDescOnce sync.Once
	file_compserv_assessmentresult_proto_rawDescData = file_compserv_assessmentresult_proto_rawDesc
)

func file_compserv_assessmentresult_proto_rawDescGZIP() []byte {
	file_compserv_assessmentresult_proto_rawDescOnce.Do(func() {
		file_compserv_assessmentresult_proto_rawDescData = protoimpl.X.CompressGZIP(file_compserv_assessmentresult_proto_rawDescData)
	})
	return file_compserv_assessmentresult_proto_rawDescData
}

var file_compserv_assessmentresult_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_compserv_assessmentresult_proto_goTypes = []interface{}{
	(*Result)(nil), // 0: Result
}
var file_compserv_assessmentresult_proto_depIdxs = []int32{
	0, // 0: AssessmentResult.AddAssessmentResult:input_type -> Result
	0, // 1: AssessmentResult.AddAssessmentResult:output_type -> Result
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_compserv_assessmentresult_proto_init() }
func file_compserv_assessmentresult_proto_init() {
	if File_compserv_assessmentresult_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_compserv_assessmentresult_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_compserv_assessmentresult_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_compserv_assessmentresult_proto_goTypes,
		DependencyIndexes: file_compserv_assessmentresult_proto_depIdxs,
		MessageInfos:      file_compserv_assessmentresult_proto_msgTypes,
	}.Build()
	File_compserv_assessmentresult_proto = out.File
	file_compserv_assessmentresult_proto_rawDesc = nil
	file_compserv_assessmentresult_proto_goTypes = nil
	file_compserv_assessmentresult_proto_depIdxs = nil
}