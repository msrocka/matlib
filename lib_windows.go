package goblapack

import (
	"syscall"
)

var nativeLib, nativeLibError = syscall.LoadLibrary("goblapack.dll")
