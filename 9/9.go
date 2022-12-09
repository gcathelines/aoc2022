package main

import (
	"bufio"
	"fmt"
	"os"
)

var mp = map[string]int{
	"0:0": 1,
}

func key(i, j int) string {
	return fmt.Sprintf("%d:%d", i, j)
}

var max = -1

var dragon = [][]int{}

func main() {
	file, err := os.Open("./9.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	head := []int{0, 0}
	tail := []int{0, 0}

	for scanner.Scan() {
		txt := scanner.Text()

		var dir string
		var count int

		if _, err := fmt.Sscanf(txt, "%s %d", &dir, &count); err != nil {
			fmt.Println("error", err)
		}

		head, tail = move(dir, count, head, tail)
	}

	fmt.Println("No 1:", len(mp))

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func move(dir string, count int, head []int, tail []int) ([]int, []int) {
	prevHead := make([]int, 2)
	copy(prevHead, head)
	for count > 0 {
		switch dir {
		case "U":
			head[1] += 1
		case "D":
			head[1] -= 1
		case "L":
			head[0] -= 1
		case "R":
			head[0] += 1
		}

		if !isTouch(head, tail) {
			copy(tail, prevHead)
			mp[key(tail[0], tail[1])]++
		}
		copy(prevHead, head)
		count--
	}

	return head, tail
}

func isTouch(head []int, tail []int) bool {
	x := head[0] - tail[0]
	y := head[1] - tail[1]
	if x > 1 || x < -1 {
		return false
	}
	if y > 1 || y < -1 {
		return false
	}
	return true
}
