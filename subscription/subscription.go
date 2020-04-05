package subscription

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #cgo LDFLAGS: -L/opt/ros/eloquent/lib -Wl,-rpath=/opt/ros/eloquent/lib -lrcl -lrcutils
// #include <rosidl_generator_c/message_type_support_struct.h>
// #include "rcl/rcl.h"
// #include <rcl/error_handling.h>
// #include "std_msgs/msg/string.h"
//#define ZERO_ALLOCATE(s) \
//  rcl_get_default_allocator().zero_allocate(s, 1, rcl_get_default_allocator().state)
// int my_rcl_take (const rcl_subscription_t * subscription, rosidl_message_type_support_t* msg,void * data){
//		if(msg == NULL || subscription == NULL || data == NULL)
//			return 1;
//		int retValue = rcl_take(subscription, msg, data, NULL);
//		return retValue;
//}
import "C"
import (
	"unsafe"

	"github.com/richardrigby/rclgo/cwrap"
	"github.com/richardrigby/rclgo/node"
	"github.com/richardrigby/rclgo/types"
)

type Subscription struct {
	RCLSubscription *cwrap.RclSubscription
}

type SubscriptionOptions struct {
	RCLSubscriptionOptions *cwrap.RclSubscriptionOptions
}

func GetZeroInitializedSubscription() Subscription {
	// zeroSubscription := C.rcl_get_zero_initialized_subscription()
	zeroSubscription := cwrap.RclGetZeroInitializedSubscription()
	return Subscription{&zeroSubscription}
}

func GetSubscriptionDefaultOptions() SubscriptionOptions {
	defOpts := cwrap.RclSubscriptionGetDefaultOptions()
	return SubscriptionOptions{&defOpts}
}

func SubscriptionInit(
	subscription Subscription,
	subscriptionOptions SubscriptionOptions,
	node node.Node,
	topicName string,
	msg types.MessageTypeSupport,
) types.RCLRetT {

	ret := cwrap.RclSubscriptionInit(
		subscription.RCLSubscription,
		node.RCLNode,
		msg.ROSIdlMessageTypeSupport,
		topicName,
		subscriptionOptions.RCLSubscriptionOptions,
	)

	return types.RCLRetT(ret)

}

func SubscriptionFini(subscription Subscription, node node.Node) types.RCLRetT {
	ret := cwrap.RclSubscriptionFini(
		subscription.RCLSubscription,
		node.RCLNode,
	)
	return types.RCLRetT(ret)
}

func MyRclTake(
	subscription *cwrap.RclSubscription,
	msg *cwrap.RmwMessageInfo,
	data unsafe.Pointer,
) int {
	if msg == nil || subscription == nil || data == nil {
		return 1
	}

	ret := cwrap.RclTake(subscription, data, msg)
	return ret
}

func TakeMessage(subscription Subscription, msg *cwrap.RmwMessageInfo, data types.MessageData) types.RCLRetT {
	ret := MyRclTake(
		subscription.RCLSubscription,
		msg,
		data.Data,
	)
	return types.RCLRetT(ret)
}
