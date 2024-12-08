package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Coord struct {
	CX int
	CY int
}

func searchGrid2(grid [][]rune, nMap map[rune][]Coord) int {
	sum := 0

	var aNodes []Coord

	for _, nodes := range nMap {
		for _, node := range nodes {
			if !slices.Contains(aNodes, node) {
				aNodes = append(aNodes, node)
				sum++
			}
			for _, antiNode := range nodes {
				if node == antiNode {
					continue
				}
				x := node.CX - antiNode.CX
				y := node.CY - antiNode.CY

				dx := node.CX + x
				dy := node.CY + y
				for {
					if dy >= len(grid) || dx >= len(grid[0]) || dx < 0 || dy < 0 {
						break
					}

					anti := Coord{CX: dx, CY: dy}
					if !slices.Contains(aNodes, anti) {
						aNodes = append(aNodes, anti)
						if dy < len(grid) && dx < len(grid[0]) && dx >= 0 && dy >= 0 {
							sum++
						}
					}

					dx += x
					dy += y
				}
			}
		}
	}

	return sum
}

func searchGrid(grid [][]rune, nMap map[rune][]Coord) int {
	sum := 0

	var aNodes []Coord
	for _, nodes := range nMap {
		for _, node := range nodes {
			for _, antiNode := range nodes {
				if node == antiNode {
					continue
				}
				x := node.CX - antiNode.CX
				y := node.CY - antiNode.CY

				dx := node.CX + x
				dy := node.CY + y
				anti := Coord{CX: dx, CY: dy}
				if slices.Contains(aNodes, anti) {
					continue
				} else {
					aNodes = append(aNodes, anti)
				}
				if dy < len(grid) && dx < len(grid[0]) && dx >= 0 && dy >= 0 {
					sum++
				}
			}
		}
	}

	return sum
}

func p2(grid [][]rune, nodes map[rune][]Coord) {
	sum := searchGrid2(grid, nodes)
	fmt.Println("P2:", sum)
}

func p1(grid [][]rune, nodes map[rune][]Coord) {
	sum := searchGrid(grid, nodes)
	fmt.Println("P1:", sum)
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
	nodes := map[rune][]Coord{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		grid = append(grid, []rune(line))
		for i, r := range line {
			if r != '.' {
				nodes[r] = append(nodes[r], Coord{CX: i, CY: len(grid) - 1})
			}
		}
	}

	p1(grid, nodes)
	p2(grid, nodes)
}
