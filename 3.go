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

	for scanner.Scan() {
		x := map[int]int{}
		txt := scanner.Text()
		firstStr := txt[:len(txt)/2]
		secondStr := txt[len(txt)/2:]

		for _, char := range firstStr {
			if x[priority(char)] == 0 {
				x[priority(char)]++
			}
		}

		for _, char := range secondStr {
			if x[priority(char)] == 1 {
				x[priority(char)]++
			}
		}

		for prio, count := range x {
			if count > 1 {
				prioSum += prio
			}
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
