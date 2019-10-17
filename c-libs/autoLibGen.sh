#!/bin/bash

# Compile map lib
gcc -c -o rclgomap.o rclgomap.c

# Compile context lib
gcc -I $COLCON_PREFIX_PATH/include -lrclgomap -c -o rclgocontext.o rclgocontext.c
ar cr librclgocontext.a rclgocontext.o rclgomap.o

# Compile types lib
gcc -I $COLCON_PREFIX_PATH/include -c -o rclgotypes.o rclgotypes.c
ar cr librclgotypes.a rclgotypes.o

# Clean temporary files
rm rclgomap.o rclgocontext.o rclgotypes.o

