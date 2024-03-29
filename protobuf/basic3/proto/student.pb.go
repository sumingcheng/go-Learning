// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: student.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Course int32

const (
	Course_GO     Course = 0
	Course_JAVA   Course = 1
	Course_PYTHON Course = 2
)

// Enum value maps for Course.
var (
	Course_name = map[int32]string{
		0: "GO",
		1: "JAVA",
		2: "PYTHON",
	}
	Course_value = map[string]int32{
		"GO":     0,
		"JAVA":   1,
		"PYTHON": 2,
	}
)

func (x Course) Enum() *Course {
	p := new(Course)
	*p = x
	return p
}

func (x Course) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Course) Descriptor() protoreflect.EnumDescriptor {
	return file_student_proto_enumTypes[0].Descriptor()
}

func (Course) Type() protoreflect.EnumType {
	return &file_student_proto_enumTypes[0]
}

func (x Course) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Course.Descriptor instead.
func (Course) EnumDescriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{0}
}

type Gender int32

const (
	Gender_MALE   Gender = 0
	Gender_FEMALE Gender = 1
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "MALE",
		1: "FEMALE",
	}
	Gender_value = map[string]int32{
		"MALE":   0,
		"FEMALE": 1,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_student_proto_enumTypes[1].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_student_proto_enumTypes[1]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{1}
}

type StudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid int64 `protobuf:"varint,1,opt,name=sid,proto3" json:"sid,omitempty"`
}

func (x *StudentRequest) Reset() {
	*x = StudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentRequest) ProtoMessage() {}

func (x *StudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentRequest.ProtoReflect.Descriptor instead.
func (*StudentRequest) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{0}
}

func (x *StudentRequest) GetSid() int64 {
	if x != nil {
		return x.Sid
	}
	return 0
}

type Scores struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interview int32 `protobuf:"varint,1,opt,name=interview,proto3" json:"interview,omitempty"`
	Written   int32 `protobuf:"varint,2,opt,name=written,proto3" json:"written,omitempty"`
}

func (x *Scores) Reset() {
	*x = Scores{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scores) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scores) ProtoMessage() {}

func (x *Scores) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scores.ProtoReflect.Descriptor instead.
func (*Scores) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{1}
}

func (x *Scores) GetInterview() int32 {
	if x != nil {
		return x.Interview
	}
	return 0
}

func (x *Scores) GetWritten() int32 {
	if x != nil {
		return x.Written
	}
	return 0
}

type StudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid   int64  `protobuf:"varint,1,opt,name=sid,proto3" json:"sid,omitempty"`
	SName string `protobuf:"bytes,2,opt,name=sName,proto3" json:"sName,omitempty"`
	Age   string `protobuf:"bytes,3,opt,name=age,proto3" json:"age,omitempty"`
	// 枚举
	Course Course `protobuf:"varint,4,opt,name=course,proto3,enum=Course" json:"course,omitempty"`
	Gender Gender `protobuf:"varint,5,opt,name=gender,proto3,enum=Gender" json:"gender,omitempty"`
	// 嵌套
	Scores *Scores `protobuf:"bytes,6,opt,name=scores,proto3" json:"scores,omitempty"`
	// 时间
	JoinTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=joinTime,proto3" json:"joinTime,omitempty"`
}

func (x *StudentResponse) Reset() {
	*x = StudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentResponse) ProtoMessage() {}

func (x *StudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentResponse.ProtoReflect.Descriptor instead.
func (*StudentResponse) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{2}
}

func (x *StudentResponse) GetSid() int64 {
	if x != nil {
		return x.Sid
	}
	return 0
}

func (x *StudentResponse) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

func (x *StudentResponse) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

func (x *StudentResponse) GetCourse() Course {
	if x != nil {
		return x.Course
	}
	return Course_GO
}

func (x *StudentResponse) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_MALE
}

func (x *StudentResponse) GetScores() *Scores {
	if x != nil {
		return x.Scores
	}
	return nil
}

func (x *StudentResponse) GetJoinTime() *timestamppb.Timestamp {
	if x != nil {
		return x.JoinTime
	}
	return nil
}

var File_student_proto protoreflect.FileDescriptor

var file_student_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x22, 0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x73, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x06, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12, 0x18, 0x0a, 0x07,
	0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x77,
	0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x22, 0xe6, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x06, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x52,
	0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x2a,
	0x26, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x4f, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x41, 0x56, 0x41, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50,
	0x59, 0x54, 0x48, 0x4f, 0x4e, 0x10, 0x02, 0x2a, 0x1e, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46,
	0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x32, 0x41, 0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_student_proto_rawDescOnce sync.Once
	file_student_proto_rawDescData = file_student_proto_rawDesc
)

func file_student_proto_rawDescGZIP() []byte {
	file_student_proto_rawDescOnce.Do(func() {
		file_student_proto_rawDescData = protoimpl.X.CompressGZIP(file_student_proto_rawDescData)
	})
	return file_student_proto_rawDescData
}

var file_student_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_student_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_student_proto_goTypes = []interface{}{
	(Course)(0),                   // 0: Course
	(Gender)(0),                   // 1: Gender
	(*StudentRequest)(nil),        // 2: StudentRequest
	(*Scores)(nil),                // 3: Scores
	(*StudentResponse)(nil),       // 4: StudentResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_student_proto_depIdxs = []int32{
	0, // 0: StudentResponse.course:type_name -> Course
	1, // 1: StudentResponse.gender:type_name -> Gender
	3, // 2: StudentResponse.scores:type_name -> Scores
	5, // 3: StudentResponse.joinTime:type_name -> google.protobuf.Timestamp
	2, // 4: StudentService.GetStudent:input_type -> StudentRequest
	4, // 5: StudentService.GetStudent:output_type -> StudentResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_student_proto_init() }
func file_student_proto_init() {
	if File_student_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_student_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentRequest); i {
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
		file_student_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scores); i {
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
		file_student_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentResponse); i {
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
			RawDescriptor: file_student_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_student_proto_goTypes,
		DependencyIndexes: file_student_proto_depIdxs,
		EnumInfos:         file_student_proto_enumTypes,
		MessageInfos:      file_student_proto_msgTypes,
	}.Build()
	File_student_proto = out.File
	file_student_proto_rawDesc = nil
	file_student_proto_goTypes = nil
	file_student_proto_depIdxs = nil
}
