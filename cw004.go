package goschool

/* func RGB(r, g, b int) string {} */
func Codewars004(r, g, b int) string {
	hexValues := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	if r > 255 {
		r = 255
	} else if r < 0 {
		r = 0
	} else if g > 255 {
		g = 255
	} else if g < 0 {
		g = 0
	} else if b > 255 {
		b = 255
	} else if b < 0 {
		b = 0
	}

	return (hexValues[(r/16)] + hexValues[(r%16)] + hexValues[(g/16)] + hexValues[(g%16)] + hexValues[(b/16)] + hexValues[(b%16)])
}

/* Write the rgb function so that passing in RGB decimal values will
   result in a hexadecimal representation being returned. Valid decimal values for RGB are 0 - 255.
   Any values that fall out of that range must be rounded to the closest valid value.

   Examples:
rgb(255, 255, 255) -- returns FFFFFF
rgb(255, 255, 300) -- returns FFFFFF
rgb(0, 0, 0) -- returns 000000
rgb(148, 0, 211) -- returns 9400D3

*/
