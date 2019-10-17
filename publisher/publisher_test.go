package publisher

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"../node"
	"../rcl"
	"../types"
	"../types/msg/stdmsgs"
)

func TestPublisherStringMsg(t *testing.T) {

	// Get commandline signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Initialization.
	var contextObject rcl.RCLContext
	// Get initialized context object.
	contextObject.GetInitializedContext()

	// Initialize contextObject and rcl. This function has to be executed always.
	result := rcl.RCLInit(contextObject.ContextKey)
	if result != 0 {
		log.Fatal(fmt.Errorf("Error initializing the context object!"))
	}

	// Obtain initialized node struct and options.
	myNode := node.GetZeroInitializedNode()
	myNodeOpts := node.GetNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	// Initialize node object and associate context object to it.
	result = node.NodeInit(myNode, "GoPublisher", "", myNodeOpts, contextObject.ContextKey)
	if result != 0 {
		log.Fatal(fmt.Errorf("Error initializing the node object!"))
	}

	// Create the publisher.
	myPub := GetZeroInitializedPublisher()
	myPubOpts := GetPublisherDefaultOptions()

	// distinct datatypes already tested
	genericMsg := &stdmsgs.String{"Example generic message."}
	//genericMsg := &stdmsgs.Int64{64}

	fmt.Printf("Initializing the publisher! \n")
	// Initializing the publisher.
	PublisherInit(myPub, myPubOpts, myNode, "/myGoTopic", genericMsg)

	index := 0
loop:
	for {
		//Update my msg.
		//genericMsg.Data = fmt.Sprintf("Generic message from GO! #%d", index)
		//genericMsg.Data = int64(index)

		// Attempting to publish a new message
		retRCL := Publish(myPub, genericMsg)

		if retRCL == types.RCL_RET_OK {
			fmt.Printf("(Publisher) Published: %+v\n", genericMsg)
		}

		// Wait some amount of time
		time.Sleep(500 * time.Millisecond)
		index++

		//Loop breaker.
		select {
		case <-sigs:
			fmt.Println("Got shutdown, exiting")
			// Break out of the outer for statement and end the program.
			break loop
		default:
		}
	}

	fmt.Printf("Shutting down!! \n")

	// Shutdown the publisher
	PublisherFini(myPub, myNode)
	// Shutdown node object.
	node.NodeFini(myNode)
	// Shutdown context object.
	rcl.RCLShutdown(contextObject.ContextKey)
}
