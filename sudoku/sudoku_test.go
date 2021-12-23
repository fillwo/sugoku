package sudoku

import (
	"testing"
)

func TestValuesInRow(t *testing.T) {
	s := LoadFromJsonFile("../sudoku1.json")
	values := s.GetValuesInRow(7) // [0,1,0,0,7,0,0,8,6]
	if values[0] != 1 {
		t.Fatalf("expected 1")
	}
	if values[1] != 7 {
		t.Fatalf("expected 7")
	}
	if values[2] != 8 {
		t.Fatalf("expected 8")
	}
	if values[3] != 6 {
		t.Fatalf("expected 6")
	}
}

func TestValuesInCol(t *testing.T) {
	s := LoadFromJsonFile("../sudoku1.json")
	values := s.GetValuesInCol(4) // 1,8,7,3
	if values[0] != 1 {
		t.Fatalf("expected 1")
	}
	if values[1] != 8 {
		t.Fatalf("expected 8")
	}
	if values[2] != 7 {
		t.Fatalf("expected 7")
	}
	if values[3] != 3 {
		t.Fatalf("expected 3")
	}
}

func TestValuesInSquare(t *testing.T) {
	s := LoadFromJsonFile("../sudoku1.json")
	values := s.GetValuesInSquare(1, 4) // 2,1,8
	if len(values) != 3 {
		t.Fatalf("length should be 3, but got %d", len(values))
	}
	if values[0] != 2 {
		t.Fatalf("expected 2")
	}
	if values[1] != 1 {
		t.Fatalf("expected 1")
	}
	if values[2] != 8 {
		t.Fatalf("expected 8")
	}
	values = s.GetValuesInSquare(8, 7) // 8,6,9,1
	if len(values) != 4 {
		t.Fatalf("length should be 4, but got %d", len(values))
	}
	if values[0] != 8 {
		t.Fatalf("expected 8")
	}
	if values[1] != 6 {
		t.Fatalf("expected 6")
	}
	if values[2] != 9 {
		t.Fatalf("expected 9")
	}
	if values[3] != 1 {
		t.Fatalf("expected 1")
	}
}

func TestGetCandidates(t *testing.T) {
	s := LoadFromJsonFile("../sudoku1.json")
	candidates := s.GetCandidates(1, 1) // 3
	if len(candidates) != 1 {
		t.Fatalf("length should be 1, but got %d", len(candidates))
	}
	if candidates[0] != 3 {
		t.Fatalf("expected 3")
	}
	candidates = s.GetCandidates(4, 8) // 1,2,3,5
	if len(candidates) != 4 {
		t.Fatalf("length should be 4, but got %d", len(candidates))
	}
	if candidates[0] != 1 {
		t.Fatalf("expected 1")
	}
	if candidates[1] != 2 {
		t.Fatalf("expected 2")
	}
	if candidates[2] != 3 {
		t.Fatalf("expected 3")
	}
	if candidates[3] != 5 {
		t.Fatalf("expected 5")
	}
}

func TestNextEmptyPosition(t *testing.T) {
	s := LoadFromJsonFile("../sudoku1.json")
	i, j, err := s.NextEmptyPosition(0, 0)
	if err != nil {
		t.Fatalf("should not fail")
	}
	if i != 0 {
		t.Fatalf("expected 0")
	}
	if j != 4 {
		t.Fatalf("expected 4")
	}
	i, j, err = s.NextEmptyPosition(5, 8)
	if err != nil {
		t.Fatalf("should not fail")
	}
	if i != 6 {
		t.Fatalf("expected 6")
	}
	if j != 3 {
		t.Fatalf("expected 3")
	}
	i, j, err = s.NextEmptyPosition(8, 8)
	if err == nil {
		t.Fatalf("this is supposed to fail")
	}
}

func TestPreviousEmptyPosition(t *testing.T) {
	s := LoadFromJsonFile("../sudoku1.json")
	i, j, err := s.PreviousEmptyPosition(0, 4)
	if err != nil {
		t.Fatalf("should not fail")
	}
	if i != 0 {
		t.Fatalf("expected 0")
	}
	if j != 0 {
		t.Fatalf("expected 0")
	}
	i, j, err = s.PreviousEmptyPosition(6, 2)
	if err != nil {
		t.Fatalf("should not fail")
	}
	if i != 5 {
		t.Fatalf("expected 5")
	}
	if j != 6 {
		t.Fatalf("expected 6")
	}
	i, j, err = s.PreviousEmptyPosition(0, 0)
	if err == nil {
		t.Fatalf("this is supposed to fail")
	}
}
