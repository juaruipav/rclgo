package rclgo

import (
	"github.com/richardrigby/rclgo/err"
	cwrap "github.com/richardrigby/rclgo/internal"
	"github.com/richardrigby/rclgo/types"
)

type WaitSet struct {
	rclWaitSet *cwrap.RclWaitSet
}

func NewZeroInitializedWaitSet() WaitSet {
	zeroWaitset := cwrap.RclGetZeroInitializedWaitSet()
	return WaitSet{&zeroWaitset}
}

func (w *WaitSet) WaitSetInit(
	numSubs, numGuards, numTimers, numClients, numServices, numEvents int,
	ctx cwrap.RclContextPtr,
	allo cwrap.RclAllocator,
) error {
	ret := cwrap.RclWaitSetInit(
		w.rclWaitSet,
		numSubs,
		numGuards,
		numTimers,
		numClients,
		numServices,
		numEvents,
		ctx,
		allo,
	)
	if ret != types.RCL_RET_OK {
		return err.NewErr("RclWaitSetInit", ret)
	}

	return nil
}

func (w *WaitSet) Fini() error {
	ret := cwrap.RclWaitSetFini(w.rclWaitSet)
	if ret != types.RCL_RET_OK {
		return err.NewErr("RclWaitSetFini", ret)
	}

	return nil
}

func (w *WaitSet) GetAllocator(allocator *cwrap.RclAllocator) error {
	ret := cwrap.RclWaitSetGetAllocator(w.rclWaitSet, allocator)
	if ret != types.RCL_RET_OK {
		return err.NewErr("RclWaitSetGetAllocator", ret)
	}

	return nil
}

func (w *WaitSet) WaitSetAddsubscription(subscription Subscription) error {
	ret := cwrap.RclWaitSetAddSubscription(w.rclWaitSet, subscription.rclSubscription, nil)
	if ret != types.RCL_RET_OK {
		return err.NewErr("RclWaitSetAddSubscription", ret)
	}

	return nil
}

func (w *WaitSet) WaitSetClearSubscriptions() error {
	ret := cwrap.RclWaitSetClear(w.rclWaitSet)
	if ret != types.RCL_RET_OK {
		return err.NewErr("RclWaitSetClear", ret)
	}

	return nil
}
