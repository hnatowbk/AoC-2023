package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var count int = 0
var lines int = 0

func main() {
	file, err := os.Open("../inputs/day-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = cleanLine(line)
		var first string = ""
		var second string = ""
		for i := 0; i < len(line); i++ {
			fmt.Println(isValidNumber(string(line[i])))
			if isValidNumber(string(line[i])) {
				if first == "" {
					first = string(line[i])
				} else {
					second = string(line[i])
				}
			}
		}
		lines++
		total := first + second
		num, err := strconv.Atoi(total)
		fmt.Println(num)
		if err != nil {
			panic(err)
		}
		count = count + num
	}
	fmt.Println(lines)
	fmt.Println(count)
}

func isValidNumber(category string) bool {
	switch category {
	case
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9":
		return true
	}
	return false
}

func cleanLine(line string) string {
	var cleanedLine string
	cleanedLine = strings.ReplaceAll(line, "one", "1")
	cleanedLine = strings.Replace(line, "two", "2", -1)
	cleanedLine = strings.Replace(line, "three", "3", -1)
	cleanedLine = strings.Replace(line, "four", "4", -1)
	cleanedLine = strings.Replace(line, "five", "5", -1)
	cleanedLine = strings.Replace(line, "six", "6", -1)
	cleanedLine = strings.Replace(line, "seven", "7", -1)
	cleanedLine = strings.Replace(line, "eight", "8", -1)
	cleanedLine = strings.Replace(line, "nine", "9", -1)
	return cleanedLine
}
