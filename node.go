package rclgo

import (
	"github.com/richardrigby/rclgo/err"
	cwrap "github.com/richardrigby/rclgo/internal"
)

//
type Node struct {
	rclNode *cwrap.RclNode
}

//
type NodeOptions struct {
	rclNodeOptions *cwrap.RclNodeOptions
}

//
func NewZeroInitializedNode() Node {
	zeroNode := cwrap.RclGetZeroInitializedNode()
	return Node{rclNode: &zeroNode}
}

//
func NewNodeDefaultOptions() NodeOptions {
	defOpts := cwrap.RclNodeGetDefaultOptions()
	return NodeOptions{rclNodeOptions: &defOpts}
}

//
func (n *Node) Init(name string, namespace string, ctx Context, nodeOptions NodeOptions) error {
	ret := cwrap.RclNodeInit(
		n.rclNode,
		name,
		namespace,
		ctx.rclContext,
		nodeOptions.rclNodeOptions,
	)
	if ret != 0 {
		return err.NewErr("RclNodeInit", ret)
	}

	return nil
}

//
func (n *Node) Fini() error {
	ret := cwrap.RclNodeFini(n.rclNode)
	if ret != 0 {
		return err.NewErr("RclNodeFini", ret)
	}

	return nil
}
