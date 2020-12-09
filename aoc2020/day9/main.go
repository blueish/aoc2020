package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)


func main() {
	data, e := ioutil.ReadFile("C:/Users/sam/dev/advent/aoc2020/day9/input.txt")
	if e != nil {
		log.Fatal(e)
	}

	splits := strings.Split(string(data), "\r\n")

	nums := make([]int, 0)

	for _, line := range splits {
		if line == "" {
			break
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, i)
	}


	PREAMBLE := 25

	var secretNum int

	for i := PREAMBLE; i < len(nums); i++ {
		window := nums[i-PREAMBLE:i]
		num := nums[i]

		if !canMatch(window, num) {
			secretNum = nums[i]
			break
		}
	}

	fmt.Println("Part one: ", secretNum)

	if secretNum  == 0 {
		log.Fatal("Could not find a secret number, input is invalid")
	}


	cont := findContiguous(nums, secretNum)

	sort.Ints(cont)

	fmt.Println("Part two: ", cont[0] + cont[len(cont) - 1])
}

func canMatch(ints []int, needle int) bool {
	for i := 0; i < len(ints); i++ {
		for j := i +1; j < len(ints); j++ {
			if ints[i] + ints[j] == needle {
				return true
			}
		}
	}

	return false
}

func findContiguous(nums []int, secretNum int) []int {
	cont := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {

			cont = append(cont, nums[j])

			sum := sumAll(cont)

			if sum == secretNum {
				return cont
			}

			if sum > secretNum {
				cont = make([]int, 0)
				break
			}
		}
	}

	return cont
}

func sumAll(ints []int ) int {
	sum := 0
	for _, asdf := range ints {
		sum += asdf
	}

	return sum
}
