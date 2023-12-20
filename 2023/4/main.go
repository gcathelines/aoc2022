package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	scratchs := map[int]int{}
	idx := 1
	for scanner.Scan() {
		txt := scanner.Text()
		scratchs[idx] += 1

		str := strings.Split(txt, ":")
		cards := strings.Split(str[1], "|")
		winning := strings.Split(cards[0], " ")
		win := map[string]bool{}
		for _, w := range winning {
			if w == "" {
				continue
			}
			win[w] = false
		}
		have := strings.Split(cards[1], " ")
		for _, h := range have {
			if _, ok := win[h]; ok {
				win[h] = true
			}
		}

		// part 1
		// pow := -1
		match := 1
		for _, isWin := range win {
			if isWin {
				// part 1
				// pow++
				scratchs[idx+match] += 1 * scratchs[idx]
				match++
			}
		}
		// part 1
		// if pow >= 0 {
		// 	sum += int(math.Pow(2, float64(pow)))
		// }

		idx++
	}

	fmt.Println(scratchs)

	for k, v := range scratchs {
		if k > idx {
			break
		}
		sum += v
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
