package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/andrea3bianchi/aoc2020/5/partitioner"
)

func main() {
	f, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var highest int
	seats := make(map[int]bool)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		seatID, row := getSeatID(scanner.Text())
		if seatID > highest {
			highest = seatID
		}
		if row != 0 && row != 127 {
			seats[seatID] = true
		}
	}
	fmt.Println(highest)
	for i := 0; i < highest; i++ {
		if !seats[i] &&
			seats[i-1] &&
			seats[i+1] {
			fmt.Println(i)
		}
	}
}

func getSeatID(s string) (int, int) {
	row := getRow(s[:7])
	column := getColumn(s[7:])

	return row*8 + column, row
}

func getRow(s string) int {
	p := partitioner.New(127)
	for _, r := range s {
		switch r {
		case 'F':
			p = p.Under()
		case 'B':
			p = p.Over()
		}
	}
	row, err := p.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	return row
}

func getColumn(s string) int {
	p := partitioner.New(7)
	for _, r := range s {
		switch r {
		case 'R':
			p = p.Over()
		case 'L':
			p = p.Under()
		}
	}
	row, err := p.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	return row
}
