package types

import (
	"unsafe"

	cwrap "github.com/richardrigby/rclgo/internal"
)

type StdMsgsUInt32 struct {
	data    *cwrap.StdMsgs_MsgUInt32
	MsgType MessageTypeSupport
}

func (msg *StdMsgsUInt32) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsUInt32) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsUInt32) InitMessage() {
	msg.data = cwrap.InitStdMsgsMsgUInt32()
	msg.MsgType = GetMessageTypeFromStdMsgsUInt32()
}

func (msg *StdMsgsUInt32) SetUInt32(data uint32) {
	//TODO: to implement the setter
	msg.data.Set(data)
}

func (msg *StdMsgsUInt32) GetUInt32() uint32 {
	return uint32(msg.data.Data())
}

func (msg *StdMsgsUInt32) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue := ""
	return myRetValue
}

func (msg *StdMsgsUInt32) DestroyMessage() {
	cwrap.DestroyStdMsgsMsgUInt32(msg.data)
}

func GetMessageTypeFromStdMsgsUInt32() MessageTypeSupport {
	ret := cwrap.GetMessageTypeFromStdMsgsMsgUInt32()
	return MessageTypeSupport{ret}
}
