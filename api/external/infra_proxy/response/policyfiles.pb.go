// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/infra_proxy/response/policyfiles.proto

package response

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

type Policyfiles struct {
	// List of policy files.
	Policies             []*PolicyfileListItem `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Policyfiles) Reset()         { *m = Policyfiles{} }
func (m *Policyfiles) String() string { return proto.CompactTextString(m) }
func (*Policyfiles) ProtoMessage()    {}
func (*Policyfiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{0}
}

func (m *Policyfiles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policyfiles.Unmarshal(m, b)
}
func (m *Policyfiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policyfiles.Marshal(b, m, deterministic)
}
func (m *Policyfiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policyfiles.Merge(m, src)
}
func (m *Policyfiles) XXX_Size() int {
	return xxx_messageInfo_Policyfiles.Size(m)
}
func (m *Policyfiles) XXX_DiscardUnknown() {
	xxx_messageInfo_Policyfiles.DiscardUnknown(m)
}

var xxx_messageInfo_Policyfiles proto.InternalMessageInfo

func (m *Policyfiles) GetPolicies() []*PolicyfileListItem {
	if m != nil {
		return m.Policies
	}
	return nil
}

type PolicyfileListItem struct {
	// Name of the policy file.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Revision ID of the policy file.
	RevisionId string `protobuf:"bytes,3,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	// Policy group of the policy file.
	PolicyGroup          string   `protobuf:"bytes,2,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyfileListItem) Reset()         { *m = PolicyfileListItem{} }
func (m *PolicyfileListItem) String() string { return proto.CompactTextString(m) }
func (*PolicyfileListItem) ProtoMessage()    {}
func (*PolicyfileListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{1}
}

func (m *PolicyfileListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyfileListItem.Unmarshal(m, b)
}
func (m *PolicyfileListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyfileListItem.Marshal(b, m, deterministic)
}
func (m *PolicyfileListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyfileListItem.Merge(m, src)
}
func (m *PolicyfileListItem) XXX_Size() int {
	return xxx_messageInfo_PolicyfileListItem.Size(m)
}
func (m *PolicyfileListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyfileListItem.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyfileListItem proto.InternalMessageInfo

func (m *PolicyfileListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PolicyfileListItem) GetRevisionId() string {
	if m != nil {
		return m.RevisionId
	}
	return ""
}

func (m *PolicyfileListItem) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

type Policyfile struct {
	// Name of the policy file.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Name of the policy group.
	PolicyGroup string `protobuf:"bytes,2,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty"`
	// Revision Id of the policy.
	RevisionId string `protobuf:"bytes,3,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	// Run list associated with the policy.
	RunList []string `protobuf:"bytes,4,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty"`
	// Named the run list associated with the policy.
	NamedRunList []*NamedRunList `protobuf:"bytes,5,rep,name=named_run_list,json=namedRunList,proto3" json:"named_run_list,omitempty"`
	// Included policy locks files.
	IncludedPolicyLocks []*IncludedPolicyLock `protobuf:"bytes,6,rep,name=included_policy_locks,json=includedPolicyLocks,proto3" json:"included_policy_locks,omitempty"`
	// List of cookbook locks under this policy.
	CookbookLocks []*CookbookLock `protobuf:"bytes,7,rep,name=cookbook_locks,json=cookbookLocks,proto3" json:"cookbook_locks,omitempty"`
	// Stringified json of the default attributes.
	DefaultAttributes string `protobuf:"bytes,8,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty"`
	// Stringified json of the override attributes.
	OverrideAttributes   string   `protobuf:"bytes,9,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policyfile) Reset()         { *m = Policyfile{} }
func (m *Policyfile) String() string { return proto.CompactTextString(m) }
func (*Policyfile) ProtoMessage()    {}
func (*Policyfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{2}
}

func (m *Policyfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policyfile.Unmarshal(m, b)
}
func (m *Policyfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policyfile.Marshal(b, m, deterministic)
}
func (m *Policyfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policyfile.Merge(m, src)
}
func (m *Policyfile) XXX_Size() int {
	return xxx_messageInfo_Policyfile.Size(m)
}
func (m *Policyfile) XXX_DiscardUnknown() {
	xxx_messageInfo_Policyfile.DiscardUnknown(m)
}

var xxx_messageInfo_Policyfile proto.InternalMessageInfo

func (m *Policyfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Policyfile) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

func (m *Policyfile) GetRevisionId() string {
	if m != nil {
		return m.RevisionId
	}
	return ""
}

func (m *Policyfile) GetRunList() []string {
	if m != nil {
		return m.RunList
	}
	return nil
}

func (m *Policyfile) GetNamedRunList() []*NamedRunList {
	if m != nil {
		return m.NamedRunList
	}
	return nil
}

func (m *Policyfile) GetIncludedPolicyLocks() []*IncludedPolicyLock {
	if m != nil {
		return m.IncludedPolicyLocks
	}
	return nil
}

func (m *Policyfile) GetCookbookLocks() []*CookbookLock {
	if m != nil {
		return m.CookbookLocks
	}
	return nil
}

func (m *Policyfile) GetDefaultAttributes() string {
	if m != nil {
		return m.DefaultAttributes
	}
	return ""
}

func (m *Policyfile) GetOverrideAttributes() string {
	if m != nil {
		return m.OverrideAttributes
	}
	return ""
}

type IncludedPolicyLock struct {
	// Name of the included policy file.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Revision id of the included policy file.
	RevisionId string `protobuf:"bytes,2,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	// Source options of the included policy file.
	SourceOptions        *SourceOptions `protobuf:"bytes,3,opt,name=source_options,json=sourceOptions,proto3" json:"source_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *IncludedPolicyLock) Reset()         { *m = IncludedPolicyLock{} }
func (m *IncludedPolicyLock) String() string { return proto.CompactTextString(m) }
func (*IncludedPolicyLock) ProtoMessage()    {}
func (*IncludedPolicyLock) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{3}
}

func (m *IncludedPolicyLock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncludedPolicyLock.Unmarshal(m, b)
}
func (m *IncludedPolicyLock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncludedPolicyLock.Marshal(b, m, deterministic)
}
func (m *IncludedPolicyLock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncludedPolicyLock.Merge(m, src)
}
func (m *IncludedPolicyLock) XXX_Size() int {
	return xxx_messageInfo_IncludedPolicyLock.Size(m)
}
func (m *IncludedPolicyLock) XXX_DiscardUnknown() {
	xxx_messageInfo_IncludedPolicyLock.DiscardUnknown(m)
}

var xxx_messageInfo_IncludedPolicyLock proto.InternalMessageInfo

func (m *IncludedPolicyLock) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IncludedPolicyLock) GetRevisionId() string {
	if m != nil {
		return m.RevisionId
	}
	return ""
}

func (m *IncludedPolicyLock) GetSourceOptions() *SourceOptions {
	if m != nil {
		return m.SourceOptions
	}
	return nil
}

type CookbookLock struct {
	// Name of the cookbook.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Version of the cookbook.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Identifier for the cookbook.
	Identifier string `protobuf:"bytes,3,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Decimal number identifier for the cookbook.
	DottedIdentifier string `protobuf:"bytes,4,opt,name=dotted_identifier,json=dottedIdentifier,proto3" json:"dotted_identifier,omitempty"`
	// Source of the cookbook.
	Source string `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
	// Cache key for the cookbook.
	CacheKey string `protobuf:"bytes,6,opt,name=cache_key,json=cacheKey,proto3" json:"cache_key,omitempty"`
	// SCM detail
	SCMDetail *SCMDetail `protobuf:"bytes,7,opt,name=SCMDetail,proto3" json:"SCMDetail,omitempty"`
	// Source path of the cookbook.
	SourceOptions        *SourceOptions `protobuf:"bytes,8,opt,name=source_options,json=sourceOptions,proto3" json:"source_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CookbookLock) Reset()         { *m = CookbookLock{} }
func (m *CookbookLock) String() string { return proto.CompactTextString(m) }
func (*CookbookLock) ProtoMessage()    {}
func (*CookbookLock) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{4}
}

func (m *CookbookLock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CookbookLock.Unmarshal(m, b)
}
func (m *CookbookLock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CookbookLock.Marshal(b, m, deterministic)
}
func (m *CookbookLock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CookbookLock.Merge(m, src)
}
func (m *CookbookLock) XXX_Size() int {
	return xxx_messageInfo_CookbookLock.Size(m)
}
func (m *CookbookLock) XXX_DiscardUnknown() {
	xxx_messageInfo_CookbookLock.DiscardUnknown(m)
}

var xxx_messageInfo_CookbookLock proto.InternalMessageInfo

func (m *CookbookLock) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CookbookLock) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CookbookLock) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *CookbookLock) GetDottedIdentifier() string {
	if m != nil {
		return m.DottedIdentifier
	}
	return ""
}

func (m *CookbookLock) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *CookbookLock) GetCacheKey() string {
	if m != nil {
		return m.CacheKey
	}
	return ""
}

func (m *CookbookLock) GetSCMDetail() *SCMDetail {
	if m != nil {
		return m.SCMDetail
	}
	return nil
}

func (m *CookbookLock) GetSourceOptions() *SourceOptions {
	if m != nil {
		return m.SourceOptions
	}
	return nil
}

type SCMDetail struct {
	// Name of the SCM.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Remote location of the SCM.
	Remote string `protobuf:"bytes,2,opt,name=remote,proto3" json:"remote,omitempty"`
	// Revision detail for the SCM.
	Revision string `protobuf:"bytes,3,opt,name=revision,proto3" json:"revision,omitempty"`
	// Boolean that denotes whether or not the working tree is cleaned.
	WorkingTreeClean bool `protobuf:"varint,4,opt,name=working_tree_clean,json=workingTreeClean,proto3" json:"working_tree_clean,omitempty"`
	// Published info of the source.
	Published bool `protobuf:"varint,5,opt,name=published,proto3" json:"published,omitempty"`
	// List of the synchronized remote branches.
	SynchronizedRemoteBranches []string `protobuf:"bytes,6,rep,name=synchronized_remote_branches,json=synchronizedRemoteBranches,proto3" json:"synchronized_remote_branches,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *SCMDetail) Reset()         { *m = SCMDetail{} }
func (m *SCMDetail) String() string { return proto.CompactTextString(m) }
func (*SCMDetail) ProtoMessage()    {}
func (*SCMDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{5}
}

func (m *SCMDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SCMDetail.Unmarshal(m, b)
}
func (m *SCMDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SCMDetail.Marshal(b, m, deterministic)
}
func (m *SCMDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SCMDetail.Merge(m, src)
}
func (m *SCMDetail) XXX_Size() int {
	return xxx_messageInfo_SCMDetail.Size(m)
}
func (m *SCMDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_SCMDetail.DiscardUnknown(m)
}

var xxx_messageInfo_SCMDetail proto.InternalMessageInfo

func (m *SCMDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SCMDetail) GetRemote() string {
	if m != nil {
		return m.Remote
	}
	return ""
}

func (m *SCMDetail) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *SCMDetail) GetWorkingTreeClean() bool {
	if m != nil {
		return m.WorkingTreeClean
	}
	return false
}

func (m *SCMDetail) GetPublished() bool {
	if m != nil {
		return m.Published
	}
	return false
}

func (m *SCMDetail) GetSynchronizedRemoteBranches() []string {
	if m != nil {
		return m.SynchronizedRemoteBranches
	}
	return nil
}

type SourceOptions struct {
	// Path of the source options.
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourceOptions) Reset()         { *m = SourceOptions{} }
func (m *SourceOptions) String() string { return proto.CompactTextString(m) }
func (*SourceOptions) ProtoMessage()    {}
func (*SourceOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{6}
}

func (m *SourceOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceOptions.Unmarshal(m, b)
}
func (m *SourceOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceOptions.Marshal(b, m, deterministic)
}
func (m *SourceOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceOptions.Merge(m, src)
}
func (m *SourceOptions) XXX_Size() int {
	return xxx_messageInfo_SourceOptions.Size(m)
}
func (m *SourceOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceOptions.DiscardUnknown(m)
}

var xxx_messageInfo_SourceOptions proto.InternalMessageInfo

func (m *SourceOptions) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type NamedRunList struct {
	// Name of the run list.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Run list associated with the policy.
	RunList              []string `protobuf:"bytes,2,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamedRunList) Reset()         { *m = NamedRunList{} }
func (m *NamedRunList) String() string { return proto.CompactTextString(m) }
func (*NamedRunList) ProtoMessage()    {}
func (*NamedRunList) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{7}
}

func (m *NamedRunList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamedRunList.Unmarshal(m, b)
}
func (m *NamedRunList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamedRunList.Marshal(b, m, deterministic)
}
func (m *NamedRunList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamedRunList.Merge(m, src)
}
func (m *NamedRunList) XXX_Size() int {
	return xxx_messageInfo_NamedRunList.Size(m)
}
func (m *NamedRunList) XXX_DiscardUnknown() {
	xxx_messageInfo_NamedRunList.DiscardUnknown(m)
}

var xxx_messageInfo_NamedRunList proto.InternalMessageInfo

func (m *NamedRunList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamedRunList) GetRunList() []string {
	if m != nil {
		return m.RunList
	}
	return nil
}

func init() {
	proto.RegisterType((*Policyfiles)(nil), "chef.automate.api.infra_proxy.response.Policyfiles")
	proto.RegisterType((*PolicyfileListItem)(nil), "chef.automate.api.infra_proxy.response.PolicyfileListItem")
	proto.RegisterType((*Policyfile)(nil), "chef.automate.api.infra_proxy.response.Policyfile")
	proto.RegisterType((*IncludedPolicyLock)(nil), "chef.automate.api.infra_proxy.response.IncludedPolicyLock")
	proto.RegisterType((*CookbookLock)(nil), "chef.automate.api.infra_proxy.response.CookbookLock")
	proto.RegisterType((*SCMDetail)(nil), "chef.automate.api.infra_proxy.response.SCMDetail")
	proto.RegisterType((*SourceOptions)(nil), "chef.automate.api.infra_proxy.response.SourceOptions")
	proto.RegisterType((*NamedRunList)(nil), "chef.automate.api.infra_proxy.response.NamedRunList")
}

func init() {
	proto.RegisterFile("api/external/infra_proxy/response/policyfiles.proto", fileDescriptor_64fca7030dbc5441)
}

var fileDescriptor_64fca7030dbc5441 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x56, 0xd2, 0x34, 0x71, 0x26, 0x69, 0xd5, 0x6e, 0x75, 0x2a, 0x9f, 0x9e, 0xea, 0x50, 0x82,
	0x84, 0x2a, 0x01, 0xb6, 0x68, 0xe1, 0xa6, 0x02, 0x09, 0x5a, 0x24, 0x14, 0x51, 0x28, 0x72, 0x11,
	0x17, 0x05, 0xc9, 0x72, 0xd6, 0x93, 0x66, 0x15, 0x67, 0xd7, 0xda, 0x5d, 0x97, 0x86, 0x27, 0xe0,
	0x4d, 0x78, 0x2e, 0x1e, 0x80, 0x77, 0x40, 0x5e, 0xdb, 0xb1, 0xab, 0x44, 0x6a, 0x2a, 0x71, 0xe7,
	0x99, 0xf9, 0xbe, 0xf9, 0xf9, 0x76, 0x3d, 0x0b, 0x87, 0x41, 0xcc, 0x5c, 0xbc, 0xd6, 0x28, 0x79,
	0x10, 0xb9, 0x8c, 0x0f, 0x65, 0xe0, 0xc7, 0x52, 0x5c, 0x4f, 0x5d, 0x89, 0x2a, 0x16, 0x5c, 0xa1,
	0x1b, 0x8b, 0x88, 0xd1, 0xe9, 0x90, 0x45, 0xa8, 0x9c, 0x58, 0x0a, 0x2d, 0xc8, 0x43, 0x3a, 0xc2,
	0xa1, 0x13, 0x24, 0x5a, 0x4c, 0x02, 0x8d, 0x4e, 0x10, 0x33, 0xa7, 0xc2, 0x74, 0x0a, 0x66, 0x0f,
	0xa1, 0xf3, 0xb1, 0x24, 0x93, 0xcf, 0x60, 0x99, 0x5c, 0x0c, 0x95, 0x5d, 0xdb, 0x5b, 0xd9, 0xef,
	0x1c, 0x1c, 0x39, 0xcb, 0x65, 0x72, 0xca, 0x34, 0xa7, 0x4c, 0xe9, 0xbe, 0xc6, 0x89, 0x37, 0xcb,
	0xd5, 0x8b, 0x80, 0xcc, 0xc7, 0x09, 0x81, 0x06, 0x0f, 0x26, 0x68, 0xd7, 0xf6, 0x6a, 0xfb, 0x6d,
	0xcf, 0x7c, 0x93, 0x7b, 0xd0, 0x91, 0x78, 0xc5, 0x14, 0x13, 0xdc, 0x67, 0xa1, 0xbd, 0x62, 0x42,
	0x50, 0xb8, 0xfa, 0x21, 0xb9, 0x0f, 0xdd, 0x6c, 0x5c, 0xff, 0x52, 0x8a, 0x24, 0xb6, 0xeb, 0x06,
	0xd1, 0xc9, 0x7c, 0x6f, 0x53, 0x57, 0xef, 0x47, 0x03, 0xa0, 0x2c, 0xb7, 0xb0, 0xcc, 0xed, 0x59,
	0x6e, 0xef, 0xe4, 0x5f, 0xb0, 0x64, 0xc2, 0xfd, 0x88, 0x29, 0x6d, 0x37, 0xf6, 0x56, 0xf6, 0xdb,
	0x5e, 0x4b, 0x26, 0x3c, 0x9d, 0x8e, 0x5c, 0xc0, 0x7a, 0x5a, 0x26, 0xf4, 0x67, 0x80, 0x55, 0xa3,
	0xe6, 0xb3, 0x65, 0xd5, 0xfc, 0x90, 0xb2, 0xbd, 0x2c, 0x9b, 0xd7, 0xe5, 0x15, 0x8b, 0x70, 0xf8,
	0x87, 0x71, 0x1a, 0x25, 0x21, 0x86, 0x7e, 0x3e, 0x43, 0x24, 0xe8, 0x58, 0xd9, 0xcd, 0xbb, 0x1d,
	0x58, 0x3f, 0x4f, 0x92, 0x29, 0x75, 0x2a, 0xe8, 0xd8, 0xdb, 0x62, 0x73, 0x3e, 0x45, 0xbe, 0xc0,
	0x3a, 0x15, 0x62, 0x3c, 0x10, 0x62, 0x9c, 0x17, 0x6a, 0xdd, 0x6d, 0x96, 0x93, 0x9c, 0x6d, 0x4a,
	0xac, 0xd1, 0x8a, 0xa5, 0xc8, 0x13, 0x20, 0x21, 0x0e, 0x83, 0x24, 0xd2, 0x7e, 0xa0, 0xb5, 0x64,
	0x83, 0x44, 0xa3, 0xb2, 0x2d, 0xa3, 0xf5, 0x66, 0x1e, 0x79, 0x3d, 0x0b, 0x10, 0x17, 0xb6, 0xc4,
	0x15, 0x4a, 0xc9, 0x42, 0xac, 0xe2, 0xdb, 0x06, 0x4f, 0x8a, 0x50, 0x49, 0xe8, 0xfd, 0xac, 0x01,
	0x99, 0x1f, 0x74, 0x99, 0x9b, 0x57, 0x9f, 0x3b, 0xef, 0xaf, 0xb0, 0xae, 0x44, 0x22, 0x29, 0xfa,
	0x22, 0xd6, 0x4c, 0x70, 0x65, 0xee, 0x44, 0xe7, 0xe0, 0xf9, 0xb2, 0x42, 0x9c, 0x1b, 0xf6, 0x59,
	0x46, 0xf6, 0xd6, 0x54, 0xd5, 0xec, 0xfd, 0xae, 0x43, 0xb7, 0xaa, 0xd4, 0xc2, 0x1e, 0x6d, 0x68,
	0x5d, 0xa1, 0x4c, 0xfb, 0xc9, 0xfb, 0x2b, 0x4c, 0xf2, 0x3f, 0x00, 0x0b, 0x91, 0x6b, 0x36, 0x64,
	0x28, 0x8b, 0xcb, 0x5a, 0x7a, 0xc8, 0x23, 0xd8, 0x0c, 0x85, 0xd6, 0x18, 0xfa, 0x15, 0x58, 0xc3,
	0xc0, 0x36, 0xb2, 0x40, 0xbf, 0x04, 0x6f, 0x43, 0x33, 0x6b, 0xce, 0x5e, 0x35, 0x88, 0xdc, 0x22,
	0xff, 0x41, 0x9b, 0x06, 0x74, 0x84, 0xfe, 0x18, 0xa7, 0x76, 0xd3, 0x84, 0x2c, 0xe3, 0x78, 0x87,
	0x53, 0x72, 0x06, 0xed, 0xf3, 0x93, 0xf7, 0x6f, 0x50, 0x07, 0x2c, 0xb2, 0x5b, 0x46, 0x99, 0xa7,
	0x4b, 0x2b, 0x53, 0x10, 0xbd, 0x32, 0xc7, 0x02, 0xbd, 0xad, 0xbf, 0xa8, 0xf7, 0xaf, 0x5a, 0xa5,
	0xdf, 0x85, 0x62, 0x6f, 0x43, 0x53, 0xe2, 0x44, 0x68, 0xcc, 0xb5, 0xce, 0x2d, 0xb2, 0x03, 0x56,
	0x71, 0x2b, 0x72, 0xa1, 0x67, 0x36, 0x79, 0x0c, 0xe4, 0x9b, 0x90, 0x63, 0xc6, 0x2f, 0x7d, 0x2d,
	0x11, 0x7d, 0x1a, 0x61, 0xc0, 0x8d, 0xce, 0x96, 0xb7, 0x91, 0x47, 0x3e, 0x49, 0xc4, 0x93, 0xd4,
	0x4f, 0x76, 0xa1, 0x1d, 0x27, 0x83, 0x88, 0xa9, 0x11, 0x86, 0x46, 0x6a, 0xcb, 0x2b, 0x1d, 0xe4,
	0x15, 0xec, 0xaa, 0x29, 0xa7, 0x23, 0x29, 0x38, 0xfb, 0x9e, 0xee, 0x12, 0x53, 0xde, 0x1f, 0xc8,
	0x80, 0xd3, 0x11, 0x66, 0xff, 0x7b, 0xdb, 0xdb, 0xa9, 0x62, 0x3c, 0x03, 0x39, 0xce, 0x11, 0xbd,
	0x07, 0xb0, 0x76, 0x43, 0x83, 0x74, 0xcc, 0x38, 0xd0, 0xa3, 0x62, 0xcc, 0xf4, 0xbb, 0xf7, 0x12,
	0xba, 0xd5, 0x6d, 0xb3, 0x50, 0x8a, 0xea, 0xaa, 0xab, 0xdf, 0x58, 0x75, 0xc7, 0x2f, 0x2e, 0x8e,
	0x2e, 0x99, 0x1e, 0x25, 0x03, 0x87, 0x8a, 0x89, 0x9b, 0x9e, 0x8c, 0x5b, 0x9c, 0x8c, 0x7b, 0xeb,
	0xcb, 0x35, 0x68, 0x9a, 0xe7, 0xea, 0xf0, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x9f, 0xca,
	0x4d, 0xe5, 0x06, 0x00, 0x00,
}
