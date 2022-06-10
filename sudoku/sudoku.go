package sudoku

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

type Sudoku [9][9]int

func LoadFromJsonFile(filepath string) *Sudoku {
	var s Sudoku
	content, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(content, &s)
	return &s
}

func (s *Sudoku) String() string {
	var res string
	res = "\n-------------------------------------\n"
	for i, row := range *s {
		res += "|"
		for j, v := range row {
			res += fmt.Sprintf(" %d ", v)
			if (j+1)%3 == 0 {
				res += "|"
			} else {
				res += " "
			}
		}
		if (i+1)%3 == 0 {
			res += "\n-------------------------------------\n"
		} else {
			res += "\n"
		}
	}
	return res
}

func (s *Sudoku) SaveToJsonFile(filepath string) error {
	if data, err := json.Marshal(*s); err != nil {
		return err
	} else {
		if err := os.WriteFile(filepath, data, 0644); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func (s *Sudoku) IsEqual(other *Sudoku) bool {
	for i, row := range s {
		for j, v := range row {
			if v != other[i][j] {
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) CellIsEmpty(i int, j int) bool {
	return s[i][j] == 0
}

func (s *Sudoku) GetValuesInSquare(i int, j int) []int {
	var res []int
	istart := (i / 3) * 3
	jstart := (j / 3) * 3

	for idx := 0; idx < 9; idx++ {
		for jdx := 0; jdx < 9; jdx++ {
			if idx >= istart && idx < (istart+3) {
				if jdx >= jstart && jdx < (jstart+3) {
					if s[idx][jdx] > 0 {
						res = append(res, s[idx][jdx])
					}
				}
			}
		}
	}

	return res
}

func (s *Sudoku) GetValuesInRow(i int) []int {
	var res []int

	for idx, row := range s {
		if idx == i {
			for _, v := range row {
				if v > 0 {
					res = append(res, v)
				}
			}
		}
	}
	return res
}

func (s *Sudoku) GetValuesInCol(j int) []int {
	var res []int

	for _, row := range s {
		for idx, v := range row {

			if idx == j && v > 0 {
				res = append(res, v)
			}
		}
	}
	return res
}

func (s *Sudoku) GetCandidates(i int, j int) []int {
	var res []int

	vInSquare := s.GetValuesInSquare(i, j)
	vInRow := s.GetValuesInRow(i)
	vInCol := s.GetValuesInCol(j)

	for c := 1; c < 10; c++ {
		if !contains(vInSquare, c) && !contains(vInRow, c) && !contains(vInCol, c) {
			res = append(res, c)
		}
	}
	return res
}

func (s *Sudoku) NextEmptyPosition(i int, j int) (int, int, error) {
	var ni int
	var nj int

	if i > 8 || j > 8 || i < 0 || j < 0 {
		return 0, 0, errors.New("position out of range")
	}

	if i == 8 && j == 8 {
		return 0, 0, errors.New("last position reached")
	}

	pos := i*9 + j
	newPos := pos + 1
	ni = newPos / 9
	nj = newPos % 9

	if s.CellIsEmpty(ni, nj) {
		return ni, nj, nil
	} else {
		return s.NextEmptyPosition(ni, nj)
	}
}

func (s *Sudoku) PreviousEmptyPosition(i int, j int) (int, int, error) {
	var ni int
	var nj int

	if i > 8 || j > 8 || i < 0 || j < 0 {
		return 0, 0, errors.New("position out of range")
	}

	if i == 0 && j == 0 {
		return 0, 0, errors.New("first position reached")
	}

	pos := i*9 + j
	newPos := pos - 1
	ni = newPos / 9
	nj = newPos % 9

	if s.CellIsEmpty(ni, nj) {
		return ni, nj, nil
	} else {
		return s.PreviousEmptyPosition(ni, nj)
	}
}

func (s *Sudoku) LastEmptyPosition() (int, int) {
	var li int
	var lj int

	for i, row := range s {
		for j, v := range row {
			if v == 0 {
				li = i
				lj = j
			}
		}
	}

	return li, lj
}

func (s *Sudoku) IsSolved() bool {
	for _, row := range s {
		for _, v := range row {
			if v == 0 {
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) DeepCopy() *Sudoku {
	var newSudoku Sudoku

	for i, row := range s {
		for j, v := range row {
			newSudoku[i][j] = v
		}
	}
	return &newSudoku
}

func (s *Sudoku) NewTopDownSwitch() *Sudoku {
	var result Sudoku
	// top row to bottom row
	for i := 0; i < 3; i++ {
		for j := 0; j < 9; j++ {
			result[i+6][j] = s[i][j]
		}
	}
	// keep middle row
	for i := 3; i < 6; i++ {
		for j := 0; j < 9; j++ {
			result[i][j] = s[i][j]
		}
	}
	// bottom row to top row
	for i := 6; i < 9; i++ {
		for j := 0; j < 9; j++ {
			result[i-6][j] = s[i][j]
		}
	}
	return &result
}

func (s *Sudoku) NewLeftRightSwitch() *Sudoku {
	var result Sudoku
	// left column to right column
	for i := 0; i < 9; i++ {
		for j := 0; j < 3; j++ {
			result[i][j+6] = s[i][j]
		}
	}
	// middle columns stays
	for i := 0; i < 9; i++ {
		for j := 3; j < 6; j++ {
			result[i][j] = s[i][j]
		}
	}
	// right column to left column
	for i := 0; i < 9; i++ {
		for j := 6; j < 9; j++ {
			result[i][j-6] = s[i][j]
		}
	}
	return &result
}

func (s *Sudoku) NewFullSwitch() *Sudoku {
	s1 := s.NewTopDownSwitch()
	s2 := s1.NewLeftRightSwitch()
	return s2
}

func (s *Sudoku) SolveScore() (score int) {
	for i, row := range s {
		for j, v := range row {
			idx := i*9 + j
			if v > 0 {
				score += (80 - idx)
			}
		}
	}
	return score
}

func (s *Sudoku) SwitchScore() (original int, topDown int, leftRight int, full int) {
	td := s.NewTopDownSwitch()
	lr := s.NewLeftRightSwitch()
	fu := s.NewFullSwitch()
	return s.SolveScore(), td.SolveScore(), lr.SolveScore(), fu.SolveScore()
}

func (s *Sudoku) FindBestSwitch() (switchFunc func(_ *Sudoku) *Sudoku) {
	orig := s.SolveScore()
	topdown := s.NewTopDownSwitch().SolveScore()
	leftright := s.NewLeftRightSwitch().SolveScore()
	full := s.NewFullSwitch().SolveScore()

	if full >= leftright && full >= topdown && full >= orig {
		return (*Sudoku).NewFullSwitch
	}

	if leftright >= full && leftright >= topdown && leftright >= orig {
		return (*Sudoku).NewLeftRightSwitch
	}

	if topdown >= full && topdown >= leftright && topdown >= orig {
		return (*Sudoku).NewTopDownSwitch
	}
	return (*Sudoku).DeepCopy
}
