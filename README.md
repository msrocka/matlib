#goblapack
goblapack is a thin Go layer for calling functions of the high performance
math library [OpenBLAS](https://github.com/xianyi/OpenBLAS) via system calls.
It does not require cgo; just put the native goblapack library (goblapack.dll on
Windows, goblapack.so on Linux, Mac OS is currently not supported) next to your
binary and it should work.

On Windows we statically link all required libraries in the goblapack.dll. On
Linux you have to install ##

See also the the descriptions in the `native` folder if you want to compile the
native library.
