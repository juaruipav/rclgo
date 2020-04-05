package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #cgo LDFLAGS: -L/opt/ros/eloquent/lib -Wl,-rpath=/opt/ros/eloquent/lib -lrcl -lrcutils
// #include "rcl/rcl.h"
import "C"

import (
	"unsafe"
)

//
type ROSIdlMessageTypeSupport C.rosidl_message_type_support_t

//
type RclContext C.rcl_context_t

//
type RclInitOptions C.rcl_init_options_t

//
type RclAllocator C.rcl_allocator_t

//
func GetZeroInitializedContext() RclContext {
	var ctx C.rcl_context_t = C.rcl_get_zero_initialized_context()
	return RclContext(ctx)
}

//
func RclGetZeroInitializedNode() RclNode {
	var zeroNode C.rcl_node_t = C.rcl_get_zero_initialized_node()
	return RclNode(zeroNode)
}

//
func RclGetZeroInitializedInitOptions() RclInitOptions {
	var ret C.rcl_init_options_t = C.rcl_get_zero_initialized_init_options()
	return RclInitOptions(ret)
}

//
func RclGetDefaultAllocator() RclAllocator {
	var ret = C.rcl_get_default_allocator()
	return RclAllocator(ret)
}

//
func RclInitOptionsInit(opt *RclInitOptions, allocator RclAllocator) int {
	var ret C.int32_t = C.rcl_init_options_init(
		(*C.rcl_init_options_t)(opt),
		*(*C.rcl_allocator_t)(&allocator),
	)
	return int(ret)
}

//
func RclInitOptionsFini(opt *RclInitOptions) int {
	var ret C.int32_t = C.rcl_init_options_fini(
		(*C.rcl_init_options_t)(opt),
	)
	return int(ret)
}

//
func RclInit(argc int, argv []string, opt *RclInitOptions, ctx *RclContext) int {

	var arg **C.char = nil
	if len(argv) != 0 {
		cArgv := make([]*C.char, len(argv))
		for i := range argv {
			cstr := C.CString(argv[i])
			defer C.free(unsafe.Pointer(cstr))
			cArgv[i] = cstr
		}
		arg = &cArgv[0]
	}

	var ret = C.rcl_init(
		C.int(argc),
		arg,
		(*C.rcl_init_options_t)(opt),
		(*C.rcl_context_t)(ctx),
	)

	return int(ret)
}

// RclShutdown represents Signal global shutdown of rcl.
func RclShutdown(ctx *RclContext) {
	cCtx := (*C.rcl_context_t)(ctx)
	C.rcl_shutdown(cCtx)
}

//
type RmwMessageInfo C.rmw_message_info_t

//
func RclTake(
	subscription *RclSubscription,
	msg unsafe.Pointer,
	msgInfo *RmwMessageInfo,
) int {
	var ret = C.rcl_take(
		(*C.rcl_subscription_t)(subscription),
		msg,
		(*C.rmw_message_info_t)(msgInfo),
		nil,
	)

	return int(ret)
}
