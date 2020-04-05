package node

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #cgo LDFLAGS: -L/opt/ros/eloquent/lib -Wl,-rpath=/opt/ros/eloquent/lib -lrcl -lrcutils
// #include "rcl/rcl.h"
// #include "rcl/node.h"
import "C"
import (
	"github.com/richardrigby/rclgo/cwrap"
	"github.com/richardrigby/rclgo/types"
)

type Node struct {
	RCLNode *cwrap.RclNode
}

type NodeOptions struct {
	RCLNodeOptions *cwrap.RclNodeOptions
}

func GetZeroInitializedNode() Node {
	zeroNode := cwrap.RclGetZeroInitializedNode()
	return Node{RCLNode: &zeroNode}
}

func GetNodeDefaultOptions() NodeOptions {
	defOpts := cwrap.RclNodeGetDefaultOptions()
	return NodeOptions{RCLNodeOptions: &defOpts}
}

func NodeInit(node Node, name string, namespace string, ctx types.Context, nodeOptions NodeOptions) types.RCLRetT {
	nodeInit := cwrap.RclNodeInit(
		node.RCLNode,
		name,
		namespace,
		&ctx.RCLContext,
		nodeOptions.RCLNodeOptions,
	)
	return types.RCLRetT(
		nodeInit,
	)
}

func NodeFini(node Node) {
	cwrap.RclNodeFini(node.RCLNode)
}
