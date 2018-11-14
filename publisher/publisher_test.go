package publisher

import (
	"fmt"
	"os"
	"os/signal"
	"rclgo/node"
	"rclgo/rcl"
	"rclgo/types"
	"strconv"
	"syscall"
	"testing"
	"time"
)

func TestPublishershit(t *testing.T) {

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

	//Create the publisher
	myPub := GetZeroInitializedPublisher()
	myPubOpts := GetPublisherDefaultOptions()

	//Creating the msg type
	var myMsg types.StdMsgsString
	myMsg.InitMessage()
	myMsg.SetText("Me voy al gym!")

	//Initializing the publisher
	fmt.Printf("Creating the publisher! \n")
	retValue2 := PublisherInit(myPub, myPubOpts, myNode, "/chatter", myMsg.GetMessage())
	fmt.Printf("Ret value from pub init is %d\n", retValue2)

	index := 0

loop:
	for {
		//Update my msg
		myMsg.SetText("My msg #" + strconv.Itoa(index))
		//Publish the message
		retRCL := Publish(myPub, myMsg.GetMessage(), myMsg.GetData())

		fmt.Printf("(Publish) Ret value is %d\n", retRCL)
		time.Sleep(1000 * time.Millisecond)
		index++
		select {
		case <-sigs:
			fmt.Println("Got shutdown, exiting")
			// Break out of the outer for statement and end the program
			break loop
		case <-msg:

		}
	}

	fmt.Printf("Shutting down!! \n")

	// SubscriptionFisni(mySub, myNode)
	node.NodeFini(myNode)
	rcl.Shutdown()

}
