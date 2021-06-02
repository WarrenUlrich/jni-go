package internal

/*
	#include "jvmti_wrapper.h"
	#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

//JVMTIEnvironment ...
type JVMTIEnvironment = C.jvmtiEnv

type point struct {
	x, y int
}

//GetLoadedClasses ...
func GetLoadedClasses(env *JVMTIEnvironment) ([]Class, error) {
	var count C.long
	var buf *Class

	jvmtiErr := C.jvmti_get_loaded_classes(env, &count, &buf)

	if jvmtiErr != C.JVMTI_ERROR_NONE {
		return nil, errors.New("failed to get loaded classes")
	}

	//c array to go array without copy.
	classes := (*[1 << 28]Class)(unsafe.Pointer(buf))[:count:count]

	return classes, nil
}

//Deallocate ...
func Deallocate(env *JVMTIEnvironment, ptr unsafe.Pointer) error {
	jvmtiErr := C.jvmti_deallocate(env, ptr)

	if jvmtiErr != C.JVMTI_ERROR_NONE {
		return errors.New("jvmti failed to deallocate")
	}

	return nil
}

//GetClassSignature ...
func GetClassSignature(env *JVMTIEnvironment, cls Class) (string, string, error) {
	var csig *C.char
	var cgeneric *C.char

	jvmtiErr := C.jvmti_get_class_signature(env, cls, &csig, &cgeneric)
	if jvmtiErr != C.JVMTI_ERROR_NONE {
		fmt.Println("err: ", jvmtiErr)
		return "", "", errors.New("jvmti failed to get class signature")
	}
	var sig string
	var generic string

	if csig != nil {
		sig = C.GoString(csig)
		Deallocate(env, unsafe.Pointer(csig))
		if cgeneric != nil {
			generic = C.GoString(cgeneric)
			Deallocate(env, unsafe.Pointer(cgeneric))
		}

	}
	
	return sig, generic, nil
}