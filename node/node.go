package node

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #cgo LDFLAGS: -L/opt/ros/eloquent/lib -lrcl -lrcutils
// #include "rcl/rcl.h"
// #include "rcl/node.h"
import "C"
import (
	"unsafe"

	"github.com/richardrigby/rclgo/types"
)

type Node struct {
	RCLNode *C.rcl_node_t
}

type NodeOptions struct {
	RCLNodeOptions *C.rcl_node_options_t
}

func GetZeroInitializedNode() Node {
	zeroNode := C.rcl_get_zero_initialized_node()
	return Node{&zeroNode}
}

func GetNodeDefaultOptions() NodeOptions {
	defOpts := C.rcl_node_get_default_options()
	return NodeOptions{&defOpts}
}

func NodeInit(node Node, name string, namespace string, ctx types.Context, nodeOptions NodeOptions) types.RCLRetT {

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cNameSpace := C.CString(namespace)
	defer C.free(unsafe.Pointer(cNameSpace))
	return types.RCLRetT(C.rcl_node_init(node.RCLNode, cName, cNameSpace, *C.rcl_context_t(ctx.RCLContext), nodeOptions.RCLNodeOptions))
}

func NodeFini(node Node) {
	C.rcl_node_fini(node.RCLNode)
}
