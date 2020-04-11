package publisher

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #cgo LDFLAGS: -L/opt/ros/eloquent/lib -Wl,-rpath=/opt/ros/eloquent/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include <rosidl_generator_c/message_type_support_struct.h>
// #include "rcl/rcl.h"
// #include <std_msgs/msg/string.h>
// #include <std_msgs/msg/string__functions.h>
// #include <rosidl_generator_c/string_functions.h>
// int publish (const rcl_publisher_t * publisher,  rosidl_message_type_support_t* msg, void * data){
//		if(msg == NULL || publisher == NULL || data == NULL)
//			return 1;
//		// Payload assignation to the message
//		msg->data = data;
//		int retValue = rcl_publish(publisher, msg->data, NULL);
//		return retValue;
//}
import "C"
import (
	"github.com/richardrigby/rclgo/cwrap"
	"github.com/richardrigby/rclgo/node"
	"github.com/richardrigby/rclgo/rcl"
	"github.com/richardrigby/rclgo/types"
)

type Publisher struct {
	RCLPublisher *cwrap.RclPublisher
}

type PublisherOptions struct {
	RCLPublisherOptions *cwrap.RclPublisherOptions
}

func GetZeroInitializedPublisher() Publisher {
	zeroPublisher := cwrap.RclGetZeroInitializedPublisher()
	return Publisher{&zeroPublisher}
}

func GetPublisherDefaultOptions() PublisherOptions {
	defOpts := cwrap.RclPublisherGetDefaultOptions()
	return PublisherOptions{&defOpts}
}

func GetTopicName(publisher Publisher) string {
	return cwrap.RclPublisherGetTopicName(publisher.RCLPublisher)
}

func PublisherInit(
	publisher Publisher,
	publisherOptions PublisherOptions,
	node node.Node,
	topicName string,
	msg types.MessageTypeSupport,
) error {

	ret := cwrap.RclPublisherInit(
		publisher.RCLPublisher,
		node.RCLNode,
		msg.ROSIdlMessageTypeSupport,
		topicName,
		publisherOptions.RCLPublisherOptions,
	)

	if ret != 0 {
		return rcl.NewErr("cwrap.RclInitOptionsInit", int(ret))
	}

	return nil
}

func PublisherFini(publisher Publisher, node node.Node) error {
	ret := cwrap.RclPublisherFini(publisher.RCLPublisher, node.RCLNode)
	if ret != 0 {
		return rcl.NewErr("C.rcl_publisher_fini", int(ret))
	}

	return nil
}

func Publish(publisher Publisher, msg types.MessageTypeSupport, data types.MessageData) error {

	ret := cwrap.RclPublish(
		publisher.RCLPublisher,
		msg.ROSIdlMessageTypeSupport,
		data.Data,
	)

	if ret != 0 {
		return rcl.NewErr("cwrap.Publish", ret)
	}

	return nil
}

func IsValid(publisher Publisher) bool {
	return cwrap.RclPublisherIsValid(publisher.RCLPublisher)
}
