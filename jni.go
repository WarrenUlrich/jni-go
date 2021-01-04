package jni

/*
	#cgo 386 CFLAGS: -IC/ -IC/openjdk8-win32/include -IC/openjdk8-win32/include/win32
	#cgo 386 LDFLAGS: -L/C/openjdk8-win32/lib -ljvm

	#cgo amd64 CFLAGS: -IC/openjdk8-win64/include -IC/openjdk8-win64/include/win32
	#cgo amd64 LDFLAGS: -L/C/openjdk8-win64/lib -ljvm

	#include "jvm_wrapper.h"
*/
import "C"

import "unsafe"

//GetCreatedVM returns a javaVM, and an error if there is no created VM.
func GetCreatedVM() (JavaVM, error) {
	var jvm unsafe.Pointer

	res := C.jvm_get_created_vm(&jvm)

	if res != 0 {
		return JavaVM{}, &Error{Code: ErrorCode(res)}
	}

	return JavaVM{jvm}, nil
}
