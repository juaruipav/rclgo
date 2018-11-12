package types

// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include "msg_types.h"
import "C"
import "unsafe"

type MessageTypeSupport struct {
	ROSIdlMessageTypeSupport *C.rosidl_message_type_support_t
}

type Message interface {
	GetMessage() unsafe.Pointer
	SetMessage()
	InitMessage()
	DestroyMessage()
}

type StdMsgsStringMsg struct {
	Data *C.struct_std_msg__msg__String
}

func (msg StdMsgsStringMsg) GetMessage() unsafe.Pointer {
	return unsafe.Pointer(msg.Data)
}

func (msg StdMsgsStringMsg) InitMessage() {
	msg.Data = (*C.struct_std_msg__msg__String)(unsafe.Pointer(C.init_std_msgs_msg_String()))
}

func (msg StdMsgsStringMsg) SetMessage(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	C.rosidl_generator_c__String__assign(nil, cText)
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
