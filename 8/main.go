package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	cmd      string
	register int
}

type instructions []instruction

func main() {
	d, _ := ioutil.ReadFile("data")
	dd := strings.Split(string(d), "\n")
	var instr instructions
	for _, ddd := range dd {
		dddd := strings.Split(ddd, " ")
		register, _ := strconv.Atoi(dddd[1])
		instr = append(instr, instruction{
			register: register,
			cmd:      dddd[0],
		})
	}
	firstCount, _, indexesExecuted := instr.FinalRegister()
	fmt.Println(firstCount)
	// we only need to try changin nop and jmp for the cmd's
	// that are executed before entering the infinite loop
	for executed := range indexesExecuted {
		switch instr[executed].cmd {
		case "nop":
			instr[executed].cmd = "jmp"
			counter, finished, _ := instr.FinalRegister()
			if finished {
				fmt.Println(counter)
				return
			}
			instr[executed].cmd = "nop"
		case "jmp":
			instr[executed].cmd = "nop"
			counter, finished, _ := instr.FinalRegister()
			if finished {
				fmt.Println(counter)
				return
			}
			instr[executed].cmd = "jmp"
		}
	}

}

// FinalRegister returns the final register value, whether it has exited successfully
// or entered an infinite loop and the instruction indexes that have been executed
func (instr instructions) FinalRegister() (int, bool, map[int]bool) {
	indexesExecuted := make(map[int]bool)
	var registerTotal, i int
	for {
		if i == len(instr) {
			return registerTotal, true, indexesExecuted
		}
		if indexesExecuted[i] {
			return registerTotal, false, indexesExecuted
		}
		indexesExecuted[i] = true
		switch instr[i].cmd {
		case "jmp":
			i += instr[i].register
		case "acc":
			registerTotal += instr[i].register
			i++
		case "nop":
			i++
		}
	}
}
