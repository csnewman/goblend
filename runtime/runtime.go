package runtime

import (
	_ "unsafe"
)

//go:linkname _iscgo runtime.iscgo
var _iscgo bool = true

//go:linkname goblend_cgo_init goblend_cgo_init
//go:linkname _cgo_init _cgo_init
var goblend_cgo_init byte
var _cgo_init = &goblend_cgo_init

//go:linkname goblend_cgo_thread_start goblend_cgo_thread_start
//go:linkname _cgo_thread_start _cgo_thread_start
var goblend_cgo_thread_start byte
var _cgo_thread_start = &goblend_cgo_thread_start

//go:linkname _cgo_pthread_key_created _cgo_pthread_key_created
var goblend_cgo_pthread_key_created uintptr
var _cgo_pthread_key_created = &goblend_cgo_pthread_key_created

//go:linkname x_crosscall2_ptr x_crosscall2_ptr
//go:linkname _crosscall2_ptr _crosscall2_ptr
var x_crosscall2_ptr byte
var _crosscall2_ptr = &x_crosscall2_ptr

func set_crosscall2()

//go:linkname _set_crosscall2 runtime.set_crosscall2
var _set_crosscall2 = set_crosscall2

//go:linkname goblend_cgo_bindm goblend_cgo_bindm
//go:linkname _cgo_bindm _cgo_bindm
var goblend_cgo_bindm byte
var _cgo_bindm = &goblend_cgo_bindm

//go:linkname goblend_cgo_notify_runtime_init_done goblend_cgo_notify_runtime_init_done
//go:linkname _cgo_notify_runtime_init_done _cgo_notify_runtime_init_done
var goblend_cgo_notify_runtime_init_done byte
var _cgo_notify_runtime_init_done = &goblend_cgo_notify_runtime_init_done

//go:linkname goblend_cgo_set_context_function goblend_cgo_set_context_function
//go:linkname _cgo_set_context_function _cgo_set_context_function
var goblend_cgo_set_context_function byte
var _cgo_set_context_function = &goblend_cgo_set_context_function

//go:linkname goblend_cgo_getstackbound goblend_cgo_getstackbound
//go:linkname _cgo_getstackbound _cgo_getstackbound
var goblend_cgo_getstackbound byte
var _cgo_getstackbound = &goblend_cgo_getstackbound
