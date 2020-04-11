package subscription

import (
	cwrap "github.com/richardrigby/rclgo/internal"
	"github.com/richardrigby/rclgo/node"
	"github.com/richardrigby/rclgo/rcl"
	"github.com/richardrigby/rclgo/types"
)

type Subscription struct {
	RCLSubscription *cwrap.RclSubscription
}

type SubscriptionOptions struct {
	RCLSubscriptionOptions *cwrap.RclSubscriptionOptions
}

func GetZeroInitializedSubscription() Subscription {
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
) error {

	ret := cwrap.RclSubscriptionInit(
		subscription.RCLSubscription,
		node.RCLNode,
		msg.ROSIdlMessageTypeSupport,
		topicName,
		subscriptionOptions.RCLSubscriptionOptions,
	)

	if ret != types.RCL_RET_OK {
		return rcl.NewErr("RclSubscriptionInit", ret)
	}

	return nil
}

func SubscriptionFini(subscription Subscription, node node.Node) error {
	ret := cwrap.RclSubscriptionFini(
		subscription.RCLSubscription,
		node.RCLNode,
	)

	if ret != types.RCL_RET_OK {
		return rcl.NewErr("RclSubscriptionFini", ret)
	}

	return nil
}

//
func TakeMessage(subscription Subscription, msg *cwrap.RmwMessageInfo, data types.MessageData) error {
	if msg == nil || subscription.RCLSubscription == nil || data.Data == nil {
		return rcl.NewErr("nil", types.RCL_RET_ERROR)
	}

	ret := cwrap.RclTake(subscription.RCLSubscription, data.Data, msg)

	if ret != types.RCL_RET_OK {
		return rcl.NewErr("MyRclTake %s", ret)
	}

	return nil
}
