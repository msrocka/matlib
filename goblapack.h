#ifndef _GOBLAPACK_H_
#define _GOBLAPACK_H_

#ifdef __cplusplus
extern "C" {
#endif


int goblapack_invert(int n, double* a);

void goblapack_mmult(int rowsA, int colsB, int k, double *a, double *b, double *c);


#ifdef __cplusplus
}
#endif

#endif /* _GOBLAPACK_H_ */
