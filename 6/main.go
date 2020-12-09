package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type answers map[rune]int
type group []answers

func main() {
	f, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	all := make([]group, 1)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		// check for the separating newlines, add an empty group,
		// below always work on latest group
		if len(t) == 0 {
			all = append(all, group{})
			continue
		}
		currentGroupIdx := len(all) - 1
		personAnswers := make(answers)
		all[currentGroupIdx] = append(all[currentGroupIdx], personAnswers)
		for _, r := range t {
			personAnswers[r]++
		}
	}
	var (
		totalAnswers  int
		totalAnswersB int
	)
	for _, group := range all {
		groupAnswers := make(answers)
		for _, personAnswers := range group {
			for question := range personAnswers {
				groupAnswers[question]++
			}
		}
		for _, all := range groupAnswers {
			totalAnswers++
			if len(group) == all {
				totalAnswersB++
			}
		}
	}
	fmt.Println(totalAnswers)
	fmt.Println(totalAnswersB)

}
