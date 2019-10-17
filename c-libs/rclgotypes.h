#include <string.h>
#include <std_msgs/msg/string.h>
#include <std_msgs/msg/int64.h>


// msg/stdmsgs typesupports
const rosidl_message_type_support_t* get_typesupport_std_msgs_msg_string();
const rosidl_message_type_support_t* get_typesupport_std_msgs_msg_int64();

// Function to get the correct typesupport
const rosidl_message_type_support_t* get_typesupport_by_name(char*);