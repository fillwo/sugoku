package main

import (
	"fmt"

	// "sugoku/generator"
	"sugoku/solver"
	"sugoku/sudoku"
)

func main() {
	fmt.Println("sugoku ...")
	// var fails, success int

	s := sudoku.LoadFromJsonFile("sudoku1.json")

	sudokuSolver := solver.NewSolver(s)

	result := sudokuSolver.Solve()

	fmt.Printf("%v\n", result)

	fmt.Println("check if solution is unique:")

	fmt.Println(result.IsOnlySolution())

	// generate filled sudokus
	// _, fails, success = generator.GenerateFinishedSudokus(81, 5)
	// fmt.Println(fails, success)
}
