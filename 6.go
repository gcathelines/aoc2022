package main

import (
	"bufio"
	"fmt"
	"os"
)

// 4 / 14
const rnge = 14

func main() {
	file, err := os.Open("./6-test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	start := 0
	for scanner.Scan() {
		txt := scanner.Text()
		subArr := txt[start : rnge+start]
		m := map[string]int{}
		for _, s := range subArr {
			m[string(s)]++
		}

		for start+rnge <= len(txt) {
			found := true
			for _, v := range m {
				if v > 1 {
					found = false
				}
			}

			if found {
				break
			}

			start++

			m[string(txt[start-1])]--
			m[string(txt[rnge+start-1])]++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(start + rnge)
}
