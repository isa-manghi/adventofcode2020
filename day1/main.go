package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

//function to go through each input line for each line that searchnum= 2020 - inputnumber
//call function on each line y that checks if y number == searchnum

func main() {
	num1, num2 := numberloop()
	fmt.Println("num1: %x, num2: %x", num1, num2)
	mult := num1 * num2
	fmt.Println("multiplied = %x", mult)
}

//func fileloop() {
//for line x in file:
//2020 - x
//numberloop(x)
//}

// readLines reads a whole file into memory
// and returns a slice of its lines.
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

func numberloop() (int, int) {
	lines, err := readLines("/Users/isabellamanghi/advent-of-code/advent-of-code/2020/Paul+IsabellaSub/day1/numbers.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for _, line := range lines {
		intnum, _ := strconv.Atoi(line)
		searchnum := 2020 - intnum
		for _, innerline := range lines {
			innernum, _ := strconv.Atoi(innerline)
			if searchnum == innernum {
				return searchnum, intnum
			}
		}
	}
	return 1, 2
}
