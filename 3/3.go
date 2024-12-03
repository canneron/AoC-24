package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func evalMuls(strs []string) int {
	sum := 0
	for _, str := range strs {
		sub := str[4 : len(str)-1]
		nums := strings.Split(sub, ",")

		v1, _ := strconv.Atoi(string(nums[0]))
		v2, _ := strconv.Atoi(string(nums[1]))
		sum += (v1 * v2)
	}

	return sum
}

func part2(input string) {
	valid := strings.Split(input, "do()")

	sum := 0
	for _, val := range valid {
		v := strings.Split(val, "don't()")[0]
		regex, _ := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
		mul := regex.FindAllString(v, -1)
		sum += evalMuls(mul)
	}

	fmt.Println("total: ", sum)
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input3 := string(data)
	// sum := 0

	regex, _ := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	mul := regex.FindAllString(input3, -1)

	fmt.Println(mul)
	evalMuls(mul)
	part2(input3)
}
