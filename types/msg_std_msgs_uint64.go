package types

import (
	"unsafe"

	cwrap "github.com/richardrigby/rclgo/internal"
)

type StdMsgsUInt64 struct {
	data    *cwrap.StdMsgs_MsgUInt64
	MsgType MessageTypeSupport
}

func (msg *StdMsgsUInt64) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsUInt64) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsUInt64) InitMessage() {
	msg.data = cwrap.InitStdMsgsMsgUInt64()
	msg.MsgType = GetMessageTypeFromStdMsgsUInt64()
}

func (msg *StdMsgsUInt64) SetUInt64(data uint64) {
	//TODO: to implement the setter
	msg.data.Set(data)
}

func (msg *StdMsgsUInt64) GetUInt64() uint64 {
	return uint64(msg.data.Data())
}

func (msg *StdMsgsUInt64) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue := ""
	return myRetValue
}

func (msg *StdMsgsUInt64) DestroyMessage() {
	cwrap.DestroyStdMsgsMsgUInt64(msg.data)
}

func GetMessageTypeFromStdMsgsUInt64() MessageTypeSupport {
	ret := cwrap.GetMessageTypeFromStdMsgsMsgUInt64()
	return MessageTypeSupport{ret}
}
