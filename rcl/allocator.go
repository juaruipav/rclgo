package rcl

// #include "rcl/rcl.h"
import "C"

// Allocator to use on rcl functions. It used malloc(), free() and realloc() to alloc memory.
// allocate -> Allocate memory, given a size and the state pointer.
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
