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

	counter := 0

	for true {
		counter++
		res, err := sudokuSolver.SolveStep()
		if err != nil {
			fmt.Printf("res: %v\n", res)
			fmt.Printf("err: %v\n", err)
			break
		}
	}
	fmt.Printf("position: %v\n", sudokuSolver.Position)
	fmt.Println("counter:", counter)
	fmt.Println(sudokuSolver.Current)
}
