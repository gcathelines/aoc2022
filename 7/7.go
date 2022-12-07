package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	fle    = 0
	folder = 1
)

type dir struct {
	typ      int
	filesize int
	name     string
	content  map[string]*dir
	root     *dir
}

var maxSum int

func main() {
	file, err := os.Open("./7.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	root := &dir{
		typ:     folder,
		name:    "root",
		root:    nil,
		content: make(map[string]*dir),
	}
	curr := root
	var isLS bool
	for scanner.Scan() {
		txt := scanner.Text()

		if isLS && string(txt[0]) != "$" {
			if !strings.HasPrefix(txt, "dir") {
				var size int
				var name string
				if _, err := fmt.Sscanf(txt, "%d %s", &size, &name); err != nil {
					panic(err)
				}
				if _, ok := curr.content[name]; !ok {
					dir := &dir{
						typ:      fle,
						name:     name,
						root:     curr,
						filesize: size,
					}
					curr.content[name] = dir
				}
			}
		}

		if string(txt[0]) == "$" {
			if isLS {
				isLS = false
			}

			var input string
			if _, err := fmt.Sscanf(txt, "$ %s", &input); err != nil {
				panic(err)
			}

			switch input {
			case "cd":
				var name string
				if _, err := fmt.Sscanf(txt, "$ cd %s", &name); err != nil {
					panic(err)

				}
				switch name {
				case "..":
					curr = curr.root
				case "/":
					curr = root
				default:
					if _, ok := curr.content[name]; !ok {
						dir := &dir{
							typ:     folder,
							name:    name,
							root:    curr,
							content: make(map[string]*dir),
						}
						curr.content[name] = dir
						curr = dir
					} else {
						curr = curr.content[name]
					}
				}
			case "ls":
				isLS = true
			}
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var y int
	sum(root)

	fmt.Println("RootSum", rootSum(root, &y), "\nTotal Below 100000:", y)

	need := (root.filesize + 30000000) - 70000000
	min := root
	findNearest(root, need, min)
	fmt.Println("To be purged:", min.filesize)
}

func findNearest(root *dir, n int, min *dir) {
	if root.typ == folder {
		if root.filesize >= n && root.filesize < min.filesize {
			fmt.Println("masuk", root.filesize, min.filesize)
			min = root
		}
		for _, c := range root.content {
			if c != nil {
				findNearest(c, n, min)
			}
		}
	}
}

func rootSum(root *dir, x *int) int {
	var dirSum int
	if root.typ == fle {
		return root.filesize
	}

	for _, c := range root.content {
		dirSum += rootSum(c, x)
	}

	if dirSum <= 100000 {
		*x += dirSum
	}

	return dirSum
}

func sum(root *dir) int {
	for _, c := range root.content {
		var res int
		if c.typ == folder {
			res = sum(c)
		} else {
			res = c.filesize
		}

		root.filesize += res
	}

	return root.filesize
}
