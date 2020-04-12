package rcl

import (
	cwrap "github.com/richardrigby/rclgo/internal"
)

type Allocator struct {
	Allocator *cwrap.RclAllocator
}

func GetZeroInitializedAllocator() Allocator {
	zeroAllocator := cwrap.RcutilsGetZeroInitializedAllocator()
	return Allocator{&zeroAllocator}
}

func GetDefaultAllocator() Allocator {
	defAllocator := cwrap.RcutilsGetDefaultAllocator()
	return Allocator{&defAllocator}
}
