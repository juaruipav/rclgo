package rcl

// #cgo CFLAGS: -I/opt/ros/crystal/include
// #cgo LDFLAGS: -L/opt/ros/crystal/lib -lrcl -lrcutils -L../c-libs -lrclgocontext
// #include "rcl/rcl.h"
// #include "../c-libs/rclgocontext.h"
import "C"
import (
	"unsafe"

	"../types"
)

// ROS2 Context struct
type RCLContext struct {
	ContextKey uint64
}

// RCLInit represents the global initialization of rcl. This must be called before using any rcl functions.
// It changes the context object state in order to be associated to a node on node initialization.
func RCLInit(contextObjectKey uint64) types.RCLRetT {
	argv := make([]*C.char, 1)
	argv[0] = C.CString("")
	defer C.free(unsafe.Pointer(argv[0]))

	var cOptions C.rcl_init_options_t
	cOptions = C.rcl_get_zero_initialized_init_options()

	C.rcl_init_options_init(&cOptions, C.rcl_get_default_allocator())

	return types.RCLRetT(C.rcl_init(1, &argv[0], &cOptions, C.getContextPointerByKey(C.uint64_t(contextObjectKey))))
}

// RCLShutdown represents Signal global shutdown of rcl. Also it calls a function to free the memory allocated in
// C code at GetInitializedContext() function.
func RCLShutdown(contextObjectKey uint64) {
	// Change state of context object to shutdown.
	C.rcl_shutdown(C.getContextPointerByKey(C.uint64_t(contextObjectKey)))
	// Free the memory previously allocated with C function GetInitializedContext().
	defer FreeAllocatedMemoryToContext(contextObjectKey)
}

// GetInitializedContext obtain a pointer to a zero initialized context object. This object is stored in memory
// Allocated with C code on function C.getInitializedRCLContext().
func (contextObject *RCLContext) GetInitializedContext() {
	// getInitializedRCLContext() returns a key (string) to the new rcl context. Convert string to GoString before return.
	contextObject.ContextKey = uint64(C.getInitializedRCLContext())
}

// FreeAllocatedMemoryToContext is used to free the memory previously allocated to the context object.
// This function is used in RCLShutdown().
func FreeAllocatedMemoryToContext(contextObjectKey uint64) {
	C.freeAllocatedMemoryToContext(C.uint64_t(contextObjectKey))
}
