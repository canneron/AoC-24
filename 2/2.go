package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func performCheck(s1 float64, s2 float64, safe bool, increase bool) bool {
	absVal := math.Abs(s1 - s2)
	if absVal >= 1.0 && absVal <= 3.0 {
		if (!increase && s1 < s2) || (increase && s1 > s2) {
			safe = false
		}
	} else {
		safe = false
	}

	return safe
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		lineNums := strings.Split(line, " ")

		increase := false

		if len(lineNums) >= 2 {
			l1, _ := strconv.ParseFloat(lineNums[0], 64)
			l2, _ := strconv.ParseFloat(lineNums[1], 64)
			if l1 < l2 {
				increase = true

			} else {
				increase = false
			}
		}

		safe := true

		for i := 0; i < len(lineNums)-1; i++ {
			s1, _ := strconv.ParseFloat(lineNums[i], 64)
			s2, _ := strconv.ParseFloat(lineNums[i+1], 64)

			safe = performCheck(s1, s2, safe, increase)

			// Part 2 code (remove for part 1 answer)
			if !safe && (i+2 < len(lineNums)) {
				safe = true
				s3, _ := strconv.ParseFloat(lineNums[i+2], 64)
				safe = performCheck(s1, s3, safe, increase)
			}
		}

		if safe {
			sum++
		}
	}

	fmt.Println(sum)
}
