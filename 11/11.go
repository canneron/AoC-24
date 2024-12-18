package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func blinkP1(input []int) {
	for range 25 {
		var newStones []int

		for _, num := range input {
			if num == 0 {
				newStones = append(newStones, 1)
			} else if digits := strconv.Itoa(num); len(digits)%2 == 0 {
				a, _ := strconv.Atoi(digits[:len(digits)/2])
				b, _ := strconv.Atoi(digits[len(digits)/2:])
				newStones = append(newStones, a)
				newStones = append(newStones, b)
			} else {
				newStones = append(newStones, num*2024)
			}
		}

		input = newStones
	}

	fmt.Println(len(input))
}

func p1(input string) {
	p1Output := strings.Split(input, " ")
	fmt.Println(len(p1Output))
}

func main() {
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	stoneArr := strings.Split(input, " ")
	var intStones []int
	for _, stone := range stoneArr {
		nStone, _ := strconv.Atoi(stone)
		intStones = append(intStones, nStone)
	}

	blinkP1(intStones)
}
