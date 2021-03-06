// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oneof.proto

package proto

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Event struct {
	// Types that are valid to be assigned to Event:
	//	*Event_EventA
	//	*Event_EventB
	//	*Event_EventC
	Event                isEvent_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_df639ee9042fde12, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

type isEvent_Event interface {
	isEvent_Event()
}

type Event_EventA struct {
	EventA *EventA `protobuf:"bytes,1,opt,name=event_a,json=eventA,proto3,oneof"`
}

type Event_EventB struct {
	EventB *EventB `protobuf:"bytes,2,opt,name=event_b,json=eventB,proto3,oneof"`
}

type Event_EventC struct {
	EventC *EventC `protobuf:"bytes,3,opt,name=event_c,json=eventC,proto3,oneof"`
}

func (*Event_EventA) isEvent_Event() {}

func (*Event_EventB) isEvent_Event() {}

func (*Event_EventC) isEvent_Event() {}

func (m *Event) GetEvent() isEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Event) GetEventA() *EventA {
	if x, ok := m.GetEvent().(*Event_EventA); ok {
		return x.EventA
	}
	return nil
}

func (m *Event) GetEventB() *EventB {
	if x, ok := m.GetEvent().(*Event_EventB); ok {
		return x.EventB
	}
	return nil
}

func (m *Event) GetEventC() *EventC {
	if x, ok := m.GetEvent().(*Event_EventC); ok {
		return x.EventC
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_EventA)(nil),
		(*Event_EventB)(nil),
		(*Event_EventC)(nil),
	}
}

type EventA struct {
	A                    string   `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventA) Reset()         { *m = EventA{} }
func (m *EventA) String() string { return proto.CompactTextString(m) }
func (*EventA) ProtoMessage()    {}
func (*EventA) Descriptor() ([]byte, []int) {
	return fileDescriptor_df639ee9042fde12, []int{1}
}

func (m *EventA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventA.Unmarshal(m, b)
}
func (m *EventA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventA.Marshal(b, m, deterministic)
}
func (m *EventA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventA.Merge(m, src)
}
func (m *EventA) XXX_Size() int {
	return xxx_messageInfo_EventA.Size(m)
}
func (m *EventA) XXX_DiscardUnknown() {
	xxx_messageInfo_EventA.DiscardUnknown(m)
}

var xxx_messageInfo_EventA proto.InternalMessageInfo

func (m *EventA) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

type EventB struct {
	A                    string   `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	B                    int64    `protobuf:"varint,2,opt,name=B,proto3" json:"B,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventB) Reset()         { *m = EventB{} }
func (m *EventB) String() string { return proto.CompactTextString(m) }
func (*EventB) ProtoMessage()    {}
func (*EventB) Descriptor() ([]byte, []int) {
	return fileDescriptor_df639ee9042fde12, []int{2}
}

func (m *EventB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventB.Unmarshal(m, b)
}
func (m *EventB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventB.Marshal(b, m, deterministic)
}
func (m *EventB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventB.Merge(m, src)
}
func (m *EventB) XXX_Size() int {
	return xxx_messageInfo_EventB.Size(m)
}
func (m *EventB) XXX_DiscardUnknown() {
	xxx_messageInfo_EventB.DiscardUnknown(m)
}

var xxx_messageInfo_EventB proto.InternalMessageInfo

func (m *EventB) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *EventB) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

type EventC struct {
	A                    string          `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	B                    int64           `protobuf:"varint,2,opt,name=B,proto3" json:"B,omitempty"`
	C                    *EventC_Context `protobuf:"bytes,3,opt,name=C,proto3" json:"C,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EventC) Reset()         { *m = EventC{} }
func (m *EventC) String() string { return proto.CompactTextString(m) }
func (*EventC) ProtoMessage()    {}
func (*EventC) Descriptor() ([]byte, []int) {
	return fileDescriptor_df639ee9042fde12, []int{3}
}

func (m *EventC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventC.Unmarshal(m, b)
}
func (m *EventC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventC.Marshal(b, m, deterministic)
}
func (m *EventC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventC.Merge(m, src)
}
func (m *EventC) XXX_Size() int {
	return xxx_messageInfo_EventC.Size(m)
}
func (m *EventC) XXX_DiscardUnknown() {
	xxx_messageInfo_EventC.DiscardUnknown(m)
}

var xxx_messageInfo_EventC proto.InternalMessageInfo

func (m *EventC) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *EventC) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

func (m *EventC) GetC() *EventC_Context {
	if m != nil {
		return m.C
	}
	return nil
}

type EventC_Context struct {
	X                    int64    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int64    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventC_Context) Reset()         { *m = EventC_Context{} }
func (m *EventC_Context) String() string { return proto.CompactTextString(m) }
func (*EventC_Context) ProtoMessage()    {}
func (*EventC_Context) Descriptor() ([]byte, []int) {
	return fileDescriptor_df639ee9042fde12, []int{3, 0}
}

func (m *EventC_Context) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventC_Context.Unmarshal(m, b)
}
func (m *EventC_Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventC_Context.Marshal(b, m, deterministic)
}
func (m *EventC_Context) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventC_Context.Merge(m, src)
}
func (m *EventC_Context) XXX_Size() int {
	return xxx_messageInfo_EventC_Context.Size(m)
}
func (m *EventC_Context) XXX_DiscardUnknown() {
	xxx_messageInfo_EventC_Context.DiscardUnknown(m)
}

var xxx_messageInfo_EventC_Context proto.InternalMessageInfo

func (m *EventC_Context) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *EventC_Context) GetY() int64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func init() {
	proto.RegisterType((*Event)(nil), "com.example.oneof.Event")
	proto.RegisterType((*EventA)(nil), "com.example.oneof.EventA")
	proto.RegisterType((*EventB)(nil), "com.example.oneof.EventB")
	proto.RegisterType((*EventC)(nil), "com.example.oneof.EventC")
	proto.RegisterType((*EventC_Context)(nil), "com.example.oneof.EventC.Context")
}

func init() { proto.RegisterFile("oneof.proto", fileDescriptor_df639ee9042fde12) }

var fileDescriptor_df639ee9042fde12 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcf, 0x4b, 0xcd,
	0x4f, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4c, 0xce, 0xcf, 0xd5, 0x4b, 0xad, 0x48,
	0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x03, 0x4b, 0x48, 0x89, 0x97, 0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96,
	0xa4, 0xea, 0xc3, 0x18, 0x10, 0xb5, 0x4a, 0xdb, 0x19, 0xb9, 0x58, 0x5d, 0xcb, 0x52, 0xf3, 0x4a,
	0x84, 0x4c, 0xb8, 0xd8, 0x53, 0x41, 0x8c, 0xf8, 0x44, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23,
	0x49, 0x3d, 0x0c, 0x73, 0xf4, 0xc0, 0x4a, 0x1d, 0x3d, 0x18, 0x82, 0xd8, 0xc0, 0x6a, 0x1d, 0x11,
	0xba, 0x92, 0x24, 0x98, 0xf0, 0xeb, 0x72, 0x82, 0xeb, 0x72, 0x42, 0xe8, 0x4a, 0x96, 0x60, 0xc6,
	0xaf, 0xcb, 0x19, 0xae, 0xcb, 0xd9, 0x89, 0x87, 0x8b, 0x15, 0xcc, 0x12, 0x62, 0xfe, 0xe1, 0xc4,
	0xa8, 0x24, 0xc6, 0xc5, 0x06, 0x71, 0x8d, 0x10, 0x0f, 0x17, 0xa3, 0x23, 0xd8, 0xcd, 0x9c, 0x41,
	0x8c, 0x8e, 0x4a, 0x2a, 0x50, 0x71, 0x27, 0x54, 0x71, 0x10, 0xcf, 0x09, 0xec, 0x46, 0xe6, 0x20,
	0x46, 0x27, 0xa5, 0x1a, 0xa8, 0x2a, 0x67, 0x7c, 0xaa, 0x84, 0xf4, 0xb9, 0x18, 0x9d, 0xa1, 0x2e,
	0x54, 0xc4, 0xe9, 0x42, 0x3d, 0xe7, 0xfc, 0xbc, 0x92, 0xd4, 0x8a, 0x92, 0x20, 0x46, 0x67, 0x29,
	0x55, 0x2e, 0x76, 0x28, 0x0f, 0x64, 0x52, 0x05, 0xd8, 0x5c, 0xe6, 0x20, 0xc6, 0x0a, 0x10, 0xaf,
	0x12, 0x66, 0x6e, 0xa5, 0x13, 0x7b, 0x14, 0x2b, 0x38, 0xf8, 0x93, 0xd8, 0xc0, 0x94, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0xfc, 0x50, 0xcf, 0xc0, 0x01, 0x00, 0x00,
}
