// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package go_micro_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type ListRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type LoginRequest struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UpdatePasswordRequest struct {
	Uid                  uint32   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	OldPassword          string   `protobuf:"bytes,2,opt,name=oldPassword,proto3" json:"oldPassword,omitempty"`
	NewPassword          string   `protobuf:"bytes,3,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
	ConfirmPassword      string   `protobuf:"bytes,4,opt,name=confirmPassword,proto3" json:"confirmPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePasswordRequest) Reset()         { *m = UpdatePasswordRequest{} }
func (m *UpdatePasswordRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePasswordRequest) ProtoMessage()    {}
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *UpdatePasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePasswordRequest.Unmarshal(m, b)
}
func (m *UpdatePasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePasswordRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePasswordRequest.Merge(m, src)
}
func (m *UpdatePasswordRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePasswordRequest.Size(m)
}
func (m *UpdatePasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePasswordRequest proto.InternalMessageInfo

func (m *UpdatePasswordRequest) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UpdatePasswordRequest) GetOldPassword() string {
	if m != nil {
		return m.OldPassword
	}
	return ""
}

func (m *UpdatePasswordRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

func (m *UpdatePasswordRequest) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

type Response struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.user.User")
	proto.RegisterType((*RegisterRequest)(nil), "go.micro.srv.user.RegisterRequest")
	proto.RegisterType((*ListRequest)(nil), "go.micro.srv.user.ListRequest")
	proto.RegisterType((*LoginRequest)(nil), "go.micro.srv.user.LoginRequest")
	proto.RegisterType((*UpdatePasswordRequest)(nil), "go.micro.srv.user.UpdatePasswordRequest")
	proto.RegisterType((*Response)(nil), "go.micro.srv.user.Response")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0xfd, 0x92, 0xa6, 0x1f, 0x75, 0xa2, 0xad, 0x0e, 0x16, 0x43, 0x05, 0x2d, 0x39, 0x05, 0x84,
	0x28, 0xf5, 0xe6, 0x41, 0xbc, 0x88, 0x97, 0x0a, 0x25, 0xd2, 0x83, 0xe0, 0xa5, 0x26, 0x63, 0x5c,
	0x30, 0xd9, 0xb8, 0x9b, 0xb6, 0x3f, 0xc5, 0x9f, 0xe4, 0xdf, 0x92, 0xdd, 0xb4, 0xb1, 0xad, 0x2b,
	0x05, 0x2f, 0x61, 0x66, 0xf6, 0xed, 0xdb, 0x97, 0xf7, 0x76, 0xa1, 0x5b, 0x08, 0x5e, 0xf2, 0xf3,
	0xa9, 0x24, 0xa1, 0x3f, 0xa1, 0xee, 0xf1, 0x20, 0xe5, 0x61, 0xc6, 0x62, 0xc1, 0x43, 0x29, 0x66,
	0xa1, 0x5a, 0xf0, 0x9f, 0xc0, 0x19, 0x4b, 0x12, 0xd8, 0x06, 0x9b, 0x25, 0x9e, 0xd5, 0xb7, 0x82,
	0xbd, 0xc8, 0x66, 0x09, 0x22, 0x38, 0xf9, 0x24, 0x23, 0xcf, 0xee, 0x5b, 0xc1, 0x4e, 0xa4, 0x6b,
	0x3c, 0x84, 0x66, 0xf1, 0xca, 0x73, 0xf2, 0x1a, 0x7a, 0x58, 0x35, 0xd8, 0x83, 0x56, 0x31, 0x91,
	0x72, 0xce, 0x45, 0xe2, 0x39, 0x7a, 0xa1, 0xee, 0xfd, 0x6b, 0xe8, 0x44, 0x94, 0x32, 0x59, 0x92,
	0x88, 0xe8, 0x7d, 0x4a, 0xb2, 0xc4, 0x33, 0x70, 0xd4, 0xc1, 0xfa, 0x28, 0x77, 0x70, 0x14, 0xfe,
	0x90, 0x14, 0x2a, 0x3d, 0x91, 0x06, 0xf9, 0x57, 0xe0, 0x0e, 0x99, 0x2c, 0xff, 0xb4, 0xf7, 0x06,
	0x76, 0x87, 0x3c, 0x65, 0xf9, 0x72, 0x73, 0xad, 0xde, 0xfa, 0x4d, 0xbd, 0xbd, 0xa1, 0xfe, 0xc3,
	0x82, 0xee, 0xb8, 0x48, 0x26, 0x25, 0x8d, 0x16, 0xa3, 0x25, 0xd7, 0x3e, 0x34, 0xa6, 0xb5, 0x5d,
	0xaa, 0xc4, 0x3e, 0xb8, 0xfc, 0x2d, 0x19, 0xad, 0x53, 0xad, 0x8e, 0x14, 0x22, 0xa7, 0x79, 0x8d,
	0xa8, 0x3c, 0x5c, 0x1d, 0x61, 0x00, 0x9d, 0x98, 0xe7, 0x2f, 0x4c, 0x64, 0xa3, 0x75, 0x43, 0x37,
	0xc7, 0xfe, 0x05, 0xb4, 0x22, 0x92, 0x05, 0xcf, 0x25, 0xa9, 0xa4, 0x62, 0x9e, 0x2c, 0x7f, 0x4b,
	0xd7, 0x4a, 0x5f, 0x26, 0xd3, 0x85, 0x0a, 0x55, 0x0e, 0x3e, 0x6d, 0x70, 0x95, 0x39, 0x0f, 0x24,
	0x66, 0x2c, 0x26, 0xbc, 0x57, 0x0c, 0x55, 0x32, 0xe8, 0x1b, 0x8c, 0xdc, 0x88, 0xad, 0x77, 0x6c,
	0xc4, 0x54, 0x12, 0xfc, 0x7f, 0x78, 0x07, 0x4d, 0x6d, 0x36, 0x9e, 0x1a, 0x70, 0xab, 0x31, 0x6c,
	0x23, 0x7a, 0x84, 0xf6, 0xba, 0xe5, 0x18, 0x98, 0x62, 0x36, 0xa5, 0xb2, 0x8d, 0xfa, 0x16, 0x1c,
	0x75, 0x99, 0xf0, 0xc4, 0x24, 0xf1, 0xfb, 0x96, 0x6d, 0xa1, 0x79, 0xfe, 0xaf, 0xdf, 0xd2, 0xe5,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xbe, 0xba, 0x45, 0x64, 0x03, 0x00, 0x00,
}
