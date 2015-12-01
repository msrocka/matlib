gcc -Wl,--kill-at -static -static-libgcc -O3 -L. -shared -o goblapack.dll goblapack.c -lopenblas -lgfortran

PAUSE
