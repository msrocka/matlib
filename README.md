# matlib
`matlib` is a small matrix package with a focus on data exchange and simple
math functions. This is a fork of the [GreenDelta/matlib](https://github.com/GreenDelta/matlib)
with the dependencies to the native solver library removed. It would be probably
good to merge this back later and create a package `matsolve` or similar which
contains the native parts.

## File format
This package provides methods for loading (`goblapack.Load`) and saving 
(`goblapack.Save`) matrices from files in a simple binary format:

    header 8 bytes
    4 bytes: uint32, number of rows, little endian order
    4 bytes: uint32, number of columns, little endian order
    
    content, rows * columns * 8 bytes:
    matrix data, float64, little endian and column major order
    
Here is a small Python script for writing a NumPy matrix in this format:

```python
m = numpy.load('path/to/file.npy')
rows, cols = m.shape
with open('path/to/file.bin', 'wb') as f:
    f.write(struct.pack("<i", rows))
    f.write(struct.pack("<i", cols))
    for col in range(0, cols):
        for row in range(0, rows):
            val = m[row, col]
            f.write(struct.pack("<d", val))
```
            
            
