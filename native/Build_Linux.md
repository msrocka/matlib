# Building matlib.so on Linux

## Building OpenBLAS
The first step is to compile the [OpenBLAS](https://github.com/xianyi/OpenBLAS)
library. To do that, get the source from:

    git clone https://github.com/xianyi/OpenBLAS

Then checkout the latest stable version which you can find from the tags:

    git tag

It is a good idea to checkout the version tag directly into a new branch, e.g.:

    git checkout tags/v0.2.19 -b build_0.2.19

Open the file `Makefile.rule` and adopt the build settings to your needs. For
the goblapack.so in the current release I took the following settings:

    DYNAMIC_ARCH = 1
    CC = gcc
    FC = gfortran
    BINARY=64
    USE_THREAD = 1
    NUM_THREADS = 2
    NO_CBLAS = 1
    NO_LAPACKE = 1

To run the build you need to have make, gcc, gfortran etc. installed. Typically,
you have to install gfortran. When everything is ready, just run make:

    make

## Building matlib.so
To build the shared matlib.so library on Linux copy the static library from
the OpenBLAS build `libopenblasp_*.a` as `libopenblas.a` to the `native` folder.
After this you can run the `build.sh` script (you may have to make it executable
via `chmod +x`) and it should generate the goblapack.so file in top-project
folder. 
