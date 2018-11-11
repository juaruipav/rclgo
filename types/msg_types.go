package types

// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrosidl_typesupport_c -lstd_msgs__rosidl_typesupport_c
// #include "msg_types.h"
import "C"

type MessageTypeSupport struct {
	ROSIdlMessageTypeSupport *C.rosidl_message_type_support_t
}

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
