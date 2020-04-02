package rcl

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #cgo LDFLAGS: -L/opt/ros/eloquent/lib -lrcl -lrcutils
// #include "rcl/rcl.h"
import "C"
import (
	"unsafe"

	"github.com/richardrigby/rclgo/types"
)

//Init represents the global initialization of rcl. This must be called before using any rcl functions.
func Init() types.RCLRetT {

	argv := make([]*C.char, 1)
	argv[0] = C.CString("")
	defer C.free(unsafe.Pointer(argv[0]))
	return types.RCLRetT(C.rcl_init(0, &argv[0], C.rcl_get_zero_initialized_init_options(), C.rcl_get_zero_initialized_context()))
}

func GetZeroInitializedContext() types.Context {
	return types.Context{RCLContext: C.rcl_get_zero_initialized_context()}
}

//Shutdown represents Signal global shutdown of rcl.
func Shutdown(ctx types.Context) {
	C.rcl_shutdown(*C.rcl_context_t(ctx.RCLContext))
}

// //GetInstanceID returns an uint64_t number that is unique for the latest rcl_init call.
// func GetInstanceID() uint {
// 	return uint(C.rcl_get_instance_id())
// }

// //OK returns true if rcl is currently initialized, otherwise false.
// func OK() bool {
// 	return bool(C.rcl_ok())
// }
