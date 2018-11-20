package types

// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include "msg_types.h"
import "C"
import "unsafe"
import "strconv"

type StdMsgsInt8 struct {
	data    *C.std_msgs__msg__Int8
	MsgType MessageTypeSupport
}

func (msg *StdMsgsInt8) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsInt8) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsInt8) InitMessage() {
	msg.data = C.init_std_msgs_msg_Int8()
	msg.MsgType = GetMessageTypeFromStdMsgsInt8()
}

func (msg *StdMsgsInt8) SetInt8(number int8) {
	msg.data.data = C.schar(number)
}

func (msg *StdMsgsInt8) GetInt8() int8 {
	return int8(msg.data.data)
}

func (msg *StdMsgsInt8) GetDataAsString() string {
	return strconv.Itoa(int(msg.GetInt8()))
}

func (msg *StdMsgsInt8) DestroyMessage() {
	C.destroy_std_msgs_msg_Int8(msg.data)
}

func GetMessageTypeFromStdMsgsInt8() MessageTypeSupport {
	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_Int8()}
}
