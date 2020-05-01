package rclgo

import (
	"github.com/richardrigby/rclgo/err"
	cwrap "github.com/richardrigby/rclgo/internal"
	"github.com/richardrigby/rclgo/types"
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
	if ret != types.Ok {
		return err.NewErr("RclNodeInit", ret)
	}

	return nil
}

// Fini finalizes an RclNode.
func (n *Node) Fini() error {
	ret := cwrap.RclNodeFini(n.rclNode)
	if ret != types.Ok {
		return err.NewErr("RclNodeFini", ret)
	}

	return nil
}

// IsValid returns `true` if the node is valid, else `false`.
func (n *Node) IsValid() bool {
	ret := cwrap.RclNodeIsValid(n.rclNode)
	return ret
}

// IsValidExceptContext returns true if node is valid, except for the context
// being valid.
func (n *Node) IsValidExceptContext() bool {
	ret := cwrap.RclNodeIsValidExceptContext(n.rclNode)
	return ret
}

// GetName returns the name of the node.
func (n *Node) GetName() string {
	ret := cwrap.RclNodeGetName(n.rclNode)
	return ret
}

// GetNamespace returns the namespace of the node.
func (n *Node) GetNamespace() string {
	ret := cwrap.RclNodeGetNamespace(n.rclNode)
	return ret
}

// GetFullyQualifiedName returns the fully qualified name of the node.
func (n *Node) GetFullyQualifiedName() string {
	ret := cwrap.RclNodeGetFullyQualifiedName(n.rclNode)
	return ret
}

// GetOptions returns the rcl node options.
func (n *Node) GetOptions() NodeOptions {
	opts := cwrap.RclNodeGetOptions(n.rclNode)
	return NodeOptions{opts}
}

// GetDomainID returns the ROS domain ID that the node is using.
func (n *Node) GetDomainID() (uint, error) {
	var domainID uint
	ret := cwrap.RclNodeGetDomainID(n.rclNode, &domainID)
	if ret != types.Ok {
		return domainID, err.NewErr("RclNodeGetDomainID", ret)
	}

	return domainID, nil
}

// AssertLiveliness manually asserts that this node is alive
// (for RMW_QOS_POLICY_LIVELINESS_MANUAL_BY_NODE)
func (n *Node) AssertLiveliness() error {
	ret := cwrap.RclNodeAssertLiveliness(n.rclNode)
	if ret != types.Ok {
		return err.NewErr("RclNodeAssertLiveliness", ret)
	}

	return nil
}

// GetRmwHandle returns the rmw node handle.
// func (n *Node) GetRmwHandle() RmwHandle {
// 	ret := cwrap.RclNodeGetRmwHandle(n.rclNode)
// 	return RmwHandle{ret}
// }

// GetRclInstanceID returns the associated rcl instance id.
func (n *Node) GetRclInstanceID() uint64 {
	var ret uint64 = cwrap.RclNodeGetRclInstanceID(n.rclNode)
	return ret
}

// GetGraphGuardCondition returns a guard condition which is triggered when the
// ROS graph changes.
// func (n *Node) GetGraphGuardCondition() GuardCondition {
// 	guard := cwrap.RclNodeGetGraphGuardCondition(n.rclNode)
// 	return GuardCondition{guard}
// }

// GetLoggerName returns the logger name of the node.
func (n *Node) GetLoggerName() string {
	var ret string = cwrap.RclNodeGetLoggerName(n.rclNode)
	return ret
}
