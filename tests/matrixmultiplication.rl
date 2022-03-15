// definition of the matrix multiplication
// input: two matrices (two-dimensional arrays) of same size
// output: matrix (two-dimensional array) of same size as input matrices
// error if input matrices are not of same size


// two-by-two integer matrix
multiply_matrices([[1,2],[3,4]], [[4,3],[2,1]])

// three-by-three float matrix
multiply_matrices([[1.0,0.5, 10459.1],[145037.54, -4.6, 0.0003], [5.6, -2.2222, -0.03]]  * [[4.3, 3.2221, -0.5],[8456291.5, 1000001.3, -0.5],[984381765.006, 564.2, -65.9]])

// ten-by-ten float matrix
multiply_matrices([[2.28 ,3.94 ,8.69 ,6.38 ,2.85 ,6.67 ,6.52 ,0.28 ,4.23 ,3.46],[2.84 ,7.24 ,0.74 ,6.67 ,8.60 ,2.57 ,4.14 ,0.23 ,4.25 ,7.42],[8.20 ,0.13 ,8.04 ,3.21 ,1.72 ,3.20 ,2.50 ,5.50 ,8.80 ,6.57],[6.44 ,7.31 ,0.35 ,0.83 ,2.57 ,3.26 ,7.38 ,0.91 ,7.39 ,0.15],[5.00 ,0.76 ,6.69 ,1.73 ,7.05 ,4.85 ,4.37 ,1.92 ,0.66 ,7.85],[8.95 ,6.58 ,2.87 ,1.61 ,8.67 ,6.30 ,4.05 ,2.20 ,7.31 ,8.52],[4.54 ,3.91 ,7.44 ,3.86 ,6.25 ,3.91 ,2.07 ,3.00 ,7.39 ,0.50],[2.15 ,5.68 ,3.23 ,8.21 ,0.59 ,7.28 ,7.83 ,4.33 ,6.34 ,7.38],[4.73 ,4.91 ,7.90 ,1.47 ,1.19 ,8.74 ,4.96 ,0.70 ,7.93 ,6.60],[4.15 ,8.19 ,5.27 ,2.68 ,0.35 ,3.29 ,1.50 ,0.05 ,8.12 ,2.93]], [[646.50, -268.17, -75.39, -471.19, -516.62, -120.21, -809.15, -461.32, -335.42, -906.73],[-347.33, 876.62, -132.13, -729.11, -199.48, 968.92, 715.93, 251.54, -769.80, -745.08],[-821.48, 182.50, -48.34, 375.97, -141.38, -468.16, -21.86, 336.71, 763.57, 280.22],[-471.54, 607.34, -955.30, -961.89, 983.99, -841.99, -269.07, -455.87, -707.85, -35.74],[472.32, -327.35, 111.09, 990.30, -479.17, 406.90, -491.34, 926.89, 965.49, -785.27],[929.32, 828.57, 464.39, 51.40, -114.47, -813.24, -506.75, 141.71, 714.26, -356.53],[-330.09, 430.30, -694.28, -292.95, -350.52, 201.82, -974.00, -577.57, -270.05, -373.44],[-399.12, 781.26, 359.06, 211.84, -255.98, -270.18, -27.79, 525.19, -677.66, -440.94],[768.53, -196.75, 380.02, -481.39, -796.96, -601.72, 867.81, -162.40, 674.06, -700.33],[573.41, 420.90, -118.85, 926.01, 647.92, -515.92, 534.54, -133.00, -732.24, -759.45]])

// two matrices of different size
multiply_matrices([[11.1,88.3],[0.3,9.4]], [[1,2,3],[0,6,3],[8,4,1]])

// wrong number of inputs for the function
multiply_matrices([[1,0,5], [6.4,3.1,9.6], [0.4,3,1]])
multiply_matrices([[1,0,5], [6.4,3.1,9.6], [0.4,3,1]], [[1,0,5], [6.4,3.1,9.6], [0.4,3,1]], [[1,0,5], [6.4,3.1,9.6], [0.4,3,1]])