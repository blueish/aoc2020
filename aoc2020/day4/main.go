package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var REQUIRED_FIELDS []string = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func isWithin(num, low, high int) bool {
	return low <= num && num <= high
}

var FIELD_VALIDATORS map[string]func(string) bool = map[string]func(string) bool{
	"byr": func(s string) bool {
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return isWithin(i, 1920, 2002)
	},
	"iyr": func(s string) bool {
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return isWithin(i, 2010, 2020)
	},
	"eyr": func(s string) bool {
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return isWithin(i, 2020, 2030)
	},
	"hgt": func(s string) bool {
		measurement := s[len(s)-2:]
		heightStr := s[:len(s)-2]
		height, err := strconv.Atoi(heightStr)
		if err != nil {
			return false
		}

		if measurement == "cm" {
			return isWithin(height, 150, 193)
		} else if measurement == "in" {
			return isWithin(height, 59, 76)
		} else {
			fmt.Println("Err, measurement invalid: ", measurement)
			return false
		}
	},
	"hcl": func(s string) bool {
		matched, _ := regexp.MatchString(`^#[\d|abcdef]{6}$`, s)
		return matched
	},
	"ecl": func(s string) bool {
		matched, _ := regexp.MatchString(`^amb|blu|brn|gry|grn|hzl|oth$`, s)
		return matched
	},
	"pid": func(s string) bool {
		matched, _ := regexp.MatchString(`^[0-9]{9}$`, s)
		return matched
	},
}

func main() {

	file, e := os.Open("C:/Users/sam/dev/advent/aoc2020/day4/input.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	validPassports := 0

	passport := make(map[string]string, 0)

	var line string

	for scanner.Scan() {
		line = scanner.Text()
		fmt.Println("line", line)
		if line == "" {
			if passportHasValidatedData(passport) {
				validPassports++
			}
			passport = make(map[string]string, 0)
		} else {
			fields := strings.Split(line, " ")
			for _, field := range fields {
				v := strings.Split(field, ":")
				passport[v[0]] = v[1]
			}
		}
	}

	fmt.Println("valid just fields", validPassports)
}

func passportIsValid(passport map[string]string) bool {
	for _, field := range REQUIRED_FIELDS {
		if _, prs := passport[field]; !prs {
			return false
		}
	}

	return true
}

func passportHasValidatedData(passport map[string]string) bool {
	for field, validator := range FIELD_VALIDATORS {
		fieldValue, prs := passport[field]
		if !prs || !validator(fieldValue) {
			return false
		}
	}

	return true
}
