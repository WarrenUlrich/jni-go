package jni

import (
	"runtime"
	"github.com/warrenulrich/jni-go/pkg/internal"
)

type methodCacheKey struct {
	class Class
	name  string
	sig   string
}

type fieldCacheKey struct {
	class Class
	name  string
	sig   string
}

var (
	//JavaVM ...
	JavaVM      *internal.JavaVM
	classCache  map[string]*Class           = make(map[string]*Class)
	methodCache map[methodCacheKey]MethodID = make(map[methodCacheKey]MethodID)
	fieldCache  map[fieldCacheKey]FieldID   = make(map[fieldCacheKey]FieldID)
)

func getEnvironment() *internal.JNIEnvironment {
	env, err := internal.GetEnv(JavaVM)
	if err != nil {
		if err == internal.ErrDetached {
			env, _ = internal.AttachCurrentThread(JavaVM)
		}
	}

	return env
}

//LoadJVMLibrary ...
func LoadJVMLibrary(path string) error {
	return internal.LoadJVMLibrary(path)
}

//InitializeFromCreatedVM ...
func InitializeFromCreatedVM() error {
	jvm, err := internal.GetCreatedVM()
	if err != nil {
		return err
	}

	JavaVM = jvm
	_, err = internal.AttachCurrentThread(JavaVM)
	if err != nil {
		return err
	}

	return nil
}

//GetVersion ...
func GetVersion() int {
	return internal.GetVersion(getEnvironment())
}

//FindClass ...
func FindClass(name string) *Class {
	if c, found := classCache[name]; found {
		return c
	}

	var clsLocal Class = (Class)(internal.FindClass(getEnvironment(), name))

	clsGlobal := NewGlobalRef((Object)(clsLocal))
	DeleteLocalRef((Object)(clsLocal))
	classCache[name] = (*Class)(clsGlobal)
	return (*Class)(clsGlobal)
}

//NewGlobalRef ...
func NewGlobalRef(obj Object) *Object {
	o := (Object)(internal.NewGlobalRef(getEnvironment(), (internal.Object)(obj)))

	runtime.SetFinalizer(&o, func(objPtr *Object) {
		DeleteGlobalRef(*objPtr)
	})

	return &o
}

//DeleteGlobalRef ...
func DeleteGlobalRef(obj Object) {
	internal.DeleteGlobalRef(getEnvironment(), (internal.Object)(obj))
}

//DeleteLocalRef  ...
func DeleteLocalRef(obj Object) {
	internal.DeleteLocalRef(getEnvironment(), (internal.Object)(obj))
}

//IsSameObject ...
func IsSameObject(obj1, obj2 Object) bool {
	return internal.IsSameObject(
		getEnvironment(),
		(internal.Object)(obj1),
		(internal.Object)(obj2),
	)
}

//GetMethodID ...
func GetMethodID(cls Class, name, sig string) MethodID {
	key := methodCacheKey{cls, name, sig}
	if id, found := methodCache[key]; found {
		return id
	}

	id := MethodID(internal.GetMethodID(
		getEnvironment(),
		(internal.Class)(cls),
		name,
		sig,
	))

	methodCache[key] = id
	return id
}

//CallObjectMethod ...
func CallObjectMethod(obj Object, id MethodID, args ...interface{}) *Object {
	local := internal.CallObjectMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	)
	defer DeleteLocalRef((Object)(local))

	return NewGlobalRef((Object)(local))
}

//CallBooleanMethod ...
func CallBooleanMethod(obj Object, id MethodID, args ...interface{}) bool {
	return internal.CallBooleanMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	)
}

//CallByteMethod ...
func CallByteMethod(obj Object, id MethodID, args ...interface{}) byte {
	return internal.CallByteMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	)
}

//CallCharMethod ...
func CallCharMethod(obj Object, id MethodID, args ...interface{}) rune {
	return internal.CallCharMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	)
}

//CallShortMethod ...
func CallShortMethod(obj Object, id MethodID, args ...interface{}) int16 {
	return internal.CallShortMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	)
}

//CallIntMethod ...
func CallIntMethod(obj Object, id MethodID, args ...interface{}) int {
	return internal.CallIntMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	)
}

//CallLongMethod ...
func CallLongMethod(obj Object, id MethodID, args ...interface{}) int64 {
	return internal.CallLongMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	)
}

//CallFloatMethod ...
func CallFloatMethod(obj Object, id MethodID, args ...interface{}) float32 {
	return internal.CallFloatMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	)
}

//CallDoubleMethod ...
func CallDoubleMethod(obj Object, id MethodID, args ...interface{}) float64 {
	return internal.CallDoubleMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	)
}

//CallVoidMethod ...
func CallVoidMethod(obj Object, id MethodID, args ...interface{}) {
	internal.CallVoidMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	)
}

//CallNonVirtualObjectMethod ...
func CallNonVirtualObjectMethod(obj Object, cls Class, id MethodID, args ...interface{}) *Object {
	local := Object(internal.CallNonVirtualObjectMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	))

	global := NewGlobalRef(local)
	DeleteLocalRef(local)
	return global
}

//CallNonVirtualBooleanMethod ...
func CallNonVirtualBooleanMethod(obj Object, cls Class, id MethodID, args ...interface{}) bool {
	return internal.CallNonVirtualBooleanMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	)
}

//CallNonVirtualByteMethod ...
func CallNonVirtualByteMethod(obj Object, cls Class, id MethodID, args ...interface{}) byte {
	return internal.CallNonVirtualByteMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	)
}

//CallNonVirtualCharMethod ...
func CallNonVirtualCharMethod(obj Object, cls Class, id MethodID, args ...interface{}) rune {
	return internal.CallNonVirtualCharMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	)
}

//CallNonVirtualShortMethod ...
func CallNonVirtualShortMethod(obj Object, cls Class, id MethodID, args ...interface{}) int16 {
	return internal.CallNonVirtualShortMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	)
}

//CallNonVirtualIntMethod ...
func CallNonVirtualIntMethod(obj Object, cls Class, id MethodID, args ...interface{}) int {
	return internal.CallNonVirtualIntMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	)
}

//CallNonVirtualLongMethod ...
func CallNonVirtualLongMethod(obj Object, cls Class, id MethodID, args ...interface{}) int64 {
	return internal.CallNonVirtualLongMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	)
}

//CallNonVirtualFloatMethod ...
func CallNonVirtualFloatMethod(obj Object, cls Class, id MethodID, args ...interface{}) float32 {
	return internal.CallNonVirtualFloatMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	)
}

//CallNonVirtualDoubleMethod ...
func CallNonVirtualDoubleMethod(obj Object, cls Class, id MethodID, args ...interface{}) float64 {
	return internal.CallNonVirtualDoubleMethod(
		getEnvironment(),
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	)
}

//GetFieldID ...
func GetFieldID(cls Class, name, sig string) FieldID {
	key := fieldCacheKey{cls, name, sig}
	if id, found := fieldCache[key]; found {
		return id
	}

	id := FieldID(internal.GetFieldID(
		getEnvironment(),
		(internal.Class)(cls),
		name,
		sig,
	))

	fieldCache[key] = id
	return id
}

//GetObjectField ...
func GetObjectField(obj Object, id FieldID) *Object {
	local := Object(internal.GetObjectField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
	))

	global := NewGlobalRef(local)
	DeleteLocalRef(local)
	return global
}

//GetBooleanField ...
func GetBooleanField(obj Object, id FieldID) bool {
	return internal.GetBooleanField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
	)
}

//GetByteField ...
func GetByteField(obj Object, id FieldID) byte {
	return internal.GetByteField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
	)
}

//GetCharField ...
func GetCharField(obj Object, id FieldID) rune {
	return internal.GetCharField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
	)
}

//GetShortField ...
func GetShortField(obj Object, id FieldID) int16 {
	return internal.GetShortField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
	)
}

//GetIntField ...
func GetIntField(obj Object, id FieldID) int {
	return internal.GetIntField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
	)
}

//GetLongField ...
func GetLongField(obj Object, id FieldID) int64 {
	return internal.GetLongField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
	)
}

//GetFloatField ...
func GetFloatField(obj Object, id FieldID) float32 {
	return internal.GetFloatField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
	)
}

//GetDoubleField ...
func GetDoubleField(obj Object, id FieldID) float64 {
	return internal.GetDoubleField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
	)
}

//SetObjectField ...
func SetObjectField(obj Object, id FieldID, val Object) {
	internal.SetObjectField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
		internal.Object(val),
	)
}

//SetBooleanField ...
func SetBooleanField(obj Object, id FieldID, val bool) {
	internal.SetBooleanField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
		val,
	)
}

//SetByteField ...
func SetByteField(obj Object, id FieldID, val byte) {
	internal.SetByteField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
		val,
	)
}

//SetCharField ...
func SetCharField(obj Object, id FieldID, val rune) {
	internal.SetCharField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
		val,
	)
}

//SetShortField ...
func SetShortField(obj Object, id FieldID, val int16) {
	internal.SetShortField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
		val,
	)
}


//SetIntField ...
func SetIntField(obj Object, id FieldID, val int) {
	internal.SetIntField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
		val,
	)
}

//SetLongField ...
func SetLongField(obj Object, id FieldID, val int64) {
	internal.SetLongField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
		val,
	)
}

//SetFloatField ...
func SetFloatField(obj Object, id FieldID, val float32) {
	internal.SetFloatField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
		val,
	)
}

//SetDoubleField ...
func SetDoubleField(obj Object, id FieldID, val float64) {
	internal.SetDoubleField(
		getEnvironment(),
		internal.Object(obj),
		internal.FieldID(id),
		val,
	)
}

//GetStaticMethodID ...
func GetStaticMethodID(cls Class, name, sig string) MethodID {
	key := methodCacheKey{cls, name, sig}
	if id, found := methodCache[key]; found {
		return id
	}

	id := MethodID(internal.GetStaticMethodID(
		getEnvironment(),
		(internal.Class)(cls),
		name,
		sig,
	))

	methodCache[key] = id
	return id
}

//CallStaticObjectMethod ...
func CallStaticObjectMethod(cls Class, id MethodID, args ...interface{}) *Object {
	local := Object(internal.CallStaticObjectMethod(
		getEnvironment(),
		internal.Class(cls),
		internal.MethodID(id),
		args...
	))

	global := NewGlobalRef(local)
	DeleteLocalRef(local)
	return global
}

//CallStaticBooleanMethod ...
func CallStaticBooleanMethod(cls Class, id MethodID, args ...interface{}) bool {
	return internal.CallStaticBooleanMethod(
		getEnvironment(),
		internal.Class(cls),
		internal.MethodID(id),
		args...
	)
}

//CallStaticByteMethod ...
func CallStaticByteMethod(cls Class, id MethodID, args ...interface{}) byte {
	return internal.CallStaticByteMethod(
		getEnvironment(),
		internal.Class(cls),
		internal.MethodID(id),
		args...
	)
}

//CallStaticCharMethod ...
func CallStaticCharMethod(cls Class, id MethodID, args ...interface{}) rune {
	return internal.CallStaticCharMethod(
		getEnvironment(),
		internal.Class(cls),
		internal.MethodID(id),
		args...
	)
}

//CallStaticShortMethod ...
func CallStaticShortMethod(cls Class, id MethodID, args ...interface{}) int16 {
	return internal.CallStaticShortMethod(
		getEnvironment(),
		internal.Class(cls),
		internal.MethodID(id),
		args...
	)
}

//CallStaticIntMethod ...
func CallStaticIntMethod(cls Class, id MethodID, args ...interface{}) int {
	return internal.CallStaticIntMethod(
		getEnvironment(),
		internal.Class(cls),
		internal.MethodID(id),
		args...
	)
}

//CallStaticLongMethod ...
func CallStaticLongMethod(cls Class, id MethodID, args ...interface{}) int64 {
	return internal.CallStaticLongMethod(
		getEnvironment(),
		internal.Class(cls),
		internal.MethodID(id),
		args...
	)
}

//CallStaticFloatMethod ...
func CallStaticFloatMethod(cls Class, id MethodID, args ...interface{}) float32 {
	return internal.CallStaticFloatMethod(
		getEnvironment(),
		internal.Class(cls),
		internal.MethodID(id),
		args...
	)
}

//CallStaticDoubleMethod ...
func CallStaticDoubleMethod(cls Class, id MethodID, args ...interface{}) float64 {
	return internal.CallStaticDoubleMethod(
		getEnvironment(),
		internal.Class(cls),
		internal.MethodID(id),
		args...
	)
}
 
//GetStaticFieldID ...
func GetStaticFieldID(cls Class, name, sig string) FieldID {
	key := fieldCacheKey{cls, name, sig}
	if id, found := fieldCache[key]; found {
		return id
	}

	id := FieldID(internal.GetStaticFieldID(
		getEnvironment(),
		(internal.Class)(cls),
		name,
		sig,
	))

	fieldCache[key] = id
	return id
}

//GetStaticObjectField ...
func GetStaticObjectField(cls Class, id FieldID) *Object {
	local := Object(internal.GetStaticObjectField(
		getEnvironment(),
		internal.Class(cls),
		internal.FieldID(id),
	))

	global := NewGlobalRef(local)
	DeleteLocalRef(local)
	return global
}

//GetStaticBooleanField ...
func GetStaticBooleanField(cls Class, id FieldID) bool {
	return internal.GetStaticBooleanField(
		getEnvironment(),
		(internal.Class)(cls),
		(internal.FieldID)(id),
	)
}

//GetStaticByteField ...
func GetStaticByteField(cls Class, id FieldID) byte {
	return internal.GetStaticByteField(
		getEnvironment(),
		(internal.Class)(cls),
		(internal.FieldID)(id),
	)
}

//GetStaticCharField ...
func GetStaticCharField(cls Class, id FieldID) rune {
	return internal.GetStaticCharField(
		getEnvironment(),
		(internal.Class)(cls),
		(internal.FieldID)(id),
	)
}

//GetStaticShortField ...
func GetStaticShortField(cls Class, id FieldID) int16 {
	return internal.GetStaticShortField(
		getEnvironment(),
		(internal.Class)(cls),
		(internal.FieldID)(id),
	)
}

//GetStaticIntField ...
func GetStaticIntField(cls Class, id FieldID) int {
	return internal.GetStaticIntField(
		getEnvironment(),
		(internal.Class)(cls),
		(internal.FieldID)(id),
	)
}

//GetStaticDoubleField ...
func GetStaticDoubleField(cls Class, id FieldID) float64 {
	return internal.GetStaticDoubleField(
		getEnvironment(),
		(internal.Class)(cls),
		(internal.FieldID)(id),
	)
}

//SetStaticObjectField ...
func SetStaticObjectField(cls Class, id FieldID, val Object) {
	internal.SetStaticObjectField(
		getEnvironment(),
		internal.Class(cls),
		internal.FieldID(id),
		internal.Object(val),
	)
}

//SetStaticBooleanField ...
func SetStaticBooleanField(cls Class, id FieldID, val bool) {
	internal.SetStaticBooleanField(
		getEnvironment(),
		internal.Class(cls),
		internal.FieldID(id),
		val,
	)
}

//SetStaticByteField ...
func SetStaticByteField(cls Class, id FieldID, val byte) {
	internal.SetStaticByteField(
		getEnvironment(),
		internal.Class(cls),
		internal.FieldID(id),
		val,
	)
}

//SetStaticCharField ...
func SetStaticCharField(cls Class, id FieldID, val rune) {
	internal.SetStaticCharField(
		getEnvironment(),
		internal.Class(cls),
		internal.FieldID(id),
		val,
	)
}

//SetStaticShortField ...
func SetStaticShortField(cls Class, id FieldID, val int16) {
	internal.SetStaticShortField(
		getEnvironment(),
		internal.Class(cls),
		internal.FieldID(id),
		val,
	)
}

//SetStaticIntField ...
func SetStaticIntField(cls Class, id FieldID, val int) {
	internal.SetStaticIntField(
		getEnvironment(),
		internal.Class(cls),
		internal.FieldID(id),
		val,
	)
}

//SetStaticFloatField ...
func SetStaticFloatField(cls Class, id FieldID, val float32) {
	internal.SetStaticFloatField(
		getEnvironment(),
		internal.Class(cls),
		internal.FieldID(id),
		val,
	)
}

//SetStaticDoubleField ...
func SetStaticDoubleField(cls Class, id FieldID, val float64) {
	internal.SetStaticDoubleField(
		getEnvironment(),
		internal.Class(cls),
		internal.FieldID(id),
		val,
	)
}

//NewString ...
func NewString(str string) *String {
	localStr := String(internal.NewString(getEnvironment(), str))

	globalStr := NewGlobalRef((Object)(localStr))

	DeleteLocalRef((Object)(localStr))

	return (*String)(globalStr)
}

//GetStringLength ...
func GetStringLength(str String) int {
	return internal.GetStringLength(
		getEnvironment(),
		internal.String(str),
	)
}

//GetStringChars ...
func GetStringChars(str String) string {
	return internal.GetStringChars(
		getEnvironment(),
		internal.String(str),
	)
}