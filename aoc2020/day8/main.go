package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type InstructionType = string

type Instruction struct {
	instr InstructionType
	val   int
}

func main() {
	data, e := ioutil.ReadFile("C:/Users/sam/dev/advent/aoc2020/day8/input.txt")
	if e != nil {
		log.Fatal(e)
	}

	splits := strings.Split(string(data), "\r\n")

	code := make([]Instruction, 0)

	for _, line := range splits {
		if line == "" {
			break
		}

		instruction := parseInstruction(line)
		code = append(code, instruction)
	}

	for i := 0; i < len(code); i++ {
		newCode, err := invertedInstructionCode(code, i)
		if err != nil {
			continue
		}

		accValue, hitLoop := runInstructions(newCode)
		if !hitLoop {
			fmt.Printf("Modified %d with an opposite, got %d \n", i, accValue)
		}
	}
}

func parseInstruction(line string) Instruction {
	split := strings.Split(line, " ")
	instr := split[0]
	i, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatalf("Unable to parse string to int %v", split[1])
	}

	return Instruction{
		instr: instr,
		val:   i,
	}
}

func runInstructions(code []Instruction) (accValue int, hitLoop bool) {
	executionCount := make(map[int]int)

	i := 0
	acc := 0
	for {
		if i >= len(code) {
			break
		}
		instruction := code[i]

		_, wasExecuted := executionCount[i]
		if wasExecuted {
			return acc, true
		} else {
			executionCount[i] = 1
		}

		if instruction.instr == "nop" {
			i++
			continue
		} else if instruction.instr == "acc" {
			i++
			acc += instruction.val
		} else if instruction.instr == "jmp" {
			i += instruction.val
		}
	}

	return acc, false
}

func invertedInstructionCode(code []Instruction, i int) ([]Instruction, error) {
	var opposite string

	if code[i].instr == "nop" {
		opposite = "jmp"
	} else if code[i].instr == "jmp" {
		opposite = "nop"
	} else {
		return nil, errors.New("cannot invert a non-jmp or nop instruction")
	}

	newPrepend := make([]Instruction, i)
	copy(newPrepend, code[0:i])
	newAppend := make([]Instruction, len(code[i+1:]))
	copy(newAppend, code[i+1:])

	newCode := append(newPrepend, Instruction{
		instr: opposite,
		val:   code[i].val,
	})
	newCode = append(newCode, newAppend...)

	return newCode, nil
}
