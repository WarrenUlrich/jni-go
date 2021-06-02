package jni

//String ...
type String Object

//Length ...
func (s String) Length() (int, error) {
	return GetStringLength(s)
}

//String ...
func (s String) GetString() (string, error) {
	return GetStringChars(s)
}