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
	ctx := types.NewZeroInitializedContext()
	err := rcl.Init(&ctx)
	if err != nil {
		t.Fatalf("rcl.Init failed: %s\n", err)
	}

	myNode := node.NewZeroInitializedNode()
	myNodeOpts := node.NewNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	err = myNode.Init("fakeNameForNode", "", ctx, myNodeOpts)
	if err != nil {
		t.Fatalf("NodeInit failed: %s\n", err)
	}

	time.Sleep(5 * time.Second) // or runtime.Gosched() or similar per @misterbee

	err = myNode.Fini()
	if err != nil {
		t.Fatalf("NodeFini failed: %s\n", err)
	}

	err = rcl.Shutdown(ctx)
	if err != nil {
		t.Fatalf("rcl.Shutdown failed: %s\n", err)
	}
}
