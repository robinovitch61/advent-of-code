package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) []string {
	var contents []string

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return contents
}

func p1(directions []string) int {
	var horizontal, depth int = 0, 0
	for _, direction := range directions {
		split := strings.Split(direction, " ")
		order := split[0]
		value, _ := strconv.Atoi(split[1])

		switch order {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}
	return horizontal * depth
}

func p2(directions []string) int {
	var horizontal, depth, aim int = 0, 0, 0
	for _, direction := range directions {
		split := strings.Split(direction, " ")
		order := split[0]
		value, _ := strconv.Atoi(split[1])

		switch order {
		case "forward":
			horizontal += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}
	return horizontal * depth
}

func main() {
	directions := ReadFile("./02_input.txt")
	fmt.Println(p1(directions))
	fmt.Println(p2(directions))
}
