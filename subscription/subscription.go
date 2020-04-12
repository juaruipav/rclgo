package subscription

import (
	cwrap "github.com/richardrigby/rclgo/internal"
	"github.com/richardrigby/rclgo/node"
	"github.com/richardrigby/rclgo/rcl"
	"github.com/richardrigby/rclgo/types"
)

type Subscription struct {
	RclSubscription *cwrap.RclSubscription
}

type SubscriptionOptions struct {
	RclSubscriptionOptions *cwrap.RclSubscriptionOptions
}

func NewZeroInitializedSubscription() Subscription {
	zeroSubscription := cwrap.RclGetZeroInitializedSubscription()
	return Subscription{&zeroSubscription}
}

func NewSubscriptionDefaultOptions() SubscriptionOptions {
	defOpts := cwrap.RclSubscriptionGetDefaultOptions()
	return SubscriptionOptions{&defOpts}
}

func (s *Subscription) Init(
	subscriptionOptions SubscriptionOptions,
	node node.Node,
	topicName string,
	msg types.MessageTypeSupport,
) error {

	ret := cwrap.RclSubscriptionInit(
		s.RclSubscription,
		node.RclNode,
		msg.ROSIdlMessageTypeSupport,
		topicName,
		subscriptionOptions.RclSubscriptionOptions,
	)

	if ret != types.RCL_RET_OK {
		return rcl.NewErr("RclSubscriptionInit", ret)
	}

	return nil
}

func (s *Subscription) SubscriptionFini(node node.Node) error {
	ret := cwrap.RclSubscriptionFini(
		s.RclSubscription,
		node.RclNode,
	)

	if ret != types.RCL_RET_OK {
		return rcl.NewErr("RclSubscriptionFini", ret)
	}

	return nil
}

//
func (s *Subscription) TakeMessage(msg *cwrap.RmwMessageInfo, data types.MessageData) error {
	if msg == nil || s.RclSubscription == nil || data.Data == nil {
		return rcl.NewErr("nil", types.RCL_RET_ERROR)
	}

	ret := cwrap.RclTake(s.RclSubscription, data.Data, msg)

	if ret != types.RCL_RET_OK {
		return rcl.NewErr("MyRclTake %s", ret)
	}

	return nil
}
