package publisher

// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include <rosidl_generator_c/message_type_support_struct.h>
// #include "rcl/rcl.h"
// #include <std_msgs/msg/string.h>
// #include  <std_msgs/msg/string__functions.h>
// #include <rosidl_generator_c/string_functions.h>
// int publish (const rcl_publisher_t * publisher,  rosidl_message_type_support_t* msg, void * data){
//		if(msg == NULL || publisher == NULL || data == NULL)
//			return 1;
//		// Payload assignation to the message
//		msg->data = data;
//		int retValue = rcl_publish(publisher, msg->data);
//		return retValue;
//}
import "C"
import (
	"rclgo/node"
	"rclgo/types"
	"unsafe"
)

type Publisher struct {
	RCLPublisher *C.rcl_publisher_t
}

type PublisherOptions struct {
	RCLPublisherOptions *C.rcl_publisher_options_t
}

func GetZeroInitializedPublisher() Publisher {
	zeroPublisher := C.rcl_get_zero_initialized_publisher()
	return Publisher{&zeroPublisher}
}

func GetPublisherDefaultOptions() PublisherOptions {
	defOpts := C.rcl_publisher_get_default_options()
	return PublisherOptions{&defOpts}
}

func PublisherInit(publisher Publisher, publisherOptions PublisherOptions, node node.Node, topicName string, msg types.MessageTypeSupport) types.RCLRetT {

	tName := C.CString(topicName)
	defer C.free(unsafe.Pointer(tName))

	return types.RCLRetT(C.rcl_publisher_init(publisher.RCLPublisher,
		(*C.struct_rcl_node_t)(unsafe.Pointer(node.RCLNode)),
		(*C.rosidl_message_type_support_t)(unsafe.Pointer(msg.ROSIdlMessageTypeSupport)),
		tName,
		publisherOptions.RCLPublisherOptions))

}

func Publish(publisher Publisher, msg types.MessageTypeSupport, data types.MessageData) types.RCLRetT {

	retValue := C.publish(publisher.RCLPublisher,
		(*C.rosidl_message_type_support_t)(unsafe.Pointer(msg.ROSIdlMessageTypeSupport)),
		data.Data)

	return types.RCLRetT(retValue)
}
