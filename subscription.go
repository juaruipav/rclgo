package rclgo

import (
	"github.com/richardrigby/rclgo/err"
	cwrap "github.com/richardrigby/rclgo/internal"
	"github.com/richardrigby/rclgo/types"
)

type Subscription struct {
	rclSubscription *cwrap.RclSubscription
}

type SubscriptionOptions struct {
	rclSubscriptionOptions *cwrap.RclSubscriptionOptions
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
	node Node,
	topicName string,
	msg types.MessageTypeSupport,
) error {

	ret := cwrap.RclSubscriptionInit(
		s.rclSubscription,
		node.rclNode,
		msg.ROSIdlMessageTypeSupport,
		topicName,
		subscriptionOptions.rclSubscriptionOptions,
	)

	if ret != types.RCL_RET_OK {
		return err.NewErr("RclSubscriptionInit", ret)
	}

	return nil
}

func (s *Subscription) SubscriptionFini(node Node) error {
	ret := cwrap.RclSubscriptionFini(
		s.rclSubscription,
		node.rclNode,
	)

	if ret != types.RCL_RET_OK {
		return err.NewErr("RclSubscriptionFini", ret)
	}

	return nil
}

//
func (s *Subscription) TakeMessage(msg *cwrap.RmwMessageInfo, data types.MessageData) error {
	if msg == nil || s.rclSubscription == nil || data.Data == nil {
		return err.NewErr("nil", types.RCL_RET_ERROR)
	}

	ret := cwrap.RclTake(s.rclSubscription, data.Data, msg)

	if ret != types.RCL_RET_OK {
		return err.NewErr("RclTake", ret)
	}

	return nil
}
