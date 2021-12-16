package monitor

import "testing"

func Test_stringSetUnion(t *testing.T) {
	first := []string{"a", "d", "e"}
	second := []string{"b", "d", "c"}

	result := stringSetUnion(first, second)

	if len(result) != 5 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", len(result), 5)
	}

	compare_result := []string{"a", "b", "c", "d", "e"}

	for i := range compare_result {
		found := false
		for j := range result {
			if compare_result[i] == result[j] {
				found = true
			}
		}
		if !found {
			t.Errorf("%v was not found in the resulting slice", compare_result[i])
		}
	}
}

func Test_stringSetSubtract(t *testing.T) {
	first := []string{"a", "b", "c", "d", "e"}
	second := []string{"b", "d", "f"}

	result := stringSetSubtract(first, second)

	if len(result) != 3 {
		t.Errorf("Size was incorrect, got: %d, want: %d.", len(result), 5)
	}

	compare_result := []string{"a", "c", "e"}

	for i := range compare_result {
		found := false
		for j := range result {
			if compare_result[i] == result[j] {
				found = true
			}
		}
		if !found {
			t.Errorf("%v was not found in the resulting slice", compare_result[i])
		}
	}
}
