package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"

//
func RclContextFini(ctx RclContextPtr) int {
	var ret C.int32_t = C.rcl_context_fini(
		(*C.rcl_context_t)(ctx),
	)
	return int(ret)
}

//
func RclContextIsValid(ctx RclContextPtr) bool {
	var ret C.bool = C.rcl_context_is_valid(
		(*C.rcl_context_t)(ctx),
	)
	return bool(ret)
}
