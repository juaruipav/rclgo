package types

// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include "msg_types.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type MessageTypeSupport struct {
	ROSIdlMessageTypeSupport *C.rosidl_message_type_support_t
}

type Message interface {
	GetMessage() unsafe.Pointer
	InitMessage()
	// DestroyMessage()
}

type StdMsgsStringMsg struct {
	Data  *C.std_msgs__msg__String
	Data2 MessageTypeSupport
	Text  string
}

func (msg StdMsgsStringMsg) GetMessage() unsafe.Pointer {
	// return unsafe.Pointer(msg.Data)
	return unsafe.Pointer(msg.Data2.ROSIdlMessageTypeSupport)
}

func (msg StdMsgsStringMsg) InitMessage() {
	msg.Data = C.init_std_msgs_msg_String()
	msg.Data2 = GetMessageTypeFromStdMsgsString()
	cText := C.CString(msg.Text)
	defer C.free(unsafe.Pointer(cText))
	retValue := C.rosidl_generator_c__String__assign(&msg.Data.data, cText)
	fmt.Printf("Ret value from assign is %t\n", retValue)
}

// func (msg StdMsgsStringMsg) destroyMessage() {
// 	C.destroy_std_msgs_msg_String(unsafe.Pointer(msg.Data)
// }

func GetMessageTypeFromStdMsgsBool() MessageTypeSupport {
	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_Bool()}
}

func GetMessageTypeFromStdMsgsByte() MessageTypeSupport {
	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_Byte()}
}

func GetMessageTypeFromStdMsgsChar() MessageTypeSupport {
	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_Char()}
}

func GetMessageTypeFromStdMsgsColorRGBA() MessageTypeSupport {
	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_ColorRGBA()}
}

func GetMessageTypeFromStdMsgsString() MessageTypeSupport {
	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_String()}
}
