// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/test_public.proto

package test

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
)

type PublicImportMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicImportMessage) ProtoReflect() protoreflect.Message {
	return xxx_File_test_test_public_proto_messageTypes[0].MessageOf(m)
}
func (m *PublicImportMessage) Reset()         { *m = PublicImportMessage{} }
func (m *PublicImportMessage) String() string { return protoimpl.X.MessageStringOf(m) }
func (*PublicImportMessage) ProtoMessage()    {}

// Deprecated: Use PublicImportMessage.ProtoReflect.Type instead.
func (*PublicImportMessage) Descriptor() ([]byte, []int) {
	return xxx_File_test_test_public_proto_rawdesc_gzipped, []int{0}
}

var xxx_File_test_test_public_proto_rawdesc = []byte{
	// 125 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x16, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74,
}

var xxx_File_test_test_public_proto_rawdesc_gzipped = protoimpl.X.CompressGZIP(xxx_File_test_test_public_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_test_test_public_proto protoreflect.FileDescriptor

var xxx_File_test_test_public_proto_messageTypes = make([]protoimpl.MessageType, 1)
var xxx_File_test_test_public_proto_goTypes = []interface{}{
	(*PublicImportMessage)(nil), // 0: goproto.proto.test.PublicImportMessage
}
var xxx_File_test_test_public_proto_depIdxs = []int32{}

func init() { xxx_File_test_test_public_proto_init() }
func xxx_File_test_test_public_proto_init() {
	if File_test_test_public_proto != nil {
		return
	}
	File_test_test_public_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_test_test_public_proto_rawdesc,
		GoTypes:            xxx_File_test_test_public_proto_goTypes,
		DependencyIndexes:  xxx_File_test_test_public_proto_depIdxs,
		MessageOutputTypes: xxx_File_test_test_public_proto_messageTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	xxx_File_test_test_public_proto_goTypes = nil
	xxx_File_test_test_public_proto_depIdxs = nil
}