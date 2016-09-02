#include <stdio.h>
#include <stdlib.h>

/*
 * BLAS function DGEMM
 *
 * perform one of the matrix-matrix operations C := alpha*op( A )*op( B ) + beta*C
 *
 * see http://www.math.utah.edu/software/lapack/lapack-blas/dgemm.html
 */
 extern void dgemm_(char *TRANSA, char *TRANSB, int *M, int *N, int *K,
    double *ALPHA, double *A, int *LDA, double *B, int *LDB, double *BETA,
    double *C, int *LDC);

/*
 * LAPACK function DGETRF
 *
 * Compute an LU factorization of a general M-by-N matrix A using partial
 * pivoting with row interchanges
 *
 * see http://www.math.utah.edu/software/lapack/lapack-d/dgetrf.html
 */
extern void dgetrf_(int* M, int *N, double* A, int* LDA, int* IPIV, int* INFO);

/*
 * LAPACK function DGETRI
 *
 * Compute the inverse of a matrix using the LU factorization computed by DGETRF
 *
 * see http://www.math.utah.edu/software/lapack/lapack-d/dgetri.html
 */
extern void dgetri_(int* N, double* A, int* LDA, int* IPIV, double* WORK,
    int* LWORK, int* INFO);

/*
 * Calculates the inverse of the given matrix.
 *
 * n: size of the matrix (rows and columns)
 * a: matrix data in column-major-order
 */
int matlib_invert(int n, double* a) {
	int *ipiv = (int *) malloc(n * sizeof (int));
    int lwork = 64 * 2 * n;
    double *work = (double *) malloc(lwork * sizeof (double));
    int info;
    dgetrf_(&n, &n, a, &n, ipiv, &info);
    if (info) {
        return info;
	}
    dgetri_(&n, a, &n, ipiv, work, &lwork, &info);
    free(ipiv);
    free(work);
    return info;
}

/*
 * Matrix-matrix multiplication: C := A * B
 *
 * rowsA: number of rows of matrix A
 * colsB: number of columns of matrix B
 * k: number of columns of matrix A and number of rows of matrix B
 * a: matrix A (size = rowsA*k)
 * b: matrix B (size = k * colsB)
 * c: matrix C (size = rowsA * colsB)
 */
void matlib_mmult(int rowsA, int colsB, int k, double *a, double *b,
    double *c) {
    char trans = 'N';
    double alpha = 1;
    double beta = 0;
    dgemm_(&trans, &trans, &rowsA, &colsB, &k, &alpha, a, &rowsA, b, &k,
            &beta, c, &rowsA);
}
