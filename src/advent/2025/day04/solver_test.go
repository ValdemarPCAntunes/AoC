package day4

import (
	"testing"
)


func TestSolvePart1Full(t *testing.T) {
    r, _ := Solve(true, false)
    e := 0
    if r.paper_rolls != e {
        t.Errorf("Expected part 1 result %d, but got %d", e, r.paper_rolls)
        t.FailNow()
    }
}

func TestSolvePart1(t *testing.T) {

    data := []string{
        "..@@.@@@@.",
        "@@@.@.@.@@",
        "@@@@@.@.@@",
        "@.@@@@..@.",
        "@@.@@@@.@@",
        ".@@@@@@@.@",
        ".@.@.@.@@@",
        "@.@@@.@@@@",
        ".@@@@@@@@.",
        "@.@.@@@.@.",
    }

    r := SolvePart1(data)
    e := 13

    if r.paper_rolls != e {
        t.Errorf("Expected %d, but got %d", e, r.paper_rolls)
        t.FailNow()
    }
    t.Logf("TestSolvePart1 passed")
}

func TestSolvePart2Full(t *testing.T) {
    _, r2 := Solve(false, true)
    e := 0
    if r2.paper_rolls != e {
        t.Errorf("Expected part 2 result %d, but got %d", e, r2.paper_rolls)
        t.FailNow()
    }
}

func TestSolvePart2(t *testing.T) {
    data := []string{
        "..@@.@@@@.",
        "@@@.@.@.@@",
        "@@@@@.@.@@",
        "@.@@@@..@.",
        "@@.@@@@.@@",
        ".@@@@@@@.@",
        ".@.@.@.@@@",
        "@.@@@.@@@@",
        ".@@@@@@@@.",
        "@.@.@@@.@.",
    }

    r := SolvePart2(data)
    e := 0

    if r.paper_rolls != e {
        t.Errorf("Expected %d, but got %d", e, r.paper_rolls)
        t.FailNow()
    }
    t.Logf("TestSolvePart1 passed")
}

func TestSolveBoth(t *testing.T) {
    TestSolvePart1Full(t)
    TestSolvePart2Full(t)
}