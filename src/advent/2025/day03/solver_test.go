package day03

import (
	"testing"
)


func TestSolvePart1Full(t *testing.T) {
    r, _ := Solve(true, false)
    e := 17155
    if r.jolts != e {
        t.Errorf("Expected part 1 result %d, but got %d", e, r.jolts)
        t.FailNow()
    }
}

func TestSolvePart1(t *testing.T) {

    data := []string{
        "987654321111111",
        "811111111111119",
        "234234234234278",
        "818181911112111",
    }

    r := SolvePart1(data)
    e := 357

    if r.jolts != e {
        t.Errorf("Expected %d, but got %d", e, r.jolts)
        t.FailNow()
    }
    t.Logf("TestSolvePart1 passed")
}

func TestSolvePart2Full(t *testing.T) {
    _, r2 := Solve(false, true)
    e := 0
    if r2.jolts != e {
        t.Errorf("Expected part 2 result %d, but got %d", e, r2.jolts)
        t.FailNow()
    }
}

func TestSolvePart2(t *testing.T) {
    data := []string{
        
    }

    r := SolvePart2(data)
    e := 0

    if r.jolts != e {
        t.Errorf("Expected %d, but got %d", e, r.jolts)
        t.FailNow()
    }
    t.Logf("TestSolvePart1 passed")
}

func TestSolveBoth(t *testing.T) {
    r1, r2 := Solve(true, true)
    e1, e2 := 17155, 0
    if r1.jolts != e1 {
        t.Errorf("Expected part 1 result %d, but got %d", e1, r1.jolts)
        t.FailNow()
    }
    if r2.jolts != e2 {
        t.Errorf("Expected part 2 result %d, but got %d", e2, r2.jolts)
        t.FailNow()
    }
}