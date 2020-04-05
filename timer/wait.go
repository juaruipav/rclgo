package timer

// #cgo CFLAGS: -I/home/travers/Projects/go/src/github.com/richardrigby/rclgo/ros2-linux/include
// #cgo LDFLAGS: -L/opt/ros/eloquent/lib -Wl,-rpath=/opt/ros/eloquent/lib -lrcl -lrcutils
// #include <rcl/rcl.h>
import "C"
import (
	"github.com/richardrigby/rclgo/cwrap"
	"github.com/richardrigby/rclgo/subscription"
	"github.com/richardrigby/rclgo/types"
)

type WaitSet struct {
	RCLWaitSet *cwrap.RclWaitSet
}

func GetZeroInitializedWaitSet() WaitSet {
	zeroWaitset := cwrap.RclGetZeroInitializedWaitSet()
	return WaitSet{&zeroWaitset}
}

func WaitSetInit(
	waitSet WaitSet,
	numSubs, numGuards, numTimers, numClients, numServices, numEvents int,
	ctx *cwrap.RclContext,
	allo cwrap.RclAllocator,
) types.RCLRetT {

	// allocatorPtr := (*C.struct_rcutils_allocator_t)(unsafe.Pointer(allo.Allocator))
	retValue := cwrap.RclWaitSetInit(
		waitSet.RCLWaitSet,
		numSubs,
		numGuards,
		numTimers,
		numClients,
		numServices,
		numEvents,
		ctx,
		allo,
	)
	return types.RCLRetT(retValue)
}

func WaitSetFini(waitSet WaitSet) types.RCLRetT {
	ret := cwrap.RclWaitSetFini(waitSet.RCLWaitSet)
	return types.RCLRetT(ret)
}

func WaitSetGetAllocator(waitSet WaitSet, allocator *cwrap.RclAllocator) types.RCLRetT {
	// allocatorPtr := (*C.struct_rcutils_allocator_t)(unsafe.Pointer(allocator.Allocator))
	// return types.RCLRetT(C.rcl_wait_set_get_allocator(waitset.RCLWaitSet, allocatorPtr))

	ret := cwrap.RclWaitSetGetAllocator(waitSet.RCLWaitSet, allocator)
	return types.RCLRetT(ret)
}

// Store a pointer to the given subscription in the next empty spot in the set.
/**
 * This function does not guarantee that the subscription is not already in the
 * wait set.
 *
 * Also add the rmw representation to the underlying rmw array and increment
 * the rmw array count.
 *
 * <hr>
 * Attribute          | Adherence
 * ------------------ | -------------
 * Allocates Memory   | Yes
 * Thread-Safe        | No
 * Uses Atomics       | No
 * Lock-Free          | Yes
 *
 * \param[inout] wait_set struct in which the subscription is to be stored
 * \param[in] subscription the subscription to be added to the wait set
 * \param[out] index the index of the added subscription in the storage container.
 *   This parameter is optional and can be set to `NULL` to be ignored.
 * \return `RCL_RET_OK` if added successfully, or
 * \return `RCL_RET_INVALID_ARGUMENT` if any arguments are invalid, or
 * \return `RCL_RET_WAIT_SET_INVALID` if the wait set is zero initialized, or
 * \return `RCL_RET_WAIT_SET_FULL` if the subscription set is full, or
 * \return `RCL_RET_ERROR` if an unspecified error occurs.
 */
func WaitSetAddsubscription(waitSet WaitSet, subscription subscription.Subscription) types.RCLRetT {
	// subscriptionPtr := (*C.rcl_subscription_t)(unsafe.Pointer(subscription.RCLSubscription))
	// return types.RCLRetT(C.rcl_wait_set_add_subscription(waitset.RCLWaitSet, subscriptionPtr, nil))
	ret := cwrap.RclWaitSetAddSubscription(waitSet.RCLWaitSet, subscription.RCLSubscription, nil)
	return types.RCLRetT(ret)
}

func WaitSetClearSubscriptions(waitSet WaitSet) types.RCLRetT {
	ret := cwrap.RclWaitSetClear(waitSet.RCLWaitSet)
	return types.RCLRetT(ret)
}
