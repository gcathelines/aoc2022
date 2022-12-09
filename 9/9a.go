package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var mp = map[string]int{}

func key(i, j int) string {
	return fmt.Sprintf("%d:%d", i, j)
}

const max = 35

var dragon = [][]int{
	{10, 10},
	{10, 10},
	{10, 10},
	{10, 10},
	{10, 10},
	{10, 10},
	{10, 10},
	{10, 10},
	{10, 10},
	{10, 10},
}

func main() {
	file, err := os.Open("./9.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		txt := scanner.Text()
		var dir string
		var count int

		if _, err := fmt.Sscanf(txt, "%s %d", &dir, &count); err != nil {
			fmt.Println("error", err)
		}

		move(dir, count, dragon[0])
	}

	fmt.Println("No 2:", len(mp))

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func move(dir string, count int, head []int) {
	for count > 0 {
		switch dir {
		case "U":
			head[0] -= 1
		case "D":
			head[0] += 1
		case "L":
			head[1] -= 1
		case "R":
			head[1] += 1
		}

		for idx := range dragon {
			if idx == len(dragon)-1 {
				mp[key(dragon[idx][0], dragon[idx][1])]++
				continue
			}

			x, y := delta(dragon[idx], dragon[idx+1])
			if x < -1 || x > 1 || y < -1 || y > 1 {
				a := float64(dragon[idx][0]+dragon[idx+1][0]) / float64(2)
				b := float64(dragon[idx][1]+dragon[idx+1][1]) / float64(2)
				if a == 0 {
					dragon[idx+1][0] = 0
				}
				if x > 0 {
					dragon[idx+1][0] = int(math.Ceil(a))
				}
				if x < 0 {
					dragon[idx+1][0] = int(math.Floor(a))
				}

				if b == 0 {
					dragon[idx+1][1] = 0
				}
				if y > 0 {
					dragon[idx+1][1] = int(math.Ceil(b))
				}
				if y < 0 {
					dragon[idx+1][1] = int(math.Floor(b))
				}
			}
		}

		// simulation
		// for i := 0; i < max; i++ {
		// 	for j := 0; j < max; j++ {
		// 		var found bool
		// 		for _, v := range dragon {
		// 			if v[0]+10 == i && v[1]+10 == j {
		// 				found = true
		// 				break
		// 			}
		// 		}
		// 		if found {
		// 			fmt.Print("X")
		// 		} else {
		// 			if i-10 == 0 && j-10 == 0 {
		// 				fmt.Print("s")
		// 			} else {
		// 				fmt.Print(".")
		// 			}
		// 		}
		// 	}
		// 	fmt.Println()
		// }
		count--
	}
}

func delta(head []int, tail []int) (int, int) {
	x := head[0] - tail[0]
	y := head[1] - tail[1]
	return x, y
}
