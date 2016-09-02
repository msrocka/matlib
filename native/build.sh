#!/bin/bash

gcc -fPIC -O3 -L. -shared -o ../libmatlib.so matlib.c -lopenblas -lgfortran -lgcc -lpthread
