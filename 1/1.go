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

	sort.Float64s(l1)
	sort.Float64s(l2)

	sum := 0.0
	for i := 0; i < len(l1); i++ {
		sum += (math.Abs(l1[i] - l2[i]))
	}

	fmt.Println(int(sum))
}
