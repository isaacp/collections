package set

import "testing"

func TestAddMembers(t *testing.T) {
	var addTests = []struct {
		name  string
		input []int
		want  int
	}{
		{"no duplication", []int{10, 10, 10}, 1},
		{"correct length", []int{10, 10, 10, 23, 12, 34, 54, 65, 10, 34, 65, 34, 12, 43}, 7},
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
