package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

import (
	"unsafe"
)

//
type RclSubscription C.rcl_subscription_t

//
type RclSubscriptionOptions C.rcl_subscription_options_t

//
func RclGetZeroInitializedSubscription() RclSubscription {
	var ret C.rcl_subscription_t = C.rcl_get_zero_initialized_subscription()
	return RclSubscription(ret)
}

//
func RclSubscriptionInit(
	subscription *RclSubscription,
	node *RclNode,
	typeSupport *ROSIdlMessageTypeSupport,
	topicName string,
	options *RclSubscriptionOptions,
) int {

	cTopicName := C.CString(topicName)
	defer C.free(unsafe.Pointer(cTopicName))

	var ret C.int32_t = C.rcl_subscription_init(
		(*C.rcl_subscription_t)(subscription),
		(*C.struct_rcl_node_t)(node),
		(*C.rosidl_message_type_support_t)(typeSupport),
		cTopicName,
		(*C.rcl_subscription_options_t)(options),
	)

	return int(ret)
}

//
func RclSubscriptionFini(subscription *RclSubscription, node *RclNode) int {
	var ret C.int32_t = C.rcl_subscription_fini(
		(*C.rcl_subscription_t)(subscription),
		(*C.rcl_node_t)(node),
	)
	return int(ret)
}

//
func RclSubscriptionGetDefaultOptions() RclSubscriptionOptions {
	var ret C.rcl_subscription_options_t = C.rcl_subscription_get_default_options()
	return RclSubscriptionOptions(ret)
}
