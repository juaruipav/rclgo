package rcl

// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrcutils
// #include "rcl/rcl.h"
import "C"

type Allocator struct {
	Allocator *C.rcl_allocator_t
}

func GetZeroInitializedAllocator() Allocator {
	zeroAllocator := C.rcutils_get_zero_initialized_allocator()
	return Allocator{&zeroAllocator}
}

func GetDefaultAllocator() Allocator {
	defAllocator := C.rcutils_get_default_allocator()
	return Allocator{&defAllocator}
}
