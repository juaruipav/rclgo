package node

import (
	cwrap "github.com/richardrigby/rclgo/internal"
	"github.com/richardrigby/rclgo/rcl"
	"github.com/richardrigby/rclgo/types"
)

//
type Node struct {
	RclNode *cwrap.RclNode
}

//
type NodeOptions struct {
	rclNodeOptions *cwrap.RclNodeOptions
}

//
func NewZeroInitializedNode() Node {
	zeroNode := cwrap.RclGetZeroInitializedNode()
	return Node{RclNode: &zeroNode}
}

//
func NewNodeDefaultOptions() NodeOptions {
	defOpts := cwrap.RclNodeGetDefaultOptions()
	return NodeOptions{rclNodeOptions: &defOpts}
}

//
func (n *Node) Init(name string, namespace string, ctx types.Context, nodeOptions NodeOptions) error {
	ret := cwrap.RclNodeInit(
		n.RclNode,
		name,
		namespace,
		ctx.RCLContext,
		nodeOptions.rclNodeOptions,
	)
	if ret != 0 {
		return rcl.NewErr("cwrap.RclNodeInit", ret)
	}

	return nil
}

//
func (n *Node) Fini() error {
	ret := cwrap.RclNodeFini(n.RclNode)
	if ret != 0 {
		return rcl.NewErr("cwrap.RclNodeFini", ret)
	}

	return nil
}
