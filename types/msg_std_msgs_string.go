package types

import (
	"unsafe"

	cwrap "github.com/richardrigby/rclgo/internal"
)

type StdMsgsString struct {
	data *cwrap.StdMsgs_MsgString
	StdMsgsBase
}

func (msg *StdMsgsString) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsString) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsString) GetDataAsString() string {
	return msg.data.String()
}

func (msg *StdMsgsString) InitMessage() {
	msg.data = cwrap.InitStdMsgsMsgString()
	msg.MsgType = GetMessageTypeFromStdMsgsString()
}

func (msg *StdMsgsString) DestroyMessage() {
	cwrap.DestroyStdMsgsMsgString(msg.data)
}

func (msg *StdMsgsString) SetText(text string) {
	msg.data.Set(text)
}

func GetMessageTypeFromStdMsgsString() MessageTypeSupport {
	ret := cwrap.GetMessageTypeFromStdMsgsMsgString()
	return MessageTypeSupport{ret}
}
