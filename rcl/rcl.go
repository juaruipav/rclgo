package rcl

// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrcutils
// #include "rcl/rcl.h"
import "C"
import (
	"rclgo/types"
	"unsafe"
)

//Init represents the global initialization of rcl. This must be called before using any rcl functions.
func Init() types.RCLRetT {

	argv := make([]*C.char, 1)
	argv[0] = C.CString("")
	defer C.free(unsafe.Pointer(argv[0]))
	return types.RCLRetT(C.rcl_init(0, &argv[0], C.rcl_get_default_allocator()))
}

//Shutdown represents Signal global shutdown of rcl.
func Shutdown() {
	C.rcl_shutdown()
}

//GetInstanceID returns an uint64_t number that is unique for the latest rcl_init call.
func GetInstanceID() uint {
	return uint(C.rcl_get_instance_id())
}

//OK returns true if rcl is currently initialized, otherwise false.
func OK() bool {
	return bool(C.rcl_ok())
}
