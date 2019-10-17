package node

// #cgo CFLAGS: -I/opt/ros/crystal/include
// #cgo LDFLAGS: -L/opt/ros/crystal/lib -lrcl -lrcutils -L../c-libs -lrclgocontext
// #include "rcl/rcl.h"
// #include "rcl/node.h"
// #include "../c-libs/rclgocontext.h"
import "C"
import (
	"unsafe"

	"../types"
)

// ROS2 Node struct
type Node struct {
	RCLNode *C.rcl_node_t
}

// Node Options to initialize the node.
type NodeOptions struct {
	RCLNodeOptions *C.rcl_node_options_t
}

// GetZeroInitializedNode obtains a zero initialized node struct.
func GetZeroInitializedNode() Node {
	zeroNode := C.rcl_get_zero_initialized_node()
	return Node{&zeroNode}
}

// GetNodeDefaultOptions obtains a default initialized node options.
func GetNodeDefaultOptions() NodeOptions {
	defOpts := C.rcl_node_get_default_options()
	return NodeOptions{&defOpts}
}

// NodeInit set the node parameters and associates the Context object to the node object.
func NodeInit(node Node, name string, namespace string, nodeOptions NodeOptions, contextObjectKey uint64) types.RCLRetT {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cNameSpace := C.CString(namespace)
	defer C.free(unsafe.Pointer(cNameSpace))

	return types.RCLRetT(C.rcl_node_init(node.RCLNode, cName, cNameSpace, C.getContextPointerByKey(C.uint64_t(contextObjectKey)), nodeOptions.RCLNodeOptions))
}

// NodeFini erase the node information and change the context object state associated with the node object.
func NodeFini(node Node) {
	C.rcl_node_fini(node.RCLNode)
}
