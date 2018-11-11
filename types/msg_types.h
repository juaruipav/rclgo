 #include <rcl/rcl.h>
 #include <rosidl_generator_c/message_type_support_struct.h>
#include <rcl/error_handling.h>

// MACRO to generate the proper function 
#define GET_MSG_TYPE_SUPPORT(x,y,z) const rosidl_message_type_support_t* get_message_type_from_## x ##_## y ##_## z (){ \
  return ROSIDL_GET_MSG_TYPE_SUPPORT(x,y,z); \
}

#define ALLOCATE(s) \
  rcl_get_default_allocator().allocate(s, rcl_get_default_allocator().state)
#define DEALLOCATE(ptr) \
  rcl_get_default_allocator().deallocate(ptr, rcl_get_default_allocator().state)
#define REALLOCATE(ptr, s) \
  rcl_get_default_allocator().reallocate(ptr, s, rcl_get_default_allocator().state)
#define ZERO_ALLOCATE(s) \
  rcl_get_default_allocator().zero_allocate(s, 1, rcl_get_default_allocator().state)

#define PRINT_RCL_ERROR(rclc, rcl) \
  do { \
    fprintf(stderr, "[" #rclc "] error in " #rcl ": %s\n", rcl_get_error_string_safe()); \
    rcl_reset_error(); \
  } while (0)


///////////////////////////////////////
///// STD MSGS
//////////////////////////////////////

// #include <std_msgs/msg/.h>
// GET_MSG_TYPE_SUPPORT(std_msgs,msg,)

#include <std_msgs/msg/bool.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,Bool)

#include <std_msgs/msg/byte.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,Byte)

#include <std_msgs/msg/char.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,Char)

#include <std_msgs/msg/color_rgba.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,ColorRGBA)

#include <std_msgs/msg/empty.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,Empty)

#include <std_msgs/msg/float32.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,Float32)

#include <std_msgs/msg/float64.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,Float64)

#include <std_msgs/msg/header.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,Header)

#include <std_msgs/msg/int16.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,Int16)

#include <std_msgs/msg/int32.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,Int32)

#include <std_msgs/msg/int64.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,Int64)

#include <std_msgs/msg/int8.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,Int8)

#include <std_msgs/msg/multi_array_dimension.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,MultiArrayDimension)

#include <std_msgs/msg/multi_array_layout.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,MultiArrayLayout)

#include <std_msgs/msg/string.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,String)

#include <std_msgs/msg/u_int16.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,UInt16)

#include <std_msgs/msg/u_int32.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,UInt32)

#include <std_msgs/msg/u_int64.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,UInt64)

#include <std_msgs/msg/u_int8.h>
GET_MSG_TYPE_SUPPORT(std_msgs,msg,UInt8)
