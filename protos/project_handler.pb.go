// Code generated by protoc-gen-go. DO NOT EDIT.
// source: project_handler.proto

package project_creator

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateProjectRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Deadline             string   `protobuf:"bytes,4,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Private              bool     `protobuf:"varint,5,opt,name=private,proto3" json:"private,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProjectRequest) Reset()         { *m = CreateProjectRequest{} }
func (m *CreateProjectRequest) String() string { return proto.CompactTextString(m) }
func (*CreateProjectRequest) ProtoMessage()    {}
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78af44de10dd809e, []int{0}
}

func (m *CreateProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProjectRequest.Unmarshal(m, b)
}
func (m *CreateProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProjectRequest.Marshal(b, m, deterministic)
}
func (m *CreateProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProjectRequest.Merge(m, src)
}
func (m *CreateProjectRequest) XXX_Size() int {
	return xxx_messageInfo_CreateProjectRequest.Size(m)
}
func (m *CreateProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProjectRequest proto.InternalMessageInfo

func (m *CreateProjectRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateProjectRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateProjectRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateProjectRequest) GetDeadline() string {
	if m != nil {
		return m.Deadline
	}
	return ""
}

func (m *CreateProjectRequest) GetPrivate() bool {
	if m != nil {
		return m.Private
	}
	return false
}

type CreateProjectResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProjectResponse) Reset()         { *m = CreateProjectResponse{} }
func (m *CreateProjectResponse) String() string { return proto.CompactTextString(m) }
func (*CreateProjectResponse) ProtoMessage()    {}
func (*CreateProjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78af44de10dd809e, []int{1}
}

func (m *CreateProjectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProjectResponse.Unmarshal(m, b)
}
func (m *CreateProjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProjectResponse.Marshal(b, m, deterministic)
}
func (m *CreateProjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProjectResponse.Merge(m, src)
}
func (m *CreateProjectResponse) XXX_Size() int {
	return xxx_messageInfo_CreateProjectResponse.Size(m)
}
func (m *CreateProjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProjectResponse proto.InternalMessageInfo

func (m *CreateProjectResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*CreateProjectRequest)(nil), "project_creator.CreateProjectRequest")
	proto.RegisterType((*CreateProjectResponse)(nil), "project_creator.CreateProjectResponse")
}

func init() { proto.RegisterFile("project_handler.proto", fileDescriptor_78af44de10dd809e) }

var fileDescriptor_78af44de10dd809e = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4a, 0x03, 0x31,
	0x18, 0x85, 0x49, 0xb5, 0x3a, 0x46, 0x54, 0x8c, 0x2d, 0xc4, 0xc1, 0xc5, 0x30, 0xa0, 0x14, 0x17,
	0x33, 0xa8, 0x3b, 0x77, 0xd2, 0x95, 0xbb, 0x32, 0x17, 0x90, 0x98, 0xf9, 0xa9, 0x91, 0x98, 0xc4,
	0xe4, 0xaf, 0xe0, 0xd6, 0x03, 0xb8, 0x71, 0xe7, 0xb5, 0xbc, 0x82, 0x07, 0x91, 0x49, 0x3a, 0xa2,
	0x45, 0x70, 0xf9, 0xde, 0xff, 0xbd, 0x9f, 0xc7, 0xa3, 0x63, 0xe7, 0xed, 0x3d, 0x48, 0xbc, 0xb9,
	0x13, 0xa6, 0xd5, 0xe0, 0x2b, 0xe7, 0x2d, 0x5a, 0xb6, 0xd7, 0xdb, 0xd2, 0x83, 0x40, 0xeb, 0xf3,
	0xa3, 0xb9, 0xb5, 0x73, 0x0d, 0xb5, 0x70, 0xaa, 0x16, 0xc6, 0x58, 0x14, 0xa8, 0xac, 0x09, 0x09,
	0x2f, 0xdf, 0x09, 0x1d, 0x4d, 0x3b, 0x12, 0x66, 0x29, 0xd7, 0xc0, 0xe3, 0x02, 0x02, 0xb2, 0x11,
	0x1d, 0xc2, 0x83, 0x50, 0x9a, 0x93, 0x82, 0x4c, 0xb6, 0x9a, 0x24, 0x3a, 0x17, 0x15, 0x6a, 0xe0,
	0x83, 0xe4, 0x46, 0xc1, 0x0a, 0xba, 0xdd, 0x42, 0x90, 0x5e, 0xb9, 0xee, 0x35, 0x5f, 0x8b, 0xb7,
	0x9f, 0x16, 0xcb, 0x69, 0xd6, 0x82, 0x68, 0xb5, 0x32, 0xc0, 0xd7, 0xe3, 0xf9, 0x5b, 0x33, 0x4e,
	0x37, 0x9d, 0x57, 0x4f, 0x02, 0x81, 0x0f, 0x0b, 0x32, 0xc9, 0x9a, 0x5e, 0x96, 0x67, 0x74, 0xbc,
	0xd2, 0x2d, 0x38, 0x6b, 0x42, 0x8c, 0x84, 0x85, 0x94, 0x10, 0x42, 0xac, 0x97, 0x35, 0xbd, 0x3c,
	0x7f, 0x25, 0x74, 0x7f, 0x49, 0x4f, 0xd3, 0x00, 0x57, 0xb3, 0x6b, 0xf6, 0x4c, 0x77, 0x7e, 0x3d,
	0x62, 0xc7, 0xd5, 0xca, 0x4c, 0xd5, 0x5f, 0x23, 0xe4, 0x27, 0xff, 0x61, 0xa9, 0x4f, 0x79, 0xf8,
	0xf2, 0xf1, 0xf9, 0x36, 0x38, 0x28, 0x77, 0xeb, 0xc8, 0xc1, 0x32, 0x75, 0x49, 0x4e, 0x6f, 0x37,
	0xe2, 0xce, 0x17, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xbd, 0xf4, 0x22, 0xaf, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProjectCreatorAPIClient is the client API for ProjectCreatorAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProjectCreatorAPIClient interface {
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error)
}

type projectCreatorAPIClient struct {
	cc *grpc.ClientConn
}

func NewProjectCreatorAPIClient(cc *grpc.ClientConn) ProjectCreatorAPIClient {
	return &projectCreatorAPIClient{cc}
}

func (c *projectCreatorAPIClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error) {
	out := new(CreateProjectResponse)
	err := c.cc.Invoke(ctx, "/project_creator.ProjectCreatorAPI/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectCreatorAPIServer is the server API for ProjectCreatorAPI service.
type ProjectCreatorAPIServer interface {
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error)
}

func RegisterProjectCreatorAPIServer(s *grpc.Server, srv ProjectCreatorAPIServer) {
	s.RegisterService(&_ProjectCreatorAPI_serviceDesc, srv)
}

func _ProjectCreatorAPI_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectCreatorAPIServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project_creator.ProjectCreatorAPI/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectCreatorAPIServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProjectCreatorAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "project_creator.ProjectCreatorAPI",
	HandlerType: (*ProjectCreatorAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _ProjectCreatorAPI_CreateProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "project_handler.proto",
}
