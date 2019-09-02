package subscription

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
)

func TestSubscription(t *testing.T) {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	msg := make(chan string, 1)
	go func() {
		// Receive input in a loop.
		for {
			var s string
			fmt.Scan(&s)
			// Send what we read over the channel.
			msg <- s
		}
	}()
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
	result = node.NodeInit(myNode, "GoSubscriber", "", myNodeOpts, contextObject.ContextKey)
	if result != 0 {
		log.Fatal(fmt.Errorf("Error initializing the node object!"))
	}
	// Create the subscriptor.
	mySub := GetZeroInitializedSubscription()
	mySubOpts := GetSubscriptionDefaultOptions()

	// Creating the type.
	msgType := types.GetMessageTypeFromStdMsgsString()

	fmt.Printf("Creating the subscriber! \n")
	SubscriptionInit(mySub, mySubOpts, myNode, "/myGoTopic", msgType)

	// Creating the msg type.
	var myMsg types.StdMsgsString
	myMsg.InitMessage()

loop:
	for {

		retRCL := TakeMessage(mySub, myMsg.GetMessage(), myMsg.GetData())

		if retRCL == types.RCL_RET_OK {
			fmt.Printf("(Suscriber) Received %s\n", myMsg.GetDataAsString())
		}

		time.Sleep(100 * time.Millisecond)
		select {
		case <-sigs:
			fmt.Println("Got shutdown, exiting")
			break loop
		case <-msg:
		}
	}

	fmt.Printf("Shutting down!! \n")

	myMsg.DestroyMessage()
	SubscriptionFini(mySub, myNode)
	// Shutdown node object.
	node.NodeFini(myNode)
	// Shutdown context object.
	rcl.RCLShutdown(contextObject.ContextKey)

}
