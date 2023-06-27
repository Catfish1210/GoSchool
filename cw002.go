package goschool

/* func FindOdd(seq []int) int {} */
func Codewars002(seq []int) int {
	var count, output int
	for _, v := range seq {
		for _, v2 := range seq {
			if v == v2 {
				count++
			}
		}
		if count%2 == 0 {
			count = 0
		} else {
			output = v
			break
		}
	}
	return output
}

/* Given an array of integers, find the one that appears an odd number of times.
   There will always be only one integer that appears an odd number of times.
*/
