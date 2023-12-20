package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("./1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	num := 0

	x := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for scanner.Scan() {
		txt := scanner.Text()
		var (
			line  string
			first int
			last  int
		)
		if _, err := fmt.Sscanf(txt, "%s", &line); err != nil {
			fmt.Println("error", err)
		}

		first = firstInt(line, x)
		last = lastInt(line, x)

		num += first*10 + last
		fmt.Println(line, first, last)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(num)
}

func firstInt(str string, m map[string]int) int {
	var (
		first int
		idx   int
		err   error
	)
	for ii, char := range str {
		if first, err = strconv.Atoi(string(char)); err == nil {
			idx = ii
			break
		}
	}

	// part 2
	for k, v := range m {
		ii := strings.Index(str, k)
		if ii == -1 {
			continue
		}
		fmt.Println("first", ii, k, v)
		if ii < idx {
			idx = ii
			first = v
		}
	}

	return first
}

func lastInt(str string, m map[string]int) int {
	var (
		last int
		idx  int
		err  error
	)
	for i := len(str) - 1; i >= 0; i-- {
		if last, err = strconv.Atoi(string(str[i])); err == nil {
			idx = i
			break
		}
	}

	// part 2
	for k, v := range m {
		ii := strings.LastIndex(str, k)
		if ii == -1 {
			continue
		}
		fmt.Println("last", ii, k, v)
		if ii > idx {
			idx = ii
			last = v
		}
	}

	return last
}
