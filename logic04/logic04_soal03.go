package logic04

import (
	array2 "github.com/aronipurwanto/golang-logic-dasar/array"
)

func Logic04Soal03(n int) {
	// create array
	array := array2.NewNumberArray(n, n)
	// initial angka dari 1
	angka := 1
	// loop baris
	for b := 0; b < n; b++ {
		// loop kolom
		for k := 0; k < n; k++ {
			// jika baris 0
			if b%4 == 0 {
				array[n-1-k][b] = int32(angka)
				angka += 3
			} else if b%4 == 1 && k == 0 {
				array[k][b] = int32(angka)
				angka += 3
			} else if b%4 == 2 {
				array[k][b] = int32(angka)
				angka += 3
			} else if b%4 == 3 && k == n-1 {
				array[k][b] = int32(angka)
				angka += 3
			}
		}
	}
	// print array
	array2.PrintNumberArrayZeroEmpty(array)
}
