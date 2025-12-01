package test

import "testing"

func TestMain(m *testing.M) {
	m.Run()
}

func TestCountFairPairs1(t *testing.T) {
	nums := []int{1,7,9,2,5}
	lower := 11
	upper := 11
	expected := int64(1)
	result := CountFairPairs(nums, lower, upper)
	if result != expected {
		t.Logf("Expected %d, got %d", expected, result)
		t.Fail()
	}
}

func TestCountFairPairs2(t *testing.T) {
	nums := []int{0,1,7,4,4,5}
	lower := 3
	upper := 6
	expected := int64(6)
	result := CountFairPairs(nums, lower, upper)
	if result != expected {
		t.Logf("Expected %d, got %d", expected, result)
		t.Fail()
	}
}

func TestCountFairPairs3(t *testing.T) {
	nums := []int{-5,-7,-5,-7,-5}
	lower := -12
	upper := -12
	expected := int64(6)
	result := CountFairPairs(nums, lower, upper)
	if result != expected {
		t.Logf("Expected %d, got %d", expected, result)
		t.Fail()
	}
}

func TestTakeCharacters1(t *testing.T) {
	s := "aabaaaacaabc"
	n := 2
	expected := 8
	result := TakeCharacters(s, n)
	if result != expected {
		t.Logf("Expected %d, got %d", expected, result)
		t.Fail()
	}
}

func TestTakeCharacters2(t *testing.T) {
	n := 1
	expected := -1
	s := "a"
	result := TakeCharacters(s, n)
	if result != expected {
		t.Logf("Expected %d, got %d", expected, result)
		t.Fail()
	}
}

func TestTakeCharacters3(t *testing.T) {
	n := 0
	expected := 0
	s := "aabaaaacaabc"
	result := TakeCharacters(s, n)
	if result != expected {
		t.Logf("Expected %d, got %d", expected, result)
		t.Fail()
	}
}

func TestTakeCharacters4(t *testing.T) {
	n := 5
	expected := -1
	s := "aababcccbabc"
	result := TakeCharacters(s, n)
	if result != expected {
		t.Logf("Expected %d, got %d", expected, result)
		t.Fail()
	}
}

func TestTakeCharacters5(t *testing.T) {
	n := 4
	expected := 12
	s := "aababcccbabc"
	result := TakeCharacters(s, n)
	if result != expected {
		t.Logf("Expected %d, got %d", expected, result)
		t.Fail()
	}
}

func TestTakeCharacters6(t *testing.T) {
	n := 1
	expected := -1
	s := "caaa"
	result := TakeCharacters(s, n)
	if result != expected {
		t.Logf("Expected %d, got %d", expected, result)
		t.Fail()
	}
}

func TestTakeCharacters7(t *testing.T) {
	n := 1
	expected := 3
	s := "abc"
	result := TakeCharacters(s, n)
	if result != expected {
		t.Logf("Expected %d, got %d", expected, result)
		t.Fail()
	}
}
func TestTakeCharacters8(t *testing.T) {
	n := 1
	expected := 3
	s := "cbbac"
	result := TakeCharacters(s, n)
	if result != expected {
		t.Logf("Expected %d, got %d", expected, result)
		t.Fail()
	}
}

func TestTakeCharacters9(t *testing.T) {
	n := 1
	expected := 3
	s := "cbabc"
	result := TakeCharacters(s, n)
	if result != expected {
		t.Logf("Expected %d, got %d", expected, result)
		t.Fail()
	}
}
func TestTakeCharacters10(t *testing.T) {
	n := 2
	expected := 6
	s := "aabbccca"
	result := TakeCharacters(s, n)
	if result != expected {
		t.Logf("Expected %d, got %d", expected, result)
		t.Fail()
	}
}

