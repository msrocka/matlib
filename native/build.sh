#!/bin/bash

gcc -fPIC -O3 -L. -shared -o ../goblapack.so goblapack.c -lopenblas -lgfortran -lgcc -lpthread
