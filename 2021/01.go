package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(path string) []string {
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

func p1(depths []int) int {
	numIncreases := 0
	for i, _ := range depths {
		if i > 0 && depths[i] > depths[i-1] {
			numIncreases += 1
		}
	}
	return numIncreases
}

func sum(input []int) int {
	sum := 0
	for i := range input {
		sum += input[i]
	}
	return sum
}

func p2(depths []int, windowSize int) int {
	numIncreases := 0
	for i, _ := range depths {
		if i >= windowSize && sum(depths[i-windowSize:i]) < sum(depths[i-windowSize+1:i+1]) {
			numIncreases += 1
		}
	}
	return numIncreases
}

func main() {
	depthStrings := readFile("./01_input.txt")
	//fmt.Println(cap(depthStrings)) // why is this 2560?
	var depths = make([]int, len(depthStrings), len(depthStrings))
	//fmt.Println(cap(depths)) // this is now 2000 as expected
	for i, depth := range depthStrings {
		intDepth, _ := strconv.Atoi(depth)
		depths[i] = intDepth
	}
	//fmt.Println(cap(depths)) // still 2000
	fmt.Println(p1(depths))
	fmt.Println(p2(depths, 3))
}
