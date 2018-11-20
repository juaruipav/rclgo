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

func TestPublisherStringMsg(t *testing.T) {

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
	node.NodeInit(myNode, "GoPublisher", "", myNodeOpts)

	//Create the publisher
	myPub := GetZeroInitializedPublisher()
	myPubOpts := GetPublisherDefaultOptions()

	//Create the msg type
	var myMsg types.StdMsgsString
	myMsg.InitMessage()

	fmt.Printf("Creating the publisher! \n")
	//Initializing the publisher
	PublisherInit(myPub, myPubOpts, myNode, "/myGoTopic", myMsg.GetMessage())

	index := 0
loop:
	for {
		//Update my msg
		myMsg.SetText("Greetings from GO! #" + strconv.Itoa(index))
		//Publish the message
		retRCL := Publish(myPub, myMsg.GetMessage(), myMsg.GetData())

		if retRCL == types.RCL_RET_OK {
			fmt.Printf("(Publisher) Published: %s\n", myMsg.GetDataAsString())
		}
		time.Sleep(500 * time.Millisecond)
		index++

		//Loop breaker
		select {
		case <-sigs:
			fmt.Println("Got shutdown, exiting")
			// Break out of the outer for statement and end the program
			break loop
		case <-msg:

		}
	}

	fmt.Printf("Shutting down!! \n")

	myMsg.DestroyMessage()
	PublisherFini(myPub, myNode)
	node.NodeFini(myNode)
	rcl.Shutdown()

}

func TestPublisherInt8Msg(t *testing.T) {

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
	node.NodeInit(myNode, "GoPublisher", "", myNodeOpts)

	//Create the publisher
	myPub := GetZeroInitializedPublisher()
	myPubOpts := GetPublisherDefaultOptions()

	//Create the msg type
	var myMsg types.StdMsgsInt8
	myMsg.InitMessage()
	myMsg.SetInt8(32)

	fmt.Printf("Creating the publisher! \n")
	//Initializing the publisher
	PublisherInit(myPub, myPubOpts, myNode, "/myGoTopic", myMsg.GetMessage())

	index := 0
loop:
	for {
		//Update my msg
		myMsg.SetInt8(myMsg.GetInt8() + 1)

		//Publish the message
		retRCL := Publish(myPub, myMsg.GetMessage(), myMsg.GetData())

		if retRCL == types.RCL_RET_OK {
			var value string
			fmt.Sprintf("%v", value)

			fmt.Printf("(Publisher) Published: %v\n", myMsg.GetInt8())
		}
		time.Sleep(500 * time.Millisecond)
		index++

		//Loop breaker
		select {
		case <-sigs:
			fmt.Println("Got shutdown, exiting")
			// Break out of the outer for statement and end the program
			break loop
		case <-msg:

		}
	}

	fmt.Printf("Shutting down!! \n")

	myMsg.DestroyMessage()
	PublisherFini(myPub, myNode)
	node.NodeFini(myNode)
	rcl.Shutdown()

}
