#!/bin/bash
gcc -fPIC -O3 -L. -shared -o libmatlib.dylib matlib.c -static -lopenblas -lc
