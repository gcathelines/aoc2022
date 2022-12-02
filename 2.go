package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	A for Rock, B for Paper, and C for Scissors
	X for Rock, Y for Paper, and Z for Scissors
	The winner of the whole tournament is the player with the highest score.
	Your total score is the sum of your scores for each round.
	The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
	plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).
*/

var score = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func rps(opp string, you string) int {
	switch opp {
	case "A":
		switch you {
		case "X":
			return 3
		case "Y":
			return 6
		case "Z":
			return 0
		}
	case "B":
		switch you {
		case "X":
			return 0
		case "Y":
			return 3
		case "Z":
			return 6
		}
	case "C":
		switch you {
		case "X":
			return 6
		case "Y":
			return 0
		case "Z":
			return 3
		}
	}
	return -1
}

func main() {
	file, err := os.Open("./2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()

		res := rps(string(str[0]), string(str[2]))
		sum += res + score[string(str[2])]
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
