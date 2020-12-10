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
	data, e := ioutil.ReadFile("C:/Users/sam/dev/advent/aoc2020/day10/input.txt")
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

	sort.Ints(nums)

	// add an extra + 3 for the last one
	nums = append(nums, nums[len(nums)-1] + 3)

	diffMap := make(map[int]int)
	diffMap[1] = 1

	for i := 1; i < len(nums); i++ {
		prev := nums[i-1]
		num := nums[i]

		difference := num - prev

		diffMap[difference]++
	}

	fmt.Println("Part 1: ", diffMap[1]* diffMap[3])

	dpTable := make(map[int]int)
	dpTable[0] = 1

	for _, num := range nums {
		dpTable[num] = dpTable[num - 1] + dpTable[num - 2] + dpTable[num - 3]
	}

	fmt.Println("Part 2: ", dpTable[nums[len(nums) - 1]])
}

