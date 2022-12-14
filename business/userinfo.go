package business

type ClientStatus struct {
	IsLogin bool
}

type UserInfo struct {
	UserID   uint32
	UserName string
	NickName string
}

var Gbs_UserMap map[uint32]UserInfo = make(map[uint32]UserInfo)
var Gbs_StatMap map[uint32]ClientStatus = make(map[uint32]ClientStatus)
