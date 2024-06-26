// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: proto/smartbalance.proto

package api

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

// Struct
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_smartbalance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smartbalance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_smartbalance_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CoolingSystem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coolinglevel     string `protobuf:"bytes,1,opt,name=coolinglevel,proto3" json:"coolinglevel,omitempty"`
	Coolingfrequency string `protobuf:"bytes,2,opt,name=coolingfrequency,proto3" json:"coolingfrequency,omitempty"`
	Coolingtype      string `protobuf:"bytes,3,opt,name=coolingtype,proto3" json:"coolingtype,omitempty"`
}

func (x *CoolingSystem) Reset() {
	*x = CoolingSystem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_smartbalance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoolingSystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoolingSystem) ProtoMessage() {}

func (x *CoolingSystem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smartbalance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoolingSystem.ProtoReflect.Descriptor instead.
func (*CoolingSystem) Descriptor() ([]byte, []int) {
	return file_proto_smartbalance_proto_rawDescGZIP(), []int{1}
}

func (x *CoolingSystem) GetCoolinglevel() string {
	if x != nil {
		return x.Coolinglevel
	}
	return ""
}

func (x *CoolingSystem) GetCoolingfrequency() string {
	if x != nil {
		return x.Coolingfrequency
	}
	return ""
}

func (x *CoolingSystem) GetCoolingtype() string {
	if x != nil {
		return x.Coolingtype
	}
	return ""
}

// Cooling System defenition
type CoolingSystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *CoolingSystem `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *CoolingSystemRequest) Reset() {
	*x = CoolingSystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_smartbalance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoolingSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoolingSystemRequest) ProtoMessage() {}

func (x *CoolingSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smartbalance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoolingSystemRequest.ProtoReflect.Descriptor instead.
func (*CoolingSystemRequest) Descriptor() ([]byte, []int) {
	return file_proto_smartbalance_proto_rawDescGZIP(), []int{2}
}

func (x *CoolingSystemRequest) GetInfo() *CoolingSystem {
	if x != nil {
		return x.Info
	}
	return nil
}

type CoolingSystemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record string `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *CoolingSystemResponse) Reset() {
	*x = CoolingSystemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_smartbalance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoolingSystemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoolingSystemResponse) ProtoMessage() {}

func (x *CoolingSystemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smartbalance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoolingSystemResponse.ProtoReflect.Descriptor instead.
func (*CoolingSystemResponse) Descriptor() ([]byte, []int) {
	return file_proto_smartbalance_proto_rawDescGZIP(), []int{3}
}

func (x *CoolingSystemResponse) GetRecord() string {
	if x != nil {
		return x.Record
	}
	return ""
}

type CoolingSystemGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record string `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *CoolingSystemGetRequest) Reset() {
	*x = CoolingSystemGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_smartbalance_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoolingSystemGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoolingSystemGetRequest) ProtoMessage() {}

func (x *CoolingSystemGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smartbalance_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoolingSystemGetRequest.ProtoReflect.Descriptor instead.
func (*CoolingSystemGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_smartbalance_proto_rawDescGZIP(), []int{4}
}

func (x *CoolingSystemGetRequest) GetRecord() string {
	if x != nil {
		return x.Record
	}
	return ""
}

type CoolingSystemGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CoolingSystemGetResponse) Reset() {
	*x = CoolingSystemGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_smartbalance_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoolingSystemGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoolingSystemGetResponse) ProtoMessage() {}

func (x *CoolingSystemGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smartbalance_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoolingSystemGetResponse.ProtoReflect.Descriptor instead.
func (*CoolingSystemGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_smartbalance_proto_rawDescGZIP(), []int{5}
}

func (x *CoolingSystemGetResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Check User
type CheckUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *User `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *CheckUserRequest) Reset() {
	*x = CheckUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_smartbalance_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserRequest) ProtoMessage() {}

func (x *CheckUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smartbalance_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserRequest.ProtoReflect.Descriptor instead.
func (*CheckUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_smartbalance_proto_rawDescGZIP(), []int{6}
}

func (x *CheckUserRequest) GetInfo() *User {
	if x != nil {
		return x.Info
	}
	return nil
}

type CheckUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CheckUserResponse) Reset() {
	*x = CheckUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_smartbalance_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserResponse) ProtoMessage() {}

func (x *CheckUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smartbalance_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserResponse.ProtoReflect.Descriptor instead.
func (*CheckUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_smartbalance_proto_rawDescGZIP(), []int{7}
}

func (x *CheckUserResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// Create User
type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *User `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_smartbalance_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smartbalance_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_smartbalance_proto_rawDescGZIP(), []int{8}
}

func (x *CreateUserRequest) GetInfo() *User {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Confirm string `protobuf:"bytes,1,opt,name=confirm,proto3" json:"confirm,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_smartbalance_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smartbalance_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_smartbalance_proto_rawDescGZIP(), []int{9}
}

func (x *CreateUserResponse) GetConfirm() string {
	if x != nil {
		return x.Confirm
	}
	return ""
}

// Dashboard
type DashboardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag string `protobuf:"bytes,1,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *DashboardRequest) Reset() {
	*x = DashboardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_smartbalance_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DashboardRequest) ProtoMessage() {}

func (x *DashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smartbalance_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DashboardRequest.ProtoReflect.Descriptor instead.
func (*DashboardRequest) Descriptor() ([]byte, []int) {
	return file_proto_smartbalance_proto_rawDescGZIP(), []int{10}
}

func (x *DashboardRequest) GetFlag() string {
	if x != nil {
		return x.Flag
	}
	return ""
}

type DashboardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*CoolingSystem `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *DashboardResponse) Reset() {
	*x = DashboardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_smartbalance_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DashboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DashboardResponse) ProtoMessage() {}

func (x *DashboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smartbalance_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DashboardResponse.ProtoReflect.Descriptor instead.
func (*DashboardResponse) Descriptor() ([]byte, []int) {
	return file_proto_smartbalance_proto_rawDescGZIP(), []int{11}
}

func (x *DashboardResponse) GetInfo() []*CoolingSystem {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_proto_smartbalance_proto protoreflect.FileDescriptor

var file_proto_smartbalance_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x70, 0x69, 0x22, 0x3e, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x0d,
	0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x66, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6f,
	0x6c, 0x69, 0x6e, 0x67, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x4a, 0x0a, 0x14, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x2f, 0x0a, 0x15, 0x43,
	0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x31, 0x0a, 0x17,
	0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22,
	0x30, 0x0a, 0x18, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x3d, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x22, 0x29, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3e, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x70, 0x69,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x2e, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x22, 0x26, 0x0a, 0x10, 0x44,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x6c, 0x61, 0x67, 0x22, 0x47, 0x0a, 0x11, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xdf, 0x03, 0x0a,
	0x13, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x21, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x41, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5e, 0x0a, 0x0d, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x12, 0x25, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41,
	0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e,
	0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x69, 0x0a, 0x12, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x28, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x70,
	0x69, 0x2e, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x09, 0x44, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x21, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_smartbalance_proto_rawDescOnce sync.Once
	file_proto_smartbalance_proto_rawDescData = file_proto_smartbalance_proto_rawDesc
)

func file_proto_smartbalance_proto_rawDescGZIP() []byte {
	file_proto_smartbalance_proto_rawDescOnce.Do(func() {
		file_proto_smartbalance_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_smartbalance_proto_rawDescData)
	})
	return file_proto_smartbalance_proto_rawDescData
}

var file_proto_smartbalance_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_smartbalance_proto_goTypes = []interface{}{
	(*User)(nil),                     // 0: smartbalanceApi.User
	(*CoolingSystem)(nil),            // 1: smartbalanceApi.CoolingSystem
	(*CoolingSystemRequest)(nil),     // 2: smartbalanceApi.CoolingSystemRequest
	(*CoolingSystemResponse)(nil),    // 3: smartbalanceApi.CoolingSystemResponse
	(*CoolingSystemGetRequest)(nil),  // 4: smartbalanceApi.CoolingSystemGetRequest
	(*CoolingSystemGetResponse)(nil), // 5: smartbalanceApi.CoolingSystemGetResponse
	(*CheckUserRequest)(nil),         // 6: smartbalanceApi.CheckUserRequest
	(*CheckUserResponse)(nil),        // 7: smartbalanceApi.CheckUserResponse
	(*CreateUserRequest)(nil),        // 8: smartbalanceApi.CreateUserRequest
	(*CreateUserResponse)(nil),       // 9: smartbalanceApi.CreateUserResponse
	(*DashboardRequest)(nil),         // 10: smartbalanceApi.DashboardRequest
	(*DashboardResponse)(nil),        // 11: smartbalanceApi.DashboardResponse
}
var file_proto_smartbalance_proto_depIdxs = []int32{
	1,  // 0: smartbalanceApi.CoolingSystemRequest.info:type_name -> smartbalanceApi.CoolingSystem
	0,  // 1: smartbalanceApi.CheckUserRequest.info:type_name -> smartbalanceApi.User
	0,  // 2: smartbalanceApi.CreateUserRequest.info:type_name -> smartbalanceApi.User
	1,  // 3: smartbalanceApi.DashboardResponse.info:type_name -> smartbalanceApi.CoolingSystem
	6,  // 4: smartbalanceApi.SmartBalanceService.CheckUser:input_type -> smartbalanceApi.CheckUserRequest
	8,  // 5: smartbalanceApi.SmartBalanceService.CreateUser:input_type -> smartbalanceApi.CreateUserRequest
	2,  // 6: smartbalanceApi.SmartBalanceService.CoolingSystem:input_type -> smartbalanceApi.CoolingSystemRequest
	4,  // 7: smartbalanceApi.SmartBalanceService.CoolingSystemCheck:input_type -> smartbalanceApi.CoolingSystemGetRequest
	10, // 8: smartbalanceApi.SmartBalanceService.Dashboard:input_type -> smartbalanceApi.DashboardRequest
	7,  // 9: smartbalanceApi.SmartBalanceService.CheckUser:output_type -> smartbalanceApi.CheckUserResponse
	9,  // 10: smartbalanceApi.SmartBalanceService.CreateUser:output_type -> smartbalanceApi.CreateUserResponse
	3,  // 11: smartbalanceApi.SmartBalanceService.CoolingSystem:output_type -> smartbalanceApi.CoolingSystemResponse
	5,  // 12: smartbalanceApi.SmartBalanceService.CoolingSystemCheck:output_type -> smartbalanceApi.CoolingSystemGetResponse
	11, // 13: smartbalanceApi.SmartBalanceService.Dashboard:output_type -> smartbalanceApi.DashboardResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_smartbalance_proto_init() }
func file_proto_smartbalance_proto_init() {
	if File_proto_smartbalance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_smartbalance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_smartbalance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoolingSystem); i {
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
		file_proto_smartbalance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoolingSystemRequest); i {
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
		file_proto_smartbalance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoolingSystemResponse); i {
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
		file_proto_smartbalance_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoolingSystemGetRequest); i {
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
		file_proto_smartbalance_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoolingSystemGetResponse); i {
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
		file_proto_smartbalance_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserRequest); i {
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
		file_proto_smartbalance_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserResponse); i {
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
		file_proto_smartbalance_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_proto_smartbalance_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_proto_smartbalance_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DashboardRequest); i {
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
		file_proto_smartbalance_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DashboardResponse); i {
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
			RawDescriptor: file_proto_smartbalance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_smartbalance_proto_goTypes,
		DependencyIndexes: file_proto_smartbalance_proto_depIdxs,
		MessageInfos:      file_proto_smartbalance_proto_msgTypes,
	}.Build()
	File_proto_smartbalance_proto = out.File
	file_proto_smartbalance_proto_rawDesc = nil
	file_proto_smartbalance_proto_goTypes = nil
	file_proto_smartbalance_proto_depIdxs = nil
}
