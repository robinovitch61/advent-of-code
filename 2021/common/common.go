package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadFile(day string) []string {
	var contents []string

	file, err := os.Open(fmt.Sprintf("./day%s/%s_input.txt", day, day))
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

func PrintDay(day int) {
	fmt.Println("---------")
	fmt.Println("Day " + strconv.Itoa(day))
}
