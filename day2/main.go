package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// readLines reads a whole file into memory
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Split splits the file into a readable and workable file.
func Split(r rune) bool {
	return r == '-' || r == ' '
}

func main() {
	lines, err := readLines("/Users/isabellamanghi/advent-of-code/advent-of-code/2020/paul&isabella/day2/passwords.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	//input := "8-9 x: xxxxxxxxrk"
	var totalvalid int
	for _, line := range lines {
		a := strings.FieldsFunc(line, Split)
		minint := a[0]
		maxint := a[1]
		searchletter := a[2]
		password := a[3]
		searchletter = strings.Trim(searchletter, ":")
		// fmt.Println(minint, maxint, searchletter, password)
		minint2, _ := strconv.Atoi(minint)
		maxint2, _ := strconv.Atoi(maxint)
		occurences := strings.Count(password, searchletter)
		if occurences >= minint2 && occurences <= maxint2 {
			totalvalid++
		}
	}
	fmt.Println(totalvalid)
}
