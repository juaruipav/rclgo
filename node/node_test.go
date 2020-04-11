package node_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/richardrigby/rclgo/node"
	"github.com/richardrigby/rclgo/rcl"
	"github.com/richardrigby/rclgo/types"
)

func TestNodeCreation(t *testing.T) {
	// Initialization
	ctx := types.GetZeroInitializedContext()
	err := rcl.Init(&ctx)
	if err != nil {
		t.Fatalf("rcl.Init failed: %s\n", err)
	}

	myNode := node.GetZeroInitializedNode()
	myNodeOpts := node.GetNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	err = node.NodeInit(myNode, "fakeNameForNode", "", ctx, myNodeOpts)
	if err != nil {
		t.Fatalf("NodeInit failed: %s\n", err)
	}

	time.Sleep(5 * time.Second) // or runtime.Gosched() or similar per @misterbee

	err = node.NodeFini(myNode)
	if err != nil {
		t.Fatalf("NodeFini failed: %s\n", err)
	}

	err = rcl.Shutdown(ctx)
	if err != nil {
		t.Fatalf("rcl.Shutdown failed: %s\n", err)
	}
}
