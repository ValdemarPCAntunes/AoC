package day7

import (
	"testing"
)


func TestSolvePart1Full(t *testing.T) {
    r, _ := Solve(true, false)
    e := 1600
    if r.splits != e {
        t.Errorf("Expected part 1 result %d, but got %d", e, r.splits)
        t.FailNow()
    }
}

func TestSolvePart1(t *testing.T) {

    data := []string{
        ".......S.......",
        "...............",
        ".......^.......",
        "...............",
        "......^.^......",
        "...............",
        ".....^.^.^.....",
        "...............",
        "....^.^...^....",
        "...............",
        "...^.^...^.^...",
        "...............",
        "..^...^.....^..",
        "...............",
        ".^.^.^.^.^...^.",
        "...............",
    }

    r := SolvePart1(data)
    e := 21

    if r.splits != e {
        t.Errorf("Expected %d, but got %d", e, r.splits)
        t.FailNow()
    }
}

func TestSolvePart2Full(t *testing.T) {
    _, r2 := Solve(false, true)
    e := 8632253783011
    if r2.timelines != e {
        t.Errorf("Expected part 2 result %d, but got %d", e, r2.timelines)
        t.FailNow()
    }
}

func TestSolvePart2(t *testing.T) {
    data := []string{
        ".......S.......",
        "...............",
        ".......^.......",
        "...............",
        "......^.^......",
        "...............",
        ".....^.^.^.....",
        "...............",
        "....^.^...^....",
        "...............",
        "...^.^...^.^...",
        "...............",
        "..^...^.....^..",
        "...............",
        ".^.^.^.^.^...^.",
        "...............",
    }

    r := SolvePart2(data)
    e := 40

    if r.timelines != e {
        t.Errorf("Expected %d, but got %d", e, r.timelines)
        t.FailNow()
    }
}

func TestSolveBoth(t *testing.T) {
    TestSolvePart1Full(t)
    TestSolvePart2Full(t)
}