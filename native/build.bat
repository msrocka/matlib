gcc -Wl,--kill-at -static -static-libgcc -O3 -L. -shared -o ../matlib.dll matlib.c -lopenblas -lgfortran

PAUSE
