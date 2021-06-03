package dllmain

/*
	#include "dllmain.h"
*/
import "C"

var ThreadDetach func()

//export threadDetach
func threadDetach() {
	ThreadDetach()
}