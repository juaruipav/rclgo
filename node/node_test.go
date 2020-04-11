package node

import (
	"fmt"
	"testing"
	"time"

	"github.com/richardrigby/rclgo/cwrap"
	"github.com/richardrigby/rclgo/rcl"
	"github.com/richardrigby/rclgo/types"
)

func TestNodeCreation(t *testing.T) {
	// Initialization
	ctx := types.Context{RCLContext: cwrap.GetZeroInitializedContextPtr()}
	err := rcl.Init(&ctx)
	if err != nil {
		t.Fatalf("rcl.Init failed: %s\n", err)
	}
	fmt.Printf("ctx.RCLContext: %v\n", ctx)
	fmt.Printf("ctx.RCLContext: %v\n", ctx.RCLContext)

	myNode := GetZeroInitializedNode()
	myNodeOpts := GetNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	err = NodeInit(myNode, "fakeNameForNode", "", ctx, myNodeOpts)
	if err != nil {
		t.Fatalf("NodeInit failed: %s\n", err)
	}

	time.Sleep(5 * time.Second) // or runtime.Gosched() or similar per @misterbee

	err = NodeFini(myNode)
	if err != nil {
		t.Fatalf("NodeFini failed: %s\n", err)
	}

	err = rcl.Shutdown(ctx)
	if err != nil {
		t.Fatalf("rcl.Shutdown failed: %s\n", err)
	}
}
