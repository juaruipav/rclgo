package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #cgo LDFLAGS: -L/opt/ros/eloquent/lib -Wl,-rpath=/opt/ros/eloquent/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include "msg_types.h"
// char * getCharFromStruct(std_msgs__msg__String* msg){
//		if(msg!=NULL)
//			return msg->data.data;
//		else
//			return "";
//}
import "C"
import "unsafe"

//
type StdMsgs_MsgBool C.std_msgs__msg__Bool

//
func (msg *StdMsgs_MsgBool) Set(data bool) {
	(*C.std_msgs__msg__Bool)(msg).data = (C.bool)(data)
}

//
func (msg *StdMsgs_MsgBool) Data() bool {
	var ret C.bool = (*C.std_msgs__msg__Bool)(msg).data
	return bool(ret)
}

//
func InitStdMsgsMsgBool() *StdMsgs_MsgBool {
	var ret *C.std_msgs__msg__Bool = C.init_std_msgs_msg_Bool()
	return (*StdMsgs_MsgBool)(ret)
}

//
func DestroyStdMsgsMsgBool(msg *StdMsgs_MsgBool) {
	cMsg := (*C.std_msgs__msg__Bool)(msg)
	C.destroy_std_msgs_msg_Bool(cMsg)
}

//
func GetMessageTypeFromStdMsgsMsgBool() *ROSIdlMessageTypeSupport {
	var ret *C.rosidl_message_type_support_t = C.get_message_type_from_std_msgs_msg_Bool()
	return (*ROSIdlMessageTypeSupport)(ret)
}

//
type StdMsgs_MsgByte C.std_msgs__msg__Byte

//
func (msg *StdMsgs_MsgByte) Set(data byte) {
	(*C.std_msgs__msg__Byte)(msg).data = (C.uchar)(data)
}

//
func (msg *StdMsgs_MsgByte) Data() byte {
	var ret C.uchar = (*C.std_msgs__msg__Byte)(msg).data
	return byte(ret)
}

//
func InitStdMsgsMsgByte() *StdMsgs_MsgByte {
	var ret *C.std_msgs__msg__Byte = C.init_std_msgs_msg_Byte()
	return (*StdMsgs_MsgByte)(ret)
}

//
func DestroyStdMsgsMsgByte(msg *StdMsgs_MsgByte) {
	cMsg := (*C.std_msgs__msg__Byte)(msg)
	C.destroy_std_msgs_msg_Byte(cMsg)
}

//
func GetMessageTypeFromStdMsgsMsgByte() *ROSIdlMessageTypeSupport {
	var ret *C.rosidl_message_type_support_t = C.get_message_type_from_std_msgs_msg_Byte()
	return (*ROSIdlMessageTypeSupport)(ret)
}

//
type StdMsgs_MsgFloat32 C.std_msgs__msg__Float32

//
func (msg *StdMsgs_MsgFloat32) Set(data float32) {
	(*C.std_msgs__msg__Float32)(msg).data = (C.float)(data)
}

//
func (msg *StdMsgs_MsgFloat32) Data() float32 {
	var ret C.float = (*C.std_msgs__msg__Float32)(msg).data
	return float32(ret)
}

//
func InitStdMsgsMsgFloat32() *StdMsgs_MsgFloat32 {
	var ret *C.std_msgs__msg__Float32 = C.init_std_msgs_msg_Float32()
	return (*StdMsgs_MsgFloat32)(ret)
}

//
func DestroyStdMsgsMsgFloat32(msg *StdMsgs_MsgFloat32) {
	cMsg := (*C.std_msgs__msg__Float32)(msg)
	C.destroy_std_msgs_msg_Float32(cMsg)
}

//
func GetMessageTypeFromStdMsgsMsgFloat32() *ROSIdlMessageTypeSupport {
	var ret *C.rosidl_message_type_support_t = C.get_message_type_from_std_msgs_msg_Float32()
	return (*ROSIdlMessageTypeSupport)(ret)
}

//
type StdMsgs_MsgFloat64 C.std_msgs__msg__Float64

//
func (msg *StdMsgs_MsgFloat64) Set(data float64) {
	(*C.std_msgs__msg__Float64)(msg).data = (C.double)(data)
}

//
func (msg *StdMsgs_MsgFloat64) Data() float64 {
	var ret C.double = (*C.std_msgs__msg__Float64)(msg).data
	return float64(ret)
}

//
func InitStdMsgsMsgFloat64() *StdMsgs_MsgFloat64 {
	var ret *C.std_msgs__msg__Float64 = C.init_std_msgs_msg_Float64()
	return (*StdMsgs_MsgFloat64)(ret)
}

//
func DestroyStdMsgsMsgFloat64(msg *StdMsgs_MsgFloat64) {
	cMsg := (*C.std_msgs__msg__Float64)(msg)
	C.destroy_std_msgs_msg_Float64(cMsg)
}

//
func GetMessageTypeFromStdMsgsMsgFloat64() *ROSIdlMessageTypeSupport {
	var ret *C.rosidl_message_type_support_t = C.get_message_type_from_std_msgs_msg_Float64()
	return (*ROSIdlMessageTypeSupport)(ret)
}

//
type StdMsgs_MsgInt8 C.std_msgs__msg__Int8

//
func (msg *StdMsgs_MsgInt8) Set(data int8) {
	(*C.std_msgs__msg__Int8)(msg).data = (C.schar)(data)
}

//
func (msg *StdMsgs_MsgInt8) Data() int8 {
	var ret C.schar = (*C.std_msgs__msg__Int8)(msg).data
	return int8(ret)
}

//
func InitStdMsgsMsgInt8() *StdMsgs_MsgInt8 {
	var ret *C.std_msgs__msg__Int8 = C.init_std_msgs_msg_Int8()
	return (*StdMsgs_MsgInt8)(ret)
}

//
func DestroyStdMsgsMsgInt8(msg *StdMsgs_MsgInt8) {
	cMsg := (*C.std_msgs__msg__Int8)(msg)
	C.destroy_std_msgs_msg_Int8(cMsg)
}

//
func GetMessageTypeFromStdMsgsMsgInt8() *ROSIdlMessageTypeSupport {
	var ret *C.rosidl_message_type_support_t = C.get_message_type_from_std_msgs_msg_Int8()
	return (*ROSIdlMessageTypeSupport)(ret)
}

//
type StdMsgs_MsgInt16 C.std_msgs__msg__Int16

//
func (msg *StdMsgs_MsgInt16) Set(data int16) {
	(*C.std_msgs__msg__Int16)(msg).data = (C.short)(data)
}

//
func (msg *StdMsgs_MsgInt16) Data() int16 {
	var ret C.short = (*C.std_msgs__msg__Int16)(msg).data
	return int16(ret)
}

//
func InitStdMsgsMsgInt16() *StdMsgs_MsgInt16 {
	var ret *C.std_msgs__msg__Int16 = C.init_std_msgs_msg_Int16()
	return (*StdMsgs_MsgInt16)(ret)
}

//
func DestroyStdMsgsMsgInt16(msg *StdMsgs_MsgInt16) {
	cMsg := (*C.std_msgs__msg__Int16)(msg)
	C.destroy_std_msgs_msg_Int16(cMsg)
}

//
func GetMessageTypeFromStdMsgsMsgInt16() *ROSIdlMessageTypeSupport {
	var ret *C.rosidl_message_type_support_t = C.get_message_type_from_std_msgs_msg_Int16()
	return (*ROSIdlMessageTypeSupport)(ret)
}

//
type StdMsgs_MsgInt32 C.std_msgs__msg__Int32

//
func (msg *StdMsgs_MsgInt32) Set(data int32) {
	(*C.std_msgs__msg__Int32)(msg).data = (C.int)(data)
}

//
func (msg *StdMsgs_MsgInt32) Data() int32 {
	var ret C.int = (*C.std_msgs__msg__Int32)(msg).data
	return int32(ret)
}

//
func InitStdMsgsMsgInt32() *StdMsgs_MsgInt32 {
	var ret *C.std_msgs__msg__Int32 = C.init_std_msgs_msg_Int32()
	return (*StdMsgs_MsgInt32)(ret)
}

//
func DestroyStdMsgsMsgInt32(msg *StdMsgs_MsgInt32) {
	cMsg := (*C.std_msgs__msg__Int32)(msg)
	C.destroy_std_msgs_msg_Int32(cMsg)
}

//
func GetMessageTypeFromStdMsgsMsgInt32() *ROSIdlMessageTypeSupport {
	var ret *C.rosidl_message_type_support_t = C.get_message_type_from_std_msgs_msg_Int32()
	return (*ROSIdlMessageTypeSupport)(ret)
}

//
type StdMsgs_MsgInt64 C.std_msgs__msg__Int64

//
func (msg *StdMsgs_MsgInt64) Set(data int64) {
	(*C.std_msgs__msg__Int64)(msg).data = (C.long)(data)
}

//
func (msg *StdMsgs_MsgInt64) Data() int64 {
	var ret C.long = (*C.std_msgs__msg__Int64)(msg).data
	return int64(ret)
}

//
func InitStdMsgsMsgInt64() *StdMsgs_MsgInt64 {
	var ret *C.std_msgs__msg__Int64 = C.init_std_msgs_msg_Int64()
	return (*StdMsgs_MsgInt64)(ret)
}

//
func DestroyStdMsgsMsgInt64(msg *StdMsgs_MsgInt64) {
	cMsg := (*C.std_msgs__msg__Int64)(msg)
	C.destroy_std_msgs_msg_Int64(cMsg)
}

//
func GetMessageTypeFromStdMsgsMsgInt64() *ROSIdlMessageTypeSupport {
	var ret *C.rosidl_message_type_support_t = C.get_message_type_from_std_msgs_msg_Int64()
	return (*ROSIdlMessageTypeSupport)(ret)
}

//
type StdMsgs_MsgString C.std_msgs__msg__String

//
func (msg *StdMsgs_MsgString) Set(s string) {
	cText := C.CString(s)
	defer C.free(unsafe.Pointer(cText))
	C.rosidl_generator_c__String__assign(&msg.data, cText)
}

//
func (msg *StdMsgs_MsgString) String() string {
	var data = (*C.std_msgs__msg__String)(msg)
	return C.GoString(C.getCharFromStruct(data))
}

//
func InitStdMsgsMsgString() *StdMsgs_MsgString {
	var ret *C.std_msgs__msg__String = C.init_std_msgs_msg_String()
	return (*StdMsgs_MsgString)(ret)
}

//
func DestroyStdMsgsMsgString(msg *StdMsgs_MsgString) {
	cMsg := (*C.std_msgs__msg__String)(msg)
	C.destroy_std_msgs_msg_String(cMsg)
}

//
func GetMessageTypeFromStdMsgsMsgString() *ROSIdlMessageTypeSupport {
	var ret *C.rosidl_message_type_support_t = C.get_message_type_from_std_msgs_msg_String()
	return (*ROSIdlMessageTypeSupport)(ret)
}

//
type StdMsgs_MsgUInt8 C.std_msgs__msg__UInt8

//
func (msg *StdMsgs_MsgUInt8) Set(data uint8) {
	(*C.std_msgs__msg__UInt8)(msg).data = (C.uchar)(data)
}

//
func (msg *StdMsgs_MsgUInt8) Data() uint8 {
	var ret C.uchar = (*C.std_msgs__msg__UInt8)(msg).data
	return uint8(ret)
}

//
func InitStdMsgsMsgUInt8() *StdMsgs_MsgUInt8 {
	var ret *C.std_msgs__msg__UInt8 = C.init_std_msgs_msg_UInt8()
	return (*StdMsgs_MsgUInt8)(ret)
}

//
func DestroyStdMsgsMsgUInt8(msg *StdMsgs_MsgUInt8) {
	cMsg := (*C.std_msgs__msg__UInt8)(msg)
	C.destroy_std_msgs_msg_UInt8(cMsg)
}

//
func GetMessageTypeFromStdMsgsMsgUInt8() *ROSIdlMessageTypeSupport {
	var ret *C.rosidl_message_type_support_t = C.get_message_type_from_std_msgs_msg_UInt8()
	return (*ROSIdlMessageTypeSupport)(ret)
}

//
type StdMsgs_MsgUInt16 C.std_msgs__msg__UInt16

//
func (msg *StdMsgs_MsgUInt16) Set(data uint16) {
	(*C.std_msgs__msg__UInt16)(msg).data = (C.ushort)(data)
}

//
func (msg *StdMsgs_MsgUInt16) Data() uint16 {
	var ret C.ushort = (*C.std_msgs__msg__UInt16)(msg).data
	return uint16(ret)
}

//
func InitStdMsgsMsgUInt16() *StdMsgs_MsgUInt16 {
	var ret *C.std_msgs__msg__UInt16 = C.init_std_msgs_msg_UInt16()
	return (*StdMsgs_MsgUInt16)(ret)
}

//
func DestroyStdMsgsMsgUInt16(msg *StdMsgs_MsgUInt16) {
	cMsg := (*C.std_msgs__msg__UInt16)(msg)
	C.destroy_std_msgs_msg_UInt16(cMsg)
}

//
func GetMessageTypeFromStdMsgsMsgUInt16() *ROSIdlMessageTypeSupport {
	var ret *C.rosidl_message_type_support_t = C.get_message_type_from_std_msgs_msg_UInt16()
	return (*ROSIdlMessageTypeSupport)(ret)
}

//
type StdMsgs_MsgUInt32 C.std_msgs__msg__UInt32

//
func (msg *StdMsgs_MsgUInt32) Set(data uint32) {
	(*C.std_msgs__msg__UInt32)(msg).data = (C.uint)(data)
}

//
func (msg *StdMsgs_MsgUInt32) Data() uint32 {
	var ret C.uint = (*C.std_msgs__msg__UInt32)(msg).data
	return uint32(ret)
}

//
func InitStdMsgsMsgUInt32() *StdMsgs_MsgUInt32 {
	var ret *C.std_msgs__msg__UInt32 = C.init_std_msgs_msg_UInt32()
	return (*StdMsgs_MsgUInt32)(ret)
}

//
func DestroyStdMsgsMsgUInt32(msg *StdMsgs_MsgUInt32) {
	cMsg := (*C.std_msgs__msg__UInt32)(msg)
	C.destroy_std_msgs_msg_UInt32(cMsg)
}

//
func GetMessageTypeFromStdMsgsMsgUInt32() *ROSIdlMessageTypeSupport {
	var ret *C.rosidl_message_type_support_t = C.get_message_type_from_std_msgs_msg_UInt32()
	return (*ROSIdlMessageTypeSupport)(ret)
}

//
type StdMsgs_MsgUInt64 C.std_msgs__msg__UInt64

//
func (msg *StdMsgs_MsgUInt64) Set(data uint64) {
	(*C.std_msgs__msg__UInt64)(msg).data = (C.ulong)(data)
}

//
func (msg *StdMsgs_MsgUInt64) Data() uint64 {
	var ret C.ulong = (*C.std_msgs__msg__UInt64)(msg).data
	return uint64(ret)
}

//
func InitStdMsgsMsgUInt64() *StdMsgs_MsgUInt64 {
	var ret *C.std_msgs__msg__UInt64 = C.init_std_msgs_msg_UInt64()
	return (*StdMsgs_MsgUInt64)(ret)
}

//
func DestroyStdMsgsMsgUInt64(msg *StdMsgs_MsgUInt64) {
	cMsg := (*C.std_msgs__msg__UInt64)(msg)
	C.destroy_std_msgs_msg_UInt64(cMsg)
}

//
func GetMessageTypeFromStdMsgsMsgUInt64() *ROSIdlMessageTypeSupport {
	var ret *C.rosidl_message_type_support_t = C.get_message_type_from_std_msgs_msg_UInt64()
	return (*ROSIdlMessageTypeSupport)(ret)
}
