package types

// #cgo CFLAGS: -I/opt/ros/crystal/include
// #include <rosidl_generator_c/message_type_support_struct.h>
import "C"
import (
	"unsafe"
)

// TypesupportManagerInterface TODO
type TypesupportManagerInterface interface {
	GetHandlerByMSGInterface(MSGInterface) (TypesupportHandlerInterface, error)
	GetHandlerByName(string) (TypesupportHandlerInterface, error)
	setHandler(string, TypesupportHandlerInterface) error
}

// TypesupportHandlerInterface TODO
type TypesupportHandlerInterface interface {
	GetROSTypesupportPointer(MSGInterface) (*C.struct_rosidl_message_type_support_t, error)
	ParseCTypesupport(MSGInterface, CTypesupport) error
	GetCTypesupport(MSGInterface) (unsafe.Pointer, error)
}

// MSGInterface TODO
type MSGInterface interface {
	ParseBytes([]byte) error
	ToString() string
}

// CTypesupport TODO
type CTypesupport C.struct_rosidl_message_type_support_t

// RCLRetT TODO
type RCLRetT int

// Typesupport TODO
type Typesupport struct {
	CTypesupport C.struct_rosidl_message_type_support_t
	Data         []byte
	Size         uint64
	Length       uint64
}

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
