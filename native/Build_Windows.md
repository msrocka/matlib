# Building the matlib.dll on Windows

## Building OpenBLAS
First, we need to build OpenBLAS for Windows as described in the 
[openLCA](https://github.com/GreenDelta/olca-modules/blob/master/doc/openblas_windows.md) 
or [OpenBLAS](https://github.com/xianyi/OpenBLAS/wiki/Installation-Guide) 
documentation.

## Building the matlib.dll
Copy the static library from the OpenBLAS build `libopenblasp_*.a` as 
`libopenblas.a` to the `native` folder. Additionally, you need to copy the
`libgfortran.a` library from the MinGW package 
(e.g. in MinGW/lib/gcc/x86_64-w64-mingw32/4.8.1) to the `native` folder. After
this you should be able to run the `build.bat` which will produce the
`matlib.dll`.