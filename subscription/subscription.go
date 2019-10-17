package subscription

// #cgo CFLAGS: -I/opt/ros/crystal/include
// #cgo LDFLAGS: -L/opt/ros/crystal/lib -lrcl -lrcutils
// #include <rosidl_generator_c/message_type_support_struct.h>
// #include "rcl/rcl.h"
import "C"
import (
	"log"
	"unsafe"

	"../node"
	"../types"
)

// ####################################################################
// FIX ME: Subscription at the time of this writing is based on
// active polling, which has its own drawbacks, should be implemented
// a non-polling (real subscription) methodology for it.
// ####################################################################

// subscriber pointer for TypesupportManager Singleton Struct
var supportManager types.TypesupportManagerInterface = types.GetTypesupportHandler()

// Subscription structure to hold the C pointer to the actual rcl_subscription object.
// TODO change to pointer functions, therefore cleaner code.
type Subscription struct {
	RCLSubscription *C.rcl_subscription_t
}

// SubscriptionOptions structure to hold the C pointer fo the actual rcl_subscription_options object.
// TODO change to "Options" ans suggested by LINT.
type SubscriptionOptions struct {
	RCLSubscriptionOptions *C.rcl_subscription_options_t
}

// GetZeroInitializedSubscription Returns a nil structure to the subscription C object
func GetZeroInitializedSubscription() Subscription {
	zeroSubscription := C.rcl_get_zero_initialized_subscription()
	return Subscription{&zeroSubscription}
}

// GetSubscriptionDefaultOptions Returns a default configured structure to the subscriptionOptions C object
func GetSubscriptionDefaultOptions() SubscriptionOptions {
	defOpts := C.rcl_subscription_get_default_options()
	return SubscriptionOptions{&defOpts}
}

// SubscriptionInit Initializes the previously obtained C object from "GetZeroInitializedSubscription"
func SubscriptionInit(subscription Subscription, subscriptionOptions SubscriptionOptions, node node.Node, topicName string, msg types.MSGInterface) types.RCLRetT {

	// Variable declaration
	var rclResult int
	var cTypesupport *C.rosidl_message_type_support_t
	var cNode *C.struct_rcl_node_t

	tName := C.CString(topicName)
	defer C.free(unsafe.Pointer(tName))

	// Load the typesupport handler
	handler, err := supportManager.GetHandlerByMSGInterface(msg)
	if err != nil {
		log.Printf("err: %s\n", err)
	}

	// Converts the Golang datatype into it's respective C typesupport handler
	cPointer, err := handler.GetROSTypesupportPointer(msg)
	if err != nil {
		log.Printf("err: %s\n", err)
	}
	cTypesupport = (*C.rosidl_message_type_support_t)(unsafe.Pointer(cPointer))

	// Convert Golang node to C pointer
	cNode = (*C.struct_rcl_node_t)(unsafe.Pointer(node.RCLNode))

	// Inits the C RCLSubscription object
	rclResult = int(C.rcl_subscription_init(subscription.RCLSubscription, cNode, cTypesupport, tName, subscriptionOptions.RCLSubscriptionOptions))

	// Converts to RCLRetT struct
	return types.RCLRetT(rclResult)

}

// SubscriptionFini Destroys the given subscription C object
func SubscriptionFini(subscription Subscription, node node.Node) types.RCLRetT {

	return types.RCLRetT(C.rcl_subscription_fini(subscription.RCLSubscription,
		(*C.struct_rcl_node_t)(unsafe.Pointer(node.RCLNode))))

}

// Take Attempts to take a new message from the given subscription with the given message type
func Take(subscription Subscription, msg types.MSGInterface) types.RCLRetT {

	// Variable declaration
	var rclResult types.RCLRetT
	var err error
	var handler types.TypesupportHandlerInterface
	var csupport types.CTypesupport
	var rclReturnCValue C.int

	// Load the typesupport handler
	handler, err = supportManager.GetHandlerByMSGInterface(msg)
	if err != nil {
		log.Printf("err: %s\n", err)
	}

	// Attempt to take new message
	rclReturnCValue = C.rcl_take(subscription.RCLSubscription, unsafe.Pointer(&csupport), nil)

	// Transform rclReturnCValue into Golang Values
	rclResult = types.RCLRetT(int(rclReturnCValue))

	handler.ParseCTypesupport(msg, csupport)

	return rclResult
}
