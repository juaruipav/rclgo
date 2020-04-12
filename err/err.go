package err

import (
	"github.com/richardrigby/rclgo/internal"
	"github.com/richardrigby/rclgo/types"
)

// Err implements the Error interface for rcl errors.
type Err struct {
	msg   string
	ret   types.RCLRetT
	state *cwrap.RcutilsErrorState
}

// NewErr gets the error state from rcutils and wraps it in an Err.
// Intended for internal use only.
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

// Unwrap returns the underlying rcl error code.
func (e Err) Unwrap() error { return e.ret }
