package jni

//ErrorCode represents a jni error code.
type ErrorCode int

func (e ErrorCode) String() string {
	switch e {
	case Ok:
		return "ok"

	case Unknown:
		return "unknown"

	case Detached:
		return "thread detached from jvm"

	case WrongVersion:
		return "wrong version"

	case OutOfMemory:
		return "out of memory"

	case VMExists:
		return "vm already created"

	case InvalidArgs:
		return "invalid args"
	}

	return ""
}

const (
	//Ok means no error occured.
	Ok ErrorCode = 0

	//Unknown means an error occured, but the type is unknown.
	Unknown ErrorCode = -1

	//Detached means the current thread is detached from the JVM.
	Detached ErrorCode = -2

	//WrongVersion means the function is not supported on the current version.
	WrongVersion ErrorCode = -3

	//OutOfMemory means the JVM is out of memory.
	OutOfMemory ErrorCode = -4

	//VMExists means a VM is already created.
	VMExists ErrorCode = -5

	//InvalidArgs means invalid arguments were passed to the function.
	InvalidArgs ErrorCode = -6
)

//Error represents a jni error.
type Error struct {
	Code ErrorCode
}

func (e Error) Error() string {
	return e.Code.String()
}
