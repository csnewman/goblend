//go:build windows

package runtime

import (
	_ "unsafe"
)

//go:cgo_import_dynamic __imp_CreateEventA CreateEventA "Kernel32.dll"
//go:cgo_import_dynamic __imp_EnterCriticalSection EnterCriticalSection "Kernel32.dll"
//go:cgo_import_dynamic __imp_InitializeCriticalSection InitializeCriticalSection "Kernel32.dll"
//go:cgo_import_dynamic __imp_LeaveCriticalSection LeaveCriticalSection "Kernel32.dll"
//go:cgo_import_dynamic __imp_SetEvent SetEvent "Kernel32.dll"
//go:cgo_import_dynamic __imp_Sleep Sleep "Kernel32.dll"
//go:cgo_import_dynamic __imp_WaitForSingleObject WaitForSingleObject "Kernel32.dll"
//go:cgo_import_dynamic __imp___acrt_iob_func __acrt_iob_func "UCRTBASE.dll"
//go:cgo_import_dynamic __imp__beginthread _beginthread "MSVCRT.dll"
//go:cgo_import_dynamic __imp__errno _errno "MSVCRT.dll"
//go:cgo_import_dynamic __imp_abort abort "MSVCRT.dll"
//go:cgo_import_dynamic __imp_fprintf fprintf "MSVCRT.dll"
//go:cgo_import_dynamic __imp_free free "MSVCRT.dll"
//go:cgo_import_dynamic __imp_fwrite fwrite "MSVCRT.dll"
//go:cgo_import_dynamic __imp_malloc malloc "MSVCRT.dll"
