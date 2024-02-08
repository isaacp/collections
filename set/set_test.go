package set

import "testing"

func TestAddMembers(t *testing.T) {
	var addTests = []struct {
		name  string
		input []int
		want  int
	}{
		{"no duplication", []int{10, 10, 10}, 1},
		{"no duplication", []int{10, 11, 10}, 2},
		{"no duplication", []int{10, 10, 10, 23, 12, 34, 54, 65, 10, 34, 65, 34, 12, 43}, 7},
	}

	for _, tt := range addTests {
		t.Run(tt.name, func(t *testing.T) {
			ans := NewSet(tt.input...).Count()
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
func TestDisjoint(t *testing.T) {
	var disjointTests = []struct {
		name  string
		input [][]int
		want  bool
	}{
		{"test1", [][]int{
			{10, 11, 10},
			{10, 11, 10},
		},
			false,
		},
	}

	for _, tt := range disjointTests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Disjoint(*NewSet(tt.input[0]), *NewSet(tt.input[1]))
			if ans != tt.want {
				t.Errorf("got %t, expected %t", ans, tt.want)
			}
		})
	}
}
