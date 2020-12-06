package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	data, e := ioutil.ReadFile("C:/Users/sam/dev/advent/aoc2020/day6/input.txt")
	if e != nil {
		log.Fatal(e)
	}

	splits := strings.Split(string(data), "\r\n")

	curr := ""

	numQuestions := 0

	for _, split := range splits {
		if split == "" {
			curr += split

			nonDupes := make(map[rune]bool)
			for _, let := range curr {
				_, prs := nonDupes[let]
				if !prs {
					nonDupes[let] = true
				}
			}

			currLen := len(nonDupes)
			numQuestions += currLen
			curr = ""
		} else {

			curr += split
		}
	}

	fmt.Println("Part 1:", numQuestions)

	currentGroup := make([]string, 0)

	allYeses := 0

	for _, split := range splits {
		if split == "" {

			howManyPeople := len(currentGroup)

			allOccurrences := make(map[rune]int)

			for _, personAnswers := range currentGroup {
				for _, let := range personAnswers {
					_, prs := allOccurrences[let]
					if prs {
						allOccurrences[let] += 1
					} else {
						allOccurrences[let] = 1
					}
				}
			}

			questionsEveryoneYesed := 0
			for _, value := range allOccurrences {
				if howManyPeople == value {
					questionsEveryoneYesed++
				}
			}

			allYeses += questionsEveryoneYesed

			currentGroup = make([]string, 0)
		} else {

			currentGroup = append(currentGroup, split)
		}
	}

	fmt.Println("Part 2 answer:", allYeses)
}
