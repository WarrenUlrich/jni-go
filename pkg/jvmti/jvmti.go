package jvmti

/*
	#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
	"errors"
	"github.com/warrenulrich/jni-go/pkg/internal"
	"github.com/warrenulrich/jni-go/pkg/jni"
)

type classKey struct {
	sig string
	gen string
}

var (
	javaVM   *internal.JavaVM
	jvmtiEnv *internal.JVMTIEnvironment

	classCache map[classKey]*jni.Class = make(map[classKey]*jni.Class)
)

func initEnv() error {
	if javaVM == nil {
		vm, err := internal.GetCreatedVM()
		if err != nil {
			return err
		}
		javaVM = vm

		env, err := internal.GetJVMTIEnv(vm)
		if err != nil {
			return err
		}
		jvmtiEnv = env
	}
	return nil
}

//GetLoadedClasses ...
func GetLoadedClasses() ([]*jni.Class, error) {
	initEnv()

	temp, err := internal.GetLoadedClasses(jvmtiEnv)
	if err != nil {
		return nil, err
	}

	result := make([]*jni.Class, len(temp))

	//Create global refs and delete the local ones.
	for i, c := range temp {
		cls, err := jni.NewGlobalRef(jni.Object(c))
		if err != nil {
			return nil, err
		}
		result[i] = (*jni.Class)(cls)
		jni.DeleteLocalRef(jni.Object(c))
	}
	
	internal.Deallocate(jvmtiEnv, unsafe.Pointer(&temp[0]))
	return result, nil
}

//GetLoadedClass ...
func GetLoadedClass(sig, gen string) (*jni.Class, error) {
	key := classKey{sig, gen}

	if c, ok := classCache[key]; ok {
		return c, nil
	}

	classes, err := GetLoadedClasses()
	if err != nil {
		return nil, err
	}

	for _, c := range classes {
		s, g, err := GetClassSignature(*c)
		if err != nil {
			return nil, err
		}

		if s == sig && g == gen {
			classCache[key] = c
			return c, nil
		}
	}

	return nil, errors.New("jvmti class not found")
}

//GetClassSignature ...
func GetClassSignature(cls jni.Class) (string, string, error) {
	err := initEnv()
	if err != nil {
		return "", "", err
	}

	return internal.GetClassSignature(jvmtiEnv, internal.Class(cls))
}
