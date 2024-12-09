//go:build linux

package runtime

import "unsafe"

//go:cgo_import_dynamic __errno_location __errno_location "libc.so.6"
//go:cgo_import_dynamic abort abort "libc.so.6"
//go:cgo_import_dynamic fatalf fatalf "libc.so.6"
//go:cgo_import_dynamic fprintf fprintf "libc.so.6"
//go:cgo_import_dynamic free free "libc.so.6"
//go:cgo_import_dynamic fwrite fwrite "libc.so.6"
//go:cgo_import_dynamic malloc malloc "libc.so.6"
//go:cgo_import_dynamic memset memset "libc.so.6"
//go:cgo_import_dynamic mmap mmap "libc.so.6"
//go:cgo_import_dynamic munmap munmap "libc.so.6"
//go:cgo_import_dynamic nanosleep nanosleep "libc.so.6"
//go:cgo_import_dynamic pthread_attr_destroy pthread_attr_destroy "libc.so.6"
//go:cgo_import_dynamic pthread_attr_getstack pthread_attr_getstack "libc.so.6"
//go:cgo_import_dynamic pthread_attr_getstacksize pthread_attr_getstacksize "libc.so.6"
//go:cgo_import_dynamic pthread_attr_init pthread_attr_init "libc.so.6"
//go:cgo_import_dynamic pthread_attr_setdetachstate pthread_attr_setdetachstate "libc.so.6"
//go:cgo_import_dynamic pthread_cond_broadcast pthread_cond_broadcast "libc.so.6"
//go:cgo_import_dynamic pthread_cond_wait pthread_cond_wait "libc.so.6"
//go:cgo_import_dynamic pthread_create pthread_create "libc.so.6"
//go:cgo_import_dynamic pthread_getattr_np pthread_getattr_np "libc.so.6"
//go:cgo_import_dynamic pthread_key_create pthread_key_create "libc.so.6"
//go:cgo_import_dynamic pthread_mutex_lock pthread_mutex_lock "libc.so.6"
//go:cgo_import_dynamic pthread_mutex_unlock pthread_mutex_unlock "libc.so.6"
//go:cgo_import_dynamic pthread_self pthread_self "libc.so.6"
//go:cgo_import_dynamic pthread_setspecific pthread_setspecific "libc.so.6"
//go:cgo_import_dynamic pthread_sigmask pthread_sigmask "libc.so.6"
//go:cgo_import_dynamic setegid setegid "libc.so.6"
//go:cgo_import_dynamic setenv setenv "libc.so.6"
//go:cgo_import_dynamic seteuid seteuid "libc.so.6"
//go:cgo_import_dynamic setgid setgid "libc.so.6"
//go:cgo_import_dynamic setgroups setgroups "libc.so.6"
//go:cgo_import_dynamic setregid setregid "libc.so.6"
//go:cgo_import_dynamic setresgid setresgid "libc.so.6"
//go:cgo_import_dynamic setresuid setresuid "libc.so.6"
//go:cgo_import_dynamic setreuid setreuid "libc.so.6"
//go:cgo_import_dynamic setuid setuid "libc.so.6"
//go:cgo_import_dynamic sigaction sigaction "libc.so.6"
//go:cgo_import_dynamic sigaddset sigaddset "libc.so.6"
//go:cgo_import_dynamic sigemptyset sigemptyset "libc.so.6"
//go:cgo_import_dynamic sigfillset sigfillset "libc.so.6"
//go:cgo_import_dynamic sigismember sigismember "libc.so.6"
//go:cgo_import_dynamic stderr stderr "libc.so.6"
//go:cgo_import_dynamic strerror strerror "libc.so.6"
//go:cgo_import_dynamic unsetenv unsetenv "libc.so.6"

//go:linkname goblend_cgo_libc_setegid goblend_cgo_libc_setegid
//go:linkname cgo_libc_setegid syscall.cgo_libc_setegid
var goblend_cgo_libc_setegid byte
var cgo_libc_setegid = unsafe.Pointer(&goblend_cgo_libc_setegid)

//go:linkname goblend_cgo_libc_seteuid goblend_cgo_libc_seteuid
//go:linkname cgo_libc_seteuid syscall.cgo_libc_seteuid
var goblend_cgo_libc_seteuid byte
var cgo_libc_seteuid = unsafe.Pointer(&goblend_cgo_libc_seteuid)

//go:linkname goblend_cgo_libc_setregid goblend_cgo_libc_setregid
//go:linkname cgo_libc_setregid syscall.cgo_libc_setregid
var goblend_cgo_libc_setregid byte
var cgo_libc_setregid = unsafe.Pointer(&goblend_cgo_libc_setregid)

//go:linkname goblend_cgo_libc_setresgid goblend_cgo_libc_setresgid
//go:linkname cgo_libc_setresgid syscall.cgo_libc_setresgid
var goblend_cgo_libc_setresgid byte
var cgo_libc_setresgid = unsafe.Pointer(&goblend_cgo_libc_setresgid)

//go:linkname goblend_cgo_libc_setresuid goblend_cgo_libc_setresuid
//go:linkname cgo_libc_setresuid syscall.cgo_libc_setresuid
var goblend_cgo_libc_setresuid byte
var cgo_libc_setresuid = unsafe.Pointer(&goblend_cgo_libc_setresuid)

//go:linkname goblend_cgo_libc_setreuid goblend_cgo_libc_setreuid
//go:linkname cgo_libc_setreuid syscall.cgo_libc_setreuid
var goblend_cgo_libc_setreuid byte
var cgo_libc_setreuid = unsafe.Pointer(&goblend_cgo_libc_setreuid)

//go:linkname goblend_cgo_libc_setgroups goblend_cgo_libc_setgroups
//go:linkname cgo_libc_setgroups syscall.cgo_libc_setgroups
var goblend_cgo_libc_setgroups byte
var cgo_libc_setgroups = unsafe.Pointer(&goblend_cgo_libc_setgroups)

//go:linkname goblend_cgo_libc_setgid goblend_cgo_libc_setgid
//go:linkname cgo_libc_setgid syscall.cgo_libc_setgid
var goblend_cgo_libc_setgid byte
var cgo_libc_setgid = unsafe.Pointer(&goblend_cgo_libc_setgid)

//go:linkname goblend_cgo_libc_setuid goblend_cgo_libc_setuid
//go:linkname cgo_libc_setuid syscall.cgo_libc_setuid
var goblend_cgo_libc_setuid byte
var cgo_libc_setuid = unsafe.Pointer(&goblend_cgo_libc_setuid)
