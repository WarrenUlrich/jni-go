package jni

import "github.com/warrenulrich/jni-go/pkg/internal"

//Object ...
type Object internal.Object

const (
	//NullReference ...
	NullReference Object = Object(0)
)

//CallObjectMethod ...
func (o *Object) CallObjectMethod(id MethodID, args... interface{}) (*Object, error) {
	return CallObjectMethod(*o, id, args...)
}

//CallBooleanMethod ...
func (o *Object) CallBooleanMethod(id MethodID, args... interface{}) (bool, error) {
	return CallBooleanMethod(*o, id, args...)
}

//CallByteMethod ...
func (o *Object) CallByteMethod(id MethodID, args... interface{}) (byte, error) {
	return CallByteMethod(*o, id, args...)
}

//CallCharMethod ...
func (o *Object) CallCharMethod(id MethodID, args... interface{}) (rune, error) {
	return CallCharMethod(*o, id, args...)
}

//CallShortMethod ...
func (o *Object) CallShortMethod(id MethodID, args... interface{}) (int16, error) {
	return CallShortMethod(*o, id, args...)
}

//CallIntMethod ...
func (o *Object) CallIntMethod(id MethodID, args... interface{}) (int32, error) {
	return CallIntMethod(*o, id, args...)
}

//CallFloatMethod ...
func (o *Object) CallFloatMethod(id MethodID, args... interface{}) (float32, error) {
	return CallFloatMethod(*o, id, args...)
}

//CallDoubleMethod ...
func (o *Object) CallDoubleMethod(id MethodID, args... interface{}) (float64, error) {
	return CallDoubleMethod(*o, id, args...)
}

//GetObjectField ...
func(o *Object) GetObjectField(id FieldID) (*Object, error) {
	return GetObjectField(*o, id)
}

//GetBooleanField ...
func(o *Object) GetBooleanField(id FieldID) (bool, error) {
	return GetBooleanField(*o, id)
}

//GetByteField ...
func(o *Object) GetByteField(id FieldID) (byte, error) {
	return GetByteField(*o, id)
}

//GetCharField ...
func(o *Object) GetCharField(id FieldID) (rune, error) {
	return GetCharField(*o, id)
}

//GetShortField ...
func(o *Object) GetShortField(id FieldID) (int16, error) {
	return GetShortField(*o, id)
}

//GetIntField ...
func(o *Object) GetIntField(id FieldID) (int32, error) {
	return GetIntField(*o, id)
}

//GetLongField ...
func(o *Object) GetLongField(id FieldID) (int64, error) {
	return GetLongField(*o, id)
}

//GetFloatField ...
func(o *Object) GetFloatField(id FieldID) (float32, error) {
	return GetFloatField(*o, id)
}

//GetDoubleField ...
func(o *Object) GetDoubleField(id FieldID) (float64, error) {
	return GetDoubleField(*o, id)
}

//SetObjectField ...
func (o *Object) SetObjectField(id FieldID, val Object) error {
	return SetObjectField(*o, id, val)
}

//SetBooleanField ...
func (o *Object) SetBooleanField(id FieldID, val bool) error {
	return SetBooleanField(*o, id, val)
}

//SetByteField ...
func (o *Object) SetByteField(id FieldID, val byte) error {
	return SetByteField(*o, id, val)
}

//SetCharField ...
func (o *Object) SetCharField(id FieldID, val rune) error {
	return SetCharField(*o, id, val)
}

//SetShortField ...
func (o *Object) SetShortField(id FieldID, val int16) error {
	return SetShortField(*o, id, val)
}

//SetIntField ...
func (o *Object) SetIntField(id FieldID, val int32) error {
	return SetIntField(*o, id, val)
}

//SetLongField ...
func (o *Object) SetLongField(id FieldID, val int64) error {
	return SetLongField(*o, id, val)
}

//SetFloatField ...
func (o *Object) SetFloatField(id FieldID, val float32) error {
	return SetFloatField(*o, id, val)
}

//SetDoubleField ...
func (o *Object) SetDoubleField(id FieldID, val float64) error {
	return SetDoubleField(*o, id, val)
}

//Null ...
func (o *Object) Null() bool {
	return *o == NullReference
}