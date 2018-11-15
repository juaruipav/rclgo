# rclgo

The goal of this repository is to create a wrapper using cgo for the C library developed for the upcoming ROS2 using the ROS Client Libraries API (RCL)

**ADVICE**: This project is **under development**. It is not totally functional yet.

## Available functionality

Some basic behaviours such as creating a Node, a Publisher or a Subscriber is already implemented. For messages, only the std_msgs/String has been implemented by now. In principle, it is intended to generate custom messages with C generator and bind our own types later in go. For the generation of messages, templates and go generate will be used.

![snapshot](ros2go.gif)

## Repository structure
Following the structure of the [RCL API](http://docs.ros2.org/latest/api/rcl/index.html), the library in go has been organized following the ROS2 concepts. Each concept correspond to a go package.

The goal of this repository is to be as closed as possible to the official C API. When binding is complete, a good option would be to bind this rclgo library with more native/complex behavior (or even with new types/interfaces)


## Pre-requisites

cgo is used for the library wrapping. Temporarly, it is assumed that you have the RCL libraries in:

```
/opt/ros/bouncy
```












