package day02

import (
	"testing"
)


func TestSolvePart1Full(t *testing.T) {
    r, _ := Solve(true, false)
    e := 31839939622
    if r.invalids != e {
        t.Errorf("Expected part 1 result %d, but got %d", e, r.invalids)
        t.FailNow()
    }
}

func TestSolvePart1(t *testing.T) {

    data := []string{
        "11-22","95-115","998-1012","1188511880-1188511890","222220-222224",
        "1698522-1698528","446443-446449","38593856-38593862","565653-565659",
        "824824821-824824827","2121212118-2121212124",
    }

    r := SolvePart1(data)
    e := 1227775554

    if r.invalids != e {
        t.Errorf("Expected %d, but got %d", e, r.invalids)
        t.FailNow()
    }
}

func TestSolvePart2Full(t *testing.T) {
    _, r2 := Solve(false, true)
    e := 41662374059
    if r2.invalids != e {
        t.Errorf("Expected part 2 result %d, but got %d", e, r2.invalids)
        t.FailNow()
    }
}

func TestSolvePart2(t *testing.T) {
    data := []string{
        "11-22","95-115","998-1012","1188511880-1188511890","222220-222224",
        "1698522-1698528","446443-446449","38593856-38593862","565653-565659",
        "824824821-824824827","2121212118-2121212124",
    }

    r := SolvePart2(data)
    e := 4174379265

    if r.invalids != e {
        t.Errorf("Expected %d, but got %d", e, r.invalids)
        t.FailNow()
    }
}

func TestSolveBoth(t *testing.T) {
    TestSolvePart1Full(t)
    TestSolvePart2Full(t)
}