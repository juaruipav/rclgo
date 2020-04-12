package rclgo

import (
	"github.com/richardrigby/rclgo/err"
	"github.com/richardrigby/rclgo/internal"
	"github.com/richardrigby/rclgo/types"
)

// Context encapsulates the non-global state of an init/shutdown cycle.
type Context struct {
	rclContext cwrap.RclContextPtr
}

// NewZeroInitializedContext returns a zero initialization context object.
func NewZeroInitializedContext() Context {
	ctxPtr := cwrap.GetZeroInitializedContextPtr()
	return Context{rclContext: ctxPtr}
}

// Init finalizes a context.
func (ctx *Context) Init() error {
	var ret int

	var opts = cwrap.RclGetZeroInitializedInitOptions()
	alloc := cwrap.RclGetDefaultAllocator()

	ret = cwrap.RclInitOptionsInit(&opts, alloc)
	if ret != types.Ok {
		return err.NewErr("RclInitOptionsInit", ret)
	}

	ret = cwrap.RclInit(
		0,
		[]string{},
		&opts,
		ctx.rclContext,
	)
	if ret != types.Ok {
		return err.NewErr("RclInit", ret)
	}

	ret = cwrap.RclInitOptionsFini(&opts)
	if ret != types.Ok {
		return err.NewErr("RclInitOptionsFini", ret)
	}

	return nil
}

// Fini finalizes a context.
func (ctx *Context) Fini() error {
	ret := cwrap.RclContextFini(ctx.rclContext)
	if ret != types.Ok {
		return err.NewErr("RclContextFini", ret)
	}

	return nil
}

// Shutdown shuts down the context.
func (ctx *Context) Shutdown() error {
	ret := cwrap.RclShutdown(ctx.rclContext)
	if ret != types.Ok {
		return err.NewErr("RclShutdown", ret)
	}

	return nil
}

// IsValid returns `true` if the context is currently valid, otherwise `false`.
func (ctx *Context) IsValid() bool {
	ret := cwrap.RclContextIsValid(ctx.rclContext)
	return bool(ret)
}
