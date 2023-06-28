package goschool

/* func DigitalRoot(n int) int {} */
func Codewars003(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	if sum > 9 {
		return Codewars003(sum)
	} else {
		return sum
	}
}

/*
  Digital root is the recursive sum of all the digits in a number.
  Given n, take the sum of the digits of n.
  If that value has more than one digit, continue reducing in this
  way until a single-digit number is produced. The input will be a non-negative integer.

  Examples:
   16  -->  1 + 6 = 7
   942  -->  9 + 4 + 2 = 15  -->  1 + 5 = 6
   132189  -->  1 + 3 + 2 + 1 + 8 + 9 = 24  -->  2 + 4 = 6
*/
