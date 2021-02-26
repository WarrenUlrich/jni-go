package internal

/*
	#include "jni_wrapper.h"
*/
import "C"
import (
	"encoding/binary"
)

//Value ...
type Value = C.jvalue

//toValue ...
func toValue(arg interface{}) Value {
	var result [8]byte
	temp := make([]byte, 8)

	switch t := arg.(type) {
	case bool:
		if t {
			temp[0] = 1
		} else {
			temp[0] = 0
		}

	case []bool:

	case int8:
		temp[0] = byte(t)

	case []int8:
		
	case uint8:
		temp[0] = byte(t)
	case []uint8:

	case int16:
		binary.BigEndian.PutUint16(temp, uint16(t))
	case []int16:

	case uint16:
		binary.BigEndian.PutUint16(temp, t)
	case []uint16:

	case int32:
		binary.BigEndian.PutUint32(temp, uint32(t))
	case []int32:

	case uint32:
		binary.BigEndian.PutUint32(temp, t)
	case []uint32:

	case int64:
		binary.BigEndian.PutUint64(temp, uint64(t))
	case []int64:

	case uint64:
		binary.BigEndian.PutUint64(temp, t)
	case []uint64:

	case Object:
		binary.BigEndian.PutUint64(temp, uint64(t))

	case *Object:
		binary.BigEndian.PutUint64(temp, uint64(*t))
	}

	copy(result[:], temp)
	return result
}

//toValues ...
func toValues(args ...interface{}) []Value {
	result := make([]Value, len(args))

	for i, arg := range args {
		result[i] = toValue(arg)
	}

	return result
}