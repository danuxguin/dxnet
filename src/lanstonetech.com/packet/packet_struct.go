package packet

//4 	UINT32	MsgId  		消息Id
//4 	UINT32	MsgVer 		消息版本号
//4 	UINT32	MsgLen 		消息体长度
//4	UINT32	MsgCpsLen		消息压缩长度

type MsgHeader struct {
	MsgID     uint32
	MsgVer    uint32
	MsgLen    uint32
	MsgCpsLen uint32
}

type PacketCommon struct {
	Account   string
	Signature string
	Token     string
}
