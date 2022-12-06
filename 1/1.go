package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var mx = -1
var curr = 0

var top3 []int

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
			app(curr)
			mx = max(mx, curr)
			curr = 0
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	x := 0
	for _, top := range top3 {
		x += top
	}

	fmt.Println(mx, x)
}

func app(x int) {
	if len(top3) < 3 {
		top3 = append(top3, x)
		return
	}

	z := -1

	for idx, cal := range top3 {
		if x > cal {
			z = idx
			break
		}
	}

	if z != -1 {
		for i := 2; i >= z; i-- {
			if i == z {
				top3[i] = x
				break
			}
			top3[i] = top3[i-1]
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
