package timer

// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrcutils
// #include "rcl/rcl.h"
import "C"
import (
	"rclgo/rcl"
	"rclgo/subscription"
	"rclgo/types"
	"unsafe"
)

type WaitSet struct {
	RCLWaitSet *C.rcl_wait_set_t
}

func GetZeroInitializedWaitSet() WaitSet {
	zeroWaitset := C.rcl_get_zero_initialized_wait_set()
	return WaitSet{&zeroWaitset}
}

func WaitSetInit(waitSet WaitSet, numSubs int, numGuards int, numTimers int, numClients int, numServices int, allo Allocator) types.RCLRetT {

	allocatorPtr := (*C.struct_rcutils_allocator_t)(unsafe.Pointer(allo.Allocator))
	retValue := C.rcl_wait_set_init(waitSet.RCLWaitSet,
		C.ulong(numSubs),
		C.ulong(numGuards),
		C.ulong(numTimers),
		C.ulong(numClients),
		C.ulong(numServices),
		*allocatorPtr)
	return types.RCLRetT(retValue)
}

func WaitSetFini(waitSet WaitSet) types.RCLRetT {
	return types.RCLRetT(C.rcl_wait_set_fini(waitSet.RCLWaitSet))
}

func WaitSetGetAllocator(waitset WaitSet, allocator rcl.Allocator) types.RCLRetT {
	allocatorPtr := (*C.struct_rcutils_allocator_t)(unsafe.Pointer(allocator.Allocator))
	return types.RCLRetT(C.rcl_wait_set_get_allocator(waitset.RCLWaitSet, allocatorPtr))
}

func WaitSetAddsubscription(waitset WaitSet, subscription subscription.Subscription) types.RCLRetT {
	subscriptionPtr := (*C.rcl_subscription_t)(unsafe.Pointer(subscription.RCLSubscription))
	return types.RCLRetT(C.rcl_wait_set_add_subscription(waitset.RCLWaitSet, subscriptionPtr))
}

func WaitSetClearSubscriptions(waitset WaitSet) types.RCLRetT {
	return types.RCLRetT(C.rcl_wait_set_clear_subscriptions(waitset.RCLWaitSet))
}
