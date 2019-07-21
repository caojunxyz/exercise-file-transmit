// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deliver.proto

package deliver

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NodeType int32

const (
	NodeType_MASTER NodeType = 0
	NodeType_LEADER NodeType = 1
	NodeType_NORMAL NodeType = 2
)

var NodeType_name = map[int32]string{
	0: "MASTER",
	1: "LEADER",
	2: "NORMAL",
}

var NodeType_value = map[string]int32{
	"MASTER": 0,
	"LEADER": 1,
	"NORMAL": 2,
}

func (x NodeType) String() string {
	return proto.EnumName(NodeType_name, int32(x))
}

func (NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3ebde617daa8954a, []int{0}
}

type Null struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Null) Reset()         { *m = Null{} }
func (m *Null) String() string { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()    {}
func (*Null) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ebde617daa8954a, []int{0}
}

func (m *Null) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Null.Unmarshal(m, b)
}
func (m *Null) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Null.Marshal(b, m, deterministic)
}
func (m *Null) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Null.Merge(m, src)
}
func (m *Null) XXX_Size() int {
	return xxx_messageInfo_Null.Size(m)
}
func (m *Null) XXX_DiscardUnknown() {
	xxx_messageInfo_Null.DiscardUnknown(m)
}

var xxx_messageInfo_Null proto.InternalMessageInfo

type Session struct {
	SessionId            int64     `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Artifact             *Artifact `protobuf:"bytes,2,opt,name=artifact,proto3" json:"artifact,omitempty"`
	NodeId               int64     `protobuf:"varint,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Group                *Group    `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	Leaders              []*Node   `protobuf:"bytes,5,rep,name=leaders,proto3" json:"leaders,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ebde617daa8954a, []int{1}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *Session) GetArtifact() *Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

func (m *Session) GetNodeId() int64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *Session) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *Session) GetLeaders() []*Node {
	if m != nil {
		return m.Leaders
	}
	return nil
}

type Node struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 NodeType `protobuf:"varint,2,opt,name=type,proto3,enum=deliver.NodeType" json:"type,omitempty"`
	Addr                 string   `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ebde617daa8954a, []int{2}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Node) GetType() NodeType {
	if m != nil {
		return m.Type
	}
	return NodeType_MASTER
}

func (m *Node) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type Group struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LeaderId             int64    `protobuf:"varint,2,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	Nodes                []*Node  `protobuf:"bytes,3,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ebde617daa8954a, []int{3}
}

func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Group) GetLeaderId() int64 {
	if m != nil {
		return m.LeaderId
	}
	return 0
}

func (m *Group) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Artifact struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Sha1                 string   `protobuf:"bytes,3,opt,name=sha1,proto3" json:"sha1,omitempty"`
	BlockSize            int64    `protobuf:"varint,4,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	BlockNum             int64    `protobuf:"varint,5,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ebde617daa8954a, []int{4}
}

func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (m *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(m, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Artifact) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Artifact) GetSha1() string {
	if m != nil {
		return m.Sha1
	}
	return ""
}

func (m *Artifact) GetBlockSize() int64 {
	if m != nil {
		return m.BlockSize
	}
	return 0
}

func (m *Artifact) GetBlockNum() int64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

type Block struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Sha1                 string   `protobuf:"bytes,3,opt,name=sha1,proto3" json:"sha1,omitempty"`
	Raw                  []byte   `protobuf:"bytes,4,opt,name=raw,proto3" json:"raw,omitempty"`
	PulledFrom           []int64  `protobuf:"varint,5,rep,packed,name=pulled_from,json=pulledFrom,proto3" json:"pulled_from,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ebde617daa8954a, []int{5}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Block) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Block) GetSha1() string {
	if m != nil {
		return m.Sha1
	}
	return ""
}

func (m *Block) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *Block) GetPulledFrom() []int64 {
	if m != nil {
		return m.PulledFrom
	}
	return nil
}

type BlockFlag struct {
	BlockId              int64    `protobuf:"varint,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Node                 *Node    `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockFlag) Reset()         { *m = BlockFlag{} }
func (m *BlockFlag) String() string { return proto.CompactTextString(m) }
func (*BlockFlag) ProtoMessage()    {}
func (*BlockFlag) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ebde617daa8954a, []int{6}
}

func (m *BlockFlag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockFlag.Unmarshal(m, b)
}
func (m *BlockFlag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockFlag.Marshal(b, m, deterministic)
}
func (m *BlockFlag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockFlag.Merge(m, src)
}
func (m *BlockFlag) XXX_Size() int {
	return xxx_messageInfo_BlockFlag.Size(m)
}
func (m *BlockFlag) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockFlag.DiscardUnknown(m)
}

var xxx_messageInfo_BlockFlag proto.InternalMessageInfo

func (m *BlockFlag) GetBlockId() int64 {
	if m != nil {
		return m.BlockId
	}
	return 0
}

func (m *BlockFlag) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type BlockAssembled struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PulledFrom           []int64  `protobuf:"varint,2,rep,packed,name=pulled_from,json=pulledFrom,proto3" json:"pulled_from,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockAssembled) Reset()         { *m = BlockAssembled{} }
func (m *BlockAssembled) String() string { return proto.CompactTextString(m) }
func (*BlockAssembled) ProtoMessage()    {}
func (*BlockAssembled) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ebde617daa8954a, []int{7}
}

func (m *BlockAssembled) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockAssembled.Unmarshal(m, b)
}
func (m *BlockAssembled) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockAssembled.Marshal(b, m, deterministic)
}
func (m *BlockAssembled) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockAssembled.Merge(m, src)
}
func (m *BlockAssembled) XXX_Size() int {
	return xxx_messageInfo_BlockAssembled.Size(m)
}
func (m *BlockAssembled) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockAssembled.DiscardUnknown(m)
}

var xxx_messageInfo_BlockAssembled proto.InternalMessageInfo

func (m *BlockAssembled) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BlockAssembled) GetPulledFrom() []int64 {
	if m != nil {
		return m.PulledFrom
	}
	return nil
}

type Progress struct {
	SessionId            int64           `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	AssembledBlocks      map[int64]int64 `protobuf:"bytes,2,rep,name=assembled_blocks,json=assembledBlocks,proto3" json:"assembled_blocks,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Progress) Reset()         { *m = Progress{} }
func (m *Progress) String() string { return proto.CompactTextString(m) }
func (*Progress) ProtoMessage()    {}
func (*Progress) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ebde617daa8954a, []int{8}
}

func (m *Progress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Progress.Unmarshal(m, b)
}
func (m *Progress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Progress.Marshal(b, m, deterministic)
}
func (m *Progress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Progress.Merge(m, src)
}
func (m *Progress) XXX_Size() int {
	return xxx_messageInfo_Progress.Size(m)
}
func (m *Progress) XXX_DiscardUnknown() {
	xxx_messageInfo_Progress.DiscardUnknown(m)
}

var xxx_messageInfo_Progress proto.InternalMessageInfo

func (m *Progress) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *Progress) GetAssembledBlocks() map[int64]int64 {
	if m != nil {
		return m.AssembledBlocks
	}
	return nil
}

func init() {
	proto.RegisterEnum("deliver.NodeType", NodeType_name, NodeType_value)
	proto.RegisterType((*Null)(nil), "deliver.Null")
	proto.RegisterType((*Session)(nil), "deliver.Session")
	proto.RegisterType((*Node)(nil), "deliver.Node")
	proto.RegisterType((*Group)(nil), "deliver.Group")
	proto.RegisterType((*Artifact)(nil), "deliver.Artifact")
	proto.RegisterType((*Block)(nil), "deliver.Block")
	proto.RegisterType((*BlockFlag)(nil), "deliver.BlockFlag")
	proto.RegisterType((*BlockAssembled)(nil), "deliver.BlockAssembled")
	proto.RegisterType((*Progress)(nil), "deliver.Progress")
	proto.RegisterMapType((map[int64]int64)(nil), "deliver.Progress.AssembledBlocksEntry")
}

func init() { proto.RegisterFile("deliver.proto", fileDescriptor_3ebde617daa8954a) }

var fileDescriptor_3ebde617daa8954a = []byte{
	// 623 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0x8e, 0x1d, 0x3b, 0x87, 0xc9, 0xdf, 0x34, 0x5d, 0x55, 0xfa, 0x43, 0x2a, 0x44, 0x31, 0xa7,
	0x0a, 0x41, 0x44, 0x83, 0x04, 0x88, 0xbb, 0x94, 0xa6, 0xc8, 0x52, 0x1b, 0xda, 0x6d, 0x6f, 0xb8,
	0x8a, 0xdc, 0xee, 0xa6, 0xb5, 0xba, 0xb6, 0xa3, 0x5d, 0xbb, 0x28, 0x3c, 0x02, 0x12, 0xcf, 0x04,
	0x8f, 0x86, 0x76, 0xd6, 0x76, 0x0f, 0xae, 0x04, 0x77, 0x33, 0xdf, 0xcc, 0x7e, 0xf3, 0xcd, 0x41,
	0x0b, 0x2b, 0x8c, 0x8b, 0xf0, 0x8a, 0xcb, 0xe1, 0x42, 0x26, 0x69, 0x42, 0x9a, 0xb9, 0xeb, 0x35,
	0xc0, 0x99, 0x66, 0x42, 0x78, 0xbf, 0x2c, 0x68, 0x1e, 0x73, 0xa5, 0xc2, 0x24, 0x26, 0x0f, 0x01,
	0x94, 0x31, 0x67, 0x21, 0xeb, 0x5b, 0x9b, 0xd6, 0x56, 0x9d, 0xb6, 0x73, 0xc4, 0x67, 0xe4, 0x35,
	0xb4, 0x02, 0x99, 0x86, 0xf3, 0xe0, 0x2c, 0xed, 0xdb, 0x9b, 0xd6, 0x56, 0x67, 0xb4, 0x36, 0x2c,
	0xd8, 0xc7, 0x79, 0x80, 0x96, 0x29, 0xe4, 0x7f, 0x68, 0xc6, 0x09, 0xe3, 0x9a, 0xaa, 0x8e, 0x54,
	0x0d, 0xed, 0xfa, 0x8c, 0x3c, 0x05, 0xf7, 0x5c, 0x26, 0xd9, 0xa2, 0xef, 0x20, 0x49, 0xb7, 0x24,
	0xf9, 0xac, 0x51, 0x6a, 0x82, 0xe4, 0x05, 0x34, 0x05, 0x0f, 0x18, 0x97, 0xaa, 0xef, 0x6e, 0xd6,
	0xb7, 0x3a, 0xa3, 0x95, 0x32, 0x6f, 0x9a, 0x30, 0x4e, 0x8b, 0xa8, 0x77, 0x04, 0x8e, 0x06, 0x48,
	0x17, 0xec, 0x52, 0xb5, 0x1d, 0x32, 0xf2, 0x0c, 0x9c, 0x74, 0xb9, 0xe0, 0x28, 0xb5, 0x7b, 0x43,
	0xaa, 0x4e, 0x3e, 0x59, 0x2e, 0x38, 0xc5, 0x30, 0x21, 0xe0, 0x04, 0x8c, 0x49, 0xd4, 0xd8, 0xa6,
	0x68, 0x7b, 0x5f, 0xc1, 0x45, 0x2d, 0x15, 0xce, 0x0d, 0x68, 0x9b, 0xb2, 0xba, 0x2b, 0x1b, 0xe1,
	0x96, 0x01, 0x7c, 0x46, 0x9e, 0x80, 0xab, 0x3b, 0x54, 0xfd, 0xfa, 0x7d, 0x7a, 0x4d, 0xcc, 0xfb,
	0x61, 0x41, 0xab, 0x18, 0x16, 0x19, 0x40, 0x6b, 0x1e, 0x0a, 0x1e, 0x07, 0x11, 0xc7, 0x22, 0x6d,
	0x5a, 0xfa, 0x5a, 0x97, 0x0a, 0xbf, 0xf3, 0xbc, 0x0a, 0xda, 0x88, 0x5d, 0x04, 0xdb, 0x85, 0x56,
	0x6d, 0xeb, 0xa5, 0x9d, 0x8a, 0xe4, 0xec, 0x72, 0x86, 0xd9, 0x8e, 0x59, 0x1a, 0x22, 0xc7, 0xfa,
	0xc9, 0x06, 0x18, 0x67, 0x16, 0x67, 0x51, 0xdf, 0x35, 0x8a, 0x11, 0x98, 0x66, 0x91, 0x27, 0xc1,
	0xdd, 0xd1, 0x76, 0xa5, 0xcf, 0x7f, 0x2d, 0xde, 0x83, 0xba, 0x0c, 0xbe, 0x61, 0xd5, 0xff, 0xa8,
	0x36, 0xc9, 0x23, 0xe8, 0x2c, 0x32, 0x21, 0x38, 0x9b, 0xcd, 0x65, 0x12, 0xe1, 0xea, 0xea, 0x14,
	0x0c, 0xb4, 0x27, 0x93, 0xc8, 0xf3, 0xa1, 0x8d, 0x35, 0xf7, 0x44, 0x70, 0x4e, 0x1e, 0x80, 0x11,
	0x73, 0x7d, 0x6f, 0x4d, 0xf4, 0x7d, 0x46, 0x1e, 0x83, 0xa3, 0x27, 0x96, 0x5f, 0xda, 0x9d, 0x61,
	0x62, 0xc8, 0x1b, 0x43, 0x17, 0xa9, 0xc6, 0x4a, 0xf1, 0xe8, 0x54, 0x70, 0x56, 0xe9, 0xe3, 0x8e,
	0x1a, 0xbb, 0xa2, 0xe6, 0xb7, 0x05, 0xad, 0x43, 0x99, 0x9c, 0x4b, 0xae, 0xd4, 0xdf, 0xee, 0xff,
	0x08, 0x7a, 0x41, 0x51, 0x69, 0x86, 0x32, 0x15, 0x32, 0x76, 0x46, 0xcf, 0x4b, 0x75, 0x05, 0xd7,
	0xb0, 0xd4, 0x84, 0x0a, 0xd5, 0x24, 0x4e, 0xe5, 0x92, 0xae, 0x06, 0xb7, 0xd1, 0xc1, 0x0e, 0xac,
	0xdf, 0x97, 0xa8, 0xe7, 0x7a, 0xc9, 0x97, 0xb9, 0x04, 0x6d, 0x92, 0x75, 0x70, 0xaf, 0x02, 0x91,
	0x15, 0x2b, 0x31, 0xce, 0x47, 0xfb, 0x83, 0xf5, 0x72, 0x08, 0xad, 0xe2, 0xa4, 0x09, 0x40, 0xe3,
	0x60, 0x7c, 0x7c, 0x32, 0xa1, 0xbd, 0x9a, 0xb6, 0xf7, 0x27, 0xe3, 0xdd, 0x09, 0xed, 0x59, 0xda,
	0x9e, 0x7e, 0xa1, 0x07, 0xe3, 0xfd, 0x9e, 0x3d, 0xfa, 0x69, 0x43, 0x73, 0xd7, 0xc8, 0x25, 0x6f,
	0xa0, 0xe3, 0xc7, 0x61, 0x5a, 0x7c, 0x00, 0xbd, 0xb2, 0x8f, 0x1c, 0x19, 0xdc, 0x98, 0xbb, 0xfe,
	0x2d, 0x6a, 0xe4, 0x1d, 0xac, 0x4e, 0x93, 0x34, 0x9c, 0x2f, 0xaf, 0x97, 0x48, 0xca, 0x9c, 0x12,
	0xab, 0xbe, 0x7b, 0x05, 0xed, 0xc3, 0x4c, 0x5d, 0xe4, 0xe7, 0x76, 0xfb, 0x45, 0x35, 0x7b, 0x5b,
	0x67, 0x0b, 0x61, 0xb2, 0xef, 0xe3, 0xbf, 0xc3, 0xe0, 0xd5, 0xc8, 0x7b, 0x58, 0xf9, 0x74, 0xc1,
	0xcf, 0x2e, 0xcb, 0x6d, 0xae, 0x55, 0x96, 0x32, 0xa8, 0x42, 0x5e, 0xed, 0xb4, 0x81, 0x3f, 0xe3,
	0xdb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x95, 0x02, 0x7b, 0xdb, 0x2a, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeliverClient is the client API for Deliver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeliverClient interface {
	InitSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Null, error)
	NotifyBlockFlag(ctx context.Context, in *BlockFlag, opts ...grpc.CallOption) (*Null, error)
	PushBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*Null, error)
	PullBlock(ctx context.Context, in *BlockFlag, opts ...grpc.CallOption) (*Block, error)
	CheckProgress(ctx context.Context, in *Progress, opts ...grpc.CallOption) (*Progress, error)
}

type deliverClient struct {
	cc *grpc.ClientConn
}

func NewDeliverClient(cc *grpc.ClientConn) DeliverClient {
	return &deliverClient{cc}
}

func (c *deliverClient) InitSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Null, error) {
	out := new(Null)
	err := c.cc.Invoke(ctx, "/deliver.Deliver/InitSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliverClient) NotifyBlockFlag(ctx context.Context, in *BlockFlag, opts ...grpc.CallOption) (*Null, error) {
	out := new(Null)
	err := c.cc.Invoke(ctx, "/deliver.Deliver/NotifyBlockFlag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliverClient) PushBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*Null, error) {
	out := new(Null)
	err := c.cc.Invoke(ctx, "/deliver.Deliver/PushBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliverClient) PullBlock(ctx context.Context, in *BlockFlag, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/deliver.Deliver/PullBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliverClient) CheckProgress(ctx context.Context, in *Progress, opts ...grpc.CallOption) (*Progress, error) {
	out := new(Progress)
	err := c.cc.Invoke(ctx, "/deliver.Deliver/CheckProgress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeliverServer is the server API for Deliver service.
type DeliverServer interface {
	InitSession(context.Context, *Session) (*Null, error)
	NotifyBlockFlag(context.Context, *BlockFlag) (*Null, error)
	PushBlock(context.Context, *Block) (*Null, error)
	PullBlock(context.Context, *BlockFlag) (*Block, error)
	CheckProgress(context.Context, *Progress) (*Progress, error)
}

func RegisterDeliverServer(s *grpc.Server, srv DeliverServer) {
	s.RegisterService(&_Deliver_serviceDesc, srv)
}

func _Deliver_InitSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverServer).InitSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deliver.Deliver/InitSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverServer).InitSession(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deliver_NotifyBlockFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockFlag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverServer).NotifyBlockFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deliver.Deliver/NotifyBlockFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverServer).NotifyBlockFlag(ctx, req.(*BlockFlag))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deliver_PushBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverServer).PushBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deliver.Deliver/PushBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverServer).PushBlock(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deliver_PullBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockFlag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverServer).PullBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deliver.Deliver/PullBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverServer).PullBlock(ctx, req.(*BlockFlag))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deliver_CheckProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Progress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverServer).CheckProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deliver.Deliver/CheckProgress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverServer).CheckProgress(ctx, req.(*Progress))
	}
	return interceptor(ctx, in, info, handler)
}

var _Deliver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deliver.Deliver",
	HandlerType: (*DeliverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitSession",
			Handler:    _Deliver_InitSession_Handler,
		},
		{
			MethodName: "NotifyBlockFlag",
			Handler:    _Deliver_NotifyBlockFlag_Handler,
		},
		{
			MethodName: "PushBlock",
			Handler:    _Deliver_PushBlock_Handler,
		},
		{
			MethodName: "PullBlock",
			Handler:    _Deliver_PullBlock_Handler,
		},
		{
			MethodName: "CheckProgress",
			Handler:    _Deliver_CheckProgress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deliver.proto",
}