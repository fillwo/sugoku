package solver

import (
	"sugoku/sudoku"
	"testing"
)

func TestSolve(t *testing.T) {
	sdoku := sudoku.LoadFromJsonFile("../sudoku1.json")
	solution := sudoku.LoadFromJsonFile("../solution1.json")

	s := NewSolver(sdoku)
	result := s.Solve()
	if !result.Success {
		t.Fatalf("sudoku could not be solved")
	}

	correct := solution.IsEqual(result.Solution)
	if !correct {
		t.Fatalf("solution is wrong")
	}

	if !result.IsOnlySolution() {
		t.Fatalf("multiple solutions detected")
	}
}

func TestDetectNotUnique(t *testing.T) {
	sdoku := sudoku.LoadFromJsonFile("../sudoku2.json")
	s := NewSolver(sdoku)
	result := s.Solve()

	if result.IsOnlySolution() {
		t.Fatalf("did not detect multiple solutions")
	}
}
