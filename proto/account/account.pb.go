// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/microhq/user-srv/proto/account/account.proto

/*
Package account is a generated protocol buffer package.

It is generated from these files:
	github.com/microhq/user-srv/proto/account/account.proto

It has these top-level messages:
	User
	Session
	CreateRequest
	CreateResponse
	DeleteRequest
	DeleteResponse
	ReadRequest
	ReadResponse
	UpdateRequest
	UpdateResponse
	UpdatePasswordRequest
	UpdatePasswordResponse
	SearchRequest
	SearchResponse
	ReadSessionRequest
	ReadSessionResponse
	LoginRequest
	LoginResponse
	LogoutRequest
	LogoutResponse
*/
package account

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Created  int64  `protobuf:"varint,4,opt,name=created" json:"created,omitempty"`
	Updated  int64  `protobuf:"varint,5,opt,name=updated" json:"updated,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *User) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type Session struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Created  int64  `protobuf:"varint,3,opt,name=created" json:"created,omitempty"`
	Expires  int64  `protobuf:"varint,4,opt,name=expires" json:"expires,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Session) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Session) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Session) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Session) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

type CreateRequest struct {
	User     *User  `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateResponse struct {
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type DeleteRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ReadRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReadResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateRequest struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpdateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateResponse struct {
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type UpdatePasswordRequest struct {
	UserId          string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	OldPassword     string `protobuf:"bytes,2,opt,name=oldPassword" json:"oldPassword,omitempty"`
	NewPassword     string `protobuf:"bytes,3,opt,name=newPassword" json:"newPassword,omitempty"`
	ConfirmPassword string `protobuf:"bytes,4,opt,name=confirmPassword" json:"confirmPassword,omitempty"`
}

func (m *UpdatePasswordRequest) Reset()                    { *m = UpdatePasswordRequest{} }
func (m *UpdatePasswordRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdatePasswordRequest) ProtoMessage()               {}
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UpdatePasswordRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
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

type UpdatePasswordResponse struct {
}

func (m *UpdatePasswordResponse) Reset()                    { *m = UpdatePasswordResponse{} }
func (m *UpdatePasswordResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdatePasswordResponse) ProtoMessage()               {}
func (*UpdatePasswordResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type SearchRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Limit    int64  `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset   int64  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *SearchRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SearchRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SearchRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type SearchResponse struct {
	Users []*User `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SearchResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type ReadSessionRequest struct {
	SessionId string `protobuf:"bytes,1,opt,name=sessionId" json:"sessionId,omitempty"`
}

func (m *ReadSessionRequest) Reset()                    { *m = ReadSessionRequest{} }
func (m *ReadSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadSessionRequest) ProtoMessage()               {}
func (*ReadSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ReadSessionRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type ReadSessionResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *ReadSessionResponse) Reset()                    { *m = ReadSessionResponse{} }
func (m *ReadSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadSessionResponse) ProtoMessage()               {}
func (*ReadSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *ReadSessionResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *LoginResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type LogoutRequest struct {
	SessionId string `protobuf:"bytes,1,opt,name=sessionId" json:"sessionId,omitempty"`
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *LogoutRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type LogoutResponse struct {
}

func (m *LogoutResponse) Reset()                    { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string            { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()               {}
func (*LogoutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func init() {
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*Session)(nil), "Session")
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "DeleteResponse")
	proto.RegisterType((*ReadRequest)(nil), "ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "UpdateResponse")
	proto.RegisterType((*UpdatePasswordRequest)(nil), "UpdatePasswordRequest")
	proto.RegisterType((*UpdatePasswordResponse)(nil), "UpdatePasswordResponse")
	proto.RegisterType((*SearchRequest)(nil), "SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "SearchResponse")
	proto.RegisterType((*ReadSessionRequest)(nil), "ReadSessionRequest")
	proto.RegisterType((*ReadSessionResponse)(nil), "ReadSessionResponse")
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "LogoutResponse")
}

func init() {
	proto.RegisterFile("github.com/microhq/user-srv/proto/account/account.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 634 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xb6, 0xe3, 0xb8, 0x69, 0xa7, 0xb5, 0x5d, 0x6d, 0x4b, 0x31, 0x06, 0x44, 0xb4, 0x12, 0x52,
	0x00, 0x75, 0x23, 0xb5, 0x07, 0x04, 0x37, 0x54, 0x84, 0x84, 0xc4, 0xa1, 0x4a, 0xd5, 0x1b, 0x17,
	0xd7, 0xde, 0xb6, 0x2b, 0xc5, 0xde, 0xd4, 0x6b, 0x53, 0x0e, 0xbc, 0x0a, 0x6f, 0xc2, 0xc3, 0xa1,
	0xfd, 0x4b, 0x6c, 0xb7, 0x01, 0xca, 0x29, 0x9a, 0x99, 0x6f, 0x7e, 0x76, 0xe6, 0xfb, 0x1c, 0x78,
	0x7b, 0xc5, 0xea, 0xeb, 0xe6, 0x82, 0x64, 0xbc, 0x98, 0x16, 0x2c, 0xab, 0xf8, 0xf5, 0xcd, 0xb4,
	0x11, 0xb4, 0x3a, 0x14, 0xd5, 0xb7, 0xe9, 0xa2, 0xe2, 0x35, 0x9f, 0xa6, 0x59, 0xc6, 0x9b, 0xb2,
	0xb6, 0xbf, 0x44, 0x79, 0xf1, 0x0f, 0x18, 0x9e, 0x0b, 0x5a, 0xa1, 0x10, 0x06, 0x2c, 0x8f, 0xdd,
	0xb1, 0x3b, 0xd9, 0x9a, 0x0d, 0x58, 0x8e, 0x12, 0xd8, 0x94, 0xf9, 0x65, 0x5a, 0xd0, 0x78, 0xa0,
	0xbc, 0x4b, 0x1b, 0xed, 0x83, 0x4f, 0x8b, 0x94, 0xcd, 0x63, 0x4f, 0x05, 0xb4, 0x81, 0x62, 0x18,
	0x65, 0x15, 0x4d, 0x6b, 0x9a, 0xc7, 0xc3, 0xb1, 0x3b, 0xf1, 0x66, 0xd6, 0x94, 0x91, 0x66, 0x91,
	0xab, 0x88, 0xaf, 0x23, 0xc6, 0xc4, 0x0c, 0x46, 0x67, 0x54, 0x08, 0xc6, 0xcb, 0x07, 0x0d, 0xd0,
	0x6a, 0xe5, 0xdd, 0x69, 0x45, 0xbf, 0x2f, 0x58, 0x45, 0x85, 0x1d, 0xc2, 0x98, 0xf8, 0x13, 0x04,
	0x27, 0x0a, 0x34, 0xa3, 0x37, 0x0d, 0x15, 0x35, 0x7a, 0x02, 0x43, 0x59, 0x50, 0xb5, 0xdc, 0x3e,
	0xf2, 0x89, 0x5c, 0xc3, 0x4c, 0xb9, 0x64, 0xef, 0x45, 0x2a, 0xc4, 0x2d, 0xaf, 0x72, 0xdb, 0xdb,
	0xda, 0x78, 0x17, 0x42, 0x5b, 0x47, 0x2c, 0x78, 0x29, 0x28, 0x7e, 0x01, 0xc1, 0x47, 0x3a, 0xa7,
	0xab, 0xca, 0xbd, 0xa7, 0xc8, 0x14, 0x0b, 0x30, 0x29, 0xcf, 0x61, 0x7b, 0x46, 0xd3, 0x7c, 0x5d,
	0xc2, 0x2b, 0xd8, 0xd1, 0x61, 0x0d, 0xff, 0xc3, 0xa8, 0xf8, 0x35, 0x04, 0xe7, 0x6a, 0x99, 0x7f,
	0x7f, 0x96, 0x9c, 0xc3, 0x62, 0xcd, 0x1c, 0x3f, 0x5d, 0x78, 0xa4, 0x5d, 0xa7, 0xe6, 0x7d, 0xb6,
	0xcc, 0x01, 0x6c, 0xc8, 0x9c, 0xcf, 0x76, 0x2c, 0x63, 0xa1, 0x31, 0x6c, 0xf3, 0x79, 0x7e, 0xda,
	0xdd, 0x4e, 0xdb, 0x25, 0x11, 0x25, 0xbd, 0x5d, 0x22, 0x34, 0x47, 0xda, 0x2e, 0x34, 0x81, 0x28,
	0xe3, 0xe5, 0x25, 0xab, 0x8a, 0x25, 0x6a, 0xa8, 0x50, 0x7d, 0x37, 0x8e, 0xe1, 0xa0, 0x3f, 0x9e,
	0x99, 0x9c, 0x43, 0x70, 0x46, 0xd3, 0x2a, 0xbb, 0xb6, 0x03, 0xb7, 0xf9, 0xe2, 0xae, 0x23, 0xec,
	0xa0, 0x4d, 0xd8, 0x7d, 0xf0, 0xe7, 0xac, 0x60, 0xb5, 0xe1, 0x90, 0x36, 0xe4, 0xc3, 0xf9, 0xe5,
	0xa5, 0xa0, 0xb5, 0x21, 0x90, 0xb1, 0xf0, 0x21, 0x84, 0xb6, 0xa1, 0xb9, 0xca, 0x53, 0xf0, 0x65,
	0x07, 0x11, 0xbb, 0x63, 0x6f, 0xb5, 0x6a, 0xed, 0xc3, 0x47, 0x80, 0xe4, 0x09, 0x0d, 0xbb, 0xed,
	0x90, 0xcf, 0x60, 0x4b, 0x68, 0xcf, 0x72, 0xb1, 0x2b, 0x07, 0x7e, 0x07, 0x7b, 0x9d, 0x1c, 0xd3,
	0x07, 0xc3, 0xc8, 0x60, 0xcc, 0x51, 0x37, 0x89, 0x85, 0xd8, 0x00, 0xfe, 0x0a, 0x3b, 0x5f, 0xf8,
	0x15, 0x2b, 0xff, 0x7f, 0x1b, 0x6d, 0xce, 0x7b, 0x3d, 0xce, 0x1f, 0x43, 0x60, 0xaa, 0x3f, 0x60,
	0xa4, 0x43, 0x95, 0xc4, 0x9b, 0xfa, 0xdf, 0x1e, 0xbf, 0x0b, 0xa1, 0x85, 0xeb, 0x26, 0x47, 0xbf,
	0x3c, 0x18, 0x7d, 0xd0, 0x1f, 0x2b, 0xf4, 0x06, 0x36, 0xb4, 0xea, 0x50, 0x48, 0x3a, 0x32, 0x4e,
	0x22, 0xd2, 0x93, 0xa3, 0x83, 0x5e, 0xc2, 0x50, 0xee, 0x11, 0xed, 0x90, 0x96, 0xc8, 0x92, 0x80,
	0xb4, 0x35, 0x85, 0x1d, 0x59, 0x53, 0x93, 0x0b, 0x85, 0xa4, 0xa3, 0xa1, 0x24, 0x22, 0x3d, 0x9d,
	0x28, 0xb0, 0xd6, 0x30, 0x0a, 0x49, 0x47, 0xed, 0x49, 0x44, 0x7a, 0xe2, 0x56, 0x60, 0xcd, 0x15,
	0x14, 0x92, 0x0e, 0x4b, 0x93, 0x88, 0x74, 0x49, 0x84, 0x1d, 0x74, 0x62, 0x55, 0xb9, 0xd4, 0xc7,
	0x01, 0xb9, 0x57, 0x93, 0xc9, 0x63, 0xb2, 0x46, 0x0c, 0x0e, 0x9a, 0x80, 0xaf, 0x2e, 0x84, 0x02,
	0xd2, 0xe6, 0x41, 0x12, 0x92, 0xce, 0xe1, 0xf4, 0x6c, 0x7a, 0xcf, 0x48, 0xc5, 0x56, 0xf7, 0x49,
	0x22, 0xd2, 0x3d, 0x00, 0x76, 0xd0, 0x7b, 0xfd, 0x9d, 0xb2, 0xdf, 0xe8, 0x3d, 0x72, 0x97, 0xd3,
	0xc9, 0x3e, 0xb9, 0x87, 0xb4, 0xd8, 0xb9, 0xd8, 0x50, 0x7f, 0x30, 0xc7, 0xbf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x54, 0xf9, 0x45, 0x4e, 0x9b, 0x06, 0x00, 0x00,
}
