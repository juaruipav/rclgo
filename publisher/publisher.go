package publisher

// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include <rosidl_generator_c/message_type_support_struct.h>
// #include "rcl/rcl.h"
// #include <std_msgs/msg/string.h>
// #include  <std_msgs/msg/string__functions.h>
//#include <rosidl_generator_c/string_functions.h>
// int publish (const rcl_publisher_t * publisher, void* msg){
//		if(msg == NULL)
//			printf("MSG IS NULL\n");
//		std_msgs__msg__String msg2;
//		int retValue = rcl_publish(publisher,&msg2);
//		return retValue;
//}
// int publish2 (const rcl_publisher_t * publisher, rosidl_message_type_support_t* msg){
//		if(msg == NULL)
//			printf("MSG IS NULL\n");
// 		std_msgs__msg__String* mymsg = std_msgs__msg__String__create();
// 		rosidl_generator_c__String__assign(&(*mymsg).data,"message");
//		//msg->typesupport_identifier = "holahola";
//		msg->data = mymsg;
//		printf("typesupport identifier %s\n",msg->typesupport_identifier);
//		int retValue = rcl_publish(publisher, msg->data);
//		return retValue;
//}
// int publish3 (const rcl_publisher_t * publisher, std_msgs__msg__String* msg){
//		if(msg == NULL)
//			printf("MSG IS NULL\n");
//		int retValue = rcl_publish(publisher, msg);
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

// func Publish(publisher Publisher, msg types.Message) types.RCLRetT {
// 	// retValue := C.publish(publisher.RCLPublisher, msg.GetMessage())
// 	retValue := C.rcl_publish(publisher.RCLPublisher, msg.GetMessage())
// 	return types.RCLRetT(retValue)
// }

func Publish(publisher Publisher, msg types.MessageTypeSupport) types.RCLRetT {
	retValue := C.publish2(publisher.RCLPublisher, (*C.rosidl_message_type_support_t)(unsafe.Pointer(msg.ROSIdlMessageTypeSupport)))
	// retValue := C.rcl_publish(publisher.RCLPublisher, msg.ROSIdlMessageTypeSupport)
	return types.RCLRetT(retValue)
}

// func SubscriptionFini(subscription Subscription, node node.Node) types.RCLRetT {

// 	return types.RCLRetT(C.rcl_subscription_fini(subscription.RCLSubscription,
// 		(*C.struct_rcl_node_t)(unsafe.Pointer(node.RCLNode))))

// }
