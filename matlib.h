#ifndef _MATLIB_H_
#define _MATLIB_H_

#ifdef __cplusplus
extern "C" {
#endif


int matlib_invert(int n, double* a);

void matlib_mmult(int rowsA, int colsB, int k, double *a, double *b, double *c);


#ifdef __cplusplus
}
#endif

#endif /* _MATLIB_H_ */
