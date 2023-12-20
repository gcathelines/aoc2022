package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var mp [][]string
var done [][]bool
var width int

func main() {
	file, err := os.Open("./3-test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	y := 0
	for scanner.Scan() {
		txt := scanner.Text()
		arr := make([]string, len(txt))
		arrDone := make([]bool, len(txt))
		for x, d := range txt {
			arr[x] = string(d)
			arrDone[x] = false
		}
		done = append(done, arrDone)
		mp = append(mp, arr)
		y++
	}

	width = len(mp[0])

	sum := 0
	for y, m := range mp {
		for x, char := range m {
			if isSymbol(char) {
				num := checkAdjacent(x, y)
				sum += num
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}

// array of coordinates that contains integer
func checkAdjacent(x, y int) int {
	// part 1 is
	// num := 0
	// and inside ifs do
	// num += getNumber(x, y-1, width)
	// return num
	nums := []int{}
	if y-1 >= 0 {
		if !done[x][y-1] {
			num := getNumber(x, y-1, width)
			if num != 0 {
				nums = append(nums, num)
			}
		}
	}

	if y+1 < width {
		if !done[x][y+1] {
			num := getNumber(x, y+1, width)
			if num != 0 {
				nums = append(nums, num)
			}
		}
	}

	if x-1 >= 0 {
		if !done[x-1][y] {
			num := getNumber(x-1, y, width)
			if num != 0 {
				nums = append(nums, num)
			}
		}
	}

	if x+1 >= 0 {
		if !done[x+1][y] {
			num := getNumber(x+1, y, width)
			if num != 0 {
				nums = append(nums, num)
			}
		}
	}

	if x-1 >= 0 && y-1 >= 0 {
		if !done[x-1][y-1] {
			num := getNumber(x-1, y-1, width)
			if num != 0 {
				nums = append(nums, num)
			}
		}
	}

	if x+1 < width && y-1 >= 0 {
		if !done[x+1][y-1] {
			num := getNumber(x+1, y-1, width)
			if num != 0 {
				nums = append(nums, num)
			}
		}
	}

	if x-1 >= 0 && y+1 >= 0 {
		if !done[x-1][y+1] {
			num := getNumber(x-1, y+1, width)
			if num != 0 {
				nums = append(nums, num)
			}
		}
	}

	if x+1 < width && y+1 >= 0 {
		if !done[x+1][y+1] {
			num := getNumber(x+1, y+1, width)
			if num != 0 {
				nums = append(nums, num)
			}
		}
	}

	if len(nums) != 2 {
		return 0
	}

	return nums[0] * nums[1]
}

func isSymbol(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return false
	}

	if str == "." {
		return false
	}

	return true
}

func getNumber(x, y, width int) int {
	if _, err := strconv.Atoi(mp[y][x]); err != nil {
		return 0
	}

	initX := x

	str := mp[y][x]
	x--
	for x >= 0 {
		if _, err := strconv.Atoi(mp[y][x]); err != nil || mp[y][x] == "." {
			break
		}
		str = fmt.Sprintf("%s%s", mp[y][x], str)
		done[x][y] = true
		x--
	}

	x = initX + 1
	for x < width {
		if _, err := strconv.Atoi(mp[y][x]); err != nil || mp[y][x] == "." {
			break
		}
		str = fmt.Sprintf("%s%s", str, mp[y][x])
		done[x][y] = true
		x++
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}
