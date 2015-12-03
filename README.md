#goblapack
goblapack is a thin Go layer for calling functions of the high performance
math library [OpenBLAS](https://github.com/xianyi/OpenBLAS) via
[cgo](https://golang.org/cmd/cgo/). It currently only contains the functions we
need. For more complete BLAS and LAPACK bindings in Go, please have a look at
the packages provided by the [gonum team](https://github.com/gonum).   




On Windows we statically link all required libraries in the goblapack.dll. On
Linux you have to install ##

See also the the descriptions in the `native` folder if you want to compile the
native library.
