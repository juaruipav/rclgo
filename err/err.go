package err

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
