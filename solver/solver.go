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
	Sudoku       *sudoku.Sudoku
	EasySolution *sudoku.Sudoku
	Current      *sudoku.Sudoku
	Position     SPosition
	STrafo       func(_ *sudoku.Sudoku) *sudoku.Sudoku
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
	return Solver{
		Sudoku:       s,
		EasySolution: s.FindBestSwitch()(s),
		Current:      s.FindBestSwitch()(s),
		Position:     SPosition{0, 0},
		STrafo:       s.FindBestSwitch(),
	}
}

func (s *Solver) EasyStep() bool {
	var numChanges int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.EasySolution[i][j] == 0 {
				candidates := s.EasySolution.GetCandidates(i, j)
				if len(candidates) == 1 {
					s.EasySolution[i][j] = candidates[0]
					numChanges++
				}
			}
		}
	}
	return numChanges > 0
}

func (s *Solver) SolveStep() (bool, error) {
	i := s.Position.I
	j := s.Position.J
	// if cell is not empty move on to the next
	if !s.EasySolution.CellIsEmpty(i, j) {
		ni, nj, err := s.EasySolution.NextEmptyPosition(i, j)
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
		ni, nj, err := s.EasySolution.NextEmptyPosition(i, j)
		if err != nil {
			return true, err
		}
		s.Position.I = ni
		s.Position.J = nj
		return false, nil
	}
	// no candidates found - need to go back
	ni, nj, err := s.EasySolution.PreviousEmptyPosition(i, j)
	// deleting current value
	s.Current[i][j] = 0
	if err != nil {
		return false, err
	}
	s.Position.I = ni
	s.Position.J = nj
	return false, nil
}

func (s *Solver) SolveWithOption(tryEasy bool) SolveResult {
	var res bool
	var err error

	counter := 0

	// try easy solution first (not when checking for unique solution)
	if tryEasy {
		for {
			counter++
			if ok := s.EasyStep(); !ok {
				break
			}
		}
		if s.EasySolution.IsSolved() {
			return SolveResult{
				Sudoku:     s.Sudoku,
				Solution:   s.STrafo(s.EasySolution),
				Iterations: counter,
				Success:    true,
			}
		}
	}

	// solution using backtracing
	for {
		counter++
		res, err = s.SolveStep()
		if err != nil {
			break
		}
	}
	// successfully solved
	if res {
		return SolveResult{
			Sudoku:     s.Sudoku,
			Solution:   s.STrafo(s.Current),
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

func (s *Solver) Solve() SolveResult {
	return s.SolveWithOption(true)
}

func (sr *SolveResult) IsOnlySolution() bool {
	current := *sr.Solution.DeepCopy()

	li, lj := sr.Sudoku.LastEmptyPosition()
	current[li][lj] = 0

	li, lj, err := sr.Sudoku.PreviousEmptyPosition(li, lj)
	if err != nil {
		panic(err)
	}

	pos := SPosition{li, lj}

	solver := Solver{
		Sudoku:       sr.Sudoku,
		Current:      &current,
		EasySolution: sr.Sudoku.DeepCopy(),
		Position:     pos,
		STrafo:       (*sudoku.Sudoku).DeepCopy,
	}

	result := solver.SolveWithOption(false)

	return !result.Success
}
