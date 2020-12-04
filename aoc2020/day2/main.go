package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("C:/Users/sam/dev/advent/aoc2020/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numValid := 0

	scnr := bufio.NewScanner(file)
	for scnr.Scan() {
		split := strings.SplitN(scnr.Text(), " ", 3)
		condition := split[0]
		letter := split[1][0:1] // it's parsed as "a:", so remove the ":" to get just the letter
		password := split[2]

		splitCond := strings.SplitN(condition, "-", 2)

		//minLetCount, err := strconv.Atoi(splitCond[0])
		//if err != nil {
		//	panic(err)
		//}
		//
		//maxLetCount, err :=  strconv.Atoi(splitCond[1])
		//if err != nil {
		//	panic(err)
		//}

		//numOccurrences := strings.Count(password, letter)
		//
		//if minLetCount <= numOccurrences && numOccurrences <= maxLetCount {
		//	numValid++
		//}

		// pt 2
		firstPlace, err := strconv.Atoi(splitCond[0])
		if err != nil {
			panic(err)
		}

		secondPlace, err := strconv.Atoi(splitCond[1])
		if err != nil {
			panic(err)
		}

		// subtract one since the "first place" is index zero
		firstValid := password[firstPlace-1:firstPlace] == letter
		secondValid := password[secondPlace-1:secondPlace] == letter


		if (firstValid || secondValid) && !(firstValid && secondValid) {
			numValid++
		}
	}

	fmt.Println("num valid", numValid)
}
