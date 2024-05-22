// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/msg/objmsg.proto

package msg

type MsgPic struct {
	SmallPicUrl    []byte `protobuf:"bytes,1,opt"`
	OriginalPicUrl []byte `protobuf:"bytes,2,opt"`
	LocalPicId     int32  `protobuf:"varint,3,opt"`
}

type ObjMsg struct {
	MsgType        int32             `protobuf:"varint,1,opt"`
	Title          []byte            `protobuf:"bytes,2,opt"`
	BytesAbstact   []byte            `protobuf:"bytes,3,opt"`
	TitleExt       []byte            `protobuf:"bytes,5,opt"`
	MsgPic         []*MsgPic         `protobuf:"bytes,6,rep"`
	MsgContentInfo []*MsgContentInfo `protobuf:"bytes,7,rep"`
	ReportIdShow   int32             `protobuf:"varint,8,opt"`
}

type MsgContentInfo struct {
	ContentInfoId []byte   `protobuf:"bytes,1,opt"`
	MsgFile       *MsgFile `protobuf:"bytes,2,opt"`
}

type MsgFile struct {
	BusId         int32  `protobuf:"varint,1,opt"`
	FilePath      []byte `protobuf:"bytes,2,opt"`
	FileSize      int64  `protobuf:"varint,3,opt"`
	FileName      string `protobuf:"bytes,4,opt"`
	Int64DeadTime int64  `protobuf:"varint,5,opt"`
	FileSha1      []byte `protobuf:"bytes,6,opt"`
	Ext           []byte `protobuf:"bytes,7,opt"`
}