package protocol

const LEN_USERNAME int = 64
const LEN_PWD int = 32
const LEN_MAC int = 24
const LEN_DEVICE int = 128
const LEN_DEVICE_ID int = 32
const LEN_JSADDR int = 32
const LEN_CHECKCODE int = 8
const LEN_IP int = 16
const LEN_HOST int = 40
const LEN_SZCOMM int = 256
const LEN_NICKNAME int = 32
const LEN_MOBILE int = 16
const LEN_AUTHENTIC int = 64

// ***************************************************************************
// 普通登录包 GENERAL_LOGIN_REQ = 0x1203
type GeneralLogin_Req struct {
	Head Header

	AccountType ui8_4 // 账号类型， 0:UserName 1:UserID 2:Mobile 3:Ticket 4:QQ 5:Sina 6:Visitor 7:Weixin 8:Baidu 9:360 10:Huawei 11:Wxpub 12:Oppo 13:Vivo 14:Mi 15:Yyb 20:手机短信 21:手机票据 22:手机一键登录， 默认为0
	LoginName   [LEN_USERNAME]byte
	Password    [LEN_PWD]byte
	LoginType   ui8_4 // 登陆类型0xAB (A,0:PC,2:手机)(B,1:锁定本机,2:第三方账号登陆,3:手机登陆)
	MacAddr1    [LEN_MAC]byte
	MacAddr2    [LEN_MAC]byte

	TransID    ui8_4 //代理transferID
	PlatID     ui8_4 //平台id
	CltVersion ui8_4 //客户端版本

	CntTypeID  ui8_4 //连接类型 pc:1100 lua:1200 h5:1300
	GameID     ui8_4 //游戏ID
	LocalIP    ui8_4 //客户端本机IP
	InternetIP ui8_4 //客户端外网IP

	DevType [LEN_DEVICE]byte //设备型号
	DevName [LEN_DEVICE]byte //设备名称
	OsType  ui8_4            //操作系统

	JsVersion ui8_4 //js版本号
	Port      ui8_4
	Reserve   ui8_4               //保留
	DevID     [LEN_DEVICE_ID]byte //设备ID

	JsAddr    [LEN_JSADDR]byte    //客户端js地址
	CheckCode [LEN_CHECKCODE]byte //一级校验码
	AccID     ui8_4
	TimeStamp ui8_4 //请求时间戳

	LongKey SuffixPkg //手机一键登录第三方key
	Txt     SuffixPkg //其它后缀数据
}

// 普通登录回包 GENERAL_LOGIN_ACK
type GeneralLogin_Ack struct {
	Head Header

	IP      ui8_4
	Port    ui8_4
	StrIp   [LEN_IP]byte
	StrPort [LEN_HOST]byte

	UserID   ui8_4
	FigureID ui8_4 //玩家头像id
	FormID   ui8_4 //游戏形象id
	SexID    ui8_4 //性别
	UserVip  ui8_4 //vip会员标识

	Money ui8_4 //银子
	Coin  ui8_4 //元宝
	Bonus ui8_4 //奖券

	Score       ui8_4 //累计积分
	MasterScore ui8_4 //大师分
	MatchCount  ui8_4 //比赛盘数
	MatchTime   ui8_4 //比赛时长

	UserName [LEN_USERNAME]byte //账号名称
	NickName [LEN_NICKNAME]byte //用户昵称
	UserFace [LEN_SZCOMM]byte   //用户自定义头像
	RealName [LEN_NICKNAME]byte //真实姓名
	CardID   [LEN_NICKNAME]byte //身份证

	Mobile    [LEN_MOBILE]byte    //隐藏中间4位安保手机密码
	Authentic [LEN_AUTHENTIC]byte //票据

	DevTrust ui8_4 //设备是否可信 0-不可信 1-可信
	MID      ui8_4 //MID

	IsNewUser ui8_4 //是否新用户
	IsUpdate  ui8_4 //客户端判断版本号 1:更新(已暂停使用)
	IsMember  ui8_4 //是否转正
	Param     ui8_4 //保留

	SuffixBuf  SuffixPkg
	SuffixType ui8_4
	Count      ui8_4
}

//***************************************************************************

// ***************************************************************************
// 登出包 LOGOUT_REQ = 0x1204
type Logout_Req struct {
	Head Header

	UserID     ui8_4
	TransferID ui8_4
	NickName   [LEN_NICKNAME]byte
}

// 登出回包 LOGOUT_ACK
type Logout_Ack struct {
	Head        Header
	AnonymousID ui8_4 //匿名访问id
}

//***************************************************************************
