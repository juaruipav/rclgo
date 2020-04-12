package subscription_test

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/richardrigby/rclgo/node"
	"github.com/richardrigby/rclgo/rcl"
	"github.com/richardrigby/rclgo/subscription"
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
	ctx := types.NewZeroInitializedContext()
	err := rcl.Init(&ctx)
	if err != nil {
		t.Fatalf("rcl.Init: %s", err)
	}
	myNode := node.NewZeroInitializedNode()
	myNodeOpts := node.NewNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	err = myNode.Init("GoSubscriber", "", ctx, myNodeOpts)
	if err != nil {
		t.Fatalf("NodeInit: %s", err)
	}

	//Create the subscriptor
	mySub := subscription.NewZeroInitializedSubscription()
	mySubOpts := subscription.NewSubscriptionDefaultOptions()

	//Creating the type
	msgType := types.GetMessageTypeFromStdMsgsString()

	fmt.Printf("Creating the subscriber! \n")
	err = mySub.Init(mySubOpts, myNode, "/myGoTopic", msgType)
	if err != nil {
		t.Fatalf("SubscriptionsInit: %s", err)
	}

	//Creating the msg type
	var myMsg types.StdMsgsString
	myMsg.InitMessage()

loop:
	for {
		err = mySub.TakeMessage(&myMsg.MsgInfo, myMsg.GetData())
		if err == nil {
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
	err = mySub.SubscriptionFini(myNode)
	if err != nil {
		t.Fatalf("SubscriptionFini: %s", err)
	}

	err = myNode.Fini()
	if err != nil {
		t.Fatalf("NodeFini: %s", err)
	}

	err = rcl.Shutdown(ctx)
	if err != nil {
		t.Fatalf("rcl.Shutdown: %s", err)
	}
}
