package goschool

import "testing"

func TestCodewars008(t *testing.T) {
	got := Codewars008(3)
	want := [][]int{
		{1, 2, 3},
		{2, 4, 6},
		{3, 6, 9},
	}
	for i, v := range got {
		for j, val := range v {
			if val != got[i][j] {
				t.Errorf("got %q, wanted %q", val, want)
			}
		}
	}
}
func TestCodewars007(t *testing.T) {
	inputs := []uint32{
		2149583361,
		32,
		0,
	}
	outputs := []string{"128.32.10.1", "0.0.0.32", "0.0.0.0"}
	for i, v := range inputs {
		got := Codewars007(v)
		want := outputs[i]
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}
func TestCodewars006(t *testing.T) {
	inputs := []string{
		"stress",
		"sTreSs",
	}
	outputs := []string{"t", "T"}
	for i, v := range inputs {
		got := Codewars006(v)
		want := outputs[i]
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}
func TestCodewars005(t *testing.T) {
	inputs := []string{
		"MCMXC",
		"MMVIII",
		"MDCLXVI",
	}
	outputs := []int{1990, 2008, 1666}
	for i, v := range inputs {
		got := Codewars005(v)
		want := outputs[i]
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

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
