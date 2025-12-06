package day6

import (
	"testing"
)


func TestSolvePart1Full(t *testing.T) {
    r, _ := Solve(true, false)
    e := 4076006202939
    if r.result != e {
        t.Errorf("Expected part 1 result %d, but got %d", e, r.result)
        t.FailNow()
    }
}

func TestSolvePart1(t *testing.T) {

    data := []string{
        "123 328 51  64",
        " 45 64  387 23",
        "  6 98  215 314",
        "*   +   *   +",
    }

    r := SolvePart1(data)
    e := 4277556

    if r.result != e {
        t.Errorf("Expected %d, but got %d", e, r.result)
        t.FailNow()
    }
    t.Logf("TestSolvePart1 passed")
}

func TestSolvePart2Full(t *testing.T) {
    _, r2 := Solve(false, true)
    e := 0
    if r2.result != e {
        t.Errorf("Expected part 2 result %d, but got %d", e, r2.result)
        t.FailNow()
    }
}

func TestSolvePart2(t *testing.T) {
    data := []string{
        "123 328 51  64 ",
        " 45 64  387 23 ",
        "  6 98  215 314",
        "*   +   *   +",
    }

    r := SolvePart2(data)
    e := 0

    if r.result != e {
        t.Errorf("Expected %d, but got %d", e, r.result)
        t.FailNow()
    }
    t.Logf("TestSolvePart1 passed")
}

func TestSolveBoth(t *testing.T) {
    TestSolvePart1Full(t)
    TestSolvePart2Full(t)
}