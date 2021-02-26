package jni

import "github.com/warrenulrich/jni-go/pkg/internal"

//Object ...
type Object internal.Object

const (
	//NullReference ...
	NullReference Object = Object(0)
)

//Null ...
func (o *Object) Null() bool {
	return *o == NullReference
}