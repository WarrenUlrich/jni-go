package jni

import "unsafe"

//Object represents a JNI object reference
type Object struct {
	pointer       unsafe.Pointer
	referenceType referenceType
	environment   *Environment
}
