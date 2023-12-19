package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./8.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := [][]int{}
	visited := [][]int{}

	i := 0
	for scanner.Scan() {
		txt := scanner.Text()
		x := []int{}
		for _, tree := range txt {
			i, err := strconv.Atoi(string(tree))
			if err != nil {
				panic(err)
			}
			x = append(x, i)
		}
		m = append(m, x)
		i++

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for _, arr := range m {
		visited = append(visited, make([]int, len(arr)))
	}

	for i, arr := range m {
		for j := range arr {
			if i == 0 || i == len(m)-1 {
				visited[i][j] = 1
			}
			if j == 0 || j == len(arr)-1 {
				visited[i][j] = 1
			}
		}
	}

	for i, arr := range m {
		for j := range arr {
			if visited[i][j] > 0 {
				cnt := 0
				prevMax := m[i][j]
				// top
				if i-1 < 0 {
					cnt = i + 1
					for cnt < len(arr)-1 {
						if m[cnt][j] > m[i][j] && m[cnt][j] > prevMax {
							visited[cnt][j] += 1
							prevMax = m[cnt][j]

						}
						cnt++
					}
				}
				// bottom
				if i+1 > len(arr)-1 {
					cnt = i - 1
					for cnt >= 0 {
						if m[cnt][j] > m[i][j] && m[cnt][j] > prevMax {
							visited[cnt][j] += 1
							prevMax = m[cnt][j]

						}
						cnt--
					}
				}
				// left
				if j-1 < 0 {
					cnt = j + 1
					for cnt < len(arr)-1 {
						if m[i][cnt] > m[i][j] && m[i][cnt] > prevMax {
							visited[i][cnt] += 1
							prevMax = m[i][cnt]

						}
						cnt++
					}
				}
				// right
				if j+1 > len(arr)-1 {
					cnt = j - 1
					for cnt >= 0 {
						if m[i][cnt] > m[i][j] && m[i][cnt] > prevMax {
							visited[i][cnt] += 1
							prevMax = m[i][cnt]

						}
						cnt--
					}
				}
			}
		}
	}

	total := 0
	for i, arr := range visited {
		for j := range arr {
			if visited[i][j] > 0 {
				total++
			}
		}
	}

	fmt.Println("no 1", total)
	var a []int
	var b [][]int

	scenic := -1

	for i, arr := range m {
		for j := range arr {
			if i == 0 || i == len(arr)-1 || j == 0 || j == len(arr)-1 {
				a = append(a, 0)
				continue
			}

			// left
			var left, right, top, bottom int
			ptr := j - 1
			for ptr >= 0 {
				left++
				if m[i][j] <= m[i][ptr] {
					break
				}
				ptr--
			}

			ptr = j + 1
			for ptr < len(arr) {
				right++
				if m[i][j] <= m[i][ptr] {
					break
				}
				ptr++
			}

			ptr = i - 1
			for ptr >= 0 {
				top++

				if m[i][j] <= m[ptr][j] {
					break
				}
				ptr--

			}

			ptr = i + 1
			for ptr < len(arr) {
				bottom++
				if m[i][j] <= m[ptr][j] {
					break
				}
				ptr++

			}

			curr := left * right * top * bottom
			if curr > scenic {
				scenic = curr
			}
			a = append(a, curr)
		}
		b = append(b, a)
		a = []int{}
	}

	fmt.Println("no 2", scenic)

}
