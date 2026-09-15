package mapTimePass

import "testing"

func Test_isAnagramSingleHashMap(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		t    string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAnagramSingleHashMap(tt.s, tt.t)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("isAnagramSingleHashMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
