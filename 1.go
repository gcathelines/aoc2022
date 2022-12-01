package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var mx = -1
var curr = 0

func main() {
	file, err := os.Open("./1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			i, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			curr += i
		} else {
			mx = max(mx, curr)
			curr = 0
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(mx)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
