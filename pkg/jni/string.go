package jni

//String ...
type String Object

//Length ...
func (s String) Length() int {
	return GetStringLength(s)
}

//String ...
func (s String) String() string {
	return GetStringChars(s)
}