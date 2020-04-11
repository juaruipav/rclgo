package types

import (
	"unsafe"

	cwrap "github.com/richardrigby/rclgo/internal"
)

type StdMsgsInt32 struct {
	data    *cwrap.StdMsgs_MsgInt32
	MsgType MessageTypeSupport
}

func (msg *StdMsgsInt32) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsInt32) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsInt32) InitMessage() {
	msg.data = cwrap.InitStdMsgsMsgInt32()
	msg.MsgType = GetMessageTypeFromStdMsgsInt32()
}

func (msg *StdMsgsInt32) SetInt32(data int32) {
	//TODO: to implement the setter
	msg.data.Set(data)
}

func (msg *StdMsgsInt32) GetInt32() int32 {
	return int32(msg.data.Data())
}

func (msg *StdMsgsInt32) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue := ""
	return myRetValue
}

func (msg *StdMsgsInt32) DestroyMessage() {
	cwrap.DestroyStdMsgsMsgInt32(msg.data)
}

func GetMessageTypeFromStdMsgsInt32() MessageTypeSupport {
	ret := cwrap.GetMessageTypeFromStdMsgsMsgInt32()
	return MessageTypeSupport{ret}
}
