package goschool

import "testing"

/*
   Write the rgb function so that passing in RGB decimal values will
   result in a hexadecimal representation being returned. Valid decimal values for RGB are 0 - 255.
   Any values that fall out of that range must be rounded to the closest valid value.

   Examples:
	rgb(255, 255, 255) -- returns FFFFFF
	rgb(255, 255, 300) -- returns FFFFFF
	rgb(0, 0, 0) -- returns 000000
	rgb(148, 0, 211) -- returns 9400D3
*/
func TestCodewars004(t *testing.T) {
	inputs := [][]int{
		{255, 255, 255},
		{255, 255, 300},
		{0, 0, 0},
		{148, 0, 211},
	}
	outputs := []string{"FFFFFF", "FFFFFF", "000000", "9400D3"}
	for i, v := range inputs {
		got := Codewars004(v[0], v[1], v[2])
		want := outputs[i]
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}
func TestCodewars003(t *testing.T) {
	inputs := []int{16, 942, 132189}
	outputs := []int{7, 6, 6}
	for i, v := range inputs {
		got := Codewars003(v)
		want := outputs[i]
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}
func TestCodewars002(t *testing.T) {
	got := Codewars002([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1})
	want := 4
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
func TestCodewars001(t *testing.T) {
	got := Codewars001([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	want := "(123) 456-7890"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
