package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	data, e := ioutil.ReadFile("C:/Users/sam/dev/advent/aoc2020/day5/input.txt")
	if e != nil {
		log.Fatal(e)
	}

	splits := strings.Split(string(data), "\r\n")

	highestId := 0

	allSeats := make(map[int]bool, 0)
	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			allSeats[i * 8 + j] = false
		}
	}

	for _, split := range splits {
		if split == "" {
			continue
		}

		seatRowDeterminer := split[0:7]
		seatColumnDeterminer := split[7:]

		rowSteps := make([]bool, 0)
		for _, let := range seatRowDeterminer {
			if string(let) == "F" {
				rowSteps = append(rowSteps, false)
			} else {
				rowSteps = append(rowSteps, true)
			}
		}

		seatRow := binarySearch(0, 127, rowSteps)

		columnSteps := make([]bool, 0)
		for _, l := range seatColumnDeterminer {
			if string(l) == "L" {
				columnSteps = append(columnSteps, false)
			} else {
				columnSteps = append(columnSteps, true)
			}
		}
		seatCol := binarySearch(0, 7, columnSteps)

		seatId := getSeatId(seatRow, seatCol)

		allSeats[seatId] = true

		if seatId > highestId {
			highestId = seatId
		}
	}

	// Part 1
	fmt.Println("Highest Seat Id: ", highestId)


	// Part 2
	seatsTaken := make([]int, 0)
	for seatId, taken := range allSeats {
		if !taken {
			seatsTaken = append(seatsTaken, seatId)
		}
	}

	sort.Ints(seatsTaken)

	// find the first non-sequential number
	for i := range seatsTaken {
		if i != 0 {
			if seatsTaken[i] != seatsTaken[i - 1] + 1 {
				fmt.Println("My seat id: ", seatsTaken[i])
				break
			}
		}
	}

}

func binarySearch(lowRange int, highRange int, lowerHigherArr []bool) int {
	if len(lowerHigherArr) == 0 {
		return lowRange
	}

	difference := ((highRange - lowRange) / 2) + 1
	higher := lowerHigherArr[0]

	if higher { // higher
		return binarySearch(lowRange + difference, highRange, lowerHigherArr[1:])
	} else {
		return binarySearch(lowRange, highRange - difference, lowerHigherArr[1:])
	}
}

func getSeatId(row, column int) int {
	return row * 8 + column
}
