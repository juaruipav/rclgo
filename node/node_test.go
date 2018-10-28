package node

import (
	"fmt"
	"rclgo/rcl"
	"testing"
	"time"
)

func TestNodeCreation(t *testing.T) {

	// * TODO for this test: check the retValues from C functions

	// Initialization
	rcl.Init()
	myNode := GetZeroInitializedNode()
	myNodeOpts := GetNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	NodeInit(myNode, "fakeNameForNode", "", myNodeOpts)
	time.Sleep(10 * time.Second) // or runtime.Gosched() or similar per @misterbee
	NodeFini(myNode)
	rcl.Shutdown()

}
