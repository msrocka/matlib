#matlib
matlib is a small package for the creation and testing of matrices for
[openLCA](https://github.com/GreenDelta/olca-app). It uses the high 
performance math library [OpenBLAS](https://github.com/xianyi/OpenBLAS).

## Usage
Checkout this repository:

    git clone https://github.com/GreenDelta/matlib.git

Build the native library for native library for your system as described
in `native/Build_<your_os>.md`.

## File format
This package provides methods for loading (`goblapack.Load`) and saving 
(`goblapack.Save`) matrices from files in a simple binary format:

    header 8 bytes
    4 bytes: uint32, number of rows, little endian order
    4 bytes: uint32, number of columns, little endian order
    
    content, rows * columns * 8 bytes:
    matrix data, float64, little endian and column major order
    
Here is a small Python script for writing a Numpy matrix in this format:

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
            
            
