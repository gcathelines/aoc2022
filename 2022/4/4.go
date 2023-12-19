package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count1 := 0
	count2 := 0

	for scanner.Scan() {
		txt := scanner.Text()
		var a1, a2, b1, b2 int

		if _, err := fmt.Sscanf(txt, "%d-%d,%d-%d", &a1, &a2, &b1, &b2); err != nil {
			fmt.Println("error", err)
		}

		if a1 <= b1 && a2 >= b2 || a1 >= b1 && a2 <= b2 {
			count1++
		}

		if a1 >= b1 && a1 <= b2 {
			count2++
			continue
		}

		if b1 >= a1 && b1 <= a2 {
			count2++
			continue
		}

		if b2 >= a1 && b2 <= a2 {
			count2++
			continue
		}

		if a2 >= b1 && a2 <= b2 {
			count2++
			continue
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println("Prob 1:", count1, "Prob 2:", count2)

}
