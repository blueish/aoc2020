package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	file, e := os.Open("C:/Users/sam/dev/advent/aoc2020/day3/input.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var grid [][]bool
	line := make([]bool, 0)

	var err error
	var r rune

	for ; err != io.EOF; r, _, err = reader.ReadRune() {
		if r == '.' {
			line = append(line, false)
		} else if r == '#' {
			line = append(line, true)
		} else if r == '\n' {
			grid = append(grid, line)
			line = make([]bool, 0)
		} else {
			fmt.Println("Unable to process rune:", r)
		}

	}
	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fmt.Println("trees hit right 1 x down 1: ", treesForSlope(grid, 1, 1))
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fmt.Println("trees hit right 3 x down 1: ", treesForSlope(grid, 3, 1))

		wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fmt.Println("trees hit right 5 x down 1: ", treesForSlope(grid, 5, 1))
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fmt.Println("trees hit right 7 x down 1: ", treesForSlope(grid, 7, 1))
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fmt.Println("trees hit right 1 x down 2: ", treesForSlope(grid, 1, 2))
		wg.Done()
	}(&wg)

	wg.Wait()
}

func treesForSlope(grid [][]bool, right int, down int) int {
	horizIdx := 0
	vertIdx := 0

	trees := 0
	for {
		if grid[vertIdx][horizIdx] {
			trees++
		}

		horizIdx += right
		horizIdx %= len(grid[0])
		vertIdx += down

		if vertIdx >= len(grid) {
			return trees
		}
	}

}
