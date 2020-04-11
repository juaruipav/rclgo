package types

import (
	"unsafe"

	cwrap "github.com/richardrigby/rclgo/internal"
)

type StdMsgsInt64 struct {
	data    *cwrap.StdMsgs_MsgInt64
	MsgType MessageTypeSupport
}

func (msg *StdMsgsInt64) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsInt64) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsInt64) InitMessage() {
	msg.data = cwrap.InitStdMsgsMsgInt64()
	msg.MsgType = GetMessageTypeFromStdMsgsInt64()
}

func (msg *StdMsgsInt64) SetInt64(data int64) {
	//TODO: to implement the setter
	msg.data.Set(data)
}

func (msg *StdMsgsInt64) GetInt64() int64 {
	return int64(msg.data.Data())
}

func (msg *StdMsgsInt64) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue := ""
	return myRetValue
}

func (msg *StdMsgsInt64) DestroyMessage() {
	cwrap.DestroyStdMsgsMsgInt64(msg.data)
}

func GetMessageTypeFromStdMsgsInt64() MessageTypeSupport {
	ret := cwrap.GetMessageTypeFromStdMsgsMsgInt64()
	return MessageTypeSupport{ret}
}
