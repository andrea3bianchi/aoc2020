package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct {
	X int
	Y int
}

func main() {
	f, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	trees := make(map[point]bool)
	var length, line int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		length = len(scanner.Text())
		for i, x := range scanner.Text() {
			if x == '#' {
				trees[point{
					X: i,
					Y: line,
				}] = true
			}
		}
		line++
	}

	fmt.Println(countTreesCrossed(3, 1, length, line, trees))

	fmt.Println(countTreesCrossed(1, 1, length, line, trees) *
		countTreesCrossed(3, 1, length, line, trees) *
		countTreesCrossed(5, 1, length, line, trees) *
		countTreesCrossed(7, 1, length, line, trees) *
		countTreesCrossed(1, 2, length, line, trees))

}

func countTreesCrossed(right, down, length, depth int, trees map[point]bool) int {
	var ourPosition point
	var treesCrossed int
	for {
		ourPosition.X += right
		ourPosition.Y += down

		positionToCheck := ourPosition
		positionToCheck.X = positionToCheck.X % length
		if trees[positionToCheck] {
			treesCrossed++
		}
		if ourPosition.Y >= depth {
			return treesCrossed
		}
	}
}
