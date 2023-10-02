// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: proto/auth.proto

package authGrpc

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

// QUERIES
type UserGetQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserGetQuestionRequest) Reset() {
	*x = UserGetQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGetQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGetQuestionRequest) ProtoMessage() {}

func (x *UserGetQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGetQuestionRequest.ProtoReflect.Descriptor instead.
func (*UserGetQuestionRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *UserGetQuestionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserGetRequest) Reset() {
	*x = UserGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGetRequest) ProtoMessage() {}

func (x *UserGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGetRequest.ProtoReflect.Descriptor instead.
func (*UserGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *UserGetRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// MUTATIONS
type UserSignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Question string `protobuf:"bytes,3,opt,name=question,proto3" json:"question,omitempty"`
	Answer   string `protobuf:"bytes,4,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *UserSignUpRequest) Reset() {
	*x = UserSignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSignUpRequest) ProtoMessage() {}

func (x *UserSignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSignUpRequest.ProtoReflect.Descriptor instead.
func (*UserSignUpRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *UserSignUpRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserSignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserSignUpRequest) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *UserSignUpRequest) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type UserSignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserSignInRequest) Reset() {
	*x = UserSignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSignInRequest) ProtoMessage() {}

func (x *UserSignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSignInRequest.ProtoReflect.Descriptor instead.
func (*UserSignInRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{3}
}

func (x *UserSignInRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserSignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserSignOutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserSignOutRequest) Reset() {
	*x = UserSignOutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSignOutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSignOutRequest) ProtoMessage() {}

func (x *UserSignOutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSignOutRequest.ProtoReflect.Descriptor instead.
func (*UserSignOutRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{4}
}

func (x *UserSignOutRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldPassword string `protobuf:"bytes,1,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword string `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	UserId      string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserChangePasswordRequest) Reset() {
	*x = UserChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserChangePasswordRequest) ProtoMessage() {}

func (x *UserChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*UserChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{5}
}

func (x *UserChangePasswordRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *UserChangePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *UserChangePasswordRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserSendAnswerAndChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer      string `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
	NewPassword string `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	UserId      string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserSendAnswerAndChangePasswordRequest) Reset() {
	*x = UserSendAnswerAndChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSendAnswerAndChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSendAnswerAndChangePasswordRequest) ProtoMessage() {}

func (x *UserSendAnswerAndChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSendAnswerAndChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*UserSendAnswerAndChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{6}
}

func (x *UserSendAnswerAndChangePasswordRequest) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *UserSendAnswerAndChangePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *UserSendAnswerAndChangePasswordRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type AccessToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *AccessToken) Reset() {
	*x = AccessToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessToken) ProtoMessage() {}

func (x *AccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessToken.ProtoReflect.Descriptor instead.
func (*AccessToken) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{7}
}

func (x *AccessToken) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type RefreshToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RefreshToken) Reset() {
	*x = RefreshToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshToken) ProtoMessage() {}

func (x *RefreshToken) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshToken.ProtoReflect.Descriptor instead.
func (*RefreshToken) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{8}
}

func (x *RefreshToken) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

// Only no token required (no auth required) request OR ValidateAuth ( RecreateTokens includes )
type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AccessToken  string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{9}
}

func (x *UserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *UserResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type UserQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question string `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
}

func (x *UserQuestionResponse) Reset() {
	*x = UserQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserQuestionResponse) ProtoMessage() {}

func (x *UserQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserQuestionResponse.ProtoReflect.Descriptor instead.
func (*UserQuestionResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{10}
}

func (x *UserQuestionResponse) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

type UserDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UserDataResponse) Reset() {
	*x = UserDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDataResponse) ProtoMessage() {}

func (x *UserDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDataResponse.ProtoReflect.Descriptor instead.
func (*UserDataResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{11}
}

func (x *UserDataResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UserChangePasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *UserChangePasswordResponse) Reset() {
	*x = UserChangePasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserChangePasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserChangePasswordResponse) ProtoMessage() {}

func (x *UserChangePasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserChangePasswordResponse.ProtoReflect.Descriptor instead.
func (*UserChangePasswordResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{12}
}

func (x *UserChangePasswordResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *UserChangePasswordResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type Nil struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Nil) Reset() {
	*x = Nil{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nil) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nil) ProtoMessage() {}

func (x *Nil) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nil.ProtoReflect.Descriptor instead.
func (*Nil) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{13}
}

var File_proto_auth_proto protoreflect.FileDescriptor

var file_proto_auth_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x61, 0x75, 0x74, 0x68, 0x47, 0x72, 0x70, 0x63, 0x22, 0x31, 0x0a, 0x16,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x29, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7f, 0x0a, 0x11, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x11, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7a, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e,
	0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x7c, 0x0a, 0x26, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x30, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x33, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6f, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x32, 0x0a, 0x14, 0x55, 0x73, 0x65,
	0x72, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2e, 0x0a,
	0x10, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x64, 0x0a,
	0x1a, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x05, 0x0a, 0x03, 0x4e, 0x69, 0x6c, 0x32, 0xb0, 0x05, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x1b, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67,
	0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x1b, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e,
	0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x47, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x07, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x12, 0x1c, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e,
	0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x69, 0x6c, 0x12, 0x3d, 0x0a, 0x0c, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x47, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x1a, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x1c, 0x52, 0x65, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x79, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x47,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x1a, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x47, 0x72,
	0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x47, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x47,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70,
	0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x42, 0x79, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x30, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x47,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x41, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5b, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x47, 0x72,
	0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x68, 0x77,
	0x79, 0x79, 0x2f, 0x4d, 0x61, 0x6b, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x47, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_auth_proto_rawDescOnce sync.Once
	file_proto_auth_proto_rawDescData = file_proto_auth_proto_rawDesc
)

func file_proto_auth_proto_rawDescGZIP() []byte {
	file_proto_auth_proto_rawDescOnce.Do(func() {
		file_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_auth_proto_rawDescData)
	})
	return file_proto_auth_proto_rawDescData
}

var file_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_proto_auth_proto_goTypes = []interface{}{
	(*UserGetQuestionRequest)(nil),                 // 0: authGrpc.UserGetQuestionRequest
	(*UserGetRequest)(nil),                         // 1: authGrpc.UserGetRequest
	(*UserSignUpRequest)(nil),                      // 2: authGrpc.UserSignUpRequest
	(*UserSignInRequest)(nil),                      // 3: authGrpc.UserSignInRequest
	(*UserSignOutRequest)(nil),                     // 4: authGrpc.UserSignOutRequest
	(*UserChangePasswordRequest)(nil),              // 5: authGrpc.UserChangePasswordRequest
	(*UserSendAnswerAndChangePasswordRequest)(nil), // 6: authGrpc.UserSendAnswerAndChangePasswordRequest
	(*AccessToken)(nil),                            // 7: authGrpc.AccessToken
	(*RefreshToken)(nil),                           // 8: authGrpc.RefreshToken
	(*UserResponse)(nil),                           // 9: authGrpc.UserResponse
	(*UserQuestionResponse)(nil),                   // 10: authGrpc.UserQuestionResponse
	(*UserDataResponse)(nil),                       // 11: authGrpc.UserDataResponse
	(*UserChangePasswordResponse)(nil),             // 12: authGrpc.UserChangePasswordResponse
	(*Nil)(nil),                                    // 13: authGrpc.Nil
}
var file_proto_auth_proto_depIdxs = []int32{
	2,  // 0: authGrpc.User.SignUp:input_type -> authGrpc.UserSignUpRequest
	3,  // 1: authGrpc.User.SignIn:input_type -> authGrpc.UserSignInRequest
	4,  // 2: authGrpc.User.SignOut:input_type -> authGrpc.UserSignOutRequest
	7,  // 3: authGrpc.User.ValidateAuth:input_type -> authGrpc.AccessToken
	8,  // 4: authGrpc.User.RecreateTokensByRefreshToken:input_type -> authGrpc.RefreshToken
	0,  // 5: authGrpc.User.GetQuestion:input_type -> authGrpc.UserGetQuestionRequest
	1,  // 6: authGrpc.User.GetUserById:input_type -> authGrpc.UserGetRequest
	6,  // 7: authGrpc.User.ChangePasswordByAnswer:input_type -> authGrpc.UserSendAnswerAndChangePasswordRequest
	5,  // 8: authGrpc.User.ChangePassword:input_type -> authGrpc.UserChangePasswordRequest
	9,  // 9: authGrpc.User.SignUp:output_type -> authGrpc.UserResponse
	9,  // 10: authGrpc.User.SignIn:output_type -> authGrpc.UserResponse
	13, // 11: authGrpc.User.SignOut:output_type -> authGrpc.Nil
	9,  // 12: authGrpc.User.ValidateAuth:output_type -> authGrpc.UserResponse
	9,  // 13: authGrpc.User.RecreateTokensByRefreshToken:output_type -> authGrpc.UserResponse
	10, // 14: authGrpc.User.GetQuestion:output_type -> authGrpc.UserQuestionResponse
	11, // 15: authGrpc.User.GetUserById:output_type -> authGrpc.UserDataResponse
	12, // 16: authGrpc.User.ChangePasswordByAnswer:output_type -> authGrpc.UserChangePasswordResponse
	12, // 17: authGrpc.User.ChangePassword:output_type -> authGrpc.UserChangePasswordResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_auth_proto_init() }
func file_proto_auth_proto_init() {
	if File_proto_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGetQuestionRequest); i {
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
		file_proto_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGetRequest); i {
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
		file_proto_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSignUpRequest); i {
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
		file_proto_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSignInRequest); i {
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
		file_proto_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSignOutRequest); i {
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
		file_proto_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChangePasswordRequest); i {
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
		file_proto_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSendAnswerAndChangePasswordRequest); i {
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
		file_proto_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessToken); i {
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
		file_proto_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshToken); i {
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
		file_proto_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
		file_proto_auth_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserQuestionResponse); i {
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
		file_proto_auth_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDataResponse); i {
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
		file_proto_auth_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChangePasswordResponse); i {
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
		file_proto_auth_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nil); i {
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
			RawDescriptor: file_proto_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_auth_proto_goTypes,
		DependencyIndexes: file_proto_auth_proto_depIdxs,
		MessageInfos:      file_proto_auth_proto_msgTypes,
	}.Build()
	File_proto_auth_proto = out.File
	file_proto_auth_proto_rawDesc = nil
	file_proto_auth_proto_goTypes = nil
	file_proto_auth_proto_depIdxs = nil
}
