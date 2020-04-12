package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
// #include <stdlib.h>
// // This is needed to avoid `panic: runtime error: cgo argument has Go pointer
// // to Go pointer` when passing an `rcl_node_t*` to a C function. `rcl_node_t`
// // contains a pointer to an `rcl_context_t` which must therefore be a C
// // pointer when calling C functions from Go. Remember to free this memory
// // later.
// rcl_context_t* get_zero_context_ptr() {
// 	rcl_context_t* ptr = (rcl_context_t*)malloc(sizeof(rcl_context_t));
// 	*ptr = rcl_get_zero_initialized_context();
// }
//
// void free_context(rcl_context_t *ctx) {
// 	free(ctx);
// }
import "C"

import "unsafe"

//
type ROSIdlMessageTypeSupport C.rosidl_message_type_support_t

//
type RclContextPtr *C.rcl_context_t

//
type RclInitOptions C.rcl_init_options_t

//
type RclAllocator C.rcl_allocator_t

//
func GetZeroInitializedContextPtr() RclContextPtr {
	var ctxPtr (*C.rcl_context_t) = C.get_zero_context_ptr()
	return RclContextPtr(ctxPtr)
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
		(C.rcl_allocator_t)(allocator),
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
func RclInit(argc int, argv []string, opt *RclInitOptions, ctx RclContextPtr) int {

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
func RclShutdown(ctx RclContextPtr) int {
	cCtx := (*C.rcl_context_t)(ctx)
	ret := C.rcl_shutdown(cCtx)
	C.free_context(cCtx)
	return int(ret)
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
