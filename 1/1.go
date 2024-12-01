package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(l1 []float64, l2 []float64) {
	sort.Float64s(l1)
	sort.Float64s(l2)

	sum := 0.0
	for i := 0; i < len(l1); i++ {
		sum += (math.Abs(l1[i] - l2[i]))
	}

	fmt.Println(int(sum))
}

func part2(l1 []float64, l2 []float64) {
	intCount := make(map[float64]int)

	for i := range l2 {
		_, ok := intCount[l2[i]]
		if ok {
			intCount[l2[i]]++
		} else {
			intCount[l2[i]] = 1
		}
	}

	sum := 0
	for num := range l1 {
		sum += int(l1[num]) * intCount[l1[num]]
	}

	fmt.Println(sum)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var l1 []float64
	var l2 []float64

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		lineNums := strings.Split(line, "   ")

		if len(lineNums) >= 2 {
			num1, err1 := strconv.ParseFloat(lineNums[0], 64)
			num2, err2 := strconv.ParseFloat(lineNums[1], 64)
			if err1 == nil && err2 == nil {
				l1 = append(l1, num1)
				l2 = append(l2, num2)
			} else {
				fmt.Printf("%s / %s", lineNums[0], lineNums[1])
				fmt.Println("Error:")
			}
		}
	}

	part1(l1, l2)
	part2(l1, l2)
}
