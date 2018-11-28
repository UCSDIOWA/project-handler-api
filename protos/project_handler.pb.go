// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/project_handler.proto

package project_creator

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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
	Xid                  string   `protobuf:"bytes,1,opt,name=xid,proto3" json:"xid,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Projectleader        string   `protobuf:"bytes,3,opt,name=projectleader,proto3" json:"projectleader,omitempty"`
	Percentdone          int32    `protobuf:"varint,4,opt,name=percentdone,proto3" json:"percentdone,omitempty"`
	Groupsize            int32    `protobuf:"varint,5,opt,name=groupsize,proto3" json:"groupsize,omitempty"`
	Isprivate            bool     `protobuf:"varint,6,opt,name=isprivate,proto3" json:"isprivate,omitempty"`
	Tags                 []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	Deadline             string   `protobuf:"bytes,8,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Calendarid           string   `protobuf:"bytes,9,opt,name=calendarid,proto3" json:"calendarid,omitempty"`
	Description          string   `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	Joinrequests         []string `protobuf:"bytes,11,rep,name=joinrequests,proto3" json:"joinrequests,omitempty"`
	Memberslist          []string `protobuf:"bytes,12,rep,name=memberslist,proto3" json:"memberslist,omitempty"`
	Milestones           []string `protobuf:"bytes,13,rep,name=milestones,proto3" json:"milestones,omitempty"`
	Announcements        []string `protobuf:"bytes,14,rep,name=announcements,proto3" json:"announcements,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProjectRequest) Reset()         { *m = CreateProjectRequest{} }
func (m *CreateProjectRequest) String() string { return proto.CompactTextString(m) }
func (*CreateProjectRequest) ProtoMessage()    {}
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b74e335e85f3fb3, []int{0}
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

func (m *CreateProjectRequest) GetXid() string {
	if m != nil {
		return m.Xid
	}
	return ""
}

func (m *CreateProjectRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateProjectRequest) GetProjectleader() string {
	if m != nil {
		return m.Projectleader
	}
	return ""
}

func (m *CreateProjectRequest) GetPercentdone() int32 {
	if m != nil {
		return m.Percentdone
	}
	return 0
}

func (m *CreateProjectRequest) GetGroupsize() int32 {
	if m != nil {
		return m.Groupsize
	}
	return 0
}

func (m *CreateProjectRequest) GetIsprivate() bool {
	if m != nil {
		return m.Isprivate
	}
	return false
}

func (m *CreateProjectRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *CreateProjectRequest) GetDeadline() string {
	if m != nil {
		return m.Deadline
	}
	return ""
}

func (m *CreateProjectRequest) GetCalendarid() string {
	if m != nil {
		return m.Calendarid
	}
	return ""
}

func (m *CreateProjectRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateProjectRequest) GetJoinrequests() []string {
	if m != nil {
		return m.Joinrequests
	}
	return nil
}

func (m *CreateProjectRequest) GetMemberslist() []string {
	if m != nil {
		return m.Memberslist
	}
	return nil
}

func (m *CreateProjectRequest) GetMilestones() []string {
	if m != nil {
		return m.Milestones
	}
	return nil
}

func (m *CreateProjectRequest) GetAnnouncements() []string {
	if m != nil {
		return m.Announcements
	}
	return nil
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
	return fileDescriptor_4b74e335e85f3fb3, []int{1}
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

type FetchProjectRequest struct {
	Xid                  string   `protobuf:"bytes,1,opt,name=xid,proto3" json:"xid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchProjectRequest) Reset()         { *m = FetchProjectRequest{} }
func (m *FetchProjectRequest) String() string { return proto.CompactTextString(m) }
func (*FetchProjectRequest) ProtoMessage()    {}
func (*FetchProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b74e335e85f3fb3, []int{2}
}

func (m *FetchProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchProjectRequest.Unmarshal(m, b)
}
func (m *FetchProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchProjectRequest.Marshal(b, m, deterministic)
}
func (m *FetchProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchProjectRequest.Merge(m, src)
}
func (m *FetchProjectRequest) XXX_Size() int {
	return xxx_messageInfo_FetchProjectRequest.Size(m)
}
func (m *FetchProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchProjectRequest proto.InternalMessageInfo

func (m *FetchProjectRequest) GetXid() string {
	if m != nil {
		return m.Xid
	}
	return ""
}

type FetchProjectResponse struct {
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Projectleader        string   `protobuf:"bytes,3,opt,name=projectleader,proto3" json:"projectleader,omitempty"`
	Percentdone          int32    `protobuf:"varint,4,opt,name=percentdone,proto3" json:"percentdone,omitempty"`
	Groupsize            int32    `protobuf:"varint,5,opt,name=groupsize,proto3" json:"groupsize,omitempty"`
	Isprivate            bool     `protobuf:"varint,6,opt,name=isprivate,proto3" json:"isprivate,omitempty"`
	Tags                 []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	Deadline             string   `protobuf:"bytes,8,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Calendarid           string   `protobuf:"bytes,9,opt,name=calendarid,proto3" json:"calendarid,omitempty"`
	Description          string   `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	Joinrequests         []string `protobuf:"bytes,11,rep,name=joinrequests,proto3" json:"joinrequests,omitempty"`
	Memberslist          []string `protobuf:"bytes,12,rep,name=memberslist,proto3" json:"memberslist,omitempty"`
	Milestones           []string `protobuf:"bytes,13,rep,name=milestones,proto3" json:"milestones,omitempty"`
	Announcements        []string `protobuf:"bytes,14,rep,name=announcements,proto3" json:"announcements,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchProjectResponse) Reset()         { *m = FetchProjectResponse{} }
func (m *FetchProjectResponse) String() string { return proto.CompactTextString(m) }
func (*FetchProjectResponse) ProtoMessage()    {}
func (*FetchProjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b74e335e85f3fb3, []int{3}
}

func (m *FetchProjectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchProjectResponse.Unmarshal(m, b)
}
func (m *FetchProjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchProjectResponse.Marshal(b, m, deterministic)
}
func (m *FetchProjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchProjectResponse.Merge(m, src)
}
func (m *FetchProjectResponse) XXX_Size() int {
	return xxx_messageInfo_FetchProjectResponse.Size(m)
}
func (m *FetchProjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchProjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchProjectResponse proto.InternalMessageInfo

func (m *FetchProjectResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *FetchProjectResponse) GetProjectleader() string {
	if m != nil {
		return m.Projectleader
	}
	return ""
}

func (m *FetchProjectResponse) GetPercentdone() int32 {
	if m != nil {
		return m.Percentdone
	}
	return 0
}

func (m *FetchProjectResponse) GetGroupsize() int32 {
	if m != nil {
		return m.Groupsize
	}
	return 0
}

func (m *FetchProjectResponse) GetIsprivate() bool {
	if m != nil {
		return m.Isprivate
	}
	return false
}

func (m *FetchProjectResponse) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *FetchProjectResponse) GetDeadline() string {
	if m != nil {
		return m.Deadline
	}
	return ""
}

func (m *FetchProjectResponse) GetCalendarid() string {
	if m != nil {
		return m.Calendarid
	}
	return ""
}

func (m *FetchProjectResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *FetchProjectResponse) GetJoinrequests() []string {
	if m != nil {
		return m.Joinrequests
	}
	return nil
}

func (m *FetchProjectResponse) GetMemberslist() []string {
	if m != nil {
		return m.Memberslist
	}
	return nil
}

func (m *FetchProjectResponse) GetMilestones() []string {
	if m != nil {
		return m.Milestones
	}
	return nil
}

func (m *FetchProjectResponse) GetAnnouncements() []string {
	if m != nil {
		return m.Announcements
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateProjectRequest)(nil), "project_creator.CreateProjectRequest")
	proto.RegisterType((*CreateProjectResponse)(nil), "project_creator.CreateProjectResponse")
	proto.RegisterType((*FetchProjectRequest)(nil), "project_creator.FetchProjectRequest")
	proto.RegisterType((*FetchProjectResponse)(nil), "project_creator.FetchProjectResponse")
}

func init() { proto.RegisterFile("protos/project_handler.proto", fileDescriptor_4b74e335e85f3fb3) }

var fileDescriptor_4b74e335e85f3fb3 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x94, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0xb5, 0xd9, 0xa6, 0x4d, 0xa6, 0x49, 0x01, 0x37, 0x48, 0x26, 0x8a, 0xd0, 0x6a, 0xd5,
	0x42, 0xc4, 0x21, 0x2b, 0xe0, 0xc6, 0x0d, 0x55, 0x42, 0xe2, 0x56, 0xed, 0x0b, 0x20, 0x77, 0x3d,
	0xa4, 0xae, 0x36, 0xf6, 0xe2, 0xf1, 0x22, 0xe0, 0x88, 0xc4, 0x13, 0xf0, 0x1e, 0x3c, 0x03, 0xef,
	0xc0, 0x2b, 0xf0, 0x20, 0xc8, 0xde, 0x2d, 0xd9, 0x2d, 0x95, 0x72, 0xe5, 0xd0, 0x9b, 0xfd, 0xcd,
	0x6f, 0xcf, 0x68, 0xe6, 0xd7, 0xc0, 0xa2, 0xb2, 0xc6, 0x19, 0xca, 0x2a, 0x6b, 0xae, 0xb0, 0x70,
	0xef, 0x2e, 0x85, 0x96, 0x25, 0xda, 0x55, 0xc0, 0xec, 0xde, 0x35, 0x2e, 0x2c, 0x0a, 0x67, 0xec,
	0x7c, 0xb1, 0x36, 0x66, 0x5d, 0x62, 0x26, 0x2a, 0x95, 0x09, 0xad, 0x8d, 0x13, 0x4e, 0x19, 0x4d,
	0x8d, 0x3c, 0xfd, 0x19, 0xc3, 0xec, 0xcc, 0x2b, 0xf1, 0xbc, 0x79, 0x97, 0xe3, 0x87, 0x1a, 0xc9,
	0xb1, 0xfb, 0x10, 0x7f, 0x52, 0x92, 0x47, 0x49, 0xb4, 0x1c, 0xe7, 0xfe, 0xc8, 0x66, 0x30, 0x74,
	0xca, 0x95, 0xc8, 0x07, 0x81, 0x35, 0x17, 0x76, 0x02, 0xd3, 0x36, 0x63, 0x89, 0x42, 0xa2, 0xe5,
	0x71, 0x88, 0xf6, 0x21, 0x4b, 0xe0, 0xb0, 0x42, 0x5b, 0xa0, 0x76, 0xd2, 0x68, 0xe4, 0x7b, 0x49,
	0xb4, 0x1c, 0xe6, 0x5d, 0xc4, 0x16, 0x30, 0x5e, 0x5b, 0x53, 0x57, 0xa4, 0xbe, 0x20, 0x1f, 0x86,
	0xf8, 0x16, 0xf8, 0xa8, 0xa2, 0xca, 0xaa, 0x8f, 0xc2, 0x21, 0xdf, 0x4f, 0xa2, 0xe5, 0x28, 0xdf,
	0x02, 0xc6, 0x60, 0xcf, 0x89, 0x35, 0xf1, 0x83, 0x24, 0x5e, 0x8e, 0xf3, 0x70, 0x66, 0x73, 0x18,
	0x49, 0x14, 0xb2, 0x54, 0x1a, 0xf9, 0x28, 0x94, 0xf4, 0xf7, 0xce, 0x1e, 0x03, 0x14, 0xa2, 0x44,
	0x2d, 0x85, 0x55, 0x92, 0x8f, 0x43, 0xb4, 0x43, 0x7c, 0xb5, 0x12, 0xa9, 0xb0, 0xaa, 0xf2, 0xad,
	0xe2, 0x10, 0x04, 0x5d, 0xc4, 0x52, 0x98, 0x5c, 0x19, 0xa5, 0x6d, 0xd3, 0x2c, 0xe2, 0x87, 0x21,
	0x73, 0x8f, 0xf9, 0x5f, 0x36, 0xb8, 0xb9, 0x40, 0x4b, 0xa5, 0x22, 0xc7, 0x27, 0x41, 0xd2, 0x45,
	0xbe, 0x8e, 0x8d, 0x2a, 0x91, 0x9c, 0xd1, 0x48, 0x7c, 0x1a, 0x04, 0x1d, 0xe2, 0x7b, 0xeb, 0x27,
	0x56, 0xeb, 0x02, 0x37, 0xa8, 0x1d, 0xf1, 0xa3, 0x20, 0xe9, 0xc3, 0xf4, 0x39, 0x3c, 0xbc, 0x31,
	0x41, 0xaa, 0x8c, 0x26, 0x64, 0x1c, 0x0e, 0xa8, 0x2e, 0x0a, 0x24, 0x0a, 0x63, 0x1c, 0xe5, 0xd7,
	0xd7, 0xf4, 0x29, 0x1c, 0xbf, 0x41, 0x57, 0x5c, 0xee, 0x9a, 0x79, 0xfa, 0x23, 0x86, 0x59, 0x5f,
	0xd9, 0xfe, 0x7d, 0x67, 0x86, 0xff, 0xd2, 0x0c, 0x2f, 0xbe, 0x0d, 0xe0, 0x41, 0x3b, 0xab, 0xb3,
	0x66, 0x01, 0xbc, 0x3e, 0x7f, 0xcb, 0x3e, 0xc3, 0xb4, 0x67, 0x11, 0x76, 0xba, 0xba, 0xb1, 0x26,
	0x56, 0xb7, 0x2d, 0x81, 0xf9, 0x93, 0x5d, 0xb2, 0xc6, 0x0d, 0xe9, 0xa3, 0xaf, 0xbf, 0x7e, 0x7f,
	0x1f, 0x1c, 0xa7, 0x47, 0x59, 0xd0, 0x61, 0xfb, 0xea, 0x55, 0xf4, 0x8c, 0xd5, 0x30, 0xe9, 0x1a,
	0x88, 0x9d, 0xfc, 0xf3, 0xe5, 0x2d, 0x4e, 0x9c, 0x9f, 0xee, 0x50, 0xb5, 0x79, 0x79, 0xc8, 0xcb,
	0xd2, 0x69, 0xf6, 0xde, 0x87, 0xb7, 0x69, 0x2f, 0xf6, 0xc3, 0x7a, 0x7b, 0xf9, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xa7, 0xdd, 0x3c, 0x10, 0x2d, 0x05, 0x00, 0x00,
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
	FetchProject(ctx context.Context, in *FetchProjectRequest, opts ...grpc.CallOption) (*FetchProjectResponse, error)
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

func (c *projectCreatorAPIClient) FetchProject(ctx context.Context, in *FetchProjectRequest, opts ...grpc.CallOption) (*FetchProjectResponse, error) {
	out := new(FetchProjectResponse)
	err := c.cc.Invoke(ctx, "/project_creator.ProjectCreatorAPI/FetchProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectCreatorAPIServer is the server API for ProjectCreatorAPI service.
type ProjectCreatorAPIServer interface {
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error)
	FetchProject(context.Context, *FetchProjectRequest) (*FetchProjectResponse, error)
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

func _ProjectCreatorAPI_FetchProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectCreatorAPIServer).FetchProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project_creator.ProjectCreatorAPI/FetchProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectCreatorAPIServer).FetchProject(ctx, req.(*FetchProjectRequest))
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
		{
			MethodName: "FetchProject",
			Handler:    _ProjectCreatorAPI_FetchProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/project_handler.proto",
}
