package node

import (
	"fmt"
	"log"
	"testing"
	"time"

	"../rcl"
)

func TestNodeCreation(t *testing.T) {
	// Initialization
	var contextObject rcl.RCLContext

	// Get initialized context object.
	contextObject.GetInitializedContext()

	// Initialize contextObject and rcl. This function have to be executed always.
	result := rcl.RCLInit(contextObject.ContextKey)
	if result != 0 {
		log.Fatal(fmt.Errorf("Error initializing the context object!"))
	}

	// Obtain initialized node struct and options.
	myNode := GetZeroInitializedNode()
	myNodeOpts := GetNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	// Initialize node object and associate context object to it.
	result = NodeInit(myNode, "fakeNameForNode", "", myNodeOpts, contextObject.ContextKey)
	if result != 0 {
		log.Fatal(fmt.Errorf("Error initializing the node object!"))
	}

	time.Sleep(3 * time.Second)

	// Shutdown node object.
	NodeFini(myNode)

	// Shutdown context object.
	rcl.RCLShutdown(contextObject.ContextKey)
}
