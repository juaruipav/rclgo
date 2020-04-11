package rcl

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #cgo LDFLAGS: -L/opt/ros/eloquent/lib -Wl,-rpath=/opt/ros/eloquent/lib -lrcl -lrcutils
// #include "rcl/rcl.h"
import "C"
import (
	"github.com/richardrigby/rclgo/internal"
	"github.com/richardrigby/rclgo/types"
)

//
type Err struct {
	msg   string
	ret   types.RCLRetT
	state *cwrap.RcutilsErrorState
}

//
func NewErr(msg string, ret int) Err {
	return Err{
		msg:   msg,
		ret:   types.RCLRetT(ret),
		state: cwrap.RcutilsGetErrorState(),
	}
}

func (e Err) Error() string {
	if e.ret == types.RCL_RET_OK {
		return e.msg
	}
	return e.msg + ": " + e.ret.String() + "\n" + e.state.Error()
}

//
func (e Err) Unwrap() error { return e.ret }

// Init represents the global initialization of rcl. This must be called before using any rcl functions.
func Init(ctx *types.Context) error {
	var ret int

	var opts = cwrap.RclGetZeroInitializedInitOptions()
	alloc := cwrap.RclGetDefaultAllocator()

	ret = cwrap.RclInitOptionsInit(&opts, alloc)
	if ret != 0 {
		return Err{"cwrap.RclInitOptionsInit", types.RCLRetT(ret), cwrap.RcutilsGetErrorState()}
	}

	ret = cwrap.RclInit(
		0,
		[]string{},
		&opts,
		ctx.RCLContext,
	)
	if ret != 0 {
		return Err{"cwrap.RclInit", types.RCLRetT(ret), cwrap.RcutilsGetErrorState()}
	}

	ret = cwrap.RclInitOptionsFini(&opts)
	if ret != 0 {
		return Err{"cwrap.RclInitOptionsFini", types.RCLRetT(ret), cwrap.RcutilsGetErrorState()}
	}

	return nil
}

// Shutdown represents Signal global shutdown of rcl.
func Shutdown(ctx types.Context) error {
	ret := cwrap.RclShutdown(ctx.RCLContext)
	if ret != 0 {
		return Err{"cwrap.RclShutdown", types.RCLRetT(ret), cwrap.RcutilsGetErrorState()}
	}

	return nil
}
