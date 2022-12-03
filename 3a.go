package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./3.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	prioSum := 0

	count := 0
	x := map[int]int{}
	for scanner.Scan() {
		txt := scanner.Text()

		for _, char := range txt {
			if x[priority(char)] == count {
				x[priority(char)]++
			}
		}
		count++

		if count == 3 {
			for prio, count := range x {
				if count == 3 {
					prioSum += prio
				}
			}

			count = 0
			x = map[int]int{}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(prioSum)

}

/*
Lowercase item types a through z have priorities 1 through 26.
Uppercase item types A through Z have priorities 27 through 52.
*/
func priority(x rune) int {
	if x < 97 {
		return int(x) - 38
	}

	return int(x) - 96
}
