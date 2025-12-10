package day9

import (
	"testing"
)


func TestSolvePart1Full(t *testing.T) {
    r, _ := Solve(true, false)
    e := 4773451098
    if r.largestArea != e {
        t.Errorf("Expected part 1 result %d, but got %d", e, r)
        t.FailNow()
    }
}

func TestSolvePart1(t *testing.T) {

    data := []string{
        "7,1",
        "11,1",
        "11,7",
        "9,7",
        "9,5",
        "2,5",
        "2,3",
        "7,3",
    }

    r := SolvePart1(data)
    e := 50

    if r.largestArea != e {
        t.Errorf("Expected %d, but got %d", e, r)
        t.FailNow()
    }
}

func TestSolvePart2Full(t *testing.T) {
    _, r2 := Solve(false, true)
    e := 0
    if r2.largestArea != e {
        t.Errorf("Expected part 2 result %d, but got %d", e, r2)
        t.FailNow()
    }
}

func TestSolvePart2(t *testing.T) {
    data := []string{
        "7,1",
        "11,1",
        "11,7",
        "9,7",
        "9,5",
        "2,5",
        "2,3",
        "7,3",
    }

    r := SolvePart2(data)
    e := 0

    if r.largestArea != e {
        t.Errorf("Expected %d, but got %d", e, r)
        t.FailNow()
    }
}

func TestSolveBoth(t *testing.T) {
    TestSolvePart1Full(t)
    TestSolvePart2Full(t)
}