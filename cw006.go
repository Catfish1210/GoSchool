package goschool

/* func FirstNonRepeating(str string) string {} */
func Codewars006(str string) string {
	output := ""
	if len(str) == 1 {
		return str
	}
	for i := 0; i < len(str); i++ {
		for _, v := range str {
			if v == rune(str[i]) {
				output = ""
			} else {
				output = string(v)
				break
			}
		}
	}
	return output
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
