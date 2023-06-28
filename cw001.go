package goschool

import "fmt"

/* func CreatePhoneNumber(numbers [10]uint) string {} */
func Codewars001(numbers [10]uint) string {
	var output string
	for i, v := range numbers {
		if i == 0 {
			output += "("
		} else if i == 3 {
			output += ") "
		} else if i == 6 {
			output += "-"
		}
		output += fmt.Sprint(v)
	}
	return output
}

/*
   Write a function that accepts an array of 10 integers (between 0 and 9),
   that returns a string of those numbers in the form of a phone number.

   *CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})  // returns "(123) 456-7890"
*/
