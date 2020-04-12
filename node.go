package rclgo

import (
	"github.com/richardrigby/rclgo/err"
	cwrap "github.com/richardrigby/rclgo/internal"
)

// Node is a structure that encapsulates a ROS Node.
type Node struct {
	rclNode *cwrap.RclNode
}

// NodeOptions is a structure that encapsulates the options for creating an
// RclNode.
type NodeOptions struct {
	rclNodeOptions *cwrap.RclNodeOptions
}

// NewZeroInitializedNode returns an RclNode with members initialized to `NULL`.
func NewZeroInitializedNode() Node {
	zeroNode := cwrap.RclGetZeroInitializedNode()
	return Node{rclNode: &zeroNode}
}

// NewNodeDefaultOptions returns the default node options in a RclNodeOptions.
func NewNodeDefaultOptions() NodeOptions {
	defOpts := cwrap.RclNodeGetDefaultOptions()
	return NodeOptions{rclNodeOptions: &defOpts}
}

// Init initialize a ROS node.
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

// Fini finalizes an RclNode.
func (n *Node) Fini() error {
	ret := cwrap.RclNodeFini(n.rclNode)
	if ret != 0 {
		return err.NewErr("RclNodeFini", ret)
	}

	return nil
}
