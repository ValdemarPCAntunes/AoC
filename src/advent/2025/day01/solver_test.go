package day01

import (
	"testing"
)


func TestSolvePart1Full(t *testing.T) {
    r, _ := Solve(true, false)
    e := 1034
    if r.pwd != e {
        t.Errorf("Expected part 1 result %d, but got %d", e, r.pwd)
        t.FailNow()
    }
}

func TestSolvePart1(t *testing.T) {

    data := []string{
        "L68",
        "L30",
        "R48",
        "L5",
        "R60",
        "L55",
        "L1",
        "L99",
        "R14",
        "L82",
    }

    r := SolvePart1(data)
    e := 3

    if r.pwd != e {
        t.Errorf("Expected %d, but got %d", e, r.pwd)
        t.FailNow()
    }
}

func TestSolvePart2Full(t *testing.T) {
    _, r2 := Solve(false, true)
    e := 6166
    if r2.pwd != e {
        t.Errorf("Expected part 2 result %d, but got %d", e, r2.pwd)
        t.FailNow()
    }
}

func TestSolvePart2(t *testing.T) {
    data := []string{
        "L68",
        "L30",
        "R48",
        "L5",
        "R60",
        "L55",
        "L1",
        "L99",
        "R14",
        "L82",
    }

    r := SolvePart2(data)
    e := 6

    if r.pwd != e {
        t.Errorf("Expected %d, but got %d", e, r.pwd)
        t.FailNow()
    }
}

func TestSolveBoth(t *testing.T) {
    r1, r2 := Solve(true, true)
    e1, e2 := 1034, 6166
    if r1.pwd != e1 {
        t.Errorf("Expected part 1 result %d, but got %d", e1, r1.pwd)
        t.FailNow()
    }
    if r2.pwd != e2 {
        t.Errorf("Expected part 2 result %d, but got %d", e2, r2.pwd)
        t.FailNow()
    }
}