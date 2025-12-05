package day5

import (
	"testing"
)


func TestSolvePart1Full(t *testing.T) {
    r, _ := Solve(true, false)
    e := 509
    if r.fresh_ingredients != e {
        t.Errorf("Expected part 1 result %d, but got %d", e, r.fresh_ingredients)
        t.FailNow()
    }
}

func TestSolvePart1(t *testing.T) {

    data := []string{
        "3-5",
        "10-14",
        "16-20",
        "12-18",
        "",
        "1",
        "5",
        "8",
        "11",
        "17",
        "32",
    }

    r := SolvePart1(data)
    e := 3

    if r.fresh_ingredients != e {
        t.Errorf("Expected %d, but got %d", e, r.fresh_ingredients)
        t.FailNow()
    }
    t.Logf("TestSolvePart1 passed")
}

func TestSolvePart2Full(t *testing.T) {
    _, r2 := Solve(false, true)
    e := 336790092076620
    if r2.fresh_ingredients != e {
        t.Errorf("Expected part 2 result %d, but got %d", e, r2.fresh_ingredients)
        t.FailNow()
    }
}

func TestSolvePart2(t *testing.T) {
    data := []string{
        "3-5",
        "10-14",
        "16-20",
        "12-18",
        "",
        "1",
        "5",
        "8",
        "11",
        "17",
        "32",
    }

    r := SolvePart2(data)
    e := 14

    if r.fresh_ingredients != e {
        t.Errorf("Expected %d, but got %d", e, r.fresh_ingredients)
        t.FailNow()
    }
    t.Logf("TestSolvePart1 passed")
}

func TestSolveBoth(t *testing.T) {
    TestSolvePart1Full(t)
    TestSolvePart2Full(t)
}