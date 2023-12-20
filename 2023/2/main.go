package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	red   = 12
	green = 13
	blue  = 14
)

func main() {
	file, err := os.Open("./2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	num := 0

	for scanner.Scan() {
		txt := scanner.Text()

		var (
			id    int
			balls map[string]int = make(map[string]int)
			color string
		)

		data := strings.Split(txt, ":")
		if _, err := fmt.Sscanf(data[0], "Game %d", &id); err != nil {
			fmt.Println("error a", err)
		}

		data = strings.Split(data[1], ";")

		for _, d := range data {
			games := strings.Split(d, ",")
			for _, g := range games {
				size := 0
				if _, err := fmt.Sscanf(g, "%d %s", &size, &color); err != nil {
					fmt.Println("error b", err)
				}
				if balls[color] < size {
					balls[color] = size
				}
			}
		}

		// part 1
		// if balls["red"] > red {
		// 	fmt.Println("id failed red", id, balls)
		// 	continue
		// }
		// if balls["blue"] > blue {
		// 	fmt.Println("id failed blue", id, balls)
		// 	continue
		// }
		// if balls["green"] > green {
		// 	fmt.Println("id failed green", id, balls)
		// 	continue
		// }
		// fmt.Println("id pass", id, balls)

		// part 2
		num += balls["red"] * balls["green"] * balls["blue"]
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
