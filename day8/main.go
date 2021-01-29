package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	begin()
	fmt.Println(accumulator)
}

var lines []string
var accumulator int
var indexesVisited []int //1, 2, 3
var index int

func begin() {
	var err error
	lines, err = readLines("/Users/isabellamanghi/adventofcode/advent-of-code/2020/paul&isabella/day8/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	var i int
	for true {
		//while line index we're currently on is NOT in indexes
		//index = parseInstruction(index)
		for _, b := range indexesVisited {
			if index == b {
				return
			}
		}
		//add index to indexesVisited
		indexesVisited = append(indexesVisited, index)
		index = parseInstruction(index)
		//s := fmt.Sprintf("index = %x, accumulator = %x\n", index, accumulator)
		//io.WriteString(os.Stdout, s)
		i++
		fmt.Println("------------")
	}
}

func parseInstruction(index int) int {
	str := lines[index]
	fmt.Println(str)
	var num int
	searchletter := string(str[0])
	var sb strings.Builder

	sign := string(str[4])

	length := len(lines[index])
	fmt.Println(length)

	if length == 6 {
		sb.WriteString(string(str[5]))
		num, _ = strconv.Atoi(sb.String())
		fmt.Println(sb.String())
		fmt.Println(num)
	} else if length == 7 {
		//numstring := str[5:6]
		sb.WriteString(string(str[5]))
		sb.WriteString(string(str[6]))
		num, _ = strconv.Atoi(sb.String())
		fmt.Println(sb.String())
		fmt.Println(num)
	} else if length == 8 {
		//numstring := str[5:7]
		sb.WriteString(string(str[5]))
		sb.WriteString(string(str[6]))
		sb.WriteString(string(str[7]))
		num, _ = strconv.Atoi(sb.String())
		fmt.Println(sb.String())
		fmt.Println(num)
	}

	//fmt.Println(num)
	//fmt.Println(searchletter)

	switch searchletter {
	case "a":
		if sign == "+" {
			accumulator += num
		} else {
			accumulator -= num
		}
		index++
	case "j":
		if sign == "+" {
			index += num
		} else {
			index -= num
		}
	case "n":
		index++
	default:
		fmt.Println("messed up")
	}
	return index
}
