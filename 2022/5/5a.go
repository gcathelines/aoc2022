package main

import (
	"bufio"
	"fmt"
	"os"
)

var queue1 []string
var queue2 []string
var queue3 []string

/*
[T]     [D]         [L]
[R]     [S] [G]     [P]         [H]
[G]     [H] [W]     [R] [L]     [P]
[W]     [G] [F] [H] [S] [M]     [L]
[Q]     [V] [B] [J] [H] [N] [R] [N]
[M] [R] [R] [P] [M] [T] [H] [Q] [C]
[F] [F] [Z] [H] [S] [Z] [T] [D] [S]
[P] [H] [P] [Q] [P] [M] [P] [F] [D]
 1   2   3   4   5   6   7   8   9
*/

var mp = map[int][]string{}

func main() {
	file, err := os.Open("./5.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var count, from, to int

	for scanner.Scan() {
		txt := scanner.Text()

		if len(txt) > 1 {
			if string(txt[1]) == "1" {
				continue
			}

			if string(txt[0:4]) != "move" {

				for idx := range txt {
					if (idx-1)%4 == 0 && string(txt[idx]) != " " {
						mp[(idx-1)/4] = append(mp[(idx-1)/4], string(txt[idx]))
					}
				}
			}

			if string(txt[0:4]) == "move" {

				if _, err := fmt.Sscanf(txt, "move %d from %d to %d", &count, &from, &to); err != nil {
					fmt.Println("error", err)
				}
				src := mp[from-1][0:count]
				dst := make([]string, len(src))
				copy(dst, src)

				mp[from-1] = mp[from-1][count:]
				mp[to-1] = append(dst, mp[to-1]...)

			}
		}

	}

	fmt.Println(mp)

	for key, val := range mp {
		if len(val) > 0 {
			fmt.Println(key, ":", val[0])
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
