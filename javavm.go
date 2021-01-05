package jni

/*
	#include "jni_wrapper.h"
*/
import "C"

import (
	"unsafe"
)

//JavaVM represents a reference to a java virtual
//machine instance.
type JavaVM struct {
	pointer unsafe.Pointer
}

//AttachCurrentThread attaches the current thread to the JavaVM.
//Returns a jni environment, and an error, if one occurs.
//Trying to attach a thread that is already attached is a no-op.
//A native thread cannot be attached simultaneously to two Java VMs.
//When a thread is attached to the VM, the context class loader is
//is the bootstrap loader.
func (j *JavaVM) AttachCurrentThread() (Environment, error) {
	var env unsafe.Pointer

	res := C.jni_attach_current_thread(j.pointer, &env)

	if res != 0 {
		return Environment{}, &Error{Code: ErrorCode(res)}
	}

	return Environment{env}, nil
}

//DetachCurrentThread detaches the thread from the JavaVM.
//All monitors held by this thread are released. All Java
//threads waiting for this thread to die are notified.
func (j *JavaVM) DetachCurrentThread() error {

	res := C.jni_detach_current_thread(j.pointer)

	if res != 0 {
		return Error{Code: ErrorCode(res)}
	}

	return nil
}

//GetEnvironment returns the jni environment interface for this VM.
func (j *JavaVM) GetEnvironment(version Version) (Environment, error) {
	var env unsafe.Pointer

	res := C.jni_get_env(j.pointer, &env, C.int(version))

	if res != 0 {
		return Environment{}, &Error{Code: ErrorCode(res)}
	}

	return Environment{env}, nil
}

//GetJvmtiEnvironment ...
/*func (j *JavaVM) GetJvmtiEnvironment() error {
	return nil
}*/

//AttachCurrentThreadAsDaemon ...
func (j *JavaVM) AttachCurrentThreadAsDaemon() (Environment, error) {
	var env unsafe.Pointer

	res := C.jni_attach_current_thread_as_daemon(j.pointer, &env)

	if res != 0 {
		return Environment{}, &Error{Code: ErrorCode(res)}
	}

	return Environment{env}, nil
}
