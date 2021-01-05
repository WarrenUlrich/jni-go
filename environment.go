package jni

/*
	#include <stdlib.h>
	#include "jni_wrapper.h"
*/
import "C"

import (
	"unsafe"
)

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

//DefineClass loads a class from a buffer of raw class data.
//The buffer containing the raw class data is not referenced
//by the VM after the DefineClass call returns, and it may
//be discarded if desired.
func (e *Environment) DefineClass(name string, loader Object, data []uint8) (Class, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	res := C.jni_define_class(e.pointer, cstr, loader.pointer, unsafe.Pointer(&data[0]), C.int(len(data)))

	return Class{res, local, e}, nil
}

//FindClass In JDK 1.1, this function loads a locally-defined class.
//it searches the directories and zip files specified by the CLASSPATH
//environment variable for the class with the specified name.
//Since JDK 1.2, the Java security model allows non-system classes
//to load and call native methods. FindClass locates the class loader
//associated with the current native method; that is, the class loader of the
//class that declared the native method. If the native method belongs to a
//system class, no class loader will be involved. Otherwise, the proper
//class loader will be invoked to load, link, and initialize, the named class.
//Since JDK 1.2, when FindClass is called through the Invocation Interface,
//there is no current native method or it's associated class loader. In that case,
//the result of ClassLoader.getSystemClassLoader is used. This is the class loader
//the virtual machine creates for applications, and is able to locate classes listed
//in the java.class.path property.
func (e *Environment) FindClass(name string) (Class, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	res := C.jni_find_class(e.pointer, cstr)

	return Class{res, local, e}, nil
}

//FromReflectedMethod convets a java.lang.reflect.Method or java.lang.reflect.Constructor 
//object to a method ID.
func (e *Environment) FromReflectedMethod(method Object) (uintptr, error) {
	res := C.jni_from_reflected_method(e.pointer, method.pointer)
	return uintptr(res), nil
}

//FromReflectedField converts a java.lang.reflect.Field to a FieldID
func (e *Environment) FromReflectedField(field Object) (uintptr, error) {
	res := C.jni_from_reflected_field(e.pointer, field.pointer)
	return uintptr(res), nil
}

//ToReflectedMethod converts a method ID derived from c to a java.lang.reflect.Method
//or java.lang.reflect.Constructor object. isStatic must be set to true if the method ID
//refers to a static field.
/*func (e *Environment) ToReflectedMethod(cls Class, mid MethodID, isStatic bool) (Object, error) {
	res := C.jni_to_reflected_method(e.pointer, cls.pointer, mid.pointer, u)
	return Object{res, local, e}, nil
}*/
