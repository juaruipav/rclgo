package types

/////////////////////////////////////////////////////
//// THE CONTENT OF THIS FILE HAS BEEN AUTOGENERATED
/////////////////////////////////////////////////////
// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include "msg_types.h"
import "C"
import "unsafe"

type StdMsgsByte struct {
	data    *C.std_msgs__msg__Byte
	MsgType MessageTypeSupport
}

func (msg *StdMsgsByte) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsByte) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsByte) InitMessage() {
	msg.data = C.init_std_msgs_msg_Byte()
	msg.MsgType = GetMessageTypeFromStdMsgsByte()
}

func (msg *StdMsgsByte) SetByte(data byte) {
	//TODO: to implement the setter
	msg.data.data = C.uchar(data)
}

func (msg *StdMsgsByte) GetByte() byte {
	return byte(msg.data.data)
}

func (msg *StdMsgsByte) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue := ""
	return myRetValue
}

func (msg *StdMsgsByte) DestroyMessage() {
	C.destroy_std_msgs_msg_Byte(msg.data)
}

func GetMessageTypeFromStdMsgsByte() MessageTypeSupport {
	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_Byte()}
}
