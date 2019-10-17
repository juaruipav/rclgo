package publisher

// #cgo CFLAGS: -I/opt/ros/crystal/include
// #cgo LDFLAGS: -L/opt/ros/crystal/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include <rosidl_generator_c/message_type_support_struct.h>
// #include "rcl/rcl.h"
import "C"
import (
	"log"
	"unsafe"

	"../node"
	"../types"
)

// publisher pointer for TypesupportManager Singleton Struct
var supportManager types.TypesupportManagerInterface = types.GetTypesupportHandler()

// Publisher structure to hold the C pointer to the actual rcl_publisher object.
// TODO change to pointer functions, therefore cleaner code.
type Publisher struct {
	RCLPublisher *C.rcl_publisher_t
}

// PublisherOptions structure to hold the C pointer fo the actual rcl_publisher_options object.
// TODO maybe change name to "Options" ans suggested by LINT.
type PublisherOptions struct {
	RCLPublisherOptions *C.rcl_publisher_options_t
}

// GetZeroInitializedPublisher Returns a nil structure to the publisher C object.
func GetZeroInitializedPublisher() Publisher {
	zeroPublisher := C.rcl_get_zero_initialized_publisher()
	return Publisher{&zeroPublisher}
}

// GetPublisherDefaultOptions Returns a default configured structure to the publisherOptions C object.
func GetPublisherDefaultOptions() PublisherOptions {
	defOpts := C.rcl_publisher_get_default_options()
	return PublisherOptions{&defOpts}
}

// GetTopicName returns the current publishing topic's name.
func GetTopicName(publisher Publisher) string {
	return C.GoString(C.rcl_publisher_get_topic_name(publisher.RCLPublisher))
}

// PublisherInit Initializes the previously obtained C object from "GetZeroInitializedPublisher"
// TODO maybe change name to "Init" as suggested by LINT.
func PublisherInit(publisher Publisher, publisherOptions PublisherOptions, node node.Node, topicName string, msg types.MSGInterface) types.RCLRetT {

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
	rclResult = int(C.rcl_publisher_init(publisher.RCLPublisher, cNode, cTypesupport, tName, publisherOptions.RCLPublisherOptions))

	// Converts to RCLRetT struct
	return types.RCLRetT(rclResult)

}

// PublisherFini Destroys the given publisher C object
// TODO maybe change name to "Fini" as suggested by LINT.
func PublisherFini(publisher Publisher, node node.Node) types.RCLRetT {
	return types.RCLRetT(C.rcl_publisher_fini(publisher.RCLPublisher, (*C.struct_rcl_node_t)(unsafe.Pointer(node.RCLNode))))
}

// Publish Attempts to send a new message based on the provided message to the provided publish object.
func Publish(publisher Publisher, msg types.MSGInterface) types.RCLRetT {

	// Variable declaration
	var rclReturnCValue C.int
	var rclResult types.RCLRetT
	var msgC unsafe.Pointer
	var err error

	// Load the typesupport handler
	handler, err := supportManager.GetHandlerByMSGInterface(msg)
	if err != nil {
		log.Printf("err: %s\n", err)
	}

	// Convert Golang msg datatype into C datatype
	msgC, err = handler.GetCTypesupport(msg)
	if err != nil {
		log.Printf("err: %s\n", err)
	}

	// Attempt to publish new message
	rclReturnCValue = C.rcl_publish(publisher.RCLPublisher, msgC, nil)

	// Transform rclReturnCValue into Golang Values
	rclResult = types.RCLRetT(int(rclReturnCValue))

	return rclResult
}
