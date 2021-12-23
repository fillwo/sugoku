package solver

import (
	"sugoku/sudoku"
)

type SPosition struct {
	I int
	J int
}

type SolveResult struct {
	Sudoku     *sudoku.Sudoku
	Solution   *sudoku.Sudoku
	Iterations int
	Success    bool
}

type Solver struct {
	Sudoku   *sudoku.Sudoku
	Current  *sudoku.Sudoku
	Position SPosition
}

func valuesBiggerThan(arr []int, biggerThan int) []int {
	var res []int
	for _, v := range arr {
		if v > biggerThan {
			res = append(res, v)
		}
	}
	return res
}

func NewSolver(s *sudoku.Sudoku) Solver {
	var tmp sudoku.Sudoku

	for i, row := range s {
		for j, v := range row {
			tmp[i][j] = v
		}
	}

	return Solver{
		Sudoku:   s,
		Current:  &tmp,
		Position: SPosition{0, 0},
	}
}

func (s *Solver) SolveStep() (bool, error) {
	i := s.Position.I
	j := s.Position.J
	// if cell is not empty move on to the next
	if !s.Sudoku.CellIsEmpty(i, j) {
		ni, nj, err := s.Sudoku.NextEmptyPosition(i, j)
		if err != nil {
			return false, err
		}
		s.Position.I = ni
		s.Position.J = nj
		return false, nil
	}
	// determine possible candidates
	candidates := s.Current.GetCandidates(i, j)
	candidates = valuesBiggerThan(candidates, s.Current[i][j])
	// cadidates exists
	if len(candidates) > 0 {
		s.Current[i][j] = candidates[0]
		ni, nj, err := s.Sudoku.NextEmptyPosition(i, j)
		if err != nil {
			return true, err
		}
		s.Position.I = ni
		s.Position.J = nj
		return false, nil
	}
	// no candidates found - need to go back
	ni, nj, err := s.Sudoku.PreviousEmptyPosition(i, j)
	// deleting current value
	s.Current[i][j] = 0
	if err != nil {
		return false, err
	}
	s.Position.I = ni
	s.Position.J = nj
	return false, nil
}

func (s *Solver) Solve() SolveResult {
	var res bool
	var err error
	var solution sudoku.Sudoku
	counter := 0

	for true {
		counter++
		res, err = s.SolveStep()
		if err != nil {
			break
		}
	}
	// successfully solved
	if res {
		// copy solution
		for i, row := range s.Current {
			for j, v := range row {
				solution[i][j] = v
			}
		}
		return SolveResult{
			Sudoku:     s.Sudoku,
			Solution:   &solution,
			Iterations: counter,
			Success:    res,
		}
	} else {
		// could not solve
		return SolveResult{
			Sudoku:     s.Sudoku,
			Iterations: counter,
			Success:    res,
		}
	}
}
