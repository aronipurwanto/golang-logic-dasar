package logic04

import (
	array2 "github.com/aronipurwanto/golang-logic-dasar/array"
)

func Logic04Soal05(n int) {
	// create array
	array := array2.NewNumberArray(n, n)
	// initial angka dari 1
	angka := 1
	// for untuk posisi dari 1 s.d
	for posisi := 1; posisi <= 4; posisi++ {
		for b := 1; b < n; b++ {
			if posisi == 1 { // posisi kiri dari bawah ke atas
				array[n-1-b][0] = int32(angka)
			} else if posisi == 2 { //posisi atas dari kiri kanan
				array[0][b] = int32(angka)
			} else if posisi == 3 { //posisi kanan dari atas ke bawah
				array[b][n-1] = int32(angka)
			} else if posisi == 4 { // posisi bawah dari kanan ke kiri
				array[n-1][n-1-b] = int32(angka)
			}
			angka++
		}
	}

	// print array
	array2.PrintNumberArrayZeroEmpty(array)
}

func Sample05Step01(n int) {
	kotak := array2.NewNumberArray(n, n)

	angka := 1
	for a := 1; a < n; a++ {
		kotak[n-1-a][0] = int32(angka)
		angka++
	}
}

func Sample05Step02(n int) {
	kotak := array2.NewNumberArray(n, n)

	angka := 1
	// buat untuk diulang sebanyak 4 sisi
	for sisi := 1; sisi <= 4; sisi++ {
		// isi array di ulang sebanyak sisi
		for a := 1; a < n; a++ {
			kotak[n-1-a][0] = int32(angka)
			angka++
		}
		// isi array selesai
	}
}
