// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/channel/common.proto

package channel

import (
	msg "github.com/Mrs4s/MiraiGo/client/pb/msg"
	proto "github.com/RomiChan/protobuf/proto"
)

type ChannelContentHead struct {
	Type    proto.Option[uint64] `protobuf:"varint,1,opt"`
	SubType proto.Option[uint64] `protobuf:"varint,2,opt"`
	Random  proto.Option[uint64] `protobuf:"varint,3,opt"`
	Seq     proto.Option[uint64] `protobuf:"varint,4,opt"`
	CntSeq  proto.Option[uint64] `protobuf:"varint,5,opt"`
	Time    proto.Option[uint64] `protobuf:"varint,6,opt"`
	Meta    []byte               `protobuf:"bytes,7,opt"`
}

type DirectMessageMember struct {
	Uin             proto.Option[uint64] `protobuf:"varint,1,opt"`
	Tinyid          proto.Option[uint64] `protobuf:"varint,2,opt"`
	SourceGuildId   proto.Option[uint64] `protobuf:"varint,3,opt"`
	SourceGuildName []byte               `protobuf:"bytes,4,opt"`
	NickName        []byte               `protobuf:"bytes,5,opt"`
	MemberName      []byte               `protobuf:"bytes,6,opt"`
	NotifyType      proto.Option[uint32] `protobuf:"varint,7,opt"`
}

type ChannelEvent struct {
	Type    proto.Option[uint64] `protobuf:"varint,1,opt"`
	Version proto.Option[uint64] `protobuf:"varint,2,opt"`
	OpInfo  *ChannelMsgOpInfo    `protobuf:"bytes,3,opt"`
	_       [0]func()
}

type ChannelExtInfo struct {
	FromNick            []byte                 `protobuf:"bytes,1,opt"`
	GuildName           []byte                 `protobuf:"bytes,2,opt"`
	ChannelName         []byte                 `protobuf:"bytes,3,opt"`
	Visibility          proto.Option[uint32]   `protobuf:"varint,4,opt"`
	NotifyType          proto.Option[uint32]   `protobuf:"varint,5,opt"`
	OfflineFlag         proto.Option[uint32]   `protobuf:"varint,6,opt"`
	NameType            proto.Option[uint32]   `protobuf:"varint,7,opt"`
	MemberName          []byte                 `protobuf:"bytes,8,opt"`
	Timestamp           proto.Option[uint32]   `protobuf:"varint,9,opt"`
	EventVersion        proto.Option[uint64]   `protobuf:"varint,10,opt"`
	Events              []*ChannelEvent        `protobuf:"bytes,11,rep"`
	FromRoleInfo        *ChannelRole           `protobuf:"bytes,12,opt"`
	FreqLimitInfo       *ChannelFreqLimitInfo  `protobuf:"bytes,13,opt"`
	DirectMessageMember []*DirectMessageMember `protobuf:"bytes,14,rep"`
}

type ChannelFreqLimitInfo struct {
	IsLimited      proto.Option[uint32] `protobuf:"varint,1,opt"`
	LeftCount      proto.Option[uint32] `protobuf:"varint,2,opt"`
	LimitTimestamp proto.Option[uint64] `protobuf:"varint,3,opt"`
	_              [0]func()
}

type ChannelInfo struct {
	Id    proto.Option[uint64] `protobuf:"varint,1,opt"`
	Name  []byte               `protobuf:"bytes,2,opt"`
	Color proto.Option[uint32] `protobuf:"varint,3,opt"`
	Hoist proto.Option[uint32] `protobuf:"varint,4,opt"`
}

type ChannelLoginSig struct {
	Type  proto.Option[uint32] `protobuf:"varint,1,opt"`
	Sig   []byte               `protobuf:"bytes,2,opt"`
	Appid proto.Option[uint32] `protobuf:"varint,3,opt"`
}

type ChannelMeta struct {
	FromUin  proto.Option[uint64] `protobuf:"varint,1,opt"`
	LoginSig *ChannelLoginSig     `protobuf:"bytes,2,opt"`
	_        [0]func()
}

type ChannelMsgContent struct {
	Head     *ChannelMsgHead     `protobuf:"bytes,1,opt"`
	CtrlHead *ChannelMsgCtrlHead `protobuf:"bytes,2,opt"`
	Body     *msg.MessageBody    `protobuf:"bytes,3,opt"`
	ExtInfo  *ChannelExtInfo     `protobuf:"bytes,4,opt"`
	_        [0]func()
}

type ChannelMsgCtrlHead struct {
	IncludeUin [][]byte `protobuf:"bytes,1,rep"`
	// repeated uint64 excludeUin = 2; // bytes?
	// repeated uint64 featureid = 3;
	OfflineFlag    proto.Option[uint32] `protobuf:"varint,4,opt"`
	Visibility     proto.Option[uint32] `protobuf:"varint,5,opt"`
	CtrlFlag       proto.Option[uint64] `protobuf:"varint,6,opt"`
	Events         []*ChannelEvent      `protobuf:"bytes,7,rep"`
	Level          proto.Option[uint64] `protobuf:"varint,8,opt"`
	PersonalLevels []*PersonalLevel     `protobuf:"bytes,9,rep"`
	GuildSyncSeq   proto.Option[uint64] `protobuf:"varint,10,opt"`
	MemberNum      proto.Option[uint32] `protobuf:"varint,11,opt"`
	ChannelType    proto.Option[uint32] `protobuf:"varint,12,opt"`
	PrivateType    proto.Option[uint32] `protobuf:"varint,13,opt"`
}

type ChannelMsgHead struct {
	RoutingHead *ChannelRoutingHead `protobuf:"bytes,1,opt"`
	ContentHead *ChannelContentHead `protobuf:"bytes,2,opt"`
	_           [0]func()
}

type ChannelMsgMeta struct {
	AtAllSeq proto.Option[uint64] `protobuf:"varint,1,opt"`
	_        [0]func()
}

type ChannelMsgOpInfo struct {
	OperatorTinyid proto.Option[uint64] `protobuf:"varint,1,opt"`
	OperatorRole   proto.Option[uint64] `protobuf:"varint,2,opt"`
	Reason         proto.Option[uint64] `protobuf:"varint,3,opt"`
	Timestamp      proto.Option[uint64] `protobuf:"varint,4,opt"`
	AtType         proto.Option[uint64] `protobuf:"varint,5,opt"`
	_              [0]func()
}

type PersonalLevel struct {
	ToUin proto.Option[uint64] `protobuf:"varint,1,opt"`
	Level proto.Option[uint64] `protobuf:"varint,2,opt"`
	_     [0]func()
}

type ChannelRole struct {
	Id   proto.Option[uint64] `protobuf:"varint,1,opt"`
	Info []byte               `protobuf:"bytes,2,opt"`
	Flag proto.Option[uint32] `protobuf:"varint,3,opt"`
}

type ChannelRoutingHead struct {
	GuildId           proto.Option[uint64] `protobuf:"varint,1,opt"`
	ChannelId         proto.Option[uint64] `protobuf:"varint,2,opt"`
	FromUin           proto.Option[uint64] `protobuf:"varint,3,opt"`
	FromTinyid        proto.Option[uint64] `protobuf:"varint,4,opt"`
	GuildCode         proto.Option[uint64] `protobuf:"varint,5,opt"`
	FromAppid         proto.Option[uint64] `protobuf:"varint,6,opt"`
	DirectMessageFlag proto.Option[uint32] `protobuf:"varint,7,opt"`
	_                 [0]func()
}
