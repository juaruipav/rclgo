package subscription

import (
	"fmt"
	"rclgo/node"
	"rclgo/rcl"
	"rclgo/types"
	"testing"
	"time"
)

func TestSubscriptionshit(t *testing.T) {
	// Initialization
	rcl.Init()
	myNode := node.GetZeroInitializedNode()
	myNodeOpts := node.GetNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	retNode := node.NodeInit(myNode, "fakeNameForNode", "", myNodeOpts)
	fmt.Printf("Ret value for node is %d \n", retNode)
	//Create the subscriptor

	mySub := GetZeroInitializedSubscription()
	mySubOpts := GetSubscriptionDefaultOptions()

	//Creating the type
	msgType := types.GetMessageTypeFromStdMsgsString()

	fmt.Printf("Creating the subscriber! \n")
	retValue2 := SubscriptionInit(mySub, mySubOpts, myNode, "/chatter", msgType)
	fmt.Printf("Ret value from sub init is %d\n", retValue2)
	msgType2 := types.GetMessageTypeFromStdMsgsString()

	for {

		retRCL := RCLTake(mySub, msgType2)
		fmt.Printf("(go) Ret value is %d\n", retRCL)
		time.Sleep(500 * time.Millisecond) // or runtime.Gosched() or similar per @misterbee
	}

	fmt.Printf("Shutting down!! \n")

	SubscriptionFini(mySub, myNode)
	node.NodeFini(myNode)
	rcl.Shutdown()

}
