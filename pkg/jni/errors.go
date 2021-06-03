package jni

import(
	"github.com/warrenulrich/jni-go/pkg/internal"
)

var(
	//ErrUnknown means an error occured but the type is unknown.
	ErrUnknown = internal.ErrUnknown

	//ErrDetached means the calling thread is detached from the JVM.
	ErrDetached = internal.ErrDetached

	//ErrWrongVersion means the feature is unsupported on the current JNI Version
	ErrWrongVersion = internal.ErrWrongVersion

	//ErrOutOfMemory means the JVM is out of memory
	ErrOutOfMemory = internal.ErrOutOfMemory

	//ErrInvalidArgs means the function was called in invalid arguments.
	ErrInvalidArgs = internal.ErrInvalidArgs
)