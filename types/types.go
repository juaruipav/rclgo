package types

import (
	"unsafe"

	cwrap "github.com/richardrigby/rclgo/internal"
)

// RCLRetT is a wrapper for `rcl_ret_t` that implements the error interface.
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

func (r RCLRetT) Error() string { return r.String() }

func (r RCLRetT) String() string {
	switch r {
	case 0:
		return "OK"
	case 1:
		return "Error"
	case 2:
		return "Timeout"
	case 10:
		return "Bad alloc"
	case 11:
		return "Invalid argument"

	case 100:
		return "Already init"
	case 101:
		return "Not init"
	case 102:
		return "Mismatched RMW ID"
	case 103:
		return "Topic name invalid"
	case 104:
		return "Service name invalid"
	case 105:
		return "Unknown substitution"
	case 200:
		return "Node invalid"
	case 201:
		return "Node invalid name"
	case 202:
		return "Node invalid namespace"
	case 300:
		return "Publisher invalid"
	case 400:
		return "Subscription invalid"
	case 401:
		return "Subscription take failed"
	case 500:
		return "Client invalid"
	case 501:
		return "Client take failed"
	case 600:
		return "Service invalid"
	case 601:
		return "Service take failed"
	case 800:
		return "Timer invalid"
	case 801:
		return "Timer canceled"
	case 900:
		return "Wait set invalid"
	case 901:
		return "Wait set empty"
	case 902:
		return "Wait set full"
	}

	return "Unknown"
}

type MessageTypeSupport struct {
	ROSIdlMessageTypeSupport *cwrap.ROSIdlMessageTypeSupport
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

type StdMsgsBase struct {
	MsgType MessageTypeSupport
	MsgInfo cwrap.RmwMessageInfo
}
