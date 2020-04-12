package types

import (
	"unsafe"

	cwrap "github.com/richardrigby/rclgo/internal"
)

type StdMsgsUInt8 struct {
	data    *cwrap.StdMsgs_MsgUInt8
	MsgType MessageTypeSupport
}

func (msg *StdMsgsUInt8) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsUInt8) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsUInt8) InitMessage() {
	msg.data = cwrap.InitStdMsgsMsgUInt8()
	msg.MsgType = GetMessageTypeFromStdMsgsUInt8()
}

func (msg *StdMsgsUInt8) SetUInt8(data uint8) {
	//TODO: to implement the setter
	msg.data.Set(data)
}

func (msg *StdMsgsUInt8) GetUInt8() uint8 {
	return uint8(msg.data.Data())
}

func (msg *StdMsgsUInt8) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue := ""
	return myRetValue
}

func (msg *StdMsgsUInt8) DestroyMessage() {
	cwrap.DestroyStdMsgsMsgUInt8(msg.data)
}

func GetMessageTypeFromStdMsgsUInt8() MessageTypeSupport {
	ret := cwrap.GetMessageTypeFromStdMsgsMsgUInt8()
	return MessageTypeSupport{ret}
}
