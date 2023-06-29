package goschool

/* func FirstNonRepeating(str string) string {} */
func Codewars006(str string) string {
	count := 0
	for _, v := range str {

		//check if v exists more than 1 match

		for _, v2 := range str {
			if v == v2 {
				count++
			}
		}
		if count > 1 {
			count = 0
			continue
		} else {
			return (string(v))
		}
	}
	return ""
}

/*
   Write a function named first_non_repeating_letter that takes a string input,
   and returns the first character that is not repeated anywhere in the string.

   For example, if given the input 'stress', the function should return 't',
   since the letter t only occurs once in the string, and occurs first in the string.

   As an added challenge, upper- and lowercase letters are considered the same character,
   but the function should return the correct case for the initial letter.
   For example, the input 'sTreSS' should return 'T'.

   If a string contains all repeating characters, it should return an empty string ("") or None
*/
