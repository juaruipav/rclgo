package types

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #cgo LDFLAGS: -L/opt/ros/eloquent/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
//// #include "msg_types.h"
//#include <rcl/rcl.h>
//#include <rosidl_generator_c/message_type_support_struct.h>
import "C"
import (
	"unsafe"
)

type RCLRetT int

// Constants used in the framework
const (
	RCL_RET_OK                       = 0
	RCL_RET_ERROR                    = 1
	RCL_RET_TIMEOUT                  = 2
	RCL_RET_BAD_ALLOC                = 10
	RMW_RET_INVALID_ARGUMENT         = 11
	RCL_RET_ALREADY_INIT             = 100
	RCL_RET_NOT_INIT                 = 101
	RCL_RET_MISMATCHED_RMW_ID        = 102
	RCL_RET_TOPIC_NAME_INVALID       = 103
	RCL_RET_SERVICE_NAME_INVALID     = 104
	RCL_RET_UNKNOWN_SUBSTITUTION     = 105
	RCL_RET_NODE_INVALID             = 200
	RCL_RET_NODE_INVALID_NAME        = 201
	RCL_RET_NODE_INVALID_NAMESPACE   = 202
	RCL_RET_PUBLISHER_INVALID        = 300
	RCL_RET_SUBSCRIPTION_INVALID     = 400
	RCL_RET_SUBSCRIPTION_TAKE_FAILED = 401
	RCL_RET_CLIENT_INVALID           = 500
	RCL_RET_CLIENT_TAKE_FAILED       = 501
	RCL_RET_SERVICE_INVALID          = 600
	RCL_RET_SERVICE_TAKE_FAILED      = 601
	RCL_RET_TIMER_INVALID            = 800
	RCL_RET_TIMER_CANCELED           = 801
	RCL_RET_WAIT_SET_INVALID         = 900
	RCL_RET_WAIT_SET_EMPTY           = 901
	RCL_RET_WAIT_SET_FULL            = 902
)

type MessageTypeSupport struct {
	ROSIdlMessageTypeSupport *C.rosidl_message_type_support_t
}

type MessageData struct {
	Data unsafe.Pointer
}

type Message interface {
	GetMessage() MessageTypeSupport
	GetData() MessageData
	InitMessage()
	DestroyMessage()
}

type Context struct {
	RCLContext C.rcl_context_t
}

// func GetMessageTypeFromStdMsgsBool() MessageTypeSupport {
// 	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_Bool()}
// }

// func GetMessageTypeFromStdMsgsByte() MessageTypeSupport {
// 	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_Byte()}
// }

// func GetMessageTypeFromStdMsgsChar() MessageTypeSupport {
// 	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_Char()}
// }

// func GetMessageTypeFromStdMsgsColorRGBA() MessageTypeSupport {
// 	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_ColorRGBA()}
// }
