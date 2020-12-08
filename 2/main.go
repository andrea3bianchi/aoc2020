package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const target = 2020

func main() {
	f, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var validPasswords, validPasswordsTwo int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if isPasswordValid(scanner.Text()) {
			validPasswords++
		}
		if isPasswordValidTwo(scanner.Text()) {
			validPasswordsTwo++
		}
	}
	fmt.Println(validPasswords)
	fmt.Println(validPasswordsTwo)

}

func isPasswordValid(s string) bool {
	ss := strings.Split(s, " ")
	zx := strings.Split(ss[0], "-")
	min, _ := strconv.Atoi(zx[0])
	max, _ := strconv.Atoi(zx[1])
	letter := rune(ss[1][0])
	password := ss[2]

	var counter int
	for _, l := range password {
		if l == letter {
			counter++
		}
	}

	return counter >= min && counter <= max
}

func isPasswordValidTwo(s string) bool {
	ss := strings.Split(s, " ")
	zx := strings.Split(ss[0], "-")
	min, _ := strconv.Atoi(zx[0])
	max, _ := strconv.Atoi(zx[1])
	letter := ss[1][0]
	password := ss[2]

	// golang doesn't have a XOR operator, it needs to be defined explicitly

	return (password[min-1] == letter) != (password[max-1] == letter)

	// below is a more 'naive' solution

	// var counter int
	// for pos, l := range password {
	// 	if pos+1 == min || pos+1 == max {
	// 		if l == letter {
	// 			counter++
	// 		}
	// 	}
	// }

	// return counter == 1
}
