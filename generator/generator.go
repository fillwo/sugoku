package generator

import (
	"errors"
	"fmt"
	"math/rand"
	"sugoku/solver"
	"sugoku/sudoku"
)

func randomPick(input []int) (int, error) {
	length := len(input)
	if length == 0 {
		return 0, errors.New("cannot pick from empty slice")
	}
	ridx := rand.Intn(length)
	return input[ridx], nil
}

func coordinatesToNumber(i int, j int) int {
	return 9*i + j
}

func numberToCoordinates(n int) (int, int) {
	return n / 9, n % 9
}

func randomOrder() []int {
	var res []int
	for i := 0; i < 81; i++ {
		res = append(res, i)
	}
	rand.Shuffle(len(res), func(i, j int) { res[i], res[j] = res[j], res[i] })
	return res
}

// todo:
// 1. shuffle 0,1,2,3,....80
// 2. random fill until problem
// 3. use solver

func TryGenerateFinishedSudoku(randomFillLen int) (*sudoku.Sudoku, bool) {
	var result sudoku.Sudoku
	var breakIdx int
	order := randomOrder()

	// fmt.Println(order)

	for idx, num := range order[0:randomFillLen] {
		i, j := numberToCoordinates(num)
		candidates := result.GetCandidates(i, j)
		c, err := randomPick(candidates)
		if err != nil {
			breakIdx = idx
			break
		}
		result[i][j] = c
	}

	if breakIdx > 0 {
		// fmt.Println("got stuck here")
		i, j := numberToCoordinates(order[breakIdx-1])
		result[i][j] = 0
	}

	fmt.Println("start solving ...")
	fmt.Println(&result)

	solver := solver.NewSolver(&result)
	solveResult := solver.Solve()

	fmt.Println("solved", solveResult.Success)

	// fmt.Println("solution")
	// fmt.Println(solveResult.Success)
	// fmt.Println(solveResult.Solution)

	return &result, solveResult.Success
}

func GenerateFinishedSudoku(randomFillLen int) *sudoku.Sudoku {
	for {
		result, ok := TryGenerateFinishedSudoku(randomFillLen)
		if ok {
			return result
		}
	}
}

func GenerateFinishedSudokus(randomFillLen int, amount int) (sudokus []*sudoku.Sudoku, fails int, success int) {
	for i := 0; i < amount; i++ {
		for {
			result, ok := TryGenerateFinishedSudoku(randomFillLen)
			if ok {
				sudokus = append(sudokus, result)
				success++
				break
			} else {
				fails++
			}
			fmt.Printf("success: %d, fails: %d\n", success, fails)
		}
	}
	return sudokus, fails, success
}
