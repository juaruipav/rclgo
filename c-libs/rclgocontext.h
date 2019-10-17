#include <inttypes.h>
#include <stdio.h>

#include <rcl/rcl.h>
#include "rclgomap.h"

// Define global map to use to store context pointers.
map_void_t contextMap;

// Get RCLContext object ID.
uint64_t getInitializedRCLContext();

// Obtain RCLContext pointer using contextKey.
rcl_context_t* getContextPointerByKey(uint64_t);

// Store pointer and returns the key used to store the pointer.
uint64_t storeRCLContextOnMap(rcl_context_t*);

// Remove rcl_context_t pointer from map.
void removeRCLContextFromMap(uint64_t);

// Get Zero initialized context pointer. This function allocs memory to rcl_context_t.
rcl_context_t* getZeroInitContextPointer();

// Free the allocated memory for rcl_context_t.
void freeAllocatedMemoryToContext(uint64_t);