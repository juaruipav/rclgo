package rclgo

import (
	cwrap "github.com/richardrigby/rclgo/internal"
)

type Allocator struct {
	rclAllocator *cwrap.RclAllocator
}

func NewZeroInitializedAllocator() Allocator {
	zeroAllocator := cwrap.RcutilsGetZeroInitializedAllocator()
	return Allocator{&zeroAllocator}
}

func NewDefaultAllocator() Allocator {
	defAllocator := cwrap.RcutilsGetDefaultAllocator()
	return Allocator{&defAllocator}
}
