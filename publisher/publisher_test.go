package publisher

import (
	"fmt"
	"rclgo/node"
	"rclgo/rcl"
	"rclgo/types"
	"testing"
	"time"
)

func TestPublishershit(t *testing.T) {

	// Initialization
	rcl.Init()
	myNode := node.GetZeroInitializedNode()
	myNodeOpts := node.GetNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	retNode := node.NodeInit(myNode, "fakeNameForNode", "", myNodeOpts)
	fmt.Printf("Ret value for node is %d \n", retNode)
	//Create the subscriptor

	myPub := GetZeroInitializedPublisher()
	myPubOpts := GetPublisherDefaultOptions()
	msgType := types.GetMessageTypeFromStdMsgsString()

	fmt.Printf("Creating the publisher! \n")
	retValue2 := PublisherInit(myPub, myPubOpts, myNode, "/chatter", msgType)
	fmt.Printf("Ret value from pub init is %d\n", retValue2)

	//Creating the type
	var myMsg types.StdMsgsStringMsg
	myMsg.InitMessage()
	myMsg.SetText("Me voy al gym!")

	retRCL := Publish(myPub, myMsg.GetMessage(), myMsg.GetData())

	fmt.Printf("(Publish) Ret value is %d\n", retRCL)
	time.Sleep(1000 * time.Millisecond) // or runtime.Gosched() or similar per @misterbee

	fmt.Printf("Shutting down!! \n")

	// SubscriptionFisni(mySub, myNode)
	node.NodeFini(myNode)
	rcl.Shutdown()

}
