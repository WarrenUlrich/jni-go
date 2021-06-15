package jni

//Class ...
type Class Object

//GetMethodID ...
func (c Class) GetMethodID(name, sig string) (MethodID, error) {
	return GetMethodID(c, name, sig)
}

//GetStaticMethodID ...
func (c Class) GetStaticMethodID(name, sig string) (MethodID, error) {
	return GetStaticMethodID(c, name, sig)
}

//CallStaticObjectMethod ...
func (c Class) CallStaticObjectMethod(id MethodID, args ...interface{}) (*Object, error) {
	return CallStaticObjectMethod(c, id, args...)
}

//CallStaticBooleanMethod ...
func (c Class) CallStaticBooleanMethod(id MethodID, args ...interface{}) (bool, error) {
	return CallStaticBooleanMethod(c, id, args...)
}

//CallStaticByteMethod ...
func (c Class) CallStaticByteMethod(id MethodID, args ...interface{}) (byte, error) {
	return CallStaticByteMethod(c, id, args...)
}

//CallStaticCharMethod ...
func (c Class) CallStaticCharMethod(id MethodID, args ...interface{}) (rune, error) {
	return CallStaticCharMethod(c, id, args...)
}

//CallStaticShortMethod ...
func (c Class) CallStaticShortMethod(id MethodID, args ...interface{}) (int16, error) {
	return CallStaticShortMethod(c, id, args...)
}

//CallStaticIntMethod ...
func (c Class) CallStaticIntMethod(id MethodID, args ...interface{}) (int32, error) {
	return CallStaticIntMethod(c, id, args...)
}

//CallStaticLongMethod ...
func (c Class) CallStaticLongMethod(id MethodID, args ...interface{}) (int64, error) {
	return CallStaticLongMethod(c, id, args...)
}


//CallStaticFloatMethod ...
func (c Class) CallStaticFloatMethod(id MethodID, args ...interface{}) (float32, error) {
	return CallStaticFloatMethod(c, id, args...)
}

//CallStaticDoubleMethod ...
func (c Class) CallStaticDoubleMethod(id MethodID, args ...interface{}) (float64, error) {
	return CallStaticDoubleMethod(c, id, args...)
}

//GetFieldID ...
func (c Class) GetFieldID(name, sig string) (FieldID, error) {
	return GetFieldID(c, name, sig)
}

//GetStaticFieldID ...
func (c Class) GetStaticFieldID(name, sig string) (FieldID, error) {
	return GetStaticFieldID(c, name, sig)
}

//GetSTaticObjectField ...
func (c Class) GetStaticObjectField(id FieldID) (*Object, error) {
	return GetStaticObjectField(c, id)
}

//GetStaticBooleanField ...
func (c Class) GetStaticBooleanField(id FieldID) (bool, error) {
	return GetStaticBooleanField(c, id)
}

//GetStaticByteField ...
func (c Class) GetStaticByteField(id FieldID) (byte, error) {
	return GetStaticByteField(c, id)
}

//GetStaticCharField ...
func (c Class) GetStaticCharField(id FieldID) (rune, error) {
	return GetStaticCharField(c, id)
}

//GetStaticShortField ...
func (c Class) GetStaticShortField(id FieldID) (int16, error) {
	return GetStaticShortField(c, id)
}

//GetStaticIntField ...
func (c Class) GetStaticIntField(id FieldID) (int32, error) {
	return GetStaticIntField(c, id)
}

//GetStaticLongField ...
func (c Class) GetStaticLongField(id FieldID) (int64, error) {
	return GetStaticLongField(c, id)
}

//GetStaticFloatField ...
func (c Class) GetStaticFloatField(id FieldID) (float32, error) {
	return GetStaticFloatField(c, id)
}

//GetStaticDoubleField ...
func (c Class) GetStaticDoubleField(id FieldID) (float64, error) {
	return GetStaticDoubleField(c, id)
}

//SetStaticObjectField ...
func (c Class) SetStaticObjectField(id FieldID, val Object) error {
	return SetStaticObjectField(c, id, val)
}

//SetStaticBooleanField ...
func (c Class) SetStaticBooleanField(id FieldID, val bool) error {
	return SetStaticBooleanField(c, id, val)
}

//SetStaticByteField ...
func (c Class) SetStaticByteField(id FieldID, val byte) error {
	return SetStaticByteField(c, id, val)
}

//SetStaticCharField ...
func (c Class) SetStaticCharField(id FieldID, val rune) error {
	return SetStaticCharField(c, id, val)
}

//SetStaticShortField ...
func (c Class) SetStaticShortField(id FieldID, val int16) error {
	return SetStaticShortField(c, id, val)
}

//SetStaticIntField ...
func (c Class) SetStaticIntField(id FieldID, val int32) error {
	return SetStaticIntField(c, id, val)
}

//SetStaticLongField ...
func (c Class) SetStaticLongField(id FieldID, val int64) error {
	return SetStaticLongField(c, id, val)
}

//SetStaticFloatField ...
func (c Class) SetStaticFloatField(id FieldID, val float32) error {
	return SetStaticFloatField(c, id, val)
}

//SetStaticDoubleField ...
func (c Class) SetStaticDoubleField(id FieldID, val float64) error {
	return SetStaticDoubleField(c, id, val)
}