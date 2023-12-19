package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var mp = map[int]int{}

func main() {
	file, err := os.Open("./10-test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cycle := 0
	curr := 1
	signals := []int{}

	for scanner.Scan() {
		txt := scanner.Text()
		if cycle%40 == 0 {
			fmt.Print("\n")
		}

		if strings.HasPrefix(txt, "noop") {
			curr += 0
		} else {
			var signal int
			if _, err := fmt.Sscanf(txt, "addx %d", &signal); err != nil {
				fmt.Println("error", err)
			}
			mp[cycle+1] = signal
		}

		if curr == cycle%40 || curr == (cycle%40)-1 || curr == (cycle%40)+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		cycle++

		if (cycle-20)%40 == 0 {
			signals = append(signals, curr*cycle)
		}

		for len(mp) > 0 {
			if cycle%40 == 0 {
				fmt.Print("\n")
			}
			if curr == cycle%40 || curr == (cycle%40)-1 || curr == (cycle%40)+1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			curr += mp[cycle]
			delete(mp, cycle)
			cycle++

			if (cycle-20)%40 == 0 {
				signals = append(signals, curr*cycle)
			}
		}

	}

	// fmt.Println(signals)
	val := 0
	for _, v := range signals {
		val += v
	}

	// fmt.Println(val)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
