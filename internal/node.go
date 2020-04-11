package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #cgo LDFLAGS: -L/opt/ros/eloquent/lib -Wl,-rpath=/opt/ros/eloquent/lib -lrcl -lrcutils
// #include "rcl/rcl.h"
import "C"
import "unsafe"

//
type RclNode C.rcl_node_t

//
type RclNodeOptions C.rcl_node_options_t

//
func RclNodeGetDefaultOptions() RclNodeOptions {
	var defOpts C.rcl_node_options_t = C.rcl_node_get_default_options()
	return RclNodeOptions(defOpts)
}

//
func RclGetZeroInitializedNode() RclNode {
	var zeroNode C.rcl_node_t = C.rcl_get_zero_initialized_node()
	return RclNode(zeroNode)
}

//
func RclNodeFini(node *RclNode) int {
	cNode := (*C.rcl_node_t)(node)
	ret := C.rcl_node_fini(cNode)
	return int(ret)
}

//
func RclNodeInit(
	node *RclNode,
	name string,
	namespace string,
	ctx RclContextPtr,
	options *RclNodeOptions,
) int {
	var cName *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var cNamespace *C.char = C.CString(namespace)
	defer C.free(unsafe.Pointer(cNamespace))

	var ret C.int32_t = C.rcl_node_init(
		(*C.rcl_node_t)(node),
		cName,
		cNamespace,
		(*C.rcl_context_t)(ctx),
		(*C.rcl_node_options_t)(options),
	)

	return int(ret)
}
