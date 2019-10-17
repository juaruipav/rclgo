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
	"../types/msg/stdmsgs"
)

func TestSubscription(t *testing.T) {

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
		log.Fatal(fmt.Errorf("error initializing the context object"))
	}

	// Obtain initialized node struct and options.
	myNode := node.GetZeroInitializedNode()
	myNodeOpts := node.GetNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	// Initialize node object and associate context object to it.
	result = node.NodeInit(myNode, "GoSubscriber", "", myNodeOpts, contextObject.ContextKey)
	if result != 0 {
		log.Fatal(fmt.Errorf("error initializing the node object"))
	}
	// Create the subscriptor.
	mySub := GetZeroInitializedSubscription()
	mySubOpts := GetSubscriptionDefaultOptions()

	// Create message structure
	//var genericMsg types.MSGInterface

	// distinct datatypes already tested
	genericMsg := &stdmsgs.String{}
	//genericMsg := &stdmsgs.Int64{}

	fmt.Printf("Creating the subscriber! \n")
	SubscriptionInit(mySub, mySubOpts, myNode, "/myGoTopic", genericMsg)

loop:
	for {

		// Bad implementation, it should wait for new messages instead of spamming rcl_take()
		// However at this time there is no achievable solution.
		select {
		case <-sigs:
			fmt.Println("Got shutdown, exiting")
			break loop
		default:
			rclReturn := Take(mySub, genericMsg)
			if rclReturn == types.RCL_RET_OK {
				fmt.Printf("Direct message print: %v\n", genericMsg.Data)
			}
			time.Sleep(50 * time.Millisecond)

			//case <-msg:

		}
	}

	log.Printf("Shutting down")

	//myMsg.DestroyMessage()
	SubscriptionFini(mySub, myNode)
	// Shutdown node object.
	node.NodeFini(myNode)
	// Shutdown context object.
	rcl.RCLShutdown(contextObject.ContextKey)

}
