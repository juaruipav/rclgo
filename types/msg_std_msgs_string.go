package types

// #cgo CFLAGS: -I /opt/ros/crystal/include
// #cgo LDFLAGS: -L /opt/ros/crystal/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include "msg_types.h"
// char * getCharFromStruct(std_msgs__msg__String* msg){
//		if(msg!=NULL)
//			return msg->data.data;
//		else
//			return "";
//}
import "C"
import (
	"unsafe"
)

type StdMsgsString struct {
	data    *C.std_msgs__msg__String
	MsgType MessageTypeSupport
}

func (msg *StdMsgsString) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsString) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsString) GetDataAsString() string {
	return C.GoString(C.getCharFromStruct(msg.data))
}

func (msg *StdMsgsString) InitMessage() {
	msg.data = C.init_std_msgs_msg_String()
	msg.MsgType = GetMessageTypeFromStdMsgsString()
}

func (msg *StdMsgsString) DestroyMessage() {
	C.destroy_std_msgs_msg_String(msg.data)
}

func (msg *StdMsgsString) SetText(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	C.rosidl_generator_c__String__assign(&msg.data.data, cText)
}

func GetMessageTypeFromStdMsgsString() MessageTypeSupport {
	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_String()}
}
