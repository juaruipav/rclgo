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

	// * TODO for this test: check the retValues from C functions

	// Initialization
	ctx := types.Context{RCLContext: cwrap.GetZeroInitializedContext()}
	suc := rcl.Init(ctx)
	if suc != 0 {
		t.Fatalf("rcl.Init failed: %s\n", suc)
	}
	fmt.Println(ctx)
	myNode := GetZeroInitializedNode()
	fmt.Println(myNode)
	myNodeOpts := GetNodeDefaultOptions()
	fmt.Println(myNodeOpts)

	fmt.Printf("Creating the node! \n")
	NodeInit(myNode, "fakeNameForNode", "", ctx, myNodeOpts)
	time.Sleep(10 * time.Second) // or runtime.Gosched() or similar per @misterbee
	NodeFini(myNode)
	rcl.Shutdown(ctx)
}
