package types
/////////////////////////////////////////////////////
//// THE CONTENT OF THIS FILE HAS BEEN AUTOGENERATED
/////////////////////////////////////////////////////
// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #cgo LDFLAGS: -L/opt/ros/eloquent/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include "msg_types.h"
import "C"
import "unsafe"


type StdMsgsFloat32 struct {
	data    *C.std_msgs__msg__Float32
	MsgType MessageTypeSupport
}

func (msg *StdMsgsFloat32) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsFloat32) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsFloat32) InitMessage() {
	msg.data = C.init_std_msgs_msg_Float32()
	msg.MsgType = GetMessageTypeFromStdMsgsFloat32()
}

func (msg *StdMsgsFloat32) SetFloat32(data float32) {
	//TODO: to implement the setter
	msg.data.data = C.float(data)
}

func (msg *StdMsgsFloat32) GetFloat32() float32 {
	return float32(msg.data.data)
}

func (msg *StdMsgsFloat32) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue:=""
	return myRetValue
}

func (msg *StdMsgsFloat32) DestroyMessage() {
	C.destroy_std_msgs_msg_Float32(msg.data)
}

func GetMessageTypeFromStdMsgsFloat32() MessageTypeSupport {
	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_Float32()}
}
