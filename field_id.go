package jni

import (
	"unsafe"
)

//FieldID represents a java field id,
//can be cached, and reused.
type FieldID struct {
	pointer unsafe.Pointer
}
