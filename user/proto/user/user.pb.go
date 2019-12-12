// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package lu_micro_srv_user

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

type Userinfo struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Userinfo) Reset()         { *m = Userinfo{} }
func (m *Userinfo) String() string { return proto.CompactTextString(m) }
func (*Userinfo) ProtoMessage()    {}
func (*Userinfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *Userinfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Userinfo.Unmarshal(m, b)
}
func (m *Userinfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Userinfo.Marshal(b, m, deterministic)
}
func (m *Userinfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Userinfo.Merge(m, src)
}
func (m *Userinfo) XXX_Size() int {
	return xxx_messageInfo_Userinfo.Size(m)
}
func (m *Userinfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Userinfo.DiscardUnknown(m)
}

var xxx_messageInfo_Userinfo proto.InternalMessageInfo

func (m *Userinfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Userinfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Userinfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Userinfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterRequest struct {
	User                 *Userinfo `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
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

func (m *RegisterRequest) GetUser() *Userinfo {
	if m != nil {
		return m.User
	}
	return nil
}

type ListRequest struct {
	User                 *Userinfo `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
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

func (m *ListRequest) GetUser() *Userinfo {
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
	return fileDescriptor_116e343673f7ffaf, []int{3}
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
	return fileDescriptor_116e343673f7ffaf, []int{4}
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
	return fileDescriptor_116e343673f7ffaf, []int{5}
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
	proto.RegisterType((*Userinfo)(nil), "lu.micro.srv.user.Userinfo")
	proto.RegisterType((*RegisterRequest)(nil), "lu.micro.srv.user.RegisterRequest")
	proto.RegisterType((*ListRequest)(nil), "lu.micro.srv.user.ListRequest")
	proto.RegisterType((*LoginRequest)(nil), "lu.micro.srv.user.LoginRequest")
	proto.RegisterType((*UpdatePasswordRequest)(nil), "lu.micro.srv.user.UpdatePasswordRequest")
	proto.RegisterType((*Response)(nil), "lu.micro.srv.user.Response")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0xfd, 0x92, 0xa6, 0x1f, 0x75, 0xaa, 0xad, 0x0e, 0x0a, 0xa1, 0x05, 0x2d, 0x39, 0xe5, 0x14,
	0xa5, 0xde, 0x45, 0x04, 0xf1, 0x52, 0xa1, 0x04, 0x7a, 0xf0, 0x66, 0x6d, 0xa6, 0x75, 0xa1, 0xd9,
	0x8d, 0xbb, 0x89, 0xfd, 0x29, 0xfe, 0x1d, 0x7f, 0x9a, 0xec, 0xa6, 0x89, 0x6d, 0x5d, 0x29, 0x78,
	0x9b, 0xcc, 0xbe, 0x7d, 0xf3, 0xe6, 0xbd, 0x0d, 0x40, 0xa1, 0x48, 0x46, 0x99, 0x14, 0xb9, 0xc0,
	0x93, 0x65, 0x11, 0xa5, 0x6c, 0x26, 0x45, 0xa4, 0xe4, 0x7b, 0xa4, 0x0f, 0x82, 0x67, 0x68, 0x4d,
	0x14, 0x49, 0xc6, 0xe7, 0x02, 0x3b, 0xe0, 0xb2, 0xc4, 0x77, 0x06, 0x4e, 0x78, 0x14, 0xbb, 0x2c,
	0x41, 0x04, 0x8f, 0x4f, 0x53, 0xf2, 0xdd, 0x81, 0x13, 0x1e, 0xc4, 0xa6, 0xc6, 0x53, 0x68, 0x66,
	0xaf, 0x82, 0x93, 0xdf, 0x30, 0xcd, 0xf2, 0x03, 0x7b, 0xd0, 0xca, 0xa6, 0x4a, 0xad, 0x84, 0x4c,
	0x7c, 0xcf, 0x1c, 0xd4, 0xdf, 0xc1, 0x1d, 0x74, 0x63, 0x5a, 0x30, 0x95, 0x93, 0x8c, 0xe9, 0xad,
	0x20, 0x95, 0xe3, 0x25, 0x78, 0x7a, 0xb8, 0x19, 0xd5, 0x1e, 0xf6, 0xa3, 0x1f, 0xb2, 0xa2, 0x4a,
	0x53, 0x6c, 0x80, 0xc1, 0x0d, 0xb4, 0x47, 0x4c, 0xe5, 0x7f, 0xbe, 0x7f, 0x0b, 0x87, 0x23, 0xb1,
	0x60, 0xbc, 0x22, 0xa8, 0xb7, 0x70, 0x7e, 0xdb, 0xc2, 0xdd, 0xd9, 0xe2, 0xc3, 0x81, 0xb3, 0x49,
	0x96, 0x4c, 0x73, 0x1a, 0xaf, 0x5b, 0x15, 0xd7, 0x31, 0x34, 0x8a, 0xda, 0x36, 0x5d, 0xe2, 0x00,
	0xda, 0x62, 0x99, 0x8c, 0xb7, 0xa9, 0x36, 0x5b, 0x1a, 0xc1, 0x69, 0x55, 0x23, 0x4a, 0x2f, 0x37,
	0x5b, 0x18, 0x42, 0x77, 0x26, 0xf8, 0x9c, 0xc9, 0x74, 0xbc, 0x6d, 0xec, 0x6e, 0x3b, 0xb8, 0x82,
	0x56, 0x4c, 0x2a, 0x13, 0x5c, 0x91, 0x4e, 0x6c, 0x26, 0x92, 0x6a, 0x2d, 0x53, 0x6b, 0x7d, 0xa9,
	0x5a, 0xac, 0x55, 0xe8, 0x72, 0xf8, 0xe9, 0x82, 0xa7, 0x0d, 0xc2, 0x47, 0x7d, 0xb5, 0x8c, 0x06,
	0x03, 0x8b, 0x8b, 0x3b, 0xb9, 0xf5, 0xfa, 0x56, 0x4c, 0x39, 0x3b, 0xf8, 0x87, 0x0f, 0xd0, 0x34,
	0x2e, 0xe3, 0x85, 0x05, 0xb7, 0xe9, 0xff, 0x3e, 0xa2, 0x27, 0xe8, 0x6c, 0x7b, 0x8d, 0xa1, 0x2d,
	0x63, 0x5b, 0x1c, 0xfb, 0xa8, 0xef, 0xc1, 0xd3, 0x2f, 0x09, 0xcf, 0x6d, 0x12, 0xbf, 0x9f, 0xd8,
	0x1e, 0x9a, 0x97, 0xff, 0xe6, 0x87, 0xba, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x97, 0x35,
	0xb3, 0x5e, 0x03, 0x00, 0x00,
}
