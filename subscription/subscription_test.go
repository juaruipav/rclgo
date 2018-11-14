package subscription

import (
	"fmt"
	"os"
	"os/signal"
	"rclgo/node"
	"rclgo/rcl"
	"rclgo/types"
	"syscall"
	"testing"
	"time"
)

func TestSubscriptionshit(t *testing.T) {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	msg := make(chan string, 1)
	go func() {
		// Receive input in a loop
		for {
			var s string
			fmt.Scan(&s)
			// Send what we read over the channel
			msg <- s
		}
	}()
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

	//Creating the msg type
	var myMsg types.StdMsgsString
	myMsg.InitMessage()

loop:
	for {

		retRCL := RCLTake(mySub, myMsg.GetMessage(), myMsg.GetData())

		fmt.Printf("(go) Ret value is %d and %s\n", retRCL, myMsg.GetDataAsString())
		time.Sleep(1000 * time.Millisecond)
		select {
		case <-sigs:
			fmt.Println("Got shutdown, exiting")
			break loop
		case <-msg:
		}
	}

	fmt.Printf("Shutting down!! \n")

	SubscriptionFini(mySub, myNode)
	node.NodeFini(myNode)
	rcl.Shutdown()

}
