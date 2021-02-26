package main

/*
	#include <windows.h>
*/
import "C"

import (
	"fmt"
	"os"

	"github.com/warrenulrich/jni-go/pkg/jni"
	"github.com/warrenulrich/jni-go/pkg/jvmti"
)

//export bingus
func bingus() {
	go func() {
		C.AllocConsole()
		os.Stdout = os.NewFile((uintptr)(C.GetStdHandle(C.STD_OUTPUT_HANDLE)), "")

		err := jni.LoadJVMLibrary("jvm.dll")

		if err != nil {
			fmt.Println("Err:", err)
		}

		err = jni.InitializeFromCreatedVM()
		if err != nil {
			fmt.Println("Err: ", err)
		}
		
		classes, _ := jvmti.GetLoadedClasses()

		fmt.Println(classes)
	}()
	
}

func use(i int) {

}


var (
	trinkets map[string]int = map[string]int{
		"": 5,
	}
)

func main() {

}
