package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Step struct {
	Destination int
	Source      int
	Range       int
}

func doStep(step Step, seed int) int {
	start := step.Source
	end := step.Source + step.Range - 1
	if seed >= start && seed <= end {
		return step.Destination + (seed - start)
	}

	return seed
}

func main() {
	var input = flag.String("input", "in", `use "test" if want to use test input`)
	flag.Parse()
	if *input == "test" {
		*input = "in-test"
	}
	file, err := os.Open(fmt.Sprintf("./%s.txt", *input))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		seeds []int
		steps [][]Step
	)

	step := []Step{}
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix(txt, "seeds:") {
			data := strings.Split(txt, ":")
			seedsTemp := strings.Split(data[1], " ")
			for _, s := range seedsTemp {
				if s == "" {
					continue
				}
				seed, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, seed)
			}
			continue
		}

		if txt == "" {
			continue
		}

		if strings.HasSuffix(txt, "map:") {
			if len(step) > 0 {
				steps = append(steps, step)
				step = []Step{}
			}
			continue
		}

		stepTemp := strings.Split(txt, " ")
		if len(stepTemp) != 3 {
			panic("should be 3")
		}
		dest, err := strconv.Atoi(stepTemp[0])
		if err != nil {
			panic(err)
		}
		src, err := strconv.Atoi(stepTemp[1])
		if err != nil {
			panic(err)
		}
		rng, err := strconv.Atoi(stepTemp[2])
		if err != nil {
			panic(err)
		}

		step = append(step, Step{
			Destination: dest,
			Source:      src,
			Range:       rng,
		})
	}

	if len(step) > 0 {
		steps = append(steps, step)
	}

	min := 99999999999
	for _, seed := range seeds {
		x := seed
		for _, s := range steps {
			for _, step := range s {
				res := doStep(step, x)
				if res != x {
					x = res
					fmt.Println(fmt.Sprintf("seed %d: %d -> %v", seed, x, step))
					break
				}
			}
		}
		if x < min {
			min = x
		}
		fmt.Println(fmt.Sprintf("FINAL FOR SEED %d: %d \n\n", seed, x))

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(min)

}
