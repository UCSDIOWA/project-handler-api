// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/project_handler.proto

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

type Project struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Users                []string `protobuf:"bytes,2,rep,name=Users,proto3" json:"Users,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Deadline             string   `protobuf:"bytes,4,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Private              bool     `protobuf:"varint,5,opt,name=private,proto3" json:"private,omitempty"`
	Announcements        []string `protobuf:"bytes,6,rep,name=Announcements,proto3" json:"Announcements,omitempty"`
	Size                 int32    `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	ProjectLeader        string   `protobuf:"bytes,8,opt,name=project_leader,json=projectLeader,proto3" json:"project_leader,omitempty"`
	ProgressBar          int32    `protobuf:"varint,9,opt,name=progress_bar,json=progressBar,proto3" json:"progress_bar,omitempty"`
	Done                 bool     `protobuf:"varint,10,opt,name=done,proto3" json:"done,omitempty"`
	Calendar             string   `protobuf:"bytes,11,opt,name=calendar,proto3" json:"calendar,omitempty"`
	Milestones           []string `protobuf:"bytes,12,rep,name=Milestones,proto3" json:"Milestones,omitempty"`
	Tags                 []string `protobuf:"bytes,13,rep,name=Tags,proto3" json:"Tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b74e335e85f3fb3, []int{0}
}

func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Project) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Project) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Project) GetDeadline() string {
	if m != nil {
		return m.Deadline
	}
	return ""
}

func (m *Project) GetPrivate() bool {
	if m != nil {
		return m.Private
	}
	return false
}

func (m *Project) GetAnnouncements() []string {
	if m != nil {
		return m.Announcements
	}
	return nil
}

func (m *Project) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Project) GetProjectLeader() string {
	if m != nil {
		return m.ProjectLeader
	}
	return ""
}

func (m *Project) GetProgressBar() int32 {
	if m != nil {
		return m.ProgressBar
	}
	return 0
}

func (m *Project) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Project) GetCalendar() string {
	if m != nil {
		return m.Calendar
	}
	return ""
}

func (m *Project) GetMilestones() []string {
	if m != nil {
		return m.Milestones
	}
	return nil
}

func (m *Project) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type UserProjects struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Projects             []string `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserProjects) Reset()         { *m = UserProjects{} }
func (m *UserProjects) String() string { return proto.CompactTextString(m) }
func (*UserProjects) ProtoMessage()    {}
func (*UserProjects) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b74e335e85f3fb3, []int{1}
}

func (m *UserProjects) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserProjects.Unmarshal(m, b)
}
func (m *UserProjects) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserProjects.Marshal(b, m, deterministic)
}
func (m *UserProjects) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserProjects.Merge(m, src)
}
func (m *UserProjects) XXX_Size() int {
	return xxx_messageInfo_UserProjects.Size(m)
}
func (m *UserProjects) XXX_DiscardUnknown() {
	xxx_messageInfo_UserProjects.DiscardUnknown(m)
}

var xxx_messageInfo_UserProjects proto.InternalMessageInfo

func (m *UserProjects) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserProjects) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type TagProjects struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Projects             []string `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagProjects) Reset()         { *m = TagProjects{} }
func (m *TagProjects) String() string { return proto.CompactTextString(m) }
func (*TagProjects) ProtoMessage()    {}
func (*TagProjects) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b74e335e85f3fb3, []int{2}
}

func (m *TagProjects) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagProjects.Unmarshal(m, b)
}
func (m *TagProjects) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagProjects.Marshal(b, m, deterministic)
}
func (m *TagProjects) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagProjects.Merge(m, src)
}
func (m *TagProjects) XXX_Size() int {
	return xxx_messageInfo_TagProjects.Size(m)
}
func (m *TagProjects) XXX_DiscardUnknown() {
	xxx_messageInfo_TagProjects.DiscardUnknown(m)
}

var xxx_messageInfo_TagProjects proto.InternalMessageInfo

func (m *TagProjects) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TagProjects) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type CreateProjectRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Deadline             string   `protobuf:"bytes,4,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Private              bool     `protobuf:"varint,5,opt,name=private,proto3" json:"private,omitempty"`
	Tags                 []string `protobuf:"bytes,6,rep,name=Tags,proto3" json:"Tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProjectRequest) Reset()         { *m = CreateProjectRequest{} }
func (m *CreateProjectRequest) String() string { return proto.CompactTextString(m) }
func (*CreateProjectRequest) ProtoMessage()    {}
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b74e335e85f3fb3, []int{3}
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

func (m *CreateProjectRequest) GetTags() []string {
	if m != nil {
		return m.Tags
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
	return fileDescriptor_4b74e335e85f3fb3, []int{4}
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
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchProjectRequest) Reset()         { *m = FetchProjectRequest{} }
func (m *FetchProjectRequest) String() string { return proto.CompactTextString(m) }
func (*FetchProjectRequest) ProtoMessage()    {}
func (*FetchProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b74e335e85f3fb3, []int{5}
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

func (m *FetchProjectRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *FetchProjectRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type FetchProjectResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Project              *Project `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchProjectResponse) Reset()         { *m = FetchProjectResponse{} }
func (m *FetchProjectResponse) String() string { return proto.CompactTextString(m) }
func (*FetchProjectResponse) ProtoMessage()    {}
func (*FetchProjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b74e335e85f3fb3, []int{6}
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

func (m *FetchProjectResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *FetchProjectResponse) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func init() {
	proto.RegisterType((*Project)(nil), "project_creator.Project")
	proto.RegisterType((*UserProjects)(nil), "project_creator.UserProjects")
	proto.RegisterType((*TagProjects)(nil), "project_creator.TagProjects")
	proto.RegisterType((*CreateProjectRequest)(nil), "project_creator.CreateProjectRequest")
	proto.RegisterType((*CreateProjectResponse)(nil), "project_creator.CreateProjectResponse")
	proto.RegisterType((*FetchProjectRequest)(nil), "project_creator.FetchProjectRequest")
	proto.RegisterType((*FetchProjectResponse)(nil), "project_creator.FetchProjectResponse")
}

func init() { proto.RegisterFile("protos/project_handler.proto", fileDescriptor_4b74e335e85f3fb3) }

var fileDescriptor_4b74e335e85f3fb3 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x8a, 0x13, 0x41,
	0x10, 0x66, 0xb2, 0xf9, 0xdb, 0x4a, 0xb2, 0x62, 0x6f, 0x84, 0x36, 0x2c, 0x12, 0x07, 0x57, 0x82,
	0x87, 0x04, 0xe3, 0x4d, 0x10, 0x8c, 0x0b, 0x82, 0xa0, 0xb0, 0x84, 0xf5, 0xbc, 0xf4, 0xce, 0x14,
	0xb3, 0x23, 0x93, 0xee, 0xb1, 0xab, 0x23, 0xe8, 0xd1, 0x07, 0xf0, 0xe2, 0x7b, 0xf8, 0x32, 0xbe,
	0x82, 0xaf, 0xe0, 0x5d, 0xba, 0x66, 0x3a, 0x24, 0xeb, 0xcf, 0x82, 0xb0, 0xb7, 0xfa, 0xbe, 0xa9,
	0xfa, 0xaa, 0xea, 0xeb, 0x4a, 0xe0, 0xa8, 0xb4, 0xc6, 0x19, 0x9a, 0x95, 0xd6, 0xbc, 0xc3, 0xc4,
	0x9d, 0x5f, 0x2a, 0x9d, 0x16, 0x68, 0xa7, 0x4c, 0x8b, 0x5b, 0x81, 0x4e, 0x2c, 0x2a, 0x67, 0xec,
	0xe8, 0x28, 0x33, 0x26, 0x2b, 0x70, 0xa6, 0xca, 0x7c, 0xa6, 0xb4, 0x36, 0x4e, 0xb9, 0xdc, 0x68,
	0xaa, 0xd2, 0xe3, 0x9f, 0x0d, 0xe8, 0x9c, 0x56, 0x15, 0x62, 0x08, 0x2d, 0x97, 0xbb, 0x02, 0x65,
	0x34, 0x8e, 0x26, 0xfb, 0xcb, 0x0a, 0x78, 0xf6, 0x2d, 0xa1, 0x25, 0xd9, 0x18, 0xef, 0x79, 0x96,
	0x81, 0x18, 0x43, 0x2f, 0x45, 0x4a, 0x6c, 0x5e, 0x7a, 0x35, 0xb9, 0xc7, 0x15, 0xdb, 0x94, 0x18,
	0x41, 0x37, 0x45, 0x95, 0x16, 0xb9, 0x46, 0xd9, 0xe4, 0xcf, 0x1b, 0x2c, 0x24, 0x74, 0x4a, 0x9b,
	0x7f, 0x50, 0x0e, 0x65, 0x6b, 0x1c, 0x4d, 0xba, 0xcb, 0x00, 0xc5, 0x03, 0x18, 0x2c, 0xb4, 0x36,
	0x6b, 0x9d, 0xe0, 0x0a, 0xb5, 0x23, 0xd9, 0xe6, 0xae, 0xbb, 0xa4, 0x10, 0xd0, 0xa4, 0xfc, 0x13,
	0xca, 0xce, 0x38, 0x9a, 0xb4, 0x96, 0x1c, 0x8b, 0x63, 0x38, 0x08, 0xab, 0x17, 0xa8, 0x52, 0xb4,
	0xb2, 0xcb, 0x5d, 0x07, 0x35, 0xfb, 0x9a, 0x49, 0x71, 0x1f, 0xfa, 0xa5, 0x35, 0x99, 0x45, 0xa2,
	0xf3, 0x0b, 0x65, 0xe5, 0x3e, 0x4b, 0xf4, 0x02, 0xf7, 0x42, 0x59, 0xaf, 0x9e, 0x1a, 0x8d, 0x12,
	0x78, 0x34, 0x8e, 0xfd, 0x36, 0x89, 0x2a, 0x50, 0xa7, 0xca, 0xca, 0x5e, 0xb5, 0x4d, 0xc0, 0xe2,
	0x1e, 0xc0, 0x9b, 0xbc, 0x40, 0x72, 0x46, 0x23, 0xc9, 0x3e, 0x0f, 0xbc, 0xc5, 0x78, 0xbd, 0x33,
	0x95, 0x91, 0x1c, 0xf0, 0x17, 0x8e, 0xe3, 0xe7, 0xd0, 0xf7, 0x46, 0xd6, 0xd6, 0x93, 0x77, 0x19,
	0x57, 0x2a, 0x2f, 0x82, 0xf7, 0x0c, 0x7c, 0xd7, 0x7a, 0xfa, 0x60, 0xff, 0x06, 0xc7, 0xcf, 0xa0,
	0x77, 0xa6, 0xb2, 0x8d, 0x80, 0x80, 0xa6, 0x56, 0xab, 0xf0, 0x76, 0x1c, 0xff, 0xb3, 0xfc, 0x5b,
	0x04, 0xc3, 0x13, 0x7f, 0x22, 0x58, 0x4b, 0x2c, 0xf1, 0xfd, 0x1a, 0xc9, 0xfd, 0x65, 0x92, 0xcd,
	0x6d, 0x34, 0xb6, 0x6f, 0xe3, 0xa6, 0xae, 0x20, 0x38, 0xd6, 0xde, 0x72, 0xec, 0x31, 0xdc, 0xb9,
	0x32, 0x2f, 0x95, 0x46, 0x13, 0xcb, 0xd0, 0x3a, 0x49, 0x90, 0x88, 0x47, 0xee, 0x2e, 0x03, 0x8c,
	0x17, 0x70, 0xf8, 0x12, 0x5d, 0x72, 0xf9, 0xff, 0x1b, 0xc6, 0x29, 0x0c, 0x77, 0x25, 0xae, 0x6b,
	0x2a, 0xe6, 0x7e, 0x2b, 0x4e, 0x66, 0xa5, 0xde, 0x5c, 0x4e, 0xaf, 0xfc, 0x24, 0xa7, 0x41, 0x2c,
	0x24, 0xce, 0xbf, 0x44, 0x70, 0xbb, 0x26, 0x4f, 0xaa, 0x9c, 0xc5, 0xe9, 0x2b, 0xf1, 0x11, 0x06,
	0x3b, 0x1b, 0x8b, 0xe3, 0xdf, 0x94, 0xfe, 0xf4, 0x82, 0xa3, 0x87, 0xd7, 0xa5, 0x55, 0x3b, 0xc4,
	0x77, 0x3f, 0x7f, 0xff, 0xf1, 0xb5, 0x71, 0x18, 0x1f, 0xcc, 0x38, 0x0f, 0xeb, 0xaa, 0xa7, 0xd1,
	0xa3, 0x8b, 0x36, 0xff, 0x3b, 0x3c, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xad, 0xfa, 0xa6,
	0x6c, 0x04, 0x00, 0x00,
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
	Metadata: "protos/project_handler.proto",
}
