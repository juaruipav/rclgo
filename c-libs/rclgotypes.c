#include "rclgotypes.h"

const rosidl_message_type_support_t* get_typesupport_std_msgs_msg_string(){
    return ROSIDL_GET_MSG_TYPE_SUPPORT(std_msgs, msg, String);
}

const rosidl_message_type_support_t* get_typesupport_std_msgs_msg_int64(){
    return ROSIDL_GET_MSG_TYPE_SUPPORT(std_msgs, msg, Int64);
}

const rosidl_message_type_support_t* get_typesupport_by_name(char* package){

    if (strcmp(package, "std_msgs/msg/String") == 0){
        return get_typesupport_std_msgs_msg_string();
    } 
    else if (strcmp(package, "std_msgs/msg/Int64") == 0){
        return get_typesupport_std_msgs_msg_int64();
    }
    else{
        return NULL;
    }
}