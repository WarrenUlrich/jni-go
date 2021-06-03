package internal

import (
	"errors"
)

var(
	//ErrUnknown means an error occured but the type is unknown.
	ErrUnknown = errors.New("unknown")

	//ErrDetached means the calling thread is detached from the JVM.
	ErrDetached = errors.New("thread detached")

	//ErrWrongVersion means the feature is unsupported on the current JNI Version
	ErrWrongVersion = errors.New("unsupported version")

	//ErrOutOfMemory means the JVM is out of memory
	ErrOutOfMemory = errors.New("out of memory")

	//ErrInvalidArgs means the function was called in invalid arguments.
	ErrInvalidArgs = errors.New("invalid arguments")
)

const(
	ok int = 0
	unknown int = -1
	detached int = -2
	version int = -3
	outOfMemory int = -4
	invalidArgs int = -5
)

func jniErrorFromCode(code int) error {
	switch code {
	case ok:
		return nil

	case unknown:
		return ErrUnknown

	case detached:
		return ErrDetached

	case version:
		return ErrWrongVersion

	case outOfMemory:
		return ErrOutOfMemory

	case invalidArgs:
		return ErrInvalidArgs
	}
	return nil
}
