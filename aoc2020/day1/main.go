package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("C:/Users/sam/dev/advent/aoc2020/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nums := make([]int, 0)

	scnr := bufio.NewScanner(file)
	for scnr.Scan() {
		i, err := strconv.Atoi(scnr.Text())
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}

	lowIdx := 0
	midIdx := len(nums) - 2
	highIdx := len(nums) - 1

	if false {
		// we can more easily find endpoints by sorting
		sort.Ints(nums)
		for {
			high := nums[highIdx]
			low := nums[lowIdx]

			comb := high + low

			if comb == 2020 {
				fmt.Println("res:", high*low)
				break
			} else if comb > 2020 {
				highIdx--
				lowIdx = 0
			} else {
				lowIdx++
			}
		}
	}

	// pt 2:
	if true {
		for {
			if midIdx <= lowIdx {
				midIdx--
				lowIdx = 0
				continue
			} else if midIdx == 1 {
				highIdx--
				midIdx = highIdx - 1
				lowIdx = 0
			}
			high := nums[highIdx]
			mid := nums[midIdx]
			low := nums[lowIdx]

			comb := high + mid + low

			if comb == 2020 {
				fmt.Println("res:", high*mid*low)
				break
			} else {
				lowIdx++
			}
		}
	}

}
