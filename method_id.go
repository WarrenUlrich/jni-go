package jni

import "unsafe"

//MethodID represents a java method ID,
//can be cached.
type MethodID struct {
	pointer unsafe.Pointer
}