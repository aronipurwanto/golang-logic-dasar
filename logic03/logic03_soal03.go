package logic03

import array2 "github.com/aronipurwanto/golang-logic-dasar/array"

func Logic03Soal03(n int) {
	// 1. create array
	array := array2.NewNumberArray(n, n)

	// 2. isi array
	// loop baris
	angka := 3 // type nya int
	for i := 0; i < len(array); i++ {
		nt := n / 2
		// loop kolom
		for j := 0; j < len(array[i]); j++ {
			// isi array
			if i < j && j <= nt {
				array[i][j] = int32(angka) // angka di convert ke int32
			} else if i > j && j >= nt {
				array[i][j] = int32(angka) // angka di convert ke int32
			} else if i+j < n-1 && i >= nt {
				array[i][j] = int32(angka)
			} else if i+j > n-1 && i <= nt {
				array[i][j] = int32(angka)
			}
		}

		// jika baris kurang dari nilai tengah
		if i < nt {
			angka += 3 // angka = angka + 3
		} else {
			angka -= 3
		}
	}

	// 3. print array
	array2.PrintNumberArrayZeroEmpty(array)
}
