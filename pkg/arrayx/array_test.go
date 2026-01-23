package arrayx

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	// Test mapping integers to their squares
	ints := []int{1, 2, 3, 4}
	squares := Map(ints, func(x int) int {
		return x * x
	})
	expected := []int{1, 4, 9, 16}
	for i, v := range squares {
		if v != expected[i] {
			t.Errorf("expected %d, got %d at index %d", expected[i], v, i)
		}
	}

	// Test mapping strings to their lengths
	strs := []string{"go", "chatgpt", "test"}
	lengths := Map(strs, func(s string) int {
		return len(s)
	})
	expectedLengths := []int{2, 7, 4}
	for i, v := range lengths {
		if v != expectedLengths[i] {
			t.Errorf("expected %d, got %d at index %d", expectedLengths[i], v, i)
		}
	}

	// Test mapping an empty slice
	empty := []int{}
	mappedEmpty := Map(empty, func(x int) int {
		return x * 2
	})
	if len(mappedEmpty) != 0 {
		t.Errorf("expected empty slice, got %v", mappedEmpty)
	}
}

func TestFindIf(t *testing.T) {
	// Test finding first even number
	ints := []int{1, 3, 5, 8, 10}
	found := FindIf(ints, func(x int) bool {
		return x%2 == 0
	})
	if found == nil || *found != 8 {
		t.Errorf("expected to find 8, got %v", found)
	}

	// Test finding in an empty slice
	empty := []int{}
	foundEmpty := FindIf(empty, func(x int) bool {
		return true
	})
	if foundEmpty != nil {
		t.Errorf("expected nil, got %v", foundEmpty)
	}

	// Test no matching element
	oddInts := []int{1, 3, 5, 7}
	foundNone := FindIf(oddInts, func(x int) bool {
		return x%2 == 0
	})
	if foundNone != nil {
		t.Errorf("expected nil, got %v", foundNone)
	}

	// Test finding in a slice of strings
	strs := []string{"apple", "banana", "cherry"}
	foundStr := FindIf(strs, func(s string) bool {
		return len(s) > 5
	})
	if foundStr == nil || *foundStr != "banana" {
		t.Errorf("expected 'banana', got %v", foundStr)
	}
}

func TestFilter_String(t *testing.T) {
	input := []string{"apple", "banana", "avocado", "grape"}
	startsWithA := func(s string) bool { return len(s) > 0 && s[0] == 'a' }

	expected := []string{"apple", "avocado"}
	result := Filter(input, startsWithA)

	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("expected %s at index %d, got %s", expected[i], i, result[i])
		}
	}
}

func TestFilter_EmptyInput(t *testing.T) {
	input := []int{}
	isPositive := func(x int) bool { return x > 0 }

	result := Filter(input, isPositive)
	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
}

func TestFlatten_IntSlices(t *testing.T) {
	input := []int{1, 2, 3}
	mapper := func(x int) []int {
		return []int{x, x * 10}
	}
	expected := []int{1, 10, 2, 20, 3, 30}

	result := Flatten(input, mapper)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFlatten_EmptyInput(t *testing.T) {
	input := []int{}
	mapper := func(x int) []int {
		return []int{x, x * 10}
	}
	expected := []int{}

	result := Flatten(input, mapper)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFlatten_StringToRuneSlices(t *testing.T) {
	input := []string{"go", "ai"}
	mapper := func(s string) []rune {
		return []rune(s)
	}
	expected := []rune{'g', 'o', 'a', 'i'}

	result := Flatten(input, mapper)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFlatten_NestedFlattening(t *testing.T) {
	input := [][]int{{1, 2}, {3}, {4, 5, 6}}
	mapper := func(x []int) []int {
		return x
	}
	expected := []int{1, 2, 3, 4, 5, 6}

	result := Flatten(input, mapper)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
