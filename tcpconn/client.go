package tcpconn

import (
	"autotest/basefunc"
	"autotest/business"
	"autotest/protocol"
	"errors"
	"fmt"
	"net"
	"unsafe"
)

const MaxConnect int = 50

var Gbs_clientList map[string]Client = make(map[string]Client)

type Client struct {
	Network string
	Addr    string
	NetConn net.Conn

	UserInfo *business.UserInfo
}

func BuildConnect(socketClient *Client, method string, host string) error {
	socketClient.Network = method
	socketClient.Addr = host
	socketErr := socketClient.Connect()
	if socketErr != nil {
		basefunc.Gbs_log.Printf("socker addr:%s connect fail", socketClient.Addr)
		return socketErr
	}

	err := SendOriginePkg(socketClient)
	return err
}

func (client *Client) Connect() error {
	var err error
	client.NetConn, err = net.Dial(client.Network, client.Addr)
	if err != nil {
		basefunc.Gbs_log.Print(err.Error())
		return err
	}
	basefunc.Gbs_log.Printf("build connect success,address:%s", client.Addr)
	return err
}

func (client *Client) PostMsg(bytes []byte, msgType [4]uint8, msgLen [4]uint8) error {
	len, err := client.NetConn.Write(bytes)
	if err != nil {
		client.NetConn.Close()
	}
	mt := basefunc.Uint8ArrayToUint32(msgType[:])
	pkgLen := basefunc.Uint8ArrayToUint32(msgLen[:])
	if len != int(pkgLen) {
		basefunc.Gbs_log.Printf("send msgId:%d len err, realLen:%d origineLen:%d", mt, len, pkgLen)
		return errors.New("len err")
	}

	if err != nil {
		basefunc.Gbs_log.Printf(err.Error())
		return err
	}
	basefunc.Gbs_log.Printf("send msg id:%d success", mt)
	return err
}

func (client *Client) HandleMsg(recvMsgType uint32) error {
	recv := [4 * 1024]byte{}
	len, err := client.NetConn.Read(recv[:])
	if err != nil {
		client.NetConn.Close()
		return err
	}
	if len < protocol.LEN_HEAD {
		newError := fmt.Sprintf("HandleMsg len err, len:%d", len)
		return errors.New(newError)
	}
	client.OnMsg(recv[:], recvMsgType)
	return nil
}

func (client *Client) OnMsg(data []byte, recvMsgType uint32) {
	head := *(*protocol.Header)(unsafe.Pointer(&data[0]))
	msgLen := basefunc.Uint8ArrayToUint32(head.Length[:])
	if msgLen < uint32(unsafe.Sizeof(head)) {
		basefunc.Gbs_log.Printf("OnMsg len err")
		return
	}
	msgType := basefunc.Uint8ArrayToUint32(head.Msgtype[:])
	//basefunc.Gbs_log.Printf("recv msgType:%d", msgType)
	switch msgType {
	case protocol.TRANS:
		RecvOriginePkgAck(data)
		break
	case protocol.GENERAL_LOGIN_ACK:
		RecvGeneralLoginAck(data, client)
		break
	case protocol.LOGOUT_ACK:
		RecvLogoutAck(data, client)
		break
	default:
		break
	}

	//收到的同步消息id和想要的匹配不上,就继续等待
	if msgType != recvMsgType {
		client.HandleMsg(recvMsgType)
	}
}

func (client *Client) SendMsg(bytes []byte, msgType [4]uint8, msgLen [4]uint8, recvMsgType uint32) error {
	err := client.PostMsg(bytes, msgType, msgLen)
	if err != nil {
		basefunc.Gbs_log.Printf("send msg err:%s, req msgid:%d", err.Error(), protocol.ORGINE)
		return err
	}
	err = client.HandleMsg(recvMsgType)
	if err != nil {
		basefunc.Gbs_log.Printf("recv msg err:%s, req msgid:%d", err.Error(), protocol.ORGINE)
	}
	return err
}
