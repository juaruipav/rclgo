package cwrap

// #cgo CFLAGS: -I/opt/ros/eloquent/include
// #include <rcl/rcl.h>
import "C"
import "unsafe"

//
type RclNode C.rcl_node_t

//
type RclNodeOptions C.rcl_node_options_t

//
type RmwNode C.rmw_node_t

//
type RclGuardCondition C.rcl_guard_condition_t

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

//
func RclNodeFini(node *RclNode) int {
	cNode := (*C.rcl_node_t)(node)
	ret := C.rcl_node_fini(cNode)
	return int(ret)
}

//
func RclNodeIsValid(node *RclNode) bool {
	var ret C.bool = C.rcl_node_is_valid(
		(*C.rcl_node_t)(node),
	)
	return bool(ret)
}

//
func RclNodeIsValidExceptContext(node *RclNode) bool {
	var ret C.bool = C.rcl_node_is_valid_except_context(
		(*C.rcl_node_t)(node),
	)
	return bool(ret)
}

//
func RclNodeGetName(node *RclNode) string {
	var cStr *C.char = C.rcl_node_get_name(
		(*C.rcl_node_t)(node),
	)
	return C.GoString(cStr)
}

//
func RclNodeGetNamespace(node *RclNode) string {
	var cStr = C.rcl_node_get_namespace(
		(*C.rcl_node_t)(node),
	)
	return C.GoString(cStr)
}

//
func RclNodeGetFullyQualifiedName(node *RclNode) string {
	var cStr = C.rcl_node_get_fully_qualified_name(
		(*C.rcl_node_t)(node),
	)
	return C.GoString(cStr)
}

//
func RclNodeGetOptions(node *RclNode) *RclNodeOptions {
	var opts *C.rcl_node_options_t = C.rcl_node_get_options(
		(*C.rcl_node_t)(node),
	)
	return (*RclNodeOptions)(opts)
}

//
func RclNodeGetDomainID(node *RclNode, domainID *uint) int {
	var dom C.ulong
	var ret C.int = C.rcl_node_get_domain_id(
		(*C.rcl_node_t)(node),
		&dom,
	)
	*domainID = uint(dom)
	return int(ret)
}

//
func RclNodeAssertLiveliness(node *RclNode) int {
	ret := C.rcl_node_assert_liveliness(
		(*C.rcl_node_t)(node),
	)
	return int(ret)
}

//
func RclNodeGetRmwHandle(node *RclNode) *RmwNode {
	ret := C.rcl_node_get_rmw_handle(
		(*C.rcl_node_t)(node),
	)
	return (*RmwNode)(ret)
}

//
func RclNodeGetRclInstanceID(node *RclNode) uint64 {
	ret := C.rcl_node_get_rcl_instance_id(
		(*C.rcl_node_t)(node),
	)
	return uint64(ret)
}

//
func RclNodeGetGraphGuardCondition(node *RclNode) *RclGuardCondition {
	ret := C.rcl_node_get_graph_guard_condition(
		(*C.rcl_node_t)(node),
	)
	return (*RclGuardCondition)(ret)
}

//
func RclNodeGetLoggerName(node *RclNode) string {
	var cStr *C.char = C.rcl_node_get_logger_name(
		(*C.rcl_node_t)(node),
	)
	return C.GoString(cStr)
}
