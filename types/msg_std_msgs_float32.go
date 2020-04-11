package types

import (
	"unsafe"

	cwrap "github.com/richardrigby/rclgo/internal"
)

type StdMsgsFloat32 struct {
	data    *cwrap.StdMsgs_MsgFloat32
	MsgType MessageTypeSupport
}

func (msg *StdMsgsFloat32) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsFloat32) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsFloat32) InitMessage() {
	msg.data = cwrap.InitStdMsgsMsgFloat32()
	msg.MsgType = GetMessageTypeFromStdMsgsFloat32()
}

func (msg *StdMsgsFloat32) SetFloat32(data float32) {
	//TODO: to implement the setter
	msg.data.Set(data)
}

func (msg *StdMsgsFloat32) GetFloat32() float32 {
	return float32(msg.data.Data())
}

func (msg *StdMsgsFloat32) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue := ""
	return myRetValue
}

func (msg *StdMsgsFloat32) DestroyMessage() {
	cwrap.DestroyStdMsgsMsgFloat32(msg.data)
}

func GetMessageTypeFromStdMsgsFloat32() MessageTypeSupport {
	ret := cwrap.GetMessageTypeFromStdMsgsMsgFloat32()
	return MessageTypeSupport{ret}
}
