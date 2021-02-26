package internal

/*
	#include "jni_wrapper.h"
	#include <stdlib.h>
*/
import "C"

import "unsafe"

//JNIEnvironment ...
type JNIEnvironment = C.JNIEnv

//jbool converts a C.jboolean to a go bool.
func jboolToBool(b C.jboolean) bool {
	return b == C.JNI_TRUE
}

func boolToJBool(b bool) C.jboolean {
	if b {
		return C.JNI_TRUE
	}

	return C.JNI_FALSE
}

//GetVersion ...
func GetVersion(env *JNIEnvironment) int {
	return int(C.jni_get_version(env))
}

//FindClass ...
func FindClass(env *JNIEnvironment, name string) Class {
	str := C.CString(name)
	cls := C.jni_find_class(env, str)
	C.free(unsafe.Pointer(str))
	return (Class)(cls)
}

//NewGlobalRef ...
func NewGlobalRef(env *JNIEnvironment, obj Object) Object {
	return C.jni_new_global_ref(env, obj)
}

//DeleteGlobalRef ...
func DeleteGlobalRef(env *JNIEnvironment, obj Object) {
	C.jni_delete_global_ref(env, obj)
}

//DeleteLocalRef ...
func DeleteLocalRef(env *JNIEnvironment, obj Object) {
	C.jni_delete_local_ref(env, obj)
}

//IsSameObject ...
func IsSameObject(env *JNIEnvironment, obj1, obj2 Object) bool {
	return jboolToBool(C.jni_is_same_object(env, obj1, obj2))
}

//IsInstanceOf ...
func IsInstanceOf(env *JNIEnvironment, obj Object, cls Class) bool {
	return jboolToBool(C.jni_is_instance_of(env, obj, cls))
}

//GetMethodID ...
func GetMethodID(env *JNIEnvironment, cls Class, name string, sig string) MethodID {
	cName := C.CString(name)
	cSig := C.CString(sig)

	res := C.jni_get_method_id(env, cls, cName, cSig)
	C.free(unsafe.Pointer(cName))
	C.free(unsafe.Pointer(cSig))
	return res
}

//CallObjectMethod ...
func CallObjectMethod(env *JNIEnvironment, obj Object, id MethodID, args ...interface{}) Object {
	values := toValues(args...)
	return C.jni_call_object_method(env, obj, id, &values[0])
}

//CallBooleanMethod ...
func CallBooleanMethod(env *JNIEnvironment, obj Object, id MethodID, args ...interface{}) bool {
	values := toValues(args...)
	return jboolToBool(C.jni_call_boolean_method(env, obj, id, &values[0]))
}

//CallByteMethod ...
func CallByteMethod(env *JNIEnvironment, obj Object, id MethodID, args ...interface{}) byte {
	values := toValues(args...)
	return byte(C.jni_call_byte_method(env, obj, id, &values[0]))
}

//CallCharMethod ...
func CallCharMethod(env *JNIEnvironment, obj Object, id MethodID, args ...interface{}) rune {
	values := toValues(args...)
	return rune(C.jni_call_char_method(env, obj, id, &values[0]))
}

//CallShortMethod ...
func CallShortMethod(env *JNIEnvironment, obj Object, id MethodID, args ...interface{}) int16 {
	values := toValues(args...)
	return int16(C.jni_call_short_method(env, obj, id, &values[0]))
}

//CallIntMethod ...
func CallIntMethod(env *JNIEnvironment, obj Object, id MethodID, args ...interface{}) int {
	values := toValues(args...)
	return int(C.jni_call_int_method(env, obj, id, &values[0]))
}

//CallLongMethod ...
func CallLongMethod(env *JNIEnvironment, obj Object, id MethodID, args ...interface{}) int64 {
	values := toValues(args...)
	return int64(C.jni_call_long_method(env, obj, id, &values[0]))
}

//CallFloatMethod ...
func CallFloatMethod(env *JNIEnvironment, obj Object, id MethodID, args ...interface{}) float32 {
	values := toValues(args...)
	return float32(C.jni_call_float_method(env, obj, id, &values[0]))
}

//CallDoubleMethod ...
func CallDoubleMethod(env *JNIEnvironment, obj Object, id MethodID, args ...interface{}) float64 {
	values := toValues(args...)
	return float64(C.jni_call_double_method(env, obj, id, &values[0]))
}

//CallVoidMethod ...
func CallVoidMethod(env *JNIEnvironment, obj Object, id MethodID, args ...interface{}) {
	values := toValues(args...)
	C.jni_call_void_method(env, obj, id, &values[0])
}

//CallNonVirtualObjectMethod ...
func CallNonVirtualObjectMethod(env *JNIEnvironment, obj Object, cls Class, id MethodID, args ...interface{}) Object {
	values := toValues(args...)
	return C.jni_call_non_virtual_object_method(env, obj, cls, id, &values[0])
}

//CallNonVirtualBooleanMethod ...
func CallNonVirtualBooleanMethod(env *JNIEnvironment, obj Object, cls Class, id MethodID, args ...interface{}) bool {
	values := toValues(args...)
	return jboolToBool(C.jni_call_non_virtual_boolean_method(env, obj, cls, id, &values[0]))
}

//CallNonVirtualByteMethod ...
func CallNonVirtualByteMethod(env *JNIEnvironment, obj Object, cls Class, id MethodID, args ...interface{}) byte {
	values := toValues(args...)
	return byte(C.jni_call_non_virtual_byte_method(env, obj, cls, id, &values[0]))
}

//CallNonVirtualCharMethod ...
func CallNonVirtualCharMethod(env *JNIEnvironment, obj Object, cls Class, id MethodID, args ...interface{}) rune {
	values := toValues(args...)
	return rune(C.jni_call_non_virtual_char_method(env, obj, cls, id, &values[0]))
}

//CallNonVirtualShortMethod ...
func CallNonVirtualShortMethod(env *JNIEnvironment, obj Object, cls Class, id MethodID, args ...interface{}) int16 {
	values := toValues(args...)
	return int16(C.jni_call_non_virtual_short_method(env, obj, cls, id, &values[0]))
}

//CallNonVirtualIntMethod ...
func CallNonVirtualIntMethod(env *JNIEnvironment, obj Object, cls Class, id MethodID, args ...interface{}) int {
	values := toValues(args...)
	return int(C.jni_call_non_virtual_int_method(env, obj, cls, id, &values[0]))
}

//CallNonVirtualLongMethod ...
func CallNonVirtualLongMethod(env *JNIEnvironment, obj Object, cls Class, id MethodID, args ...interface{}) int64 {
	values := toValues(args...)
	return int64(C.jni_call_non_virtual_long_method(env, obj, cls, id, &values[0]))
}

//CallNonVirtualFloatMethod ...
func CallNonVirtualFloatMethod(env *JNIEnvironment, obj Object, cls Class, id MethodID, args ...interface{}) float32 {
	values := toValues(args...)
	return float32(C.jni_call_non_virtual_float_method(env, obj, cls, id, &values[0]))
}

//CallNonVirtualDoubleMethod ...
func CallNonVirtualDoubleMethod(env *JNIEnvironment, obj Object, cls Class, id MethodID, args ...interface{}) float64 {
	values := toValues(args...)
	return float64(C.jni_call_non_virtual_double_method(env, obj, cls, id, &values[0]))
}

//GetFieldID ...
func GetFieldID(env *JNIEnvironment, cls Class, name, sig string) FieldID {
	cname := C.CString(name)
	csig := C.CString(sig)
	res := C.jni_get_field_id(env, cls, cname, csig)
	C.free(unsafe.Pointer(cname))
	C.free(unsafe.Pointer(csig))
	return res
}

//GetObjectField ...
func GetObjectField(env *JNIEnvironment, obj Object, id FieldID) Object {
	return C.jni_get_object_field(env, obj, id)
}

//GetBooleanField ...
func GetBooleanField(env *JNIEnvironment, obj Object, id FieldID) bool {
	return jboolToBool(C.jni_get_boolean_field(env, obj, id))
}

//GetByteField ...
func GetByteField(env *JNIEnvironment, obj Object, id FieldID) byte {
	return byte(C.jni_get_byte_field(env, obj, id))
}

//GetCharField ...
func GetCharField(env *JNIEnvironment, obj Object, id FieldID) rune {
	return rune(C.jni_get_char_field(env, obj, id))
}

//GetShortField ...
func GetShortField(env *JNIEnvironment, obj Object, id FieldID) int16 {
	return int16(C.jni_get_short_field(env, obj, id))
}

//GetIntField ...
func GetIntField(env *JNIEnvironment, obj Object, id FieldID) int {
	return int(C.jni_get_int_field(env, obj, id))
}

//GetLongField ...
func GetLongField(env *JNIEnvironment, obj Object, id FieldID) int64 {
	return int64(C.jni_get_long_field(env, obj, id))
}

//GetFloatField ...
func GetFloatField(env *JNIEnvironment, obj Object, id FieldID) float32 {
	return float32(C.jni_get_float_field(env, obj, id))
}

//GetDoubleField ...
func GetDoubleField(env *JNIEnvironment, obj Object, id FieldID) float64 {
	return float64(C.jni_get_double_field(env, obj, id))
}

//SetObjectField ...
func SetObjectField(env *JNIEnvironment, obj Object, id FieldID, val Object) {
	C.jni_set_object_field(env, obj, id, val)
}

//SetBooleanField ...
func SetBooleanField(env *JNIEnvironment, obj Object, id FieldID, val bool) {
	C.jni_set_boolean_field(env, obj, id, boolToJBool(val))
}

//SetByteField ...
func SetByteField(env *JNIEnvironment, obj Object, id FieldID, val byte) {
	C.jni_set_byte_field(env, obj, id, C.schar(val))
}

//SetCharField ...
func SetCharField(env *JNIEnvironment, obj Object, id FieldID, val rune) {
	C.jni_set_char_field(env, obj, id, C.ushort(val))
}

//SetShortField ...
func SetShortField(env *JNIEnvironment, obj Object, id FieldID, val int16) {
	C.jni_set_short_field(env, obj, id, C.short(val))
}

//SetIntField ...
func SetIntField(env *JNIEnvironment, obj Object, id FieldID, val int) {
	C.jni_set_int_field(env, obj, id, C.long(val))
}

//SetLongField ...
func SetLongField(env *JNIEnvironment, obj Object, id FieldID, val int64) {
	C.jni_set_long_field(env, obj, id, C.longlong(val))
}

//SetFloatField ...
func SetFloatField(env *JNIEnvironment, obj Object, id FieldID, val float32) {
	C.jni_set_float_field(env, obj, id, C.float(val))
}

//SetDoubleField ...
func SetDoubleField(env *JNIEnvironment, obj Object, id FieldID, val float64) {
	C.jni_set_double_field(env, obj, id, C.double(val))
}

//GetStaticMethodID ...
func GetStaticMethodID(env *JNIEnvironment, cls Class, name, sig string) MethodID {
	cname := C.CString(name)
	csig := C.CString(sig)
	
	res := C.jni_get_static_method_id(env, cls, cname, csig)

	C.free(unsafe.Pointer(cname))
	C.free(unsafe.Pointer(csig))
	return res
}

//CallStaticObjectMethod ...
func CallStaticObjectMethod(env *JNIEnvironment, cls Class, id MethodID, args ...interface{}) Object {
	values := toValues(args...)
	return C.jni_call_static_object_method(env, cls, id, &values[0])
}

//CallStaticBooleanMethod ...
func CallStaticBooleanMethod(env *JNIEnvironment, cls Class, id MethodID, args ...interface{}) bool {
	values := toValues(args...)
	return jboolToBool(C.jni_call_static_boolean_method(env, cls, id, &values[0]))
}

//CallStaticByteMethod ...
func CallStaticByteMethod(env *JNIEnvironment, cls Class, id MethodID, args ...interface{}) byte {
	values := toValues(args...)
	return byte(C.jni_call_static_byte_method(env, cls, id, &values[0]))
}

//CallStaticCharMethod ...
func CallStaticCharMethod(env *JNIEnvironment, cls Class, id MethodID, args ...interface{}) rune {
	values := toValues(args...)
	return rune(C.jni_call_static_char_method(env, cls, id, &values[0]))
}

//CallStaticShortMethod ...
func CallStaticShortMethod(env *JNIEnvironment, cls Class, id MethodID, args ...interface{}) int16 {
	values := toValues(args...)
	return int16(C.jni_call_static_byte_method(env, cls, id, &values[0]))
}

//CallStaticIntMethod ...
func CallStaticIntMethod(env *JNIEnvironment, cls Class, id MethodID, args ...interface{}) int {
	values := toValues(args...)
	return int(C.jni_call_static_int_method(env, cls, id, &values[0]))
}

//CallStaticLongMethod ...
func CallStaticLongMethod(env *JNIEnvironment, cls Class, id MethodID, args ...interface{}) int64 {
	values := toValues(args...)
	return int64(C.jni_call_static_long_method(env, cls, id, &values[0]))
}

//CallStaticFloatMethod ...
func CallStaticFloatMethod(env *JNIEnvironment, cls Class, id MethodID, args ...interface{}) float32 {
	values := toValues(args...)
	return float32(C.jni_call_static_float_method(env, cls, id, &values[0]))
}

//CallStaticDoubleMethod ...
func CallStaticDoubleMethod(env *JNIEnvironment, cls Class, id MethodID, args ...interface{}) float64 {
	values := toValues(args...)
	return float64(C.jni_call_static_double_method(env, cls, id, &values[0]))
}

//CallStaticVoidMethod ...
func CallStaticVoidMethod(env *JNIEnvironment, cls Class, id MethodID, args ...interface{}) {
	values := toValues(args...)
	C.jni_call_static_void_method(env, cls, id, &values[0])
}

//GetStaticFieldID ...
func GetStaticFieldID(env *JNIEnvironment, cls Class, name, sig string) FieldID {
	cname := C.CString(name)
	csig := C.CString(sig)
	res := C.jni_get_static_field_id(env, cls, cname, csig)
	
	C.free(unsafe.Pointer(cname))
	C.free(unsafe.Pointer(csig))
	return res
}

//GetStaticObjectField ...
func GetStaticObjectField(env *JNIEnvironment, cls Class, id FieldID) Object {
	return C.jni_get_static_object_field(env, cls, id)
}

//GetStaticBooleanField ...
func GetStaticBooleanField(env *JNIEnvironment, cls Class, id FieldID) bool {
	return jboolToBool(C.jni_get_static_boolean_field(env, cls, id))
}

//GetStaticByteField ...
func GetStaticByteField(env *JNIEnvironment, cls Class, id FieldID) byte {
	return byte(C.jni_get_static_byte_field(env, cls, id))
}

//GetStaticCharField ...
func GetStaticCharField(env *JNIEnvironment, cls Class, id FieldID) rune {
	return rune(C.jni_get_static_char_field(env, cls, id))
}

//GetStaticShortField ...
func GetStaticShortField(env *JNIEnvironment, cls Class, id FieldID) int16 {
	return int16(C.jni_get_static_short_field(env, cls, id))
}

//GetStaticIntField ...
func GetStaticIntField(env *JNIEnvironment, cls Class, id FieldID) int {
	return int(C.jni_get_static_int_field(env, cls, id))
}

//GetStaticLongField ...
func GetStaticLongField(env *JNIEnvironment, cls Class, id FieldID) int64 {
	return int64(C.jni_get_static_boolean_field(env, cls, id))
}

//GetStaticFloatField ...
func GetStaticFloatField(env *JNIEnvironment, cls Class, id FieldID) float32 {
	return float32(C.jni_get_static_boolean_field(env, cls, id))
}

//GetStaticDoubleField ...
func GetStaticDoubleField(env *JNIEnvironment, cls Class, id FieldID) float64 {
	return float64(C.jni_get_static_double_field(env, cls, id))
}

//SetStaticObjectField ...
func SetStaticObjectField(env *JNIEnvironment, cls Class, id FieldID, val Object) {
	C.jni_set_static_object_field(env, cls, id, val)
}

//SetStaticBooleanField ...
func SetStaticBooleanField(env *JNIEnvironment, cls Class, id FieldID, val bool) {
	C.jni_set_static_boolean_field(env, cls, id, boolToJBool(val))
}

//SetStaticByteField ...
func SetStaticByteField(env *JNIEnvironment, cls Class, id FieldID, val byte) {
	C.jni_set_static_byte_field(env, cls, id, C.jbyte(val))
}

//SetStaticCharField ...
func SetStaticCharField(env *JNIEnvironment, cls Class, id FieldID, val rune) {
	C.jni_set_static_char_field(env, cls, id, C.jchar(val))
}

//SetStaticShortField ...
func SetStaticShortField(env *JNIEnvironment, cls Class, id FieldID, val int16) {
	C.jni_set_static_short_field(env, cls, id, C.jshort(val))
}

//SetStaticIntField ...
func SetStaticIntField(env *JNIEnvironment, cls Class, id FieldID, val int) {
	C.jni_set_static_int_field(env, cls, id, C.jint(val))
}

//SetStaticLongField ...
func SetStaticLongField(env *JNIEnvironment, cls Class, id FieldID, val int64) {
	C.jni_set_static_long_field(env, cls, id, C.jlong(val))
}

//SetStaticFloatField ...
func SetStaticFloatField(env *JNIEnvironment, cls Class, id FieldID, val float32) {
	C.jni_set_static_float_field(env, cls, id, C.jfloat(val))
}

//SetStaticDoubleField ...
func SetStaticDoubleField(env *JNIEnvironment, cls Class, id FieldID, val float64) {
	C.jni_set_static_double_field(env, cls, id, C.jdouble(val))
}

//NewString ...
func NewString(env *JNIEnvironment, str string) String {
	cstr := C.CString(str)
	res := C.jni_new_string_utf(env, cstr)
	C.free(unsafe.Pointer(cstr))
	return (String)(res)
}

//GetStringLength ...
func GetStringLength(env *JNIEnvironment, str String) int {
	return int(C.jni_get_string_utf_length(env, str))
}

//GetStringChars ...
func GetStringChars(env *JNIEnvironment, str String) string {
	var isCopy C.jboolean

	cstr := C.jni_get_string_utf_chars(env, str, &isCopy)
	res := C.GoString(cstr)
	C.jni_release_string_utf_chars(env, str, cstr)
	return res
}

//GetArrayLength ...
func GetArrayLength(env *JNIEnvironment, array Array) int {
	return int(C.jni_get_array_length(env, array))
}

//NewObjectArray ...
func NewObjectArray(env *JNIEnvironment, len int, cls Class, init Object) ObjectArray {
	return C.jni_new_object_array(env, C.jsize(len), cls, init)
}

//GetObjectArrayElement ...
func GetObjectArrayElement(env *JNIEnvironment, arr ObjectArray, index int) Object {
	return C.jni_get_object_array_element(env, arr, C.jsize(index))
}

//SetObjectArrayElement ...
func SetObjectArrayElement(env *JNIEnvironment, arr ObjectArray, index int, val Object) {
	C.jni_set_object_array_element(env, arr, C.jsize(index), val)
}

//NewBooleanArray ...
func NewBooleanArray(env *JNIEnvironment, size int) BooleanArray {
	return C.jni_new_boolean_array(env, C.jsize(size))
}

//NewByteArray ...
func NewByteArray(env *JNIEnvironment, size int) ByteArray {
	return C.jni_new_byte_array(env, C.jsize(size))
}

//NewCharArray ...
func NewCharArray(env *JNIEnvironment, size int) CharArray {
	return C.jni_new_char_array(env, C.jsize(size))
}

//NewShortArray ...
func NewShortArray(env *JNIEnvironment, size int) ShortArray {
	return C.jni_new_short_array(env, C.jsize(size))
}

//NewLongArray ..
func NewLongArray(env *JNIEnvironment, size int) LongArray {
	return C.jni_new_long_array(env, C.jsize(size))
}

//NewFloatArray ...
func NewFloatArray(env *JNIEnvironment, size int) FloatArray {
	return C.jni_new_float_array(env, C.jsize(size))
}

//NewDoubleArray ...
func NewDoubleArray(env *JNIEnvironment, size int) DoubleArray {
	return C.jni_new_double_array(env, C.jsize(size))
}

//GetBooleanArrayElements ...
func GetBooleanArrayElements(env *JNIEnvironment, arr BooleanArray) []bool {
	len := GetArrayLength(env, Array(arr))
	arrPtr := C.jni_get_boolean_array_elements(env, arr, nil)
	array := (*[1 << 28]C.jboolean)(unsafe.Pointer(arrPtr))[:len:len]

	result := make([]bool, len)
	for i, e := range array {
		result[i] = jboolToBool(e)
	}
	C.jni_release_boolean_array_elements(env, arr, arrPtr, 0)
	return result
}

//GetByteArrayElements ...
func GetByteArrayElements(env *JNIEnvironment, arr ByteArray) []byte {
	len := GetArrayLength(env, Array(arr))
	arrPtr := C.jni_get_byte_array_elements(env, arr, nil)
	array := (*[1 << 28]C.jbyte)(unsafe.Pointer(arrPtr))[:len:len]

	result := make([]byte, len)
	for i, e := range array {
		result[i] = byte(e)
	}
	C.jni_release_byte_array_elements(env, arr, arrPtr, 0)
	return result
}

//GetCharArrayElements ...
func GetCharArrayElements(env *JNIEnvironment, arr CharArray) []rune {
	len := GetArrayLength(env, Array(arr))
	arrPtr := C.jni_get_char_array_elements(env, arr, nil)
	array := (*[1 << 28]C.jchar)(unsafe.Pointer(arrPtr))[:len:len]

	result := make([]rune, len)
	for i, e := range array {
		result[i] = rune(e)
	}
	C.jni_release_char_array_elements(env, arr, arrPtr, 0)
	return result
}

//GetShortArrayElements ...
func GetShortArrayElements(env *JNIEnvironment, arr ShortArray) []int16 {
	len := GetArrayLength(env, Array(arr))
	arrPtr := C.jni_get_short_array_elements(env, arr, nil)
	array := (*[1 << 28]C.jshort)(unsafe.Pointer(arrPtr))[:len:len]

	result := make([]int16, len)
	for i, e := range array {
		result[i] = int16(e)
	}
	C.jni_release_short_array_elements(env, arr, arrPtr, 0)
	return result
}

//GetIntArrayElements ...
func GetIntArrayElements(env *JNIEnvironment, arr IntArray) []int {
	len := GetArrayLength(env, Array(arr))
	arrPtr := C.jni_get_int_array_elements(env, arr, nil)
	array := (*[1 << 28]C.jint)(unsafe.Pointer(arrPtr))[:len:len]

	result := make([]int, len)
	for i, e := range array {
		result[i] = int(e)
	}
	C.jni_release_int_array_elements(env, arr, arrPtr, 0)
	return result
}

//GetLongArrayElements ...
func GetLongArrayElements(env *JNIEnvironment, arr LongArray) []int64 {
	len := GetArrayLength(env, Array(arr))
	arrPtr := C.jni_get_long_array_elements(env, arr, nil)
	array := (*[1 << 14]C.jlong)(unsafe.Pointer(arrPtr))[:len:len]

	result := make([]int64, len)
	for i, e := range array {
		result[i] = int64(e)
	}
	C.jni_release_long_array_elements(env, arr, arrPtr, 0)
	return result
}

//GetFloatArrayElements ...
func GetFloatArrayElements(env *JNIEnvironment, arr FloatArray) []float32 {
	len := GetArrayLength(env, Array(arr))
	arrPtr := C.jni_get_float_array_elements(env, arr, nil)
	array := (*[1 << 28]C.jfloat)(unsafe.Pointer(arrPtr))[:len:len]

	result := make([]float32, len)
	for i, e := range array {
		result[i] = float32(e)
	}
	C.jni_release_float_array_elements(env, arr, arrPtr, 0)
	return result
}

//GetDoubleArrayElements ...
func GetDoubleArrayElements(env *JNIEnvironment, arr DoubleArray) []float64 {
	len := GetArrayLength(env, Array(arr))
	arrPtr := C.jni_get_double_array_elements(env, arr, nil)

	result := make([]float64, len)
	array := (*[1 << 14]C.jdouble)(unsafe.Pointer(arrPtr))[:len:len]

	for i, e := range array {
		result[i] = float64(e)
	}

	C.jni_release_double_array_elements(env, arr, arrPtr, 0)
	return result
}

//GetBooleanArrayRegion ...
func GetBooleanArrayRegion(env *JNIEnvironment, arr BooleanArray, start, len int) []bool {
	temp := make([]C.jboolean, len)

	C.jni_get_boolean_array_region(
		env,
		arr,
		C.jsize(start),
		C.jsize(len),
		(*C.jboolean)(unsafe.Pointer(&temp[0])),
	)

	result := make([]bool, len)
	for i := range temp {
		result[i] = jboolToBool(temp[i])
	}

	return result
}

//GetByteArrayRegion ...
func GetByteArrayRegion(env *JNIEnvironment, arr ByteArray, start, len int) []byte {
	temp := make([]C.jbyte, len)

	C.jni_get_byte_array_region(
		env,
		arr,
		C.jsize(start),
		C.jsize(len),
		(*C.jbyte)(unsafe.Pointer(&temp[0])),
	)

	result := make([]byte, len)
	for i := range temp {
		result[i] = byte(temp[i])
	}

	return result
}

//GetCharArrayRegion ...
func GetCharArrayRegion(env *JNIEnvironment, arr CharArray, start, len int) []rune {
	temp := make([]C.jchar, len)

	C.jni_get_char_array_region(
		env,
		arr,
		C.jsize(start),
		C.jsize(len),
		(*C.jchar)(unsafe.Pointer(&temp[0])),
	)

	result := make([]rune, len)
	for i := range temp {
		result[i] = rune(temp[i])
	}

	return result
}

//GetShortArrayRegion ...
func GetShortArrayRegion(env *JNIEnvironment, arr ShortArray, start, len int) []int16 {
	temp := make([]C.jshort, len)

	C.jni_get_short_array_region(
		env,
		arr,
		C.jsize(start),
		C.jsize(len),
		(*C.jshort)(unsafe.Pointer(&temp[0])),
	)

	result := make([]int16, len)
	for i := range temp {
		result[i] = int16(temp[i])
	}

	return result
}

//GetIntArrayRegion ...
func GetIntArrayRegion(env *JNIEnvironment, arr IntArray, start, len int) []int {
	temp := make([]C.jint, len)

	C.jni_get_int_array_region(
		env,
		arr,
		C.jsize(start),
		C.jsize(len),
		(*C.jint)(unsafe.Pointer(&temp[0])),
	)

	result := make([]int, len)
	for i := range temp {
		result[i] = int(temp[i])
	}

	return result
}

//GetLongArrayRegion ...
func GetLongArrayRegion(env *JNIEnvironment, arr LongArray, start, len int) []int64 {
	temp := make([]C.jlong, len)

	C.jni_get_long_array_region(
		env,
		arr,
		C.jsize(start),
		C.jsize(len),
		(*C.jlong)(unsafe.Pointer(&temp[0])),
	)

	result := make([]int64, len)
	for i := range temp {
		result[i] = int64(temp[i])
	}

	return result
}

//GetFloatArrayRegion ...
func GetFloatArrayRegion(env *JNIEnvironment, arr FloatArray, start, len int) []float32 {
	temp := make([]C.jfloat, len)

	C.jni_get_float_array_region(
		env,
		arr,
		C.jsize(start),
		C.jsize(len),
		(*C.jfloat)(unsafe.Pointer(&temp[0])),
	)

	result := make([]float32, len)
	for i := range temp {
		result[i] = float32(temp[i])
	}

	return result
}

//GetDoubleArrayRegion ...
func GetDoubleArrayRegion(env *JNIEnvironment, arr DoubleArray, start, len int) []float64 {
	temp := make([]C.jdouble, len)

	C.jni_get_double_array_region(
		env,
		arr,
		C.jsize(start),
		C.jsize(len),
		(*C.jdouble)(unsafe.Pointer(&temp[0])),
	)

	result := make([]float64, len)
	for i := range temp {
		result[i] = float64(temp[i])
	}

	return result
}