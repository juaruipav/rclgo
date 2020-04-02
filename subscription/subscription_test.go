package subscription

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/richardrigby/rclgo/node"
	"github.com/richardrigby/rclgo/rcl"
	"github.com/richardrigby/rclgo/types"
)

func TestSubscription(t *testing.T) {

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
	node.NodeInit(myNode, "GoSubscriber", "", myNodeOpts)
	//Create the subscriptor
	mySub := GetZeroInitializedSubscription()
	mySubOpts := GetSubscriptionDefaultOptions()

	//Creating the type
	msgType := types.GetMessageTypeFromStdMsgsString()

	fmt.Printf("Creating the subscriber! \n")
	SubscriptionInit(mySub, mySubOpts, myNode, "/myGoTopic", msgType)

	//Creating the msg type
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
	node.NodeFini(myNode)
	rcl.Shutdown()

}
