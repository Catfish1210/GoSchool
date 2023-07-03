package goschool

/* func MultiplicationTable(size int) [][]int {} */
func Codewars008(size int) [][]int {
	mpTable := make([][]int, size, size)

	for i := 0; i < size; i++ {
		if i == 0 {
			for j := 1; j <= size; j++ {
				mpTable[0] = append(mpTable[0], j)
			}
			continue
		}
		mpTable[i] = append(mpTable[i], i+1)
		for len(mpTable[i]) < size {
			mpTable[i] = append(mpTable[i], 0)
		}
	}

	for i, v := range mpTable {
		if i == 0 {
			continue
		}

		for i2, v2 := range v {
			if v2 == 0 {
				mpTable[i][i2] = mpTable[0][i2] * mpTable[i][0]
			}
		}
	}
	return mpTable
}

/*
   Your task, is to create N×N multiplication table, of size provided in parameter.
   For example, when given size is 3:
   1 2 3
   2 4 6
   3 6 9
*/
