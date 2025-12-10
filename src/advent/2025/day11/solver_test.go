package day11

import (
	"testing"
)


func TestSolvePart1Full(t *testing.T) {
    r, _ := Solve(true, false)
    e := Result{}
    if r != e {
        t.Errorf("Expected part 1 result %d, but got %d", e, r)
        t.FailNow()
    }
}

func TestSolvePart1(t *testing.T) {

    data := []string{
        
    }

    r := SolvePart1(data)
    e := Result{}

    if r != e {
        t.Errorf("Expected %d, but got %d", e, r)
        t.FailNow()
    }
}

func TestSolvePart2Full(t *testing.T) {
    _, r2 := Solve(false, true)
    e := Result{}
    if r2 != e {
        t.Errorf("Expected part 2 result %d, but got %d", e, r2)
        t.FailNow()
    }
}

func TestSolvePart2(t *testing.T) {
    data := []string{
        
    }

    r := SolvePart2(data)
    e := Result{}

    if r != e {
        t.Errorf("Expected %d, but got %d", e, r)
        t.FailNow()
    }
}

func TestSolveBoth(t *testing.T) {
    TestSolvePart1Full(t)
    TestSolvePart2Full(t)
}