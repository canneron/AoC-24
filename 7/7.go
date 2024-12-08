package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func intPow(exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= 10
	}
	return result
}

func digitCount(n int) int {
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func concat(a, b int) int {
	return a*intPow(digitCount(b)) + b
}

func p2(nums []int, idx int, current int, target int) bool {
	if idx == len(nums) {
		return current == target
	}

	if p2(nums, idx+1, current+nums[idx], target) {
		return true
	}

	if p2(nums, idx+1, current*nums[idx], target) {
		return true
	}

	concatenated := concat(current, nums[idx])
	if p2(nums, idx+1, concatenated, target) {
		return true
	}

	return false
}

func p1(nums []int, idx int, current int, target int) bool {
	if idx == len(nums) {
		return current == target
	}

	if p1(nums, idx+1, current+nums[idx], target) {
		return true
	}

	if p1(nums, idx+1, current*nums[idx], target) {
		return true
	}
	return false
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	sum2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parts[0])

		numStrings := strings.Fields(parts[1])
		nums := make([]int, len(numStrings))
		for i, numStr := range numStrings {
			nums[i], _ = strconv.Atoi(numStr)
		}

		if p1(nums, 1, nums[0], target) {
			sum += target
		}

		if p2(nums, 1, nums[0], target) {
			sum2 += target
		}
	}

	fmt.Println("p1:", sum)
	fmt.Println("p2:", sum2)
}
