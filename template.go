package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./1-test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		txt := scanner.Text()

		if _, err := fmt.Sscanf(txt, "%d", &a); err != nil {
			fmt.Println("error", err)
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
