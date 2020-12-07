package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const ShinyGold = "shiny gold"

type BagContents struct {
	color string
	number int
}

type Bag struct {
	color string
	others []BagContents
}

func extractBagColor(bagDesc string, stripNum bool) string {
	// strip out the number and replace
	//color := strings.ReplaceAll(bagDesc[2:], "", "")
	if (stripNum) {
		bagDesc = bagDesc[2:]
	}
	color := strings.ReplaceAll(bagDesc, " bags", "") // strip out number, bags
	color = strings.ReplaceAll(color, " bag", "") // singular bag
	return color
}


func main() {
	data, e := ioutil.ReadFile("C:/Users/sam/dev/advent/aoc2020/day7/input.txt")
	if e != nil {
		log.Fatal(e)
	}

	splits := strings.Split(string(data), "\r\n")

	bags := make(map[string]Bag, 0)

	for _, line := range splits {
		if line == "" {
			break
		}
		line = strings.ReplaceAll(line, ".", "")
		bagSpec := strings.Split(line, " contain ")
		bagColor := extractBagColor(bagSpec[0], false)

		contents := strings.Split(bagSpec[1], ", ")
		bagContain := make([]BagContents, 0)

		for _, contents := range contents {
			if contents == "no other bags" {
				break
			}

			i, err := strconv.Atoi(contents[0:1])
			if err != nil {
				log.Fatalf("Could not convert %v, to string", contents[0:1])
			}

			color := extractBagColor(contents, true)

			bagContain = append(bagContain, BagContents{
				color:  color,
				number: i,
			})
		}

		bag := Bag{
			color:  bagColor,
			others: bagContain,
		}

		bags[bag.color] = bag
		//bags = append(bags, bag)
	}

	fmt.Println("bags", bags)

	bagColorsThatCanContainGold := make([]string, 0)
	bagColorsThatCanContainGold = append(bagColorsThatCanContainGold, ShinyGold)

	i := 0

	for _, bag := range bags {
		// look at the bag, figure out if it can contain shiny gold
		containsGold := false

		for _, other := range bag.others {
			if canContainGold(other.color, bags) {
				containsGold = true
				break
			}
		}

		if containsGold {
			i++
		}
	}

	fmt.Println("num that can contain", i)

	// part two - how many bags are needed for my shiny gold bag?

	fmt.Println("Num needed for: ", bagsNeededFor(ShinyGold, bags) - 1)
}

func bagsNeededFor(color string, bags map[string]Bag) int {
	numNeeded := 1

	for _, content := range bags[color].others {
		// check if any of its contents can
		contentColor := content.color
		numBags := content.number
		eachNeed := bagsNeededFor(contentColor, bags)

		numNeeded += numBags * eachNeed
	}

	return numNeeded
}

func canContainGold(color string, bags map[string]Bag) bool {
	if color == ShinyGold {
		return true
	}

	for _, content := range bags[color].others {
		// check if any of its contents can
		contentColor := content.color
		if canContainGold(contentColor, bags) {
			return true
		}
	}

	return false
}
