package goblapack

import (
	"syscall"
)

// FreeLibrary releases the native library. You should only call this function
// when you are sure to not call a native function again.
func FreeLibrary() error {
	if nativeLibError != nil {
		return nativeLibError
	}
	return syscall.FreeLibrary(nativeLib)
}
