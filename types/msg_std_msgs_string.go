package types

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #cgo LDFLAGS: -L/opt/ros/eloquent/lib -Wl,-rpath=/opt/ros/eloquent/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include <rcl/rcl.h>
// #include "../cwrap/msg_types.h"
import "C"

import (
	"unsafe"

	"github.com/richardrigby/rclgo/cwrap"
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
