package rclgo_test

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"testing"
	"time"

	"github.com/richardrigby/rclgo"
	"github.com/richardrigby/rclgo/types"
)

// func BenchmarkFoo(b *testing.B) {

// 	for n := 0; n < b.N; n++ {

// 		// Initialization
// 		rcl.Init()
// 		rcl.Shutdown()

// 	}

// }

// func BenchmarkPublisher(b *testing.B) {

// 	for n := 0; n < b.N; n++ {

// 		// Initialization
// 		// rcl.Init()
// 		// myNode := node.GetZeroInitializedNode()
// 		// myNodeOpts := node.GetNodeDefaultOptions()

// 		// // fmt.Printf("Creating the node! \n")
// 		// node.NodeInit(myNode, "GoPublisher", "", myNodeOpts)

// 		// //Create the publisher
// 		// myPub := GetZeroInitializedPublisher()
// 		// myPubOpts := GetPublisherDefaultOptions()

// 		//Create the msg type
// 		var myMsg types.StdMsgsString
// 		myMsg.InitMessage()

// 		// fmt.Printf("Creating the publisher! \n")
// 		//Initializing the publisher
// 		// PublisherInit(myPub, myPubOpts, myNode, "/myGoTopic", myMsg.GetMessage())

// 		// index := 0

// 		// //Update my msg
// 		// myMsg.SetText("Greetings from GO! #" + strconv.Itoa(index))
// 		// //Publish the message
// 		// Publish(myPub, myMsg.GetMessage(), myMsg.GetData())

// 		// fmt.Printf("Shutting down!! \n")

// 		myMsg.DestroyMessage()
// 		// PublisherFini(myPub, myNode)
// 		// node.NodeFini(myNode)
// 		// rcl.Shutdown()
// 	}

// }

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
	myContext := rclgo.NewZeroInitializedContext()
	err := myContext.Init()
	if err != nil {
		t.Fatalf("rcl.Init failed: %s\n", err)
	}
	myNode := rclgo.NewZeroInitializedNode()
	myNodeOpts := rclgo.NewNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	err = myNode.Init("GoPublisher", "", myContext, myNodeOpts)
	if err != nil {
		t.Fatalf("NodeInit failed: %s\n", err)
	}

	//Create the publisher
	myPub := rclgo.NewZeroInitializedPublisher()
	myPubOpts := rclgo.NewPublisherDefaultOptions()

	//Create the msg type
	var myMsg types.StdMsgsString
	myMsg.InitMessage()

	fmt.Printf("Creating the publisher! \n")
	//Initializing the publisher
	err = myPub.Init(myPubOpts, myNode, "/myGoTopic", myMsg.GetMessage())
	if err != nil {
		t.Fatalf("PublisherInit failed: %s\n", err)
	}

	index := 0
loop:
	for {
		//Update my msg
		myMsg.SetText("Greetings from GO! #" + strconv.Itoa(index))
		//Publish the message
		err := myPub.Publish(myMsg.GetMessage(), myMsg.GetData())

		if err != nil {
			t.Fatalf("Publish failed: %s\n", err)
		} else {
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
	err = myPub.PublisherFini(myNode)
	if err != nil {
		t.Fatalf("PublishFini: %s\n", err)
	}

	err = myNode.Fini()
	if err != nil {
		t.Fatalf("NodeFini: %s\n", err)
	}

	err = myContext.Shutdown()
	if err != nil {
		t.Fatalf("rcl.Shutdown: %s\n", err)
	}
}

// func TestPublisherInt8Msg(t *testing.T) {

// 	sigs := make(chan os.Signal, 1)
// 	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

// 	msg := make(chan string, 1)
// 	go func() {
// 		// Receive input in a loop
// 		for {
// 			var s string
// 			fmt.Scan(&s)
// 			// Send what we read over the channel
// 			msg <- s
// 		}
// 	}()

// 	// Initialization
// 	rcl.Init()
// 	myNode := node.GetZeroInitializedNode()
// 	myNodeOpts := node.GetNodeDefaultOptions()

// 	fmt.Printf("Creating the node! \n")
// 	node.NodeInit(myNode, "GoPublisher", "", myNodeOpts)

// 	//Create the publisher
// 	myPub := GetZeroInitializedPublisher()
// 	myPubOpts := GetPublisherDefaultOptions()

// 	//Create the msg type
// 	var myMsg types.StdMsgsInt8
// 	myMsg.InitMessage()
// 	myMsg.SetInt8(32)

// 	fmt.Printf("Creating the publisher! \n")
// 	//Initializing the publisher
// 	PublisherInit(myPub, myPubOpts, myNode, "/myGoTopic", myMsg.GetMessage())

// 	index := 0
// loop:
// 	for {
// 		//Update my msg
// 		myMsg.SetInt8(myMsg.GetInt8() + 1)

// 		//Publish the message
// 		retRCL := Publish(myPub, myMsg.GetMessage(), myMsg.GetData())

// 		if retRCL == types.Ok {
// 			var value string
// 			fmt.Sprintf("%v", value)

// 			fmt.Printf("(Publisher) Published: %v\n", myMsg.GetInt8())
// 		}
// 		time.Sleep(500 * time.Millisecond)
// 		index++

// 		//Loop breaker
// 		select {
// 		case <-sigs:
// 			fmt.Println("Got shutdown, exiting")
// 			// Break out of the outer for statement and end the program
// 			break loop
// 		case <-msg:

// 		}
// 	}

// 	fmt.Printf("Shutting down!! \n")

// 	myMsg.DestroyMessage()
// 	PublisherFini(myPub, myNode)
// 	node.NodeFini(myNode)
// 	rcl.Shutdown()

// }

// func TestPublisherFloat64Msg(t *testing.T) {

// 	sigs := make(chan os.Signal, 1)
// 	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

// 	msg := make(chan string, 1)
// 	go func() {
// 		// Receive input in a loop
// 		for {
// 			var s string
// 			fmt.Scan(&s)
// 			// Send what we read over the channel
// 			msg <- s
// 		}
// 	}()

// 	// Initialization
// 	rcl.Init()
// 	myNode := node.GetZeroInitializedNode()
// 	myNodeOpts := node.GetNodeDefaultOptions()

// 	fmt.Printf("Creating the node! \n")
// 	node.NodeInit(myNode, "GoPublisher", "", myNodeOpts)

// 	//Create the publisher
// 	myPub := GetZeroInitializedPublisher()
// 	myPubOpts := GetPublisherDefaultOptions()

// 	//Create the msg type
// 	var myMsg types.StdMsgsFloat64
// 	myMsg.InitMessage()
// 	myMsg.SetFloat64(32.0)

// 	fmt.Printf("Creating the publisher! \n")
// 	//Initializing the publisher
// 	PublisherInit(myPub, myPubOpts, myNode, "/myGoTopic", myMsg.GetMessage())

// 	index := 0
// loop:
// 	for {
// 		//Update my msg
// 		myMsg.SetFloat64(myMsg.GetFloat64() + 0.25)

// 		//Publish the message
// 		retRCL := Publish(myPub, myMsg.GetMessage(), myMsg.GetData())

// 		if retRCL == types.Ok {
// 			var value string
// 			fmt.Sprintf("%v", value)

// 			fmt.Printf("(Publisher) Published: %v\n", myMsg.GetFloat64())
// 		}
// 		time.Sleep(500 * time.Millisecond)
// 		index++

// 		//Loop breaker
// 		select {
// 		case <-sigs:
// 			fmt.Println("Got shutdown, exiting")
// 			// Break out of the outer for statement and end the program
// 			break loop
// 		case <-msg:

// 		}
// 	}

// 	fmt.Printf("Shutting down!! \n")

// 	myMsg.DestroyMessage()
// 	PublisherFini(myPub, myNode)
// 	node.NodeFini(myNode)
// 	rcl.Shutdown()

// }
