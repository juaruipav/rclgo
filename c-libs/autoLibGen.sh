#!/bin/bash

# Compile map lib
gcc -c -o maprcl.o map.c
#ar cr libmaprcl.a maprcl.o

# Compile context lib
gcc -I $COLCON_PREFIX_PATH/include -lmaprcl -c -o context.o context.c
ar cr libcontext.a context.o maprcl.o

# Clean temporary files
rm maprcl.o context.o

