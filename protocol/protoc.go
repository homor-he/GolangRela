package protocol

type ui8_2 [2]uint8
type ui8_3 [3]uint8
type ui8_4 [4]uint8

// 客户端连接类型
const (
	H5 = 1300
)

// 新包包头
type Header struct {
	Magic   ui8_4
	Serial  ui8_4
	Origine ui8_2
	Reserve ui8_2
	Msgtype ui8_4
	Param   ui8_4
	Length  ui8_4
}

// 新包包头长度
const LEN_HEAD int = 24

// 新包协议类型
const (
	REQ               = 0x0
	ACK               = 0x80000000
	ORGINE            = 0x600  //orgine包，用于建立对象
	GENERAL_LOGIN_REQ = 0x1203 //普通登录包
	GENERAL_LOGIN_ACK = GENERAL_LOGIN_REQ | ACK
	LOGOUT_REQ        = 0x1204
	LOGOUT_ACK        = LOGOUT_REQ | ACK
	TRANS             = 0x12ff //老协议转发
)

// 后缀结构体
type SuffixPkg struct {
	Offset ui8_4
	Size   ui8_4
}

type OrigineWithIPPkg struct {
	Head Header
	Ip   ui8_4
	Port ui8_2
}

// 老包包头
type OldPacket struct {
	Tag       ui8_2
	Length    ui8_2
	CheckCode uint8
	MsgVer    uint8
	Ident     ui8_4
	reserve   ui8_2
}

// 老包协议
const (
	Heart    = 0x0
	HeartAck = 0x0 | 0x80000000
)
