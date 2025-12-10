package day8

import (
	"testing"
)


func TestSolvePart1Full(t *testing.T) {
    r, _ := Solve(true, false)
    e := 42840
    if r.circuits != e {
        t.Errorf("Expected part 1 result %d, but got %d", e, r.circuits)
        t.FailNow()
    }
}

func TestSolvePart1(t *testing.T) {

    data := []string{
        "162,817,812",
        "57,618,57",
        "906,360,560",
        "592,479,940",
        "352,342,300",
        "466,668,158",
        "542,29,236",
        "431,825,988",
        "739,650,466",
        "52,470,668",
        "216,146,977",
        "819,987,18",
        "117,168,530",
        "805,96,715",
        "346,949,466",
        "970,615,88",
        "941,993,340",
        "862,61,35",
        "984,92,344",
        "425,690,689",
    }

    r := SolvePart1(data)
    e := 45

    if r.circuits != e {
        t.Errorf("Expected %d, but got %d", e, r.circuits)
        t.FailNow()
    }
}

func TestSolvePart2Full(t *testing.T) {
    _, r2 := Solve(false, true)
    e := 170629052
    if r2.lastJunctionXmul != e {
        t.Errorf("Expected part 2 result %d, but got %d", e, r2.lastJunctionXmul)
        t.FailNow()
    }
}

func TestSolvePart2(t *testing.T) {
    data := []string{
        "162,817,812",
        "57,618,57",
        "906,360,560",
        "592,479,940",
        "352,342,300",
        "466,668,158",
        "542,29,236",
        "431,825,988",
        "739,650,466",
        "52,470,668",
        "216,146,977",
        "819,987,18",
        "117,168,530",
        "805,96,715",
        "346,949,466",
        "970,615,88",
        "941,993,340",
        "862,61,35",
        "984,92,344",
        "425,690,689", 
    }

    r := SolvePart2(data)
    e := 25272

    if r.lastJunctionXmul != e {
        t.Errorf("Expected %d, but got %d", e, r.lastJunctionXmul)
        t.FailNow()
    }
}

func TestSolveBoth(t *testing.T) {
    TestSolvePart1Full(t)
    TestSolvePart2Full(t)
}