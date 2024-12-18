package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
	X, Y int
}

func p2(grid [][]rune, x, y int, visited map[Coords]bool) int {
	if grid[y][x] == '9' {
		return 1
	}

	visited[Coords{X: x, Y: y}] = true

	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	trailCount := 0

	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]

		if nx >= 0 && ny >= 0 && ny < len(grid) && nx < len(grid[0]) && !visited[Coords{X: nx, Y: ny}] {
			curr, _ := strconv.Atoi(string(grid[y][x]))
			next, _ := strconv.Atoi(string(grid[ny][nx]))
			if next-curr == 1 {
				trailCount += p2(grid, nx, ny, visited)
			}
		}
	}

	visited[Coords{X: x, Y: y}] = false

	return trailCount
}

func p1(grid [][]rune, start Coords) int {
	queue := []Coords{start}
	visited := make(map[Coords]bool)
	reachedNines := make(map[Coords]bool)
	visited[start] = true

	directions := [][2]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}

	for len(queue) > 0 {
		cords := queue[0]
		queue = queue[1:]

		x, y := cords.X, cords.Y

		if grid[y][x] == '9' {
			reachedNines[cords] = true
			continue
		}

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			next := Coords{X: nx, Y: ny}

			if nx >= 0 && ny >= 0 && ny < len(grid) && nx < len(grid[0]) && !visited[next] {
				curr, _ := strconv.Atoi(string(grid[y][x]))
				nextHeight, _ := strconv.Atoi(string(grid[ny][nx]))

				if nextHeight-curr == 1 {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
	}
	return len(reachedNines)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := [][]rune{}
	var trailheads []Coords

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		for i := 0; i < len(line); i++ {
			if line[i] == '0' {
				trailheads = append(trailheads, Coords{X: i, Y: len(grid)})
			}
		}
		grid = append(grid, []rune(line))
	}

	score := 0
	for _, start := range trailheads {
		score += p1(grid, start)
	}

	score2 := 0
	for _, trailhead := range trailheads {
		visited := make(map[Coords]bool)
		score2 += p2(grid, trailhead.X, trailhead.Y, visited)
	}

	fmt.Println("p1:", score)
	fmt.Println("p2:", score2)
}
