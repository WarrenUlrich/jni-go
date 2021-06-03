package jvmti

/*
	#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"unsafe"
	"github.com/warrenulrich/jni-go/pkg/internal"
	"github.com/warrenulrich/jni-go/pkg/jni"
)

type classKey struct {
	sig string
	gen string
}

var (
	classCache map[classKey]*jni.Class = make(map[classKey]*jni.Class)
)

func getThreadLocalEnvironment() (*internal.JVMTIEnvironment, error) {
	internal.AttachCurrentThread(jni.JVM) // Why do I need to call this here to prevent a crash?
	res, err := internal.GetJVMTIEnv(jni.JVM)

	if err != nil {
		if err == internal.ErrDetached {
			_, err = internal.AttachCurrentThread(jni.JVM)
			if err != nil {
				return nil, err
			}

			res, err := internal.GetJVMTIEnv(jni.JVM)
			if err != nil {
				return nil, err
			}
			return res, nil
		}
	}

	return res, nil
}

//GetLoadedClasses ...
func GetLoadedClasses() ([]*jni.Class, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return nil, err
	}

	temp, err := internal.GetLoadedClasses(localEnv)
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
	
	internal.Deallocate(localEnv, unsafe.Pointer(&temp[0]))
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
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return "", "", err
	}

	return internal.GetClassSignature(localEnv, internal.Class(cls))
}
