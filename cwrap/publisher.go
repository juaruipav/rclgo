package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include "rcl/rcl.h"
import "C"
import "unsafe"

//
type RclPublisher C.rcl_publisher_t

//
type RclPublisherOptions C.rcl_publisher_options_t

//
func RclGetZeroInitializedPublisher() RclPublisher {
	zeroPublisher := C.rcl_get_zero_initialized_publisher()
	return RclPublisher(zeroPublisher)
}

//
func RclPublisherGetDefaultOptions() RclPublisherOptions {
	ret := C.rcl_publisher_get_default_options()
	return RclPublisherOptions(ret)
}

//
func RclPublisherGetTopicName(publisher *RclPublisher) string {
	ret := C.rcl_publisher_get_topic_name((*C.rcl_publisher_t)(publisher))
	return C.GoString(ret)
}

//
func RclPublisherInit(
	publisher *RclPublisher,
	node *RclNode,
	typeSupport *ROSIdlMessageTypeSupport,
	topicName string,
	publisherOpts *RclPublisherOptions,
) int {
	tName := C.CString(topicName)
	defer C.free(unsafe.Pointer(tName))

	ret := C.rcl_publisher_init(
		(*C.rcl_publisher_t)(publisher),
		(*C.rcl_node_t)(node),
		(*C.rosidl_message_type_support_t)(typeSupport),
		tName,
		(*C.rcl_publisher_options_t)(publisherOpts),
	)

	return int(ret)
}

//
func RclPublisherFini(publisher *RclPublisher, node *RclNode) int {
	ret := C.rcl_publisher_fini(
		(*C.rcl_publisher_t)(publisher),
		(*C.rcl_node_t)(node),
	)

	return int(ret)
}

//
func RclPublish(publisher *RclPublisher, typeSupport *ROSIdlMessageTypeSupport, data unsafe.Pointer) int {

	(*C.rosidl_message_type_support_t)(typeSupport).data = data

	ret := C.rcl_publish(
		(*C.rcl_publisher_t)(publisher),
		data,
		nil,
	)

	return int(ret)
}

//
func RclPublisherIsValid(publisher *RclPublisher) bool {
	var ret C.bool = C.rcl_publisher_is_valid((*C.rcl_publisher_t)(publisher))
	return bool(ret)
}
