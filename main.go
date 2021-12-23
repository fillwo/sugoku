package main

import (
	"fmt"

	"sugoku/solver"
	"sugoku/sudoku"
)

func main() {
	fmt.Println("sugoku ...")

	s := sudoku.LoadFromJsonFile("sudoku1.json")
	sudokuSolver := solver.NewSolver(s)

	result := sudokuSolver.Solve()

	fmt.Printf("%v", result)
}
