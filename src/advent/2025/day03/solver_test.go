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
}

func TestSolvePart2Full(t *testing.T) {
    _, r2 := Solve(false, true)
    e := 169685670469164
    if r2.jolts != e {
        t.Errorf("Expected part 2 result %d, but got %d", e, r2.jolts)
        t.FailNow()
    }
}

func TestSolvePart2(t *testing.T) {
    data := []string{
        "987654321111111",
        "811111111111119",
        "234234234234278",
        "818181911112111",
    }

    r := SolvePart2(data)
    e := 3121910778619

    if r.jolts != e {
        t.Errorf("Expected %d, but got %d", e, r.jolts)
        t.FailNow()
    }
}

func TestSolveBoth(t *testing.T) {
    TestSolvePart1Full(t)
    TestSolvePart2Full(t)
}