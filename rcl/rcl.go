package rcl

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #cgo LDFLAGS: -L/opt/ros/eloquent/lib -Wl,-rpath=/opt/ros/eloquent/lib -lrcl -lrcutils
// #include "rcl/rcl.h"
import "C"
import (
	"fmt"

	"github.com/richardrigby/rclgo/cwrap"
	"github.com/richardrigby/rclgo/types"
)

//Init represents the global initialization of rcl. This must be called before using any rcl functions.
func Init(ctx types.Context) types.RCLRetT {

	var opts = cwrap.RclGetZeroInitializedInitOptions()
	alloc := cwrap.RclGetDefaultAllocator()
	suc := cwrap.RclInitOptionsInit(&opts, alloc)
	if suc != 0 {
		fmt.Println("suc")
		return types.RCLRetT(suc)
	}
	suc2 := cwrap.RclInitOptionsFini(&opts)
	if suc2 != 0 {
		fmt.Println("suc2")
		return types.RCLRetT(suc2)
	}

	ret := cwrap.RclInit(
		0,
		[]string{},
		&opts,
		&ctx.RCLContext,
	)
	return types.RCLRetT(ret)
}

//Shutdown represents Signal global shutdown of rcl.
func Shutdown(ctx types.Context) {
	cwrap.RclShutdown(&ctx.RCLContext)
}

// //GetInstanceID returns an uint64_t number that is unique for the latest rcl_init call.
// func GetInstanceID() uint {
// 	return uint(C.rcl_get_instance_id())
// }

// //OK returns true if rcl is currently initialized, otherwise false.
// func OK() bool {
// 	return bool(C.rcl_ok())
// }
