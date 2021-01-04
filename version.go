package jni

//GetVersion returns the major version number in
//the higher 16 bits and the minor version in the
//lower 16 bits.

//Version represents the jni version, e.g. 1.1, 1.2.
//The major version is stored in the higher 16 bits,
//and the minor version is in the lower 16.
type Version int

const (
	//V1_1 Version 1.1
	V1_1 Version = 0x00010001

	//V1_2 Version 1.2
	V1_2 Version = 0x00010002

	//V1_4 Version 1.4
	V1_4 Version = 0x00010004

	//V1_6 Version 1.6
	V1_6 Version = 0x00010006

	//V1_8 Version 1.8
	V1_8 Version = 0x00010008
)
