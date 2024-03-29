// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: schemas/schema.proto

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

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId  uint32 `protobuf:"varint,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title     string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body      string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemas_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_schemas_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_schemas_schema_proto_rawDescGZIP(), []int{0}
}

func (x *Article) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Article) GetAuthorId() uint32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Article) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type ArticleId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ArticleId) Reset() {
	*x = ArticleId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemas_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleId) ProtoMessage() {}

func (x *ArticleId) ProtoReflect() protoreflect.Message {
	mi := &file_schemas_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleId.ProtoReflect.Descriptor instead.
func (*ArticleId) Descriptor() ([]byte, []int) {
	return file_schemas_schema_proto_rawDescGZIP(), []int{1}
}

func (x *ArticleId) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ArticleParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  *uint32 `protobuf:"varint,1,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Page   *uint32 `protobuf:"varint,2,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Author *string `protobuf:"bytes,3,opt,name=author,proto3,oneof" json:"author,omitempty"`
	Query  *string `protobuf:"bytes,4,opt,name=query,proto3,oneof" json:"query,omitempty"`
}

func (x *ArticleParam) Reset() {
	*x = ArticleParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemas_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleParam) ProtoMessage() {}

func (x *ArticleParam) ProtoReflect() protoreflect.Message {
	mi := &file_schemas_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleParam.ProtoReflect.Descriptor instead.
func (*ArticleParam) Descriptor() ([]byte, []int) {
	return file_schemas_schema_proto_rawDescGZIP(), []int{2}
}

func (x *ArticleParam) GetLimit() uint32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *ArticleParam) GetPage() uint32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *ArticleParam) GetAuthor() string {
	if x != nil && x.Author != nil {
		return *x.Author
	}
	return ""
}

func (x *ArticleParam) GetQuery() string {
	if x != nil && x.Query != nil {
		return *x.Query
	}
	return ""
}

type ArticleInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId uint32 `protobuf:"varint,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Body     string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ArticleInput) Reset() {
	*x = ArticleInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemas_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleInput) ProtoMessage() {}

func (x *ArticleInput) ProtoReflect() protoreflect.Message {
	mi := &file_schemas_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleInput.ProtoReflect.Descriptor instead.
func (*ArticleInput) Descriptor() ([]byte, []int) {
	return file_schemas_schema_proto_rawDescGZIP(), []int{3}
}

func (x *ArticleInput) GetAuthorId() uint32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *ArticleInput) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ArticleInput) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Nothing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Nothing) Reset() {
	*x = Nothing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemas_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nothing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nothing) ProtoMessage() {}

func (x *Nothing) ProtoReflect() protoreflect.Message {
	mi := &file_schemas_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nothing.ProtoReflect.Descriptor instead.
func (*Nothing) Descriptor() ([]byte, []int) {
	return file_schemas_schema_proto_rawDescGZIP(), []int{4}
}

var File_schemas_schema_proto protoreflect.FileDescriptor

var file_schemas_schema_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x55, 0x77, 0x55, 0x22, 0x7f, 0x0a, 0x07, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x1b, 0x0a, 0x09,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa2, 0x01, 0x0a, 0x0c, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x55,
	0x0a, 0x0c, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x09, 0x0a, 0x07, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x32, 0x98, 0x01, 0x0a, 0x03, 0x55, 0x77, 0x55, 0x12, 0x33, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x11, 0x2e, 0x55, 0x77, 0x55,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0c, 0x2e,
	0x55, 0x77, 0x55, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x30, 0x01, 0x12, 0x2a, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x55, 0x77,
	0x55, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x55, 0x77,
	0x55, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x55, 0x77, 0x55,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0c, 0x2e,
	0x55, 0x77, 0x55, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x42, 0x0b, 0x5a, 0x09, 0x55,
	0x77, 0x55, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schemas_schema_proto_rawDescOnce sync.Once
	file_schemas_schema_proto_rawDescData = file_schemas_schema_proto_rawDesc
)

func file_schemas_schema_proto_rawDescGZIP() []byte {
	file_schemas_schema_proto_rawDescOnce.Do(func() {
		file_schemas_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_schemas_schema_proto_rawDescData)
	})
	return file_schemas_schema_proto_rawDescData
}

var file_schemas_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_schemas_schema_proto_goTypes = []interface{}{
	(*Article)(nil),      // 0: UwU.Article
	(*ArticleId)(nil),    // 1: UwU.ArticleId
	(*ArticleParam)(nil), // 2: UwU.ArticleParam
	(*ArticleInput)(nil), // 3: UwU.ArticleInput
	(*Nothing)(nil),      // 4: UwU.Nothing
}
var file_schemas_schema_proto_depIdxs = []int32{
	2, // 0: UwU.UwU.GetAllArticles:input_type -> UwU.ArticleParam
	1, // 1: UwU.UwU.GetArticle:input_type -> UwU.ArticleId
	3, // 2: UwU.UwU.CreateArticle:input_type -> UwU.ArticleInput
	0, // 3: UwU.UwU.GetAllArticles:output_type -> UwU.Article
	0, // 4: UwU.UwU.GetArticle:output_type -> UwU.Article
	4, // 5: UwU.UwU.CreateArticle:output_type -> UwU.Nothing
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_schemas_schema_proto_init() }
func file_schemas_schema_proto_init() {
	if File_schemas_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schemas_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_schemas_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleId); i {
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
		file_schemas_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleParam); i {
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
		file_schemas_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleInput); i {
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
		file_schemas_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nothing); i {
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
	file_schemas_schema_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schemas_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schemas_schema_proto_goTypes,
		DependencyIndexes: file_schemas_schema_proto_depIdxs,
		MessageInfos:      file_schemas_schema_proto_msgTypes,
	}.Build()
	File_schemas_schema_proto = out.File
	file_schemas_schema_proto_rawDesc = nil
	file_schemas_schema_proto_goTypes = nil
	file_schemas_schema_proto_depIdxs = nil
}
