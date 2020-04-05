package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #cgo LDFLAGS: -L/opt/ros/eloquent/lib -Wl,-rpath=/opt/ros/eloquent/lib -lrcl -lrcutils
// #include "rcl/rcl.h"
import "C"

//
type RclWaitSet C.rcl_wait_set_t

//
func RclGetZeroInitializedWaitSet() RclWaitSet {
	var waitSet C.rcl_wait_set_t = C.rcl_get_zero_initialized_wait_set()
	return RclWaitSet(waitSet)
}

//
func RclWaitSetInit(
	waitSet *RclWaitSet,
	numSubs, numGuards, numTimers, numClients, numServices, numEvents int,
	ctx *RclContext,
	allo RclAllocator,
) int {
	var ret C.int32_t = C.rcl_wait_set_init(
		(*C.rcl_wait_set_t)(waitSet),
		C.ulong(numSubs),
		C.ulong(numGuards),
		C.ulong(numTimers),
		C.ulong(numClients),
		C.ulong(numServices),
		C.ulong(numEvents),
		(*C.rcl_context_t)(ctx),
		C.rcl_allocator_t(allo),
	)

	return int(ret)
}

//
func RclWaitSetFini(waitSet *RclWaitSet) int {
	cWaitSet := (*C.rcl_wait_set_t)(waitSet)
	var ret C.int32_t = C.rcl_wait_set_fini(cWaitSet)
	return int(ret)
}

//
func RclWaitSetGetAllocator(waitSet *RclWaitSet, allocator *RclAllocator) int {
	var ret C.int32_t = C.rcl_wait_set_get_allocator(
		(*C.rcl_wait_set_t)(waitSet),
		(*C.rcl_allocator_t)(allocator),
	)
	return int(ret)
}

func RclWaitSetAddSubscription(waitSet *RclWaitSet, subscription *RclSubscription, index *uint64) int {
	ret := C.rcl_wait_set_add_subscription(
		(*C.rcl_wait_set_t)(waitSet),
		(*C.rcl_subscription_t)(subscription),
		(*C.ulong)(index),
	)
	return int(ret)
}

func RclWaitSetClear(waitSet *RclWaitSet) int {
	ret := C.rcl_wait_set_clear(
		(*C.rcl_wait_set_t)(waitSet),
	)
	return int(ret)
}
