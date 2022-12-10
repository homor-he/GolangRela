package basefunc

import (
	"bytes"
	"encoding/binary"
	"strings"
	//"unsafe"
)

func Uint32ToUint8(u32 uint32, u8 []uint8) {
	len := len(u8)
	for i := 0; i < len; i++ {
		u8[i] = uint8(u32 >> (i * 8))
	}
}

func Uint8ArrayToUint32(u8 []uint8) uint32 {
	return binary.LittleEndian.Uint32(u8)
}

func Encode(order binary.ByteOrder, data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, order, data)
	if err != nil {
		Gbs_log.Printf("struct to bytes encode err,err:%s", err.Error())
		return nil, err
	}
	//bytes := buf.Bytes()
	//fmt.Print(bytes)
	return buf.Bytes(), err
}

func Decode(order binary.ByteOrder, data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	err := binary.Read(buf, order, to)
	if err != nil {
		Gbs_log.Printf("bytes to struct decode err,err:%s", err.Error())
	}
	return err
}

func EncodeAndAppend(order binary.ByteOrder, data ...interface{}) ([]byte, error) {
	var err error
	dataSize := len(data)
	s := make([][]byte, dataSize)
	for i := 0; i < dataSize; i++ {
		buf := bytes.NewBuffer([]byte{})
		err := binary.Write(buf, order, data[i])
		if err != nil {
			Gbs_log.Printf("struct to bytes encode err")
			return nil, err
		}
		s[i] = buf.Bytes()
	}
	sep := []byte("")
	return bytes.Join(s, sep), err
}

func TrimSpaceTostring(data []byte) string {
	if data == nil {
		return ""
	}
	index := strings.Index(string(data), "\x00")
	if index > 0 {
		return string(data[:index])
	}
	return string(data)
}

// func Encode(data interface{}, len uint32) ([]byte, error) {
// 	// buf := bytes.NewBuffer(nil)
// 	// enc := gob.NewEncoder(buf)
// 	// err := enc.Encode(data)
// 	// if err != nil {
// 	// 	basefunc.Gbs_log.Print("encode fail")
// 	// 	return nil, err
// 	// }
// 	// return buf.Bytes(), nil
// 	var err error
// 	var x reflect.SliceHeader
// 	x.Len = int(len)
// 	x.Cap = int(len)
// 	x.Data = reflect.ValueOf(data).Pointer()
// 	return *(*[]byte)(unsafe.Pointer(&x)), err
// }

// func Decode(data []byte, to interface{}) error {
// 	buf := bytes.NewBuffer(data)
// 	dec := gob.NewDecoder(buf)
// 	return dec.Decode(to)
// }
