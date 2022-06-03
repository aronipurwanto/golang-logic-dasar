package array

// create new array with dinamic value
/*
1 2 3  => baris 0, kolom 0, 1, 2
4 5 6  => baris 1, kolom 0, 1, 2
7 8 9  => baris 2, kolom 0, 1, 2
*/

// function NewStringArray, parameter baris int, kolom int, return atau kembali nilai [][]string
// [][]string => array 2 dimensi
func NewStringArray(baris int, kolom int) [][]string {
	// membuat slice, length baris
	result := make([][]string, baris)
	for i := range result {
		result[i] = make([]string, kolom)
	}
	return result
}

func NewNumberArray(baris int, kolom int) [][]int32 {
	result := make([][]int32, baris)
	for i := range result {
		result[i] = make([]int32, kolom)
	}
	return result
}
