package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const target = 2020

func main() {
	f, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var nn []int
	numbers := make(map[int]bool)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers[num] = true
		nn = append(nn, num)
	}

	tn := time.Now()
	good(numbers)
	println("it took ", time.Now().Sub(tn).Nanoseconds(), " ns to run the dictionary based solution")
	tn = time.Now()
	naive(nn)
	println("it took ", time.Now().Sub(tn).Nanoseconds(), " ns to run the naive solution")

}

func good(numbers map[int]bool) {
	for num := range numbers {
		if numbers[target-num] {
			fmt.Println(num * (target - num))
			break
		}
	}

	for i := range numbers {
		newTarget := target - i
		for j := range numbers {
			if numbers[newTarget-j] {
				fmt.Println(i, j, newTarget-j)
				fmt.Println(i * j * (newTarget - j))
				return
			}
		}
	}

}

func naive(numbers []int) {
	for _, i := range numbers {
		for _, j := range numbers {
			if i+j == target {
				fmt.Println(j * i)
				break
			}
		}
	}

	for _, i := range numbers {
		for _, j := range numbers {
			for _, z := range numbers {
				if i+j+z == target {
					fmt.Println(j * i * z)
					return
				}
			}
		}
	}
}
