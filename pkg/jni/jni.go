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
	//JVM ...
	JVM      *internal.JavaVM
	classCache  map[string]*Class           = make(map[string]*Class)
	methodCache map[methodCacheKey]MethodID = make(map[methodCacheKey]MethodID)
	fieldCache  map[fieldCacheKey]FieldID   = make(map[fieldCacheKey]FieldID)
)

func getEnvironment() *internal.JNIEnvironment {
	env, err := internal.GetEnv(JVM)
	if err != nil {
		if err == internal.ErrDetached {
			env, _ = internal.AttachCurrentThread(JVM)
		}
	}

	return env
}

func getThreadLocalEnvironment() (*internal.JNIEnvironment, error) {
	env, err := internal.GetEnv(JVM)
	if err != nil {
		if err == internal.ErrDetached {
			env, err = internal.AttachCurrentThread(JVM)
			if err != nil {
				return nil, err
			}
		}
	}
	return env, nil
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

	JVM = jvm
	_, err = internal.AttachCurrentThread(JVM)
	if err != nil {
		return err
	}

	return nil
}

//GetVersion ...
func GetVersion() (int, error) {

	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}
	return internal.GetVersion(localEnv), nil
}

//FindClass ...
func FindClass(name string) (*Class, error) {
	if c, found := classCache[name]; found {
		return c, nil
	}

	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return nil, err
	}

	var clsLocal Class = (Class)(internal.FindClass(localEnv, name))

	clsGlobal, err := NewGlobalRef((Object)(clsLocal))
	if err != nil {
		return nil, err
	}

	DeleteLocalRef((Object)(clsLocal))
	classCache[name] = (*Class)(clsGlobal)
	return (*Class)(clsGlobal), nil
}

//NewGlobalRef ...
func NewGlobalRef(obj Object) (*Object, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return nil, err
	}

	o := (Object)(internal.NewGlobalRef(localEnv, (internal.Object)(obj)))

	runtime.SetFinalizer(&o, func(objPtr *Object) {
		DeleteGlobalRef(*objPtr)
	})

	return &o, nil
}

//DeleteGlobalRef ...
func DeleteGlobalRef(obj Object) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.DeleteGlobalRef(localEnv, (internal.Object)(obj))
	return nil
}

//DeleteLocalRef  ...
func DeleteLocalRef(obj Object) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}
	internal.DeleteLocalRef(localEnv, (internal.Object)(obj))
	return nil
}

//IsSameObject ...
func IsSameObject(obj1, obj2 Object) (bool, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return false, err
	}

	return internal.IsSameObject(
		localEnv,
		(internal.Object)(obj1),
		(internal.Object)(obj2),
	), nil
}

//GetMethodID ...
func GetMethodID(cls Class, name, sig string) (MethodID, error) {
	key := methodCacheKey{cls, name, sig}
	if id, found := methodCache[key]; found {
		return id, nil
	}

	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return nil, err
	}

	id := MethodID(internal.GetMethodID(
		localEnv,
		(internal.Class)(cls),
		name,
		sig,
	))

	methodCache[key] = id
	return id, nil
}

//CallObjectMethod ...
func CallObjectMethod(obj Object, id MethodID, args ...interface{}) (*Object, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return nil, err
	}

	local := internal.CallObjectMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	)
	defer DeleteLocalRef((Object)(local))

	return NewGlobalRef((Object)(local))
}

//CallBooleanMethod ...
func CallBooleanMethod(obj Object, id MethodID, args ...interface{}) (bool, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return false, err
	}

	return internal.CallBooleanMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	), nil
}

//CallByteMethod ...
func CallByteMethod(obj Object, id MethodID, args ...interface{}) (byte, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallByteMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	), nil
}

//CallCharMethod ...
func CallCharMethod(obj Object, id MethodID, args ...interface{}) (rune, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallCharMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	), nil
}

//CallShortMethod ...
func CallShortMethod(obj Object, id MethodID, args ...interface{}) (int16, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallShortMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	), nil
}

//CallIntMethod ...
func CallIntMethod(obj Object, id MethodID, args ...interface{}) (int, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallIntMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	), nil
}

//CallLongMethod ...
func CallLongMethod(obj Object, id MethodID, args ...interface{}) (int64, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallLongMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	), nil
}

//CallFloatMethod ...
func CallFloatMethod(obj Object, id MethodID, args ...interface{}) (float32, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0.0, err
	}

	return internal.CallFloatMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	), nil
}

//CallDoubleMethod ...
func CallDoubleMethod(obj Object, id MethodID, args ...interface{}) (float64, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0.0, err
	}

	return internal.CallDoubleMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	), nil
}

//CallVoidMethod ...
func CallVoidMethod(obj Object, id MethodID, args ...interface{}) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.CallVoidMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.MethodID)(id),
		args...,
	)

	return nil
}

//CallNonVirtualObjectMethod ...
func CallNonVirtualObjectMethod(obj Object, cls Class, id MethodID, args ...interface{}) (*Object, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return nil, err
	}

	local := Object(internal.CallNonVirtualObjectMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	))

	defer DeleteLocalRef(local)
	return NewGlobalRef(local)
}

//CallNonVirtualBooleanMethod ...
func CallNonVirtualBooleanMethod(obj Object, cls Class, id MethodID, args ...interface{}) (bool, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return false, err
	}

	return internal.CallNonVirtualBooleanMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	), nil
}

//CallNonVirtualByteMethod ...
func CallNonVirtualByteMethod(obj Object, cls Class, id MethodID, args ...interface{}) (byte, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallNonVirtualByteMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	), nil
}

//CallNonVirtualCharMethod ...
func CallNonVirtualCharMethod(obj Object, cls Class, id MethodID, args ...interface{}) (rune, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallNonVirtualCharMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	), nil
}

//CallNonVirtualShortMethod ...
func CallNonVirtualShortMethod(obj Object, cls Class, id MethodID, args ...interface{}) (int16, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallNonVirtualShortMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	), nil
}

//CallNonVirtualIntMethod ...
func CallNonVirtualIntMethod(obj Object, cls Class, id MethodID, args ...interface{}) (int, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallNonVirtualIntMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	), nil
}

//CallNonVirtualLongMethod ...
func CallNonVirtualLongMethod(obj Object, cls Class, id MethodID, args ...interface{}) (int64, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallNonVirtualLongMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	), nil
}

//CallNonVirtualFloatMethod ...
func CallNonVirtualFloatMethod(obj Object, cls Class, id MethodID, args ...interface{}) (float32, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallNonVirtualFloatMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	), nil
}

//CallNonVirtualDoubleMethod ...
func CallNonVirtualDoubleMethod(obj Object, cls Class, id MethodID, args ...interface{}) (float64, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallNonVirtualDoubleMethod(
		localEnv,
		(internal.Object)(obj),
		(internal.Class)(cls),
		(internal.MethodID)(id),
		args...,
	), nil
}

//GetFieldID ...
func GetFieldID(cls Class, name, sig string) (FieldID, error) {
	key := fieldCacheKey{cls, name, sig}
	if id, found := fieldCache[key]; found {
		return id, nil
	}

	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return nil, err
	}

	id := FieldID(internal.GetFieldID(
		localEnv,
		(internal.Class)(cls),
		name,
		sig,
	))

	fieldCache[key] = id
	return id, nil
}

//GetObjectField ...
func GetObjectField(obj Object, id FieldID) (*Object, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return nil, err
	}

	local := Object(internal.GetObjectField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
	))

	defer DeleteLocalRef(local)
	return NewGlobalRef(local)
}

//GetBooleanField ...
func GetBooleanField(obj Object, id FieldID) (bool, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return false, err
	}

	return internal.GetBooleanField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
	), err
}

//GetByteField ...
func GetByteField(obj Object, id FieldID) (byte, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.GetByteField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
	), err
}

//GetCharField ...
func GetCharField(obj Object, id FieldID) (rune, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.GetCharField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
	), err
}

//GetShortField ...
func GetShortField(obj Object, id FieldID) (int16, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.GetShortField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
	), err
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
func GetLongField(obj Object, id FieldID) (int64, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.GetLongField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
	), err
}

//GetFloatField ...
func GetFloatField(obj Object, id FieldID) (float32, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.GetFloatField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
	), err
}

//GetDoubleField ...
func GetDoubleField(obj Object, id FieldID) (float64, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.GetDoubleField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
	), err
}

//SetObjectField ...
func SetObjectField(obj Object, id FieldID, val Object) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetObjectField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
		internal.Object(val),
	)

	return nil
}

//SetBooleanField ...
func SetBooleanField(obj Object, id FieldID, val bool) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetBooleanField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
		val,
	)

	return nil
}

//SetByteField ...
func SetByteField(obj Object, id FieldID, val byte) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetByteField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
		val,
	)

	return nil
}

//SetCharField ...
func SetCharField(obj Object, id FieldID, val rune) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetCharField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
		val,
	)

	return nil
}

//SetShortField ...
func SetShortField(obj Object, id FieldID, val int16) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetShortField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
		val,
	)

	return nil
}


//SetIntField ...
func SetIntField(obj Object, id FieldID, val int) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetIntField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
		val,
	)

	return nil
}

//SetLongField ...
func SetLongField(obj Object, id FieldID, val int64) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetLongField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
		val,
	)

	return nil
}

//SetFloatField ...
func SetFloatField(obj Object, id FieldID, val float32) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetFloatField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
		val,
	)

	return nil
}

//SetDoubleField ...
func SetDoubleField(obj Object, id FieldID, val float64) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetDoubleField(
		localEnv,
		internal.Object(obj),
		internal.FieldID(id),
		val,
	)

	return nil
}

//GetStaticMethodID ...
func GetStaticMethodID(cls Class, name, sig string) (MethodID, error) {
	key := methodCacheKey{cls, name, sig}
	if id, found := methodCache[key]; found {
		return id, nil
	}

	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return nil, err
	}

	id := MethodID(internal.GetStaticMethodID(
		localEnv,
		(internal.Class)(cls),
		name,
		sig,
	))

	methodCache[key] = id
	return id, nil
}

//CallStaticObjectMethod ...
func CallStaticObjectMethod(cls Class, id MethodID, args ...interface{}) (*Object, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return nil, err
	}

	local := Object(internal.CallStaticObjectMethod(
		localEnv,
		internal.Class(cls),
		internal.MethodID(id),
		args...
	))

	defer DeleteLocalRef(local)
	return NewGlobalRef(local)
}

//CallStaticBooleanMethod ...
func CallStaticBooleanMethod(cls Class, id MethodID, args ...interface{}) (bool, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return false, err
	}

	return internal.CallStaticBooleanMethod(
		localEnv,
		internal.Class(cls),
		internal.MethodID(id),
		args...
	), nil
}

//CallStaticByteMethod ...
func CallStaticByteMethod(cls Class, id MethodID, args ...interface{}) (byte, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallStaticByteMethod(
		localEnv,
		internal.Class(cls),
		internal.MethodID(id),
		args...
	), nil
}

//CallStaticCharMethod ...
func CallStaticCharMethod(cls Class, id MethodID, args ...interface{}) (rune, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallStaticCharMethod(
		localEnv,
		internal.Class(cls),
		internal.MethodID(id),
		args...
	), nil
}

//CallStaticShortMethod ...
func CallStaticShortMethod(cls Class, id MethodID, args ...interface{}) (int16, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallStaticShortMethod(
		localEnv,
		internal.Class(cls),
		internal.MethodID(id),
		args...
	), nil
}

//CallStaticIntMethod ...
func CallStaticIntMethod(cls Class, id MethodID, args ...interface{}) (int, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallStaticIntMethod(
		localEnv,
		internal.Class(cls),
		internal.MethodID(id),
		args...
	), nil
}

//CallStaticLongMethod ...
func CallStaticLongMethod(cls Class, id MethodID, args ...interface{}) (int64, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallStaticLongMethod(
		localEnv,
		internal.Class(cls),
		internal.MethodID(id),
		args...
	), nil
}

//CallStaticFloatMethod ...
func CallStaticFloatMethod(cls Class, id MethodID, args ...interface{}) (float32, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallStaticFloatMethod(
		localEnv,
		internal.Class(cls),
		internal.MethodID(id),
		args...
	), nil
}

//CallStaticDoubleMethod ...
func CallStaticDoubleMethod(cls Class, id MethodID, args ...interface{}) (float64, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.CallStaticDoubleMethod(
		localEnv,
		internal.Class(cls),
		internal.MethodID(id),
		args...
	), nil
}
 
//GetStaticFieldID ...
func GetStaticFieldID(cls Class, name, sig string) (FieldID, error) {
	key := fieldCacheKey{cls, name, sig}
	if id, found := fieldCache[key]; found {
		return id, nil
	}

	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return nil, err
	}

	id := FieldID(internal.GetStaticFieldID(
		localEnv,
		(internal.Class)(cls),
		name,
		sig,
	))

	fieldCache[key] = id
	return id, nil
}

//GetStaticObjectField ...
func GetStaticObjectField(cls Class, id FieldID) (*Object, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return nil, err
	}

	local := Object(internal.GetStaticObjectField(
		localEnv,
		internal.Class(cls),
		internal.FieldID(id),
	))

	defer DeleteLocalRef(local)
	return NewGlobalRef(local)
}

//GetStaticBooleanField ...
func GetStaticBooleanField(cls Class, id FieldID) (bool, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return false, err
	}

	return internal.GetStaticBooleanField(
		localEnv,
		(internal.Class)(cls),
		(internal.FieldID)(id),
	), nil
}

//GetStaticByteField ...
func GetStaticByteField(cls Class, id FieldID) (byte, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.GetStaticByteField(
		localEnv,
		(internal.Class)(cls),
		(internal.FieldID)(id),
	), nil
}

//GetStaticCharField ...
func GetStaticCharField(cls Class, id FieldID) (rune, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.GetStaticCharField(
		localEnv,
		(internal.Class)(cls),
		(internal.FieldID)(id),
	), nil
}

//GetStaticShortField ...
func GetStaticShortField(cls Class, id FieldID) (int16, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.GetStaticShortField(
		localEnv,
		(internal.Class)(cls),
		(internal.FieldID)(id),
	), nil
}

//GetStaticIntField ...
func GetStaticIntField(cls Class, id FieldID) (int, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.GetStaticIntField(
		localEnv,
		(internal.Class)(cls),
		(internal.FieldID)(id),
	), nil
}

//GetStaticDoubleField ...
func GetStaticDoubleField(cls Class, id FieldID) (float64, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.GetStaticDoubleField(
		localEnv,
		(internal.Class)(cls),
		(internal.FieldID)(id),
	), nil
}

//SetStaticObjectField ...
func SetStaticObjectField(cls Class, id FieldID, val Object) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetStaticObjectField(
		localEnv,
		internal.Class(cls),
		internal.FieldID(id),
		internal.Object(val),
	)

	return nil
}

//SetStaticBooleanField ...
func SetStaticBooleanField(cls Class, id FieldID, val bool) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetStaticBooleanField(
		localEnv,
		internal.Class(cls),
		internal.FieldID(id),
		val,
	)

	return nil
}

//SetStaticByteField ...
func SetStaticByteField(cls Class, id FieldID, val byte) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetStaticByteField(
		localEnv,
		internal.Class(cls),
		internal.FieldID(id),
		val,
	)

	return nil
}

//SetStaticCharField ...
func SetStaticCharField(cls Class, id FieldID, val rune) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetStaticCharField(
		localEnv,
		internal.Class(cls),
		internal.FieldID(id),
		val,
	)

	return nil
}
//SetStaticShortField ...
func SetStaticShortField(cls Class, id FieldID, val int16) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetStaticShortField(
		localEnv,
		internal.Class(cls),
		internal.FieldID(id),
		val,
	)

	return nil
}

//SetStaticIntField ...
func SetStaticIntField(cls Class, id FieldID, val int) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetStaticIntField(
		localEnv,
		internal.Class(cls),
		internal.FieldID(id),
		val,
	)

	return nil
}

//SetStaticFloatField ...
func SetStaticFloatField(cls Class, id FieldID, val float32) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetStaticFloatField(
		localEnv,
		internal.Class(cls),
		internal.FieldID(id),
		val,
	)

	return nil
}

//SetStaticDoubleField ...
func SetStaticDoubleField(cls Class, id FieldID, val float64) error {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return err
	}

	internal.SetStaticDoubleField(
		localEnv,
		internal.Class(cls),
		internal.FieldID(id),
		val,
	)

	return nil
}

//NewString ...
func NewString(str string) (*String, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return nil, err
	}

	localStr := String(internal.NewString(localEnv, str))

	globalStr, err := NewGlobalRef((Object)(localStr))
	if err != nil {
		return nil, err
	}

	DeleteLocalRef((Object)(localStr))

	return (*String)(globalStr), nil
}

//GetStringLength ...
func GetStringLength(str String) (int, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return 0, err
	}

	return internal.GetStringLength(
		localEnv,
		internal.String(str),
	), nil
}

//GetStringChars ...
func GetStringChars(str String) (string, error) {
	localEnv, err := getThreadLocalEnvironment()
	if err != nil {
		return "", err
	}

	return internal.GetStringChars(
		localEnv,
		internal.String(str),
	), nil
}