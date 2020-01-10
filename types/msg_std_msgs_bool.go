package types
/////////////////////////////////////////////////////
//// THE CONTENT OF THIS FILE HAS BEEN AUTOGENERATED
/////////////////////////////////////////////////////
// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include "msg_types.h"
import "C"
import "unsafe"


type StdMsgsBool struct {
	data    *C.std_msgs__msg__Bool
	MsgType MessageTypeSupport
}

func (msg *StdMsgsBool) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsBool) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsBool) InitMessage() {
	msg.data = C.init_std_msgs_msg_Bool()
	msg.MsgType = GetMessageTypeFromStdMsgsBool()
}

func (msg *StdMsgsBool) SetBool(data bool) {
	//TODO: to implement the setter
	msg.data.data = C.bool(data)
}



func (msg *StdMsgsBool) GetBool() bool {
	return bool(msg.data.data)
}

func (msg *StdMsgsBool) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue:=""
	return myRetValue
}

func (msg *StdMsgsBool) DestroyMessage() {
	C.destroy_std_msgs_msg_Bool(msg.data)
}

func GetMessageTypeFromStdMsgsBool() MessageTypeSupport {
	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_Bool()}
}
