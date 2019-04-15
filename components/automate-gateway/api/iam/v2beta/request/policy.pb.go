// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/request/policy.proto

package request // import "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/request"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Does not contain type as the enduser can only create 'custom' policies.
type CreatePolicyReq struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Members              []string            `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	Statements           []*common.Statement `protobuf:"bytes,4,rep,name=statements,proto3" json:"statements,omitempty"`
	Projects             []string            `protobuf:"bytes,5,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CreatePolicyReq) Reset()         { *m = CreatePolicyReq{} }
func (m *CreatePolicyReq) String() string { return proto.CompactTextString(m) }
func (*CreatePolicyReq) ProtoMessage()    {}
func (*CreatePolicyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{0}
}
func (m *CreatePolicyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePolicyReq.Unmarshal(m, b)
}
func (m *CreatePolicyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePolicyReq.Marshal(b, m, deterministic)
}
func (dst *CreatePolicyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePolicyReq.Merge(dst, src)
}
func (m *CreatePolicyReq) XXX_Size() int {
	return xxx_messageInfo_CreatePolicyReq.Size(m)
}
func (m *CreatePolicyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePolicyReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePolicyReq proto.InternalMessageInfo

func (m *CreatePolicyReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreatePolicyReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreatePolicyReq) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *CreatePolicyReq) GetStatements() []*common.Statement {
	if m != nil {
		return m.Statements
	}
	return nil
}

func (m *CreatePolicyReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type DeletePolicyReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePolicyReq) Reset()         { *m = DeletePolicyReq{} }
func (m *DeletePolicyReq) String() string { return proto.CompactTextString(m) }
func (*DeletePolicyReq) ProtoMessage()    {}
func (*DeletePolicyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{1}
}
func (m *DeletePolicyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePolicyReq.Unmarshal(m, b)
}
func (m *DeletePolicyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePolicyReq.Marshal(b, m, deterministic)
}
func (dst *DeletePolicyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePolicyReq.Merge(dst, src)
}
func (m *DeletePolicyReq) XXX_Size() int {
	return xxx_messageInfo_DeletePolicyReq.Size(m)
}
func (m *DeletePolicyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePolicyReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePolicyReq proto.InternalMessageInfo

func (m *DeletePolicyReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListPoliciesReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPoliciesReq) Reset()         { *m = ListPoliciesReq{} }
func (m *ListPoliciesReq) String() string { return proto.CompactTextString(m) }
func (*ListPoliciesReq) ProtoMessage()    {}
func (*ListPoliciesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{2}
}
func (m *ListPoliciesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPoliciesReq.Unmarshal(m, b)
}
func (m *ListPoliciesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPoliciesReq.Marshal(b, m, deterministic)
}
func (dst *ListPoliciesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPoliciesReq.Merge(dst, src)
}
func (m *ListPoliciesReq) XXX_Size() int {
	return xxx_messageInfo_ListPoliciesReq.Size(m)
}
func (m *ListPoliciesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPoliciesReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListPoliciesReq proto.InternalMessageInfo

type AddPolicyMembersReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Members              []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPolicyMembersReq) Reset()         { *m = AddPolicyMembersReq{} }
func (m *AddPolicyMembersReq) String() string { return proto.CompactTextString(m) }
func (*AddPolicyMembersReq) ProtoMessage()    {}
func (*AddPolicyMembersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{3}
}
func (m *AddPolicyMembersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPolicyMembersReq.Unmarshal(m, b)
}
func (m *AddPolicyMembersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPolicyMembersReq.Marshal(b, m, deterministic)
}
func (dst *AddPolicyMembersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPolicyMembersReq.Merge(dst, src)
}
func (m *AddPolicyMembersReq) XXX_Size() int {
	return xxx_messageInfo_AddPolicyMembersReq.Size(m)
}
func (m *AddPolicyMembersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPolicyMembersReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddPolicyMembersReq proto.InternalMessageInfo

func (m *AddPolicyMembersReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddPolicyMembersReq) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type GetPolicyReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPolicyReq) Reset()         { *m = GetPolicyReq{} }
func (m *GetPolicyReq) String() string { return proto.CompactTextString(m) }
func (*GetPolicyReq) ProtoMessage()    {}
func (*GetPolicyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{4}
}
func (m *GetPolicyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPolicyReq.Unmarshal(m, b)
}
func (m *GetPolicyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPolicyReq.Marshal(b, m, deterministic)
}
func (dst *GetPolicyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPolicyReq.Merge(dst, src)
}
func (m *GetPolicyReq) XXX_Size() int {
	return xxx_messageInfo_GetPolicyReq.Size(m)
}
func (m *GetPolicyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPolicyReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetPolicyReq proto.InternalMessageInfo

func (m *GetPolicyReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Does not contain type as the enduser can only create 'custom' policies.
type UpdatePolicyReq struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Members              []string            `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	Statements           []*common.Statement `protobuf:"bytes,3,rep,name=statements,proto3" json:"statements,omitempty"`
	Name                 string              `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Projects             []string            `protobuf:"bytes,9,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UpdatePolicyReq) Reset()         { *m = UpdatePolicyReq{} }
func (m *UpdatePolicyReq) String() string { return proto.CompactTextString(m) }
func (*UpdatePolicyReq) ProtoMessage()    {}
func (*UpdatePolicyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{5}
}
func (m *UpdatePolicyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePolicyReq.Unmarshal(m, b)
}
func (m *UpdatePolicyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePolicyReq.Marshal(b, m, deterministic)
}
func (dst *UpdatePolicyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePolicyReq.Merge(dst, src)
}
func (m *UpdatePolicyReq) XXX_Size() int {
	return xxx_messageInfo_UpdatePolicyReq.Size(m)
}
func (m *UpdatePolicyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePolicyReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePolicyReq proto.InternalMessageInfo

func (m *UpdatePolicyReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdatePolicyReq) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *UpdatePolicyReq) GetStatements() []*common.Statement {
	if m != nil {
		return m.Statements
	}
	return nil
}

func (m *UpdatePolicyReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdatePolicyReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type UpgradeToV2Req struct {
	Flag                 common.Flag `protobuf:"varint,1,opt,name=flag,proto3,enum=chef.automate.api.iam.v2beta.Flag" json:"flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpgradeToV2Req) Reset()         { *m = UpgradeToV2Req{} }
func (m *UpgradeToV2Req) String() string { return proto.CompactTextString(m) }
func (*UpgradeToV2Req) ProtoMessage()    {}
func (*UpgradeToV2Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{6}
}
func (m *UpgradeToV2Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpgradeToV2Req.Unmarshal(m, b)
}
func (m *UpgradeToV2Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpgradeToV2Req.Marshal(b, m, deterministic)
}
func (dst *UpgradeToV2Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeToV2Req.Merge(dst, src)
}
func (m *UpgradeToV2Req) XXX_Size() int {
	return xxx_messageInfo_UpgradeToV2Req.Size(m)
}
func (m *UpgradeToV2Req) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeToV2Req.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeToV2Req proto.InternalMessageInfo

func (m *UpgradeToV2Req) GetFlag() common.Flag {
	if m != nil {
		return m.Flag
	}
	return common.Flag_VERSION_2_0
}

type GetPolicyVersionReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPolicyVersionReq) Reset()         { *m = GetPolicyVersionReq{} }
func (m *GetPolicyVersionReq) String() string { return proto.CompactTextString(m) }
func (*GetPolicyVersionReq) ProtoMessage()    {}
func (*GetPolicyVersionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{7}
}
func (m *GetPolicyVersionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPolicyVersionReq.Unmarshal(m, b)
}
func (m *GetPolicyVersionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPolicyVersionReq.Marshal(b, m, deterministic)
}
func (dst *GetPolicyVersionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPolicyVersionReq.Merge(dst, src)
}
func (m *GetPolicyVersionReq) XXX_Size() int {
	return xxx_messageInfo_GetPolicyVersionReq.Size(m)
}
func (m *GetPolicyVersionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPolicyVersionReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetPolicyVersionReq proto.InternalMessageInfo

type ResetToV1Req struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetToV1Req) Reset()         { *m = ResetToV1Req{} }
func (m *ResetToV1Req) String() string { return proto.CompactTextString(m) }
func (*ResetToV1Req) ProtoMessage()    {}
func (*ResetToV1Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{8}
}
func (m *ResetToV1Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetToV1Req.Unmarshal(m, b)
}
func (m *ResetToV1Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetToV1Req.Marshal(b, m, deterministic)
}
func (dst *ResetToV1Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetToV1Req.Merge(dst, src)
}
func (m *ResetToV1Req) XXX_Size() int {
	return xxx_messageInfo_ResetToV1Req.Size(m)
}
func (m *ResetToV1Req) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetToV1Req.DiscardUnknown(m)
}

var xxx_messageInfo_ResetToV1Req proto.InternalMessageInfo

type ListPolicyMembersReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPolicyMembersReq) Reset()         { *m = ListPolicyMembersReq{} }
func (m *ListPolicyMembersReq) String() string { return proto.CompactTextString(m) }
func (*ListPolicyMembersReq) ProtoMessage()    {}
func (*ListPolicyMembersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{9}
}
func (m *ListPolicyMembersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPolicyMembersReq.Unmarshal(m, b)
}
func (m *ListPolicyMembersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPolicyMembersReq.Marshal(b, m, deterministic)
}
func (dst *ListPolicyMembersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPolicyMembersReq.Merge(dst, src)
}
func (m *ListPolicyMembersReq) XXX_Size() int {
	return xxx_messageInfo_ListPolicyMembersReq.Size(m)
}
func (m *ListPolicyMembersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPolicyMembersReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListPolicyMembersReq proto.InternalMessageInfo

func (m *ListPolicyMembersReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReplacePolicyMembersReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Members              []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplacePolicyMembersReq) Reset()         { *m = ReplacePolicyMembersReq{} }
func (m *ReplacePolicyMembersReq) String() string { return proto.CompactTextString(m) }
func (*ReplacePolicyMembersReq) ProtoMessage()    {}
func (*ReplacePolicyMembersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{10}
}
func (m *ReplacePolicyMembersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplacePolicyMembersReq.Unmarshal(m, b)
}
func (m *ReplacePolicyMembersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplacePolicyMembersReq.Marshal(b, m, deterministic)
}
func (dst *ReplacePolicyMembersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplacePolicyMembersReq.Merge(dst, src)
}
func (m *ReplacePolicyMembersReq) XXX_Size() int {
	return xxx_messageInfo_ReplacePolicyMembersReq.Size(m)
}
func (m *ReplacePolicyMembersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplacePolicyMembersReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReplacePolicyMembersReq proto.InternalMessageInfo

func (m *ReplacePolicyMembersReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReplacePolicyMembersReq) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type RemovePolicyMembersReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Members              []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemovePolicyMembersReq) Reset()         { *m = RemovePolicyMembersReq{} }
func (m *RemovePolicyMembersReq) String() string { return proto.CompactTextString(m) }
func (*RemovePolicyMembersReq) ProtoMessage()    {}
func (*RemovePolicyMembersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{11}
}
func (m *RemovePolicyMembersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemovePolicyMembersReq.Unmarshal(m, b)
}
func (m *RemovePolicyMembersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemovePolicyMembersReq.Marshal(b, m, deterministic)
}
func (dst *RemovePolicyMembersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemovePolicyMembersReq.Merge(dst, src)
}
func (m *RemovePolicyMembersReq) XXX_Size() int {
	return xxx_messageInfo_RemovePolicyMembersReq.Size(m)
}
func (m *RemovePolicyMembersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RemovePolicyMembersReq.DiscardUnknown(m)
}

var xxx_messageInfo_RemovePolicyMembersReq proto.InternalMessageInfo

func (m *RemovePolicyMembersReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RemovePolicyMembersReq) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

// Does not contain type as the enduser can only create 'custom' roles.
type CreateRoleReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Actions              []string `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	Projects             []string `protobuf:"bytes,4,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoleReq) Reset()         { *m = CreateRoleReq{} }
func (m *CreateRoleReq) String() string { return proto.CompactTextString(m) }
func (*CreateRoleReq) ProtoMessage()    {}
func (*CreateRoleReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{12}
}
func (m *CreateRoleReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoleReq.Unmarshal(m, b)
}
func (m *CreateRoleReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoleReq.Marshal(b, m, deterministic)
}
func (dst *CreateRoleReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleReq.Merge(dst, src)
}
func (m *CreateRoleReq) XXX_Size() int {
	return xxx_messageInfo_CreateRoleReq.Size(m)
}
func (m *CreateRoleReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleReq proto.InternalMessageInfo

func (m *CreateRoleReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateRoleReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRoleReq) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *CreateRoleReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type GetRoleReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRoleReq) Reset()         { *m = GetRoleReq{} }
func (m *GetRoleReq) String() string { return proto.CompactTextString(m) }
func (*GetRoleReq) ProtoMessage()    {}
func (*GetRoleReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{13}
}
func (m *GetRoleReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoleReq.Unmarshal(m, b)
}
func (m *GetRoleReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoleReq.Marshal(b, m, deterministic)
}
func (dst *GetRoleReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoleReq.Merge(dst, src)
}
func (m *GetRoleReq) XXX_Size() int {
	return xxx_messageInfo_GetRoleReq.Size(m)
}
func (m *GetRoleReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoleReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoleReq proto.InternalMessageInfo

func (m *GetRoleReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteRoleReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRoleReq) Reset()         { *m = DeleteRoleReq{} }
func (m *DeleteRoleReq) String() string { return proto.CompactTextString(m) }
func (*DeleteRoleReq) ProtoMessage()    {}
func (*DeleteRoleReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{14}
}
func (m *DeleteRoleReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRoleReq.Unmarshal(m, b)
}
func (m *DeleteRoleReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRoleReq.Marshal(b, m, deterministic)
}
func (dst *DeleteRoleReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRoleReq.Merge(dst, src)
}
func (m *DeleteRoleReq) XXX_Size() int {
	return xxx_messageInfo_DeleteRoleReq.Size(m)
}
func (m *DeleteRoleReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRoleReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRoleReq proto.InternalMessageInfo

func (m *DeleteRoleReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateRoleReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Actions              []string `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	Projects             []string `protobuf:"bytes,4,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRoleReq) Reset()         { *m = UpdateRoleReq{} }
func (m *UpdateRoleReq) String() string { return proto.CompactTextString(m) }
func (*UpdateRoleReq) ProtoMessage()    {}
func (*UpdateRoleReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{15}
}
func (m *UpdateRoleReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRoleReq.Unmarshal(m, b)
}
func (m *UpdateRoleReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRoleReq.Marshal(b, m, deterministic)
}
func (dst *UpdateRoleReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRoleReq.Merge(dst, src)
}
func (m *UpdateRoleReq) XXX_Size() int {
	return xxx_messageInfo_UpdateRoleReq.Size(m)
}
func (m *UpdateRoleReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRoleReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRoleReq proto.InternalMessageInfo

func (m *UpdateRoleReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRoleReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRoleReq) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *UpdateRoleReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type ListRolesReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRolesReq) Reset()         { *m = ListRolesReq{} }
func (m *ListRolesReq) String() string { return proto.CompactTextString(m) }
func (*ListRolesReq) ProtoMessage()    {}
func (*ListRolesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{16}
}
func (m *ListRolesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRolesReq.Unmarshal(m, b)
}
func (m *ListRolesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRolesReq.Marshal(b, m, deterministic)
}
func (dst *ListRolesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRolesReq.Merge(dst, src)
}
func (m *ListRolesReq) XXX_Size() int {
	return xxx_messageInfo_ListRolesReq.Size(m)
}
func (m *ListRolesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRolesReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListRolesReq proto.InternalMessageInfo

type GetProjectReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProjectReq) Reset()         { *m = GetProjectReq{} }
func (m *GetProjectReq) String() string { return proto.CompactTextString(m) }
func (*GetProjectReq) ProtoMessage()    {}
func (*GetProjectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{17}
}
func (m *GetProjectReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProjectReq.Unmarshal(m, b)
}
func (m *GetProjectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProjectReq.Marshal(b, m, deterministic)
}
func (dst *GetProjectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProjectReq.Merge(dst, src)
}
func (m *GetProjectReq) XXX_Size() int {
	return xxx_messageInfo_GetProjectReq.Size(m)
}
func (m *GetProjectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProjectReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetProjectReq proto.InternalMessageInfo

func (m *GetProjectReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListProjectsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListProjectsReq) Reset()         { *m = ListProjectsReq{} }
func (m *ListProjectsReq) String() string { return proto.CompactTextString(m) }
func (*ListProjectsReq) ProtoMessage()    {}
func (*ListProjectsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{18}
}
func (m *ListProjectsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProjectsReq.Unmarshal(m, b)
}
func (m *ListProjectsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProjectsReq.Marshal(b, m, deterministic)
}
func (dst *ListProjectsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProjectsReq.Merge(dst, src)
}
func (m *ListProjectsReq) XXX_Size() int {
	return xxx_messageInfo_ListProjectsReq.Size(m)
}
func (m *ListProjectsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProjectsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListProjectsReq proto.InternalMessageInfo

type CreateProjectReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProjectReq) Reset()         { *m = CreateProjectReq{} }
func (m *CreateProjectReq) String() string { return proto.CompactTextString(m) }
func (*CreateProjectReq) ProtoMessage()    {}
func (*CreateProjectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{19}
}
func (m *CreateProjectReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProjectReq.Unmarshal(m, b)
}
func (m *CreateProjectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProjectReq.Marshal(b, m, deterministic)
}
func (dst *CreateProjectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProjectReq.Merge(dst, src)
}
func (m *CreateProjectReq) XXX_Size() int {
	return xxx_messageInfo_CreateProjectReq.Size(m)
}
func (m *CreateProjectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProjectReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProjectReq proto.InternalMessageInfo

func (m *CreateProjectReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateProjectReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UpdateProjectReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateProjectReq) Reset()         { *m = UpdateProjectReq{} }
func (m *UpdateProjectReq) String() string { return proto.CompactTextString(m) }
func (*UpdateProjectReq) ProtoMessage()    {}
func (*UpdateProjectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{20}
}
func (m *UpdateProjectReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProjectReq.Unmarshal(m, b)
}
func (m *UpdateProjectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProjectReq.Marshal(b, m, deterministic)
}
func (dst *UpdateProjectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProjectReq.Merge(dst, src)
}
func (m *UpdateProjectReq) XXX_Size() int {
	return xxx_messageInfo_UpdateProjectReq.Size(m)
}
func (m *UpdateProjectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProjectReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProjectReq proto.InternalMessageInfo

func (m *UpdateProjectReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateProjectReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteProjectReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProjectReq) Reset()         { *m = DeleteProjectReq{} }
func (m *DeleteProjectReq) String() string { return proto.CompactTextString(m) }
func (*DeleteProjectReq) ProtoMessage()    {}
func (*DeleteProjectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_a15a519e01facaa9, []int{21}
}
func (m *DeleteProjectReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProjectReq.Unmarshal(m, b)
}
func (m *DeleteProjectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProjectReq.Marshal(b, m, deterministic)
}
func (dst *DeleteProjectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProjectReq.Merge(dst, src)
}
func (m *DeleteProjectReq) XXX_Size() int {
	return xxx_messageInfo_DeleteProjectReq.Size(m)
}
func (m *DeleteProjectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProjectReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProjectReq proto.InternalMessageInfo

func (m *DeleteProjectReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*CreatePolicyReq)(nil), "chef.automate.api.iam.v2beta.CreatePolicyReq")
	proto.RegisterType((*DeletePolicyReq)(nil), "chef.automate.api.iam.v2beta.DeletePolicyReq")
	proto.RegisterType((*ListPoliciesReq)(nil), "chef.automate.api.iam.v2beta.ListPoliciesReq")
	proto.RegisterType((*AddPolicyMembersReq)(nil), "chef.automate.api.iam.v2beta.AddPolicyMembersReq")
	proto.RegisterType((*GetPolicyReq)(nil), "chef.automate.api.iam.v2beta.GetPolicyReq")
	proto.RegisterType((*UpdatePolicyReq)(nil), "chef.automate.api.iam.v2beta.UpdatePolicyReq")
	proto.RegisterType((*UpgradeToV2Req)(nil), "chef.automate.api.iam.v2beta.UpgradeToV2Req")
	proto.RegisterType((*GetPolicyVersionReq)(nil), "chef.automate.api.iam.v2beta.GetPolicyVersionReq")
	proto.RegisterType((*ResetToV1Req)(nil), "chef.automate.api.iam.v2beta.ResetToV1Req")
	proto.RegisterType((*ListPolicyMembersReq)(nil), "chef.automate.api.iam.v2beta.ListPolicyMembersReq")
	proto.RegisterType((*ReplacePolicyMembersReq)(nil), "chef.automate.api.iam.v2beta.ReplacePolicyMembersReq")
	proto.RegisterType((*RemovePolicyMembersReq)(nil), "chef.automate.api.iam.v2beta.RemovePolicyMembersReq")
	proto.RegisterType((*CreateRoleReq)(nil), "chef.automate.api.iam.v2beta.CreateRoleReq")
	proto.RegisterType((*GetRoleReq)(nil), "chef.automate.api.iam.v2beta.GetRoleReq")
	proto.RegisterType((*DeleteRoleReq)(nil), "chef.automate.api.iam.v2beta.DeleteRoleReq")
	proto.RegisterType((*UpdateRoleReq)(nil), "chef.automate.api.iam.v2beta.UpdateRoleReq")
	proto.RegisterType((*ListRolesReq)(nil), "chef.automate.api.iam.v2beta.ListRolesReq")
	proto.RegisterType((*GetProjectReq)(nil), "chef.automate.api.iam.v2beta.GetProjectReq")
	proto.RegisterType((*ListProjectsReq)(nil), "chef.automate.api.iam.v2beta.ListProjectsReq")
	proto.RegisterType((*CreateProjectReq)(nil), "chef.automate.api.iam.v2beta.CreateProjectReq")
	proto.RegisterType((*UpdateProjectReq)(nil), "chef.automate.api.iam.v2beta.UpdateProjectReq")
	proto.RegisterType((*DeleteProjectReq)(nil), "chef.automate.api.iam.v2beta.DeleteProjectReq")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2beta/request/policy.proto", fileDescriptor_policy_a15a519e01facaa9)
}

var fileDescriptor_policy_a15a519e01facaa9 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x93, 0x00, 0xcd, 0x90, 0x38, 0xc5, 0xe5, 0xc3, 0xaa, 0x2a, 0x1a, 0xf6, 0x00, 0xb9,
	0x60, 0x8b, 0x20, 0xf5, 0x48, 0x45, 0x8b, 0x08, 0x12, 0x54, 0x42, 0x4b, 0x9b, 0x03, 0xb7, 0x8d,
	0x3d, 0x75, 0x17, 0x79, 0xbd, 0x8e, 0x77, 0x53, 0xd4, 0xff, 0xc5, 0x0f, 0x44, 0xbb, 0x76, 0xac,
	0x24, 0xaa, 0x2d, 0x85, 0x8a, 0x5b, 0xc6, 0x99, 0x37, 0xf3, 0xe6, 0xbd, 0x9d, 0x81, 0x93, 0x50,
	0x8a, 0x4c, 0xa6, 0x98, 0x6a, 0x15, 0xb0, 0x85, 0x96, 0x82, 0x69, 0x7c, 0x1b, 0x33, 0x8d, 0xbf,
	0xd9, 0x6d, 0xc0, 0x32, 0x1e, 0x70, 0x26, 0x82, 0x9b, 0xf1, 0x0c, 0x35, 0x0b, 0x72, 0x9c, 0x2f,
	0x50, 0xe9, 0x20, 0x93, 0x09, 0x0f, 0x6f, 0xfd, 0x2c, 0x97, 0x5a, 0xba, 0x87, 0xe1, 0x35, 0x5e,
	0xf9, 0x4b, 0xa8, 0xcf, 0x32, 0xee, 0x73, 0x26, 0xfc, 0x02, 0x72, 0xf0, 0x61, 0x8b, 0xf2, 0xa1,
	0x14, 0x42, 0xa6, 0x6b, 0xd5, 0xc9, 0x9f, 0x1d, 0x18, 0x9c, 0xe5, 0xc8, 0x34, 0x7e, 0xb7, 0x9f,
	0x29, 0xce, 0x5d, 0x07, 0x5a, 0x3c, 0xf2, 0x76, 0x86, 0x3b, 0xa3, 0x2e, 0x6d, 0xf1, 0xc8, 0x75,
	0xa1, 0x93, 0x32, 0x81, 0x5e, 0xcb, 0x7e, 0xb1, 0xbf, 0x5d, 0x0f, 0x1e, 0x09, 0x14, 0x33, 0xcc,
	0x95, 0xd7, 0x1e, 0xb6, 0x47, 0x5d, 0xba, 0x0c, 0xdd, 0x09, 0x80, 0xd2, 0x4c, 0xa3, 0x30, 0x9c,
	0xbc, 0xce, 0xb0, 0x3d, 0x7a, 0x3c, 0x7e, 0xe3, 0x37, 0x0d, 0xe1, 0xff, 0x58, 0xe6, 0xd3, 0x15,
	0xa8, 0x7b, 0x00, 0xbb, 0x59, 0x2e, 0x7f, 0x61, 0xa8, 0x95, 0xf7, 0xc0, 0xf6, 0xa8, 0x62, 0xf2,
	0x0a, 0x06, 0x9f, 0x30, 0xc1, 0x06, 0xd6, 0xe4, 0x09, 0x0c, 0xbe, 0x71, 0xa5, 0x6d, 0x02, 0x47,
	0x45, 0x71, 0x4e, 0x4e, 0x60, 0xff, 0x63, 0x14, 0x15, 0x90, 0xf3, 0x82, 0xee, 0x5d, 0xf3, 0xae,
	0xcc, 0xd6, 0x5a, 0x9b, 0x8d, 0xbc, 0x84, 0xde, 0x04, 0x75, 0x7d, 0x4f, 0xa3, 0xe6, 0x65, 0x16,
	0x35, 0xaa, 0x59, 0x5b, 0x7d, 0x43, 0xb9, 0xf6, 0xbf, 0x2b, 0xb7, 0x34, 0x6c, 0x77, 0xc5, 0xb0,
	0x55, 0x35, 0xbb, 0x1b, 0x6a, 0x7e, 0x01, 0xe7, 0x32, 0x8b, 0x73, 0x16, 0xe1, 0x85, 0x9c, 0x8e,
	0x0d, 0xe9, 0x63, 0xe8, 0x5c, 0x25, 0x2c, 0xb6, 0xb4, 0x9d, 0x31, 0x69, 0x26, 0xf1, 0x39, 0x61,
	0x31, 0xb5, 0xf9, 0xe4, 0x19, 0xec, 0x57, 0x02, 0x4d, 0x31, 0x57, 0x5c, 0xa6, 0x46, 0x78, 0x07,
	0x7a, 0x14, 0x15, 0xea, 0x0b, 0x39, 0x7d, 0x67, 0xe2, 0xd7, 0xf0, 0xb4, 0xf2, 0xa6, 0xc1, 0x09,
	0x72, 0x06, 0x2f, 0x28, 0x66, 0x09, 0x0b, 0xf1, 0x1e, 0xa6, 0x9d, 0xc2, 0x73, 0x8a, 0x42, 0xde,
	0xdc, 0xa7, 0x06, 0x87, 0x7e, 0xb1, 0x25, 0x54, 0x26, 0xb8, 0xc5, 0x8e, 0xb0, 0x50, 0x73, 0x99,
	0x56, 0x3b, 0x52, 0x86, 0x6b, 0x66, 0x74, 0x36, 0xcc, 0x38, 0x04, 0x98, 0xa0, 0xae, 0xe9, 0x43,
	0x8e, 0xa0, 0x5f, 0x3c, 0xfc, 0xba, 0x04, 0x0e, 0xfd, 0xe2, 0x05, 0xfe, 0x7f, 0xa6, 0x0e, 0xf4,
	0x8c, 0x8b, 0xa6, 0x91, 0x5d, 0xaf, 0x23, 0xe8, 0x1b, 0xf3, 0x8b, 0xbf, 0x9b, 0x56, 0xb2, 0x2c,
	0x60, 0x30, 0xc7, 0xb0, 0x57, 0x9e, 0x9f, 0x5a, 0xd8, 0x5d, 0x8c, 0x0d, 0xae, 0x5c, 0xb4, 0xed,
	0x70, 0x04, 0xf6, 0xca, 0xc3, 0x51, 0x8b, 0x3b, 0x3d, 0xff, 0xf9, 0x35, 0xe6, 0xfa, 0x7a, 0x31,
	0xf3, 0x43, 0x29, 0x02, 0xf3, 0xf4, 0xab, 0xd3, 0x1a, 0x6c, 0x7f, 0xcd, 0x67, 0x0f, 0xed, 0xa5,
	0x7d, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0xec, 0xda, 0xd7, 0xec, 0x0a, 0x06, 0x00, 0x00,
}
