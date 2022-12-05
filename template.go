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

	for scanner.Scan() {
		txt := scanner.Text()

		if _, err := fmt.Sscanf(txt, "%d-%d,%d-%d", &a1, &a2, &b1, &b2); err != nil {
			fmt.Println("error", err)
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
