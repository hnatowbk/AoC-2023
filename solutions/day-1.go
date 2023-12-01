package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var count int = 0

func main() {
	file, err := os.Open("../inputs/day-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var first string = ""
		var second string = ""
		for i := 0; i < len(line); i++ {
			if isValidNumber(string(line[i])) {
				if first == "" {
					first = string(line[i])
					fmt.Println(first)
				} else {
					second = string(line[i])
					fmt.Println(second)
				}
			}
			total := first + second
			num, err := strconv.Atoi(total)
			if err != nil {
				panic(err)
			}
			count = count + num
		}

	}
	fmt.Println(count)
}

func isValidNumber(category string) bool {
	switch category {
	case
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return true
	}
	return false
}
