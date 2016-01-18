#!/bin/bash

gcc -fPIC -O3 -L. -shared -o ../libgoblapack.so goblapack.c -lopenblas -lgfortran -lgcc -lpthread
