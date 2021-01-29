package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	// "strconv"
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
		index1 := a[0]
		index2 := a[1]
		searchletter := a[2]
		passwordraw := a[3]
		searchletter = strings.Trim(searchletter, ":")
		// fmt.Println(minint, maxint, searchletter, password)
		index11, _ := strconv.Atoi(index1)
		index22, _ := strconv.Atoi(index2)
		password := []rune(passwordraw)
		// fmt.Println(string(password[index11]))
		// fmt.Println(string(password[index22]))
		if string(password[index11 -1]) == searchletter && string(password[index22 -1]) == searchletter {
			continue
		}
		if string(password[index11 -1]) != searchletter && string(password[index22 -1]) != searchletter {
			continue
		}
		if string(password[index11 -1]) == searchletter || string(password[index22 -1]) == searchletter {
			totalvalid++
		}
	}
	fmt.Println(totalvalid)
}
