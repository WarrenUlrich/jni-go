package jni

/*
	#cgo 386 CFLAGS: -I/openjdk/openjdk8-win32/include -Iopenjdk/openjdk8-win32/include/win32
	#cgo 386 LDFLAGS: -L/openjdk/openjdk8-win32/lib -ljvm

	#cgo amd64 CFLAGS: -I/openjdk/openjdk8-win64/include -Iopenjdk/openjdk8-win64/include/win32
	#cgo amd64 LDFLAGS: -L/openjdk/openjdk8-win64/lib -ljvm

	#include "jni_wrapper.h"
*/
import "C"

import "unsafe"

//GetCreatedVM returns a javaVM, and an error if there is no created VM.
func GetCreatedVM() (JavaVM, error) {
	var jvm unsafe.Pointer

	res := C.jni_get_created_jvm(&jvm)

	if res != 0 {
		return JavaVM{}, &Error{Code: ErrorCode(res)}
	}

	return JavaVM{jvm}, nil
}
