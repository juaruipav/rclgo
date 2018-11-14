package types

// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include "msg_types.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type StdMsgsStringMsg struct {
	Data    *C.std_msgs__msg__String
	MsgType MessageTypeSupport
}

func (msg *StdMsgsStringMsg) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsStringMsg) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.Data)}
}

func (msg *StdMsgsStringMsg) InitMessage() {
	msg.Data = C.init_std_msgs_msg_String()
	msg.MsgType = GetMessageTypeFromStdMsgsString()
}

func (msg *StdMsgsStringMsg) SetText(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	retValue := C.rosidl_generator_c__String__assign(&msg.Data.data, cText)
	fmt.Printf("Ret value from assign is %t\n", retValue)
}

func GetMessageTypeFromStdMsgsString() MessageTypeSupport {
	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_String()}
}
