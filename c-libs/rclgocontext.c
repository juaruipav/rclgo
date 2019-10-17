#include "rclgocontext.h"

// Global variable to check if map is already initialized.
// 0 = False, 1 = True.
int isMapInit = 0;

// Get RCLContext object ID.
uint64_t getInitializedRCLContext(){
    // First initialize context.
    rcl_context_t* contextPointer = getZeroInitContextPointer();

    // Store context on map and get the key.
    uint64_t contextKey = storeRCLContextOnMap(contextPointer);

    return contextKey;
}

// Obtain RCLContext pointer using contextKey.
rcl_context_t* getContextPointerByKey(uint64_t key){
    // Get RCLContext Pointer stored with the given key.
    rcl_context_t* contextPointer;
    contextPointer = *map_get(&contextMap, key);

    // If pointer exists, return. Else, return error.
    if(contextPointer){
        return contextPointer;
    }else{
        printf("ContextLib - Error getting pointer from map!");
        return NULL;
    }
}

// Store pointer and returns the key used to store the pointer.
uint64_t storeRCLContextOnMap(rcl_context_t* pointerToContext){
    // Initialize map only if it does not exists.
    if(isMapInit==0){
        map_init(&contextMap);
        isMapInit = 1;
    }

    rcl_context_instance_id_t contextID;

    // Get the unique contextID.
    contextID = rcl_context_get_instance_id(pointerToContext);

    // Store pointer on map using contextID as key.
    int mapReturnCode = map_set(&contextMap, contextID, pointerToContext);
    if(mapReturnCode != 0){
        printf("ContextLib - Error storing pointer on map!");
        return 0;
    }

    return contextID;
}

// Remove rcl_context_t pointer from map.
void removeRCLContextFromMap(uint64_t contextKey){
    // Remove the context from map.
    map_remove(&contextMap, contextKey);
}

// Get Zero initialized context pointer. This function allocs memory to rcl_context_t.
rcl_context_t* getZeroInitContextPointer(){
    rcl_context_t initContext;
    rcl_context_t* returnPointer;

    // Alloc memory to store rcl_context_t struct.
    returnPointer = (rcl_context_t*) malloc(sizeof(rcl_context_t));

    // Get zero initialized context.
    initContext = rcl_get_zero_initialized_context();

    // Copy zero initialized context struct to the allocated memory zone.
    memcpy(returnPointer, &initContext, sizeof(rcl_context_t));
    // Return pointer to the memory zone.
    return returnPointer;
}

// Free the allocated memory for rcl_context_t and remove from map.
void freeAllocatedMemoryToContext(uint64_t contextKey){
    // Get pointer to context.
    rcl_context_t* pointerToContext = getContextPointerByKey(contextKey);

    if(pointerToContext != NULL){
        free(pointerToContext);
        // Remove pointer from map.
        removeRCLContextFromMap(contextKey);
    }

    // If the map is empty, free the map.
	uint64_t key;
	int counter=0;
	map_iter_t iter = map_iter(&contextMap);

	while ((key = map_next(&contextMap, &iter))) {
  	    counter++;
	}

    if(counter == 0){
        map_deinit(&contextMap);
    }
}