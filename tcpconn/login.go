package tcpconn

import (
	"autotest/basefunc"
	"autotest/business"
	"autotest/protocol"
	"encoding/binary"
	"encoding/json"
	"errors"
	"time"
	"unsafe"
)

func SendOriginePkg(client *Client) error {
	var msg protocol.Header
	pkgLen := uint32(unsafe.Sizeof(msg))
	basefunc.Uint32ToUint8(protocol.ORGINE, msg.Msgtype[:])
	basefunc.Uint32ToUint8(pkgLen, msg.Length[:])
	basefunc.Uint32ToUint8(protocol.H5, msg.Origine[:])
	basefunc.Gbs_log.Printf("origine:%d len:%d", basefunc.Uint8ArrayToUint32(msg.Msgtype[:]), basefunc.Uint8ArrayToUint32(msg.Length[:]))
	bytes, err := basefunc.Encode(binary.LittleEndian, &msg)
	//bytes := *(*[]byte)(unsafe.Pointer(&msg))
	if err != nil {
		basefunc.Gbs_log.Printf("SendOriginePkg append err:%s", err.Error())
		return err
	}
	return client.PostMsg(bytes, msg.Msgtype, msg.Length)
}

func RecvOriginePkgAck(bytes []byte) {
	//oldPacket := *(*protocol.OldPacket)(unsafe.Pointer(&bytes[protocol.LEN_HEAD]))
	//if basefunc.Uint8ArrayToUint32(oldPacket.Ident[:]) == protocol.HeartAck {
	//basefunc.Gbs_log.Print("recv heartbeat package")
	//}
}

func PostOrigineWithIPPkg(client *Client) error {
	var msg protocol.OrigineWithIPPkg
	pkgLen := uint32(unsafe.Sizeof(msg))
	basefunc.Uint32ToUint8(protocol.ORGINE, msg.Head.Msgtype[:])
	basefunc.Uint32ToUint8(pkgLen, msg.Head.Length[:])
	basefunc.Uint32ToUint8(protocol.H5, msg.Head.Origine[:])
	basefunc.Gbs_log.Printf("origineWithIP:%d len:%d", basefunc.Uint8ArrayToUint32(msg.Head.Msgtype[:]), basefunc.Uint8ArrayToUint32(msg.Head.Length[:]))
	bytes, err := basefunc.Encode(binary.LittleEndian, msg)
	if err != nil {
		basefunc.Gbs_log.Print(err.Error())
		return err
	}
	return client.PostMsg(bytes, msg.Head.Msgtype, msg.Head.Length)
}

func SendCltHeartPkg(client *Client) error {
	for {
		var newHead protocol.Header
		var oldHead protocol.OldPacket
		newHeadSize := unsafe.Sizeof(newHead)
		oldHeadSize := unsafe.Sizeof(oldHead)
		basefunc.Uint32ToUint8(protocol.TRANS, newHead.Msgtype[:])
		basefunc.Uint32ToUint8(uint32(newHeadSize+oldHeadSize), newHead.Length[:])

		basefunc.Uint32ToUint8(protocol.Heart, oldHead.Ident[:])
		bytes, err := basefunc.EncodeAndAppend(binary.LittleEndian, &newHead, &oldHead)
		if err != nil {
			basefunc.Gbs_log.Printf("SendCltHeartPkg append err:%s", err.Error())
			return err
		}
		err = client.SendMsg(bytes, newHead.Msgtype, newHead.Length, protocol.TRANS)
		if err != nil {
			return err
		}
		//return err
		time.Sleep(time.Duration(5) * time.Second)
	}
}

func SendGeneralLoginPkg(client *Client, jsonStr string) error {
	var msg protocol.GeneralLogin_Req
	var jsonMap map[string]interface{}

	//json反序列化成map
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		basefunc.Gbs_log.Printf("json decode fail,err:%s", err.Error())
		return err
	}
	uid, ok := jsonMap["User_ID"].(string)
	if !ok {
		return errors.New("User_ID string parse fail")
	}
	pwd, ok := jsonMap["Password"].(string)
	if !ok {
		return errors.New("Password string parse fail")
	}
	copy(msg.LoginName[:], []byte(uid))
	copy(msg.Password[:], []byte(pwd))
	basefunc.Uint32ToUint8(34, msg.LoginType[:])
	basefunc.Uint32ToUint8(9730, msg.PlatID[:])
	basefunc.Uint32ToUint8(protocol.H5, msg.CntTypeID[:])
	copy(msg.DevType[:], "PC-Chrome(105.0.0.0)")
	copy(msg.DevName[:], "web")
	basefunc.Uint32ToUint8(3010, msg.OsType[:])
	basefunc.Uint32ToUint8(4112, msg.JsVersion[:])
	copy(msg.JsAddr[:], "game_center/web")

	basefunc.Uint32ToUint8(protocol.GENERAL_LOGIN_REQ, msg.Head.Msgtype[:])
	basefunc.Uint32ToUint8(uint32(unsafe.Sizeof(msg)), msg.Head.Length[:])
	bytes, err := basefunc.Encode(binary.LittleEndian, msg)
	if err != nil {
		basefunc.Gbs_log.Printf("SendGeneralLoginPkg append err:%s", err.Error())
		return err
	}
	err = client.SendMsg(bytes, msg.Head.Msgtype, msg.Head.Length, protocol.GENERAL_LOGIN_ACK)
	return err
}

func RecvGeneralLoginAck(data []byte, client *Client) {
	var msg protocol.GeneralLogin_Ack
	head := *(*protocol.Header)(unsafe.Pointer(&data[0]))
	headLen := basefunc.Uint8ArrayToUint32(head.Length[:])
	if headLen < uint32(unsafe.Sizeof(msg)) {
		basefunc.Gbs_log.Printf("RecvGeneralLoginAck len err,len:%d", headLen)
		return
	}
	basefunc.Decode(binary.LittleEndian, data, &msg)
	//basefunc.Gbs_log.Printf("%v", msg)
	ret := basefunc.Uint8ArrayToUint32(msg.Head.Param[:])
	if ret != 0 {
		basefunc.Gbs_log.Printf("login fail，ack ret:%d", ret)
		return
	}
	uid := basefunc.Uint8ArrayToUint32(msg.UserID[:])

	var userInfo = new(business.UserInfo)
	userInfo.UserID = uid
	userInfo.UserName = basefunc.TrimSpaceTostring(msg.UserName[:])
	userInfo.NickName = basefunc.TrimSpaceTostring(msg.NickName[:])
	client.UserInfo = userInfo
	basefunc.Gbs_log.Printf("uid:%d login success, username:%s nickname:%s", uid, userInfo.UserName, userInfo.NickName)
}

func SendLogout(client *Client, uid uint32) error {
	var msg protocol.Logout_Req
	basefunc.Uint32ToUint8(protocol.LOGOUT_REQ, msg.Head.Msgtype[:])
	basefunc.Uint32ToUint8(uint32(unsafe.Sizeof(msg)), msg.Head.Length[:])
	basefunc.Uint32ToUint8(uid, msg.UserID[:])

	bytes, err := basefunc.Encode(binary.LittleEndian, &msg)
	if err != nil {
		basefunc.Gbs_log.Printf("SendLogout append err:%s", err.Error())
		return err
	}
	err = client.SendMsg(bytes, msg.Head.Msgtype, msg.Head.Length, protocol.LOGOUT_ACK)
	return err
}

func RecvLogoutAck(data []byte, client *Client) {
	// var msg protocol.Logout_Ack
	// head := *(*protocol.Header)(unsafe.Pointer(&data[0]))
	// headLen := basefunc.Uint8ArrayToUint32(head.Length[:])
	// if headLen < uint32(unsafe.Sizeof(msg)) {
	// 	basefunc.Gbs_log.Printf("RecvLogoutAck len err,len:%d", headLen)
	// 	return
	// }
	client.UserInfo = nil
}
