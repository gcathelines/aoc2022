package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var mp [][]rune
var visited [][]int

func main() {
	file, err := os.Open("./12.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var starts [][]int
	var end []int
	for scanner.Scan() {
		var arr []rune
		txt := scanner.Text()
		for idx, s := range txt {
			switch string(s) {
			// remove "a" for first problem
			case "S", "a":
				starts = append(starts, []int{len(mp), idx})
				arr = append(arr, rune("a"[0]))
			case "E":
				end = []int{len(mp), idx}
				arr = append(arr, rune("z"[0]))
			default:
				arr = append(arr, s)
			}
		}

		mp = append(mp, arr)
	}

	for _, arr := range mp {
		var x []int
		for range arr {
			x = append(x, 0)
		}
		visited = append(visited, x)
	}

	min := 9999999999
	var minTail Data
	for _, x := range starts {
		visited = [][]int{}
		for _, arr := range mp {
			var x []int
			for range arr {
				x = append(x, 0)
			}
			visited = append(visited, x)
		}
		tail := findPath(Data{
			coor: x,
		}, Data{
			coor: end,
		})

		if tail.coor[0] == end[0] && tail.coor[1] == end[1] {

			count := 0

			for tail.prev != nil {
				tail = *tail.prev
				count++
			}

			if count < min {
				minTail = tail
				min = count
			}
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(min, minTail)

	// Visualization
	// for i, arr := range visited {
	// 	for j := range arr {
	// 		if i == start[0] && j == start[1] {
	// 			fmt.Print("S")
	// 			continue
	// 		}
	// 		if i == end[0] && j == end[1] {
	// 			fmt.Print("E")
	// 			continue
	// 		}
	// 		if mp[fmt.Sprintf("%d,%d", i, j)] {
	// 			fmt.Print("1")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

}

func heuristic(curr, end []int, elevation int) float64 {
	return math.Sqrt(math.Pow(float64(curr[0]-end[0]), 2)+math.Pow(float64(curr[1]-end[1]), 2)) - float64(elevation)
}

func findPath(start, end Data) Data {
	h := &IntHeap{start}

	var curr Data
	for len(*h) > 0 {
		curr = h.Pop()

		if curr.coor[0] == end.coor[0] && curr.coor[1] == end.coor[1] {
			fmt.Println(curr.coor, end.coor)
			break
		}

		*h = h.Push(getAround(curr, end)...)

		// Visualization
		// for _, arr := range visited {
		// 	for _, v := range arr {
		// 		fmt.Print(v)
		// 	}
		// 	fmt.Println()
		// }
		// time.Sleep(1 * time.Millisecond)
		// fmt.Println()
		// fmt.Println()
		// fmt.Println()
		// fmt.Println()
		// fmt.Println()
		// fmt.Println()

	}
	return curr
}

func getAround(d Data, end Data) []Data {
	curr := d.coor
	value := mp[curr[0]][curr[1]]
	if string(value) == "S" {
		value = rune("a"[0])
	}
	if string(value) == "E" {
		value = rune("z"[0])
	}

	var res []Data

	if curr[1]-1 >= 0 && visited[curr[0]][curr[1]-1] == 0 && mp[curr[0]][curr[1]-1]-value <= 1 {
		res = append(res, Data{
			coor: []int{curr[0], curr[1] - 1},
			heur: heuristic([]int{curr[0], curr[1] - 1}, end.coor, int(mp[curr[0]][curr[1]-1]-value)),
			str:  string(mp[curr[0]][curr[1]-1]),
			prev: &d,
		})
		visited[curr[0]][curr[1]-1]++
	}

	if curr[1]+1 <= len(mp[0])-1 && visited[curr[0]][curr[1]+1] == 0 && mp[curr[0]][curr[1]+1]-value <= 1 {
		res = append(res, Data{
			coor: []int{curr[0], curr[1] + 1},
			heur: heuristic([]int{curr[0], curr[1] + 1}, end.coor, int(mp[curr[0]][curr[1]+1]-value)),
			str:  string(mp[curr[0]][curr[1]+1]),
			prev: &d,
		})
		visited[curr[0]][curr[1]+1]++
	}

	if curr[0]-1 >= 0 && visited[curr[0]-1][curr[1]] == 0 && mp[curr[0]-1][curr[1]]-value <= 1 {
		res = append(res, Data{
			coor: []int{curr[0] - 1, curr[1]},
			heur: heuristic([]int{curr[0] - 1, curr[1]}, end.coor, int(mp[curr[0]-1][curr[1]]-value)),
			str:  string(mp[curr[0]-1][curr[1]]),
			prev: &d,
		})
		visited[curr[0]-1][curr[1]]++
	}

	if curr[0]+1 <= len(mp)-1 && visited[curr[0]+1][curr[1]] == 0 && mp[curr[0]+1][curr[1]]-value <= 1 {
		res = append(res, Data{
			coor: []int{curr[0] + 1, curr[1]},
			heur: heuristic([]int{curr[0] + 1, curr[1]}, end.coor, int(mp[curr[0]+1][curr[1]]-value)),
			str:  string(mp[curr[0]+1][curr[1]]),
			prev: &d,
		})
		visited[curr[0]+1][curr[1]]++
	}

	return res
}

// An IntHeap is a min-heap of ints.
type IntHeap []Data

type Data struct {
	coor []int
	heur float64
	str  string
	prev *Data
}

func (h *IntHeap) Push(x ...Data) IntHeap {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x...)
	return *h
}

func (h *IntHeap) Pop() Data {
	old := *h
	x := old[0]
	*h = old[1:]
	return x
}
