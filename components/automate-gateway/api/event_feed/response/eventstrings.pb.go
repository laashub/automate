// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/event_feed/response/eventstrings.proto

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

type EventStrings struct {
	Strings              []*EventString `protobuf:"bytes,1,rep,name=strings,proto3" json:"strings,omitempty"`
	Start                string         `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End                  string         `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	HoursBetween         int32          `protobuf:"varint,4,opt,name=hours_between,json=hoursBetween,proto3" json:"hours_between,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *EventStrings) Reset()         { *m = EventStrings{} }
func (m *EventStrings) String() string { return proto.CompactTextString(m) }
func (*EventStrings) ProtoMessage()    {}
func (*EventStrings) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc1ddd57995ef1e1, []int{0}
}

func (m *EventStrings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventStrings.Unmarshal(m, b)
}
func (m *EventStrings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventStrings.Marshal(b, m, deterministic)
}
func (m *EventStrings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStrings.Merge(m, src)
}
func (m *EventStrings) XXX_Size() int {
	return xxx_messageInfo_EventStrings.Size(m)
}
func (m *EventStrings) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStrings.DiscardUnknown(m)
}

var xxx_messageInfo_EventStrings proto.InternalMessageInfo

func (m *EventStrings) GetStrings() []*EventString {
	if m != nil {
		return m.Strings
	}
	return nil
}

func (m *EventStrings) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *EventStrings) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *EventStrings) GetHoursBetween() int32 {
	if m != nil {
		return m.HoursBetween
	}
	return 0
}

type EventString struct {
	Collection           []*EventCollection `protobuf:"bytes,1,rep,name=collection,proto3" json:"collection,omitempty"`
	EventAction          string             `protobuf:"bytes,2,opt,name=event_action,json=eventAction,proto3" json:"event_action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *EventString) Reset()         { *m = EventString{} }
func (m *EventString) String() string { return proto.CompactTextString(m) }
func (*EventString) ProtoMessage()    {}
func (*EventString) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc1ddd57995ef1e1, []int{1}
}

func (m *EventString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventString.Unmarshal(m, b)
}
func (m *EventString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventString.Marshal(b, m, deterministic)
}
func (m *EventString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventString.Merge(m, src)
}
func (m *EventString) XXX_Size() int {
	return xxx_messageInfo_EventString.Size(m)
}
func (m *EventString) XXX_DiscardUnknown() {
	xxx_messageInfo_EventString.DiscardUnknown(m)
}

var xxx_messageInfo_EventString proto.InternalMessageInfo

func (m *EventString) GetCollection() []*EventCollection {
	if m != nil {
		return m.Collection
	}
	return nil
}

func (m *EventString) GetEventAction() string {
	if m != nil {
		return m.EventAction
	}
	return ""
}

type EventCollection struct {
	EventsCount          []*EventCount `protobuf:"bytes,1,rep,name=events_count,json=eventsCount,proto3" json:"events_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EventCollection) Reset()         { *m = EventCollection{} }
func (m *EventCollection) String() string { return proto.CompactTextString(m) }
func (*EventCollection) ProtoMessage()    {}
func (*EventCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc1ddd57995ef1e1, []int{2}
}

func (m *EventCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventCollection.Unmarshal(m, b)
}
func (m *EventCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventCollection.Marshal(b, m, deterministic)
}
func (m *EventCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCollection.Merge(m, src)
}
func (m *EventCollection) XXX_Size() int {
	return xxx_messageInfo_EventCollection.Size(m)
}
func (m *EventCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCollection.DiscardUnknown(m)
}

var xxx_messageInfo_EventCollection proto.InternalMessageInfo

func (m *EventCollection) GetEventsCount() []*EventCount {
	if m != nil {
		return m.EventsCount
	}
	return nil
}

func init() {
	proto.RegisterType((*EventStrings)(nil), "chef.automate.api.event_feed.response.EventStrings")
	proto.RegisterType((*EventString)(nil), "chef.automate.api.event_feed.response.EventString")
	proto.RegisterType((*EventCollection)(nil), "chef.automate.api.event_feed.response.EventCollection")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/event_feed/response/eventstrings.proto", fileDescriptor_cc1ddd57995ef1e1)
}

var fileDescriptor_cc1ddd57995ef1e1 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xe9, 0x7f, 0xff, 0x29, 0x66, 0x13, 0x25, 0x78, 0x28, 0x9e, 0x66, 0x45, 0xe8, 0xc5,
	0x04, 0x27, 0x78, 0x15, 0x27, 0x7a, 0x1a, 0x1e, 0xaa, 0x78, 0xf0, 0x52, 0xd2, 0xec, 0x5d, 0x5b,
	0x58, 0x93, 0xd2, 0xbc, 0x75, 0xf8, 0x0d, 0xfc, 0x2a, 0x7e, 0x4b, 0x69, 0xd2, 0xce, 0xe2, 0xc9,
	0xed, 0x96, 0x3c, 0xf0, 0xfc, 0x9e, 0x1f, 0xbc, 0xe4, 0x51, 0xea, 0xa2, 0xd4, 0x0a, 0x14, 0x1a,
	0x2e, 0x6a, 0xd4, 0x85, 0x40, 0xb8, 0x4c, 0x05, 0xc2, 0x5a, 0x7c, 0x70, 0x51, 0xe6, 0x1c, 0xde,
	0x41, 0x61, 0xbc, 0x04, 0x58, 0xf0, 0x0a, 0x4c, 0xa9, 0x95, 0x01, 0x97, 0x19, 0xac, 0x72, 0x95,
	0x1a, 0x56, 0x56, 0x1a, 0x35, 0xbd, 0x90, 0x19, 0x2c, 0x59, 0x47, 0x60, 0xa2, 0xcc, 0xd9, 0x4f,
	0x93, 0x75, 0xcd, 0xd3, 0xdb, 0x9d, 0xe7, 0xdc, 0x4e, 0xf0, 0xe5, 0x91, 0xf1, 0x43, 0xf3, 0x7f,
	0x76, 0xf3, 0x74, 0x4e, 0xf6, 0x5b, 0x13, 0xdf, 0x9b, 0x0c, 0xc2, 0xd1, 0x74, 0xca, 0xfe, 0xa4,
	0xc2, 0x7a, 0x94, 0xa8, 0x43, 0xd0, 0x13, 0x32, 0x34, 0x28, 0x2a, 0xf4, 0xff, 0x4d, 0xbc, 0xf0,
	0x20, 0x72, 0x1f, 0x7a, 0x4c, 0x06, 0xa0, 0x16, 0xfe, 0xc0, 0x66, 0xcd, 0x93, 0x9e, 0x93, 0xc3,
	0x4c, 0xd7, 0x95, 0x89, 0x13, 0xc0, 0x35, 0x80, 0xf2, 0xff, 0x4f, 0xbc, 0x70, 0x18, 0x8d, 0x6d,
	0x38, 0x73, 0x59, 0xf0, 0xe9, 0x91, 0x51, 0x6f, 0x85, 0xbe, 0x12, 0x22, 0xf5, 0x6a, 0x05, 0x12,
	0x73, 0xad, 0x5a, 0xdb, 0x9b, 0x6d, 0x6c, 0xef, 0x37, 0xed, 0xa8, 0x47, 0xa2, 0x67, 0x64, 0xec,
	0x2a, 0xc2, 0x91, 0x9d, 0xfb, 0xc8, 0x66, 0x77, 0x36, 0x0a, 0x52, 0x72, 0xf4, 0x8b, 0x40, 0x5f,
	0xda, 0x96, 0x89, 0xa5, 0xae, 0x15, 0xb6, 0x3e, 0x57, 0xdb, 0xf9, 0xd4, 0x0a, 0xdb, 0x21, 0x63,
	0x3f, 0xb3, 0xa7, 0xb7, 0x79, 0x9a, 0x63, 0x56, 0x27, 0x4c, 0xea, 0x82, 0x37, 0xac, 0xcd, 0x9d,
	0xf9, 0x0e, 0xb7, 0x4f, 0xf6, 0xec, 0xd9, 0xaf, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xe8,
	0xe4, 0x85, 0xa8, 0x02, 0x00, 0x00,
}
