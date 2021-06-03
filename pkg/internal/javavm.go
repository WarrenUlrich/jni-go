package internal

/*
	#cgo CFLAGS: -I/jni/include -I/jni/include/win32
	#include "jni_wrapper.h"
	#include <jvmti.h>
*/
import "C"

import (
	"syscall"
	"unsafe"
)

//JavaVM ...
type JavaVM = C.JavaVM

var (
	jvmDll                       *syscall.DLL
	getDefaultJavaVMInitArgsProc *syscall.Proc
	createJavaVMProc             *syscall.Proc
	getCreatedJavaVMsProc        *syscall.Proc
)

//LoadJVMLibrary loads the jvm lib from the specified path,
//must be done before using anything else in the package.
func LoadJVMLibrary(path string) error {
	jvmDll, err := syscall.LoadDLL(path)

	if err != nil {
		return err
	}
	
	getDefaultJavaVMInitArgsProc, err = jvmDll.FindProc("JNI_GetDefaultJavaVMInitArgs")

	if err != nil {
		return err
	}

	createJavaVMProc, err = jvmDll.FindProc("JNI_CreateJavaVM")

	if err != nil {
		return err
	}

	getCreatedJavaVMsProc, err = jvmDll.FindProc("JNI_GetCreatedJavaVMs")

	if err != nil {
		return err
	}

	return nil
}

//GetEnv ...
func GetEnv(vm *JavaVM) (*JNIEnvironment, error) {
	var env unsafe.Pointer
	res := C.jni_get_env(vm, &env, C.JNI_VERSION_1_1)
	return (*JNIEnvironment)(env), jniErrorFromCode(int(res))
}

//GetJVMTIEnv ...
func GetJVMTIEnv(vm *JavaVM) (*JVMTIEnvironment, error) {
	var env unsafe.Pointer
	res := C.jni_get_env(vm, &env, C.JVMTI_VERSION_1)
	return (*JVMTIEnvironment)(env), jniErrorFromCode(int(res))
}

//GetCreatedVM ...
func GetCreatedVM() (*JavaVM, error) {
	var jvmPtr unsafe.Pointer

	getCreatedJavaVMsProc.Call(uintptr(unsafe.Pointer(&jvmPtr)), 1, 0)

	return (*JavaVM)(jvmPtr), nil
}

//AttachCurrentThread ...
func AttachCurrentThread(jvm *JavaVM) (*JNIEnvironment, error) {
	var env unsafe.Pointer
	res := C.jni_attach_current_thread(jvm, &env, nil)

	return (*JNIEnvironment)(env), jniErrorFromCode(int(res))
}

//DetachCurrentThread ...
func DetachCurrentThread(jvm *JavaVM) {
	C.jni_detach_current_thread(jvm)
}