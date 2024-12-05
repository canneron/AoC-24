package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parse(input string, sep string) ([][]int, error) {
	var rules [][]int

	entries := strings.Split(input, "\n")
	for _, entry := range entries {
		if entry == "" {
			continue
		}

		var convertedValues []int
		values := strings.Split(entry, sep)
		for _, value := range values {
			number, _ := strconv.Atoi(value)
			convertedValues = append(convertedValues, number)
		}

		rules = append(rules, convertedValues)
	}
	return rules, nil
}

func updateContainsRuleNumbers(update []int, rule []int) bool {
	for _, ruleNumber := range rule {
		if !slices.Contains(update, ruleNumber) {
			return false
		}
	}
	return true
}

func isUpdateValidForRules(update []int, rules [][]int) bool {
	for _, rule := range rules {
		if !updateContainsRuleNumbers(update, rule) {
			continue
		}

		if slices.Index(update, rule[0]) > slices.Index(update, rule[1]) {
			return false
		}
	}
	return true
}

func makeUpdateCompliant(update []int, rules [][]int) []int {
	for i := 0; i < len(update); i++ {

		lowest := -1
		value := update[i]
		for _, rule := range rules {
			if rule[0] != update[i] {
				continue
			}
			l := slices.Index(update, rule[1])
			if l != -1 && (lowest == -1 || l < lowest) {
				lowest = l
			}
		}
		if lowest != -1 && lowest < i {
			update = slices.Delete(update, i, i+1)
			update = slices.Insert(update, lowest, value)
		}

	}

	return update
}

func main() {

	data, _ := os.ReadFile("input.txt")

	input := strings.Split(string(data), "\n\n")
	rules, _ := parse(input[0], "|")
	updates, _ := parse(input[1], ",")

	countPartOne := 0
	countPartTwo := 0
	for _, update := range updates {
		if isUpdateValidForRules(update, rules) {
			countPartOne += update[len(update)/2]
		} else {
			fixedUpdate := makeUpdateCompliant(update, rules)
			countPartTwo += fixedUpdate[len(fixedUpdate)/2]
		}
	}

	fmt.Println(countPartOne)
	fmt.Println(countPartTwo)

}
