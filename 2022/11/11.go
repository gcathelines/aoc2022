package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operationType string

type data struct {
	no             int
	items          []int
	operationType  operationType
	operationValue int
	testValue      int
	monkeyTrue     int
	monkeyFalse    int
}

var monkeys = map[int]data{}
var monkeyInspect = map[int]int{}

const (
	divisionRule                              = 1
	round                                     = 10000
	operationTypeMultiplication operationType = "*"
	operationTypeAddition       operationType = "+"
)

func main() {
	file, err := os.Open("./11.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var monkey data
	for scanner.Scan() {
		txt := scanner.Text()
		txt = strings.TrimSpace(txt)
		if strings.HasPrefix(txt, "Monkey") {
			var monkeyNo int

			monkeys[monkey.no] = monkey

			if _, err := fmt.Sscanf(txt, "Monkey %d:", &monkeyNo); err != nil {
				fmt.Println("error e", err)
			}
			monkey = data{
				no: monkeyNo,
			}

		}

		if strings.HasPrefix(txt, "Starting items:") {
			input := strings.Split(txt, ":")
			items := input[1]

			itemsString := strings.Split(strings.TrimSpace(items), ",")
			for _, item := range itemsString {
				itemInt, err := strconv.Atoi(strings.TrimSpace(item))
				if err != nil {
					panic(err)
				}

				monkey.items = append(monkey.items, itemInt)
			}
		}

		if strings.HasPrefix(txt, "Operation:") {
			var operation string
			var operationValue string
			if _, err := fmt.Sscanf(txt, "Operation: new = old %s %s", &operation, &operationValue); err != nil {
				fmt.Println("error d", err)
			}

			monkey.operationType = operationType(operation)

			if operationValue == "old" {
				monkey.operationValue = -1
			} else {
				val, err := strconv.Atoi(operationValue)
				if err != nil {
					panic(err)
				}

				monkey.operationValue = val
			}

		}

		var test int
		if strings.HasPrefix(txt, "Test:") {
			if _, err := fmt.Sscanf(txt, "Test: divisible by %d", &test); err != nil {
				fmt.Println("error c", err)
			}

			monkey.testValue = test
		}

		var monkeyTrue int
		if strings.HasPrefix(txt, "If true:") {
			if _, err := fmt.Sscanf(txt, "If true: throw to monkey %d", &monkeyTrue); err != nil {
				fmt.Println("error b", err)
			}

			monkey.monkeyTrue = monkeyTrue
		}
		var monkeyFalse int
		if strings.HasPrefix(txt, "If false:") {
			if _, err := fmt.Sscanf(txt, "If false: throw to monkey %d", &monkeyFalse); err != nil {
				fmt.Println("error a", err)
			}

			monkey.monkeyFalse = monkeyFalse
		}
	}
	monkeys[monkey.no] = monkey

	modulo := 1
	for _, mky := range monkeys {
		modulo *= mky.testValue
	}

	// loop per round
	for i := 0; i < round; i++ {
		// loop per monkey
		for j := 0; j < len(monkeys); j++ {
			// loop per monkey item
			mky := monkeys[j]
			for _, item := range mky.items {
				monkeyInspect[j]++

				var res int
				val := mky.operationValue
				if val == -1 {
					val = item
				}
				switch mky.operationType {
				case operationTypeAddition:
					res = item + val
				case operationTypeMultiplication:
					res = item * val
				default:
					panic(mky.operationType)
				}

				res = res % modulo
				if res%mky.testValue == 0 {
					newMky := monkeys[mky.monkeyTrue]
					newMky.items = append(newMky.items, res)
					monkeys[mky.monkeyTrue] = newMky
				} else {
					newMky := monkeys[mky.monkeyFalse]
					newMky.items = append(newMky.items, res)
					monkeys[mky.monkeyFalse] = newMky
				}

				mky.items = []int{}
				monkeys[j] = mky
			}
		}
	}

	var topMonkeys []int
	fmt.Println(monkeys)
	fmt.Println(monkeyInspect)
	for _, inspect := range monkeyInspect {
		if len(topMonkeys) < 2 {
			topMonkeys = append(topMonkeys, inspect)
			continue
		}

		if inspect < topMonkeys[1] {
			continue
		}

		if inspect > topMonkeys[0] {
			topMonkeys[1] = topMonkeys[0]
			topMonkeys[0] = inspect
		}

		if inspect > topMonkeys[1] && inspect < topMonkeys[0] {
			topMonkeys[1] = inspect
			continue
		}
	}

	fmt.Println(topMonkeys, topMonkeys[0]*topMonkeys[1])

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
