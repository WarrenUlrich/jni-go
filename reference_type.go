package jni

type referenceType uint8

const (
	local referenceType = iota

	global referenceType = iota
)
