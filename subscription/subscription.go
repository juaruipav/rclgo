package subscription

// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrcutils
// #include <rosidl_generator_c/message_type_support_struct.h>
// #include "rcl/rcl.h"
// #include <rcl/error_handling.h>
// #include "std_msgs/msg/string.h"
//#define ZERO_ALLOCATE(s) \
//  rcl_get_default_allocator().zero_allocate(s, 1, rcl_get_default_allocator().state)
// int my_rcl_take (const rcl_subscription_t * subscription, rosidl_message_type_support_t* msg){
//		void * msg2 = ZERO_ALLOCATE(sizeof(std_msgs__msg__String));
//		int retValue = rcl_take(subscription, msg2, NULL);
//		return retValue;
//}
import "C"
import (
	"rclgo/node"
	"rclgo/types"
	"unsafe"
)

type Subscription struct {
	RCLSubscription *C.rcl_subscription_t
}

type SubscriptionOptions struct {
	RCLSubscriptionOptions *C.rcl_subscription_options_t
}

func GetZeroInitializedSubscription() Subscription {
	zeroSubscription := C.rcl_get_zero_initialized_subscription()
	return Subscription{&zeroSubscription}
}

func GetSubscriptionDefaultOptions() SubscriptionOptions {
	defOpts := C.rcl_subscription_get_default_options()
	return SubscriptionOptions{&defOpts}
}

func SubscriptionInit(subscription Subscription, subscriptionOptions SubscriptionOptions, node node.Node, topicName string, msg types.MessageTypeSupport) types.RCLRetT {

	tName := C.CString(topicName)
	defer C.free(unsafe.Pointer(tName))

	return types.RCLRetT(C.rcl_subscription_init(subscription.RCLSubscription,
		(*C.struct_rcl_node_t)(unsafe.Pointer(node.RCLNode)),
		(*C.rosidl_message_type_support_t)(unsafe.Pointer(msg.ROSIdlMessageTypeSupport)),
		tName,
		subscriptionOptions.RCLSubscriptionOptions))

}

func SubscriptionFini(subscription Subscription, node node.Node) types.RCLRetT {

	return types.RCLRetT(C.rcl_subscription_fini(subscription.RCLSubscription,
		(*C.struct_rcl_node_t)(unsafe.Pointer(node.RCLNode))))

}

func RCLTake(subscription Subscription, msg types.MessageTypeSupport) types.RCLRetT {

	// myType := (*C.rosidl_message_type_support_t)(unsafe.Pointer(msg.ROSIdlMessageTypeSupport))

	// cdata := C.malloc(C.size_t(unsafe.Sizeof(*myType)))
	// defer C.free(cdata)

	return types.RCLRetT(C.my_rcl_take(subscription.RCLSubscription,
		(*C.rosidl_message_type_support_t)(unsafe.Pointer(msg.ROSIdlMessageTypeSupport)),
	))

	//return types.RCLRetT(C.rcl_take(subscription.RCLSubscription,
	//	unsafe.Pointer(msg.ROSIdlMessageTypeSupport), nil))
}
