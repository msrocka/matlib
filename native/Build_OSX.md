# Building matlib.dylib on OS X

## Building OpenBLAS
Building OpenBLAS includes the same steps as for Linux. You need to install gcc
and gfortran for the build. If you use [Homebrew](http://brew.sh/) and install
the gcc package, gfortran will be also installed:

    brew install gcc

(test the installation via `gcc -v` and `gfortran -v`)
Additionally, you need to have `make` installed which is probably already the 
case (test it via `make -v`).

## Building matlib.dylib
Copy the static OpenBLAS library (`libopenblasp_*.a`) to the `native` folder and
rename it to `libopenblas.a` and run `build_osx.sh`.
  