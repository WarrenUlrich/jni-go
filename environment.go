package jni

/*
	#include "jni_wrapper.h"
*/
import "C"

import "unsafe"

//Environment represents a reference to a  JNI environment
//It is used to interop with the JVM.
type Environment struct {
	pointer unsafe.Pointer
}

//GetVersion returns the major version number in
//the higher 16 bits and the minor version in the
//lower 16 bits.
func (e *Environment) GetVersion() int {
	res := C.jni_get_version(e.pointer)
	return int(res)
}
