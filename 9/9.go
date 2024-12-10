package main

import (
	"fmt"
	"os"
	"strconv"
)

type File struct {
	Start int
	End   int
	Size  int
	ID    int
}

func checkP2(sLine []int) {
	sum := 0
	for i, num := range sLine {
		if num == -1 {
			continue
		}

		sum += num * i
	}

	fmt.Println("p2:", sum)
}

func p2(input string) {
	var sLine []int
	var files []File
	id := 0

	for i := 0; i < len(input); i++ {
		length, _ := strconv.Atoi(string(input[i]))

		if i%2 == 0 {
			var file File
			file.Start = len(sLine)
			file.ID = id
			for k := 0; k < length; k++ {
				sLine = append(sLine, id)
			}
			file.End = len(sLine) - 1
			file.Size = file.End - file.Start + 1
			files = append(files, file)
			id++
		} else {
			for j := 0; j < length; j++ {
				sLine = append(sLine, -1)
			}
		}
	}

	for i := len(files) - 1; i >= 0; i-- {
		f := files[i]
		count := 0

		for count < f.Start {
			for count < f.Start && sLine[count] != -1 {
				count++
			}

			if count == f.Start {
				break
			}

			dots := 0
			startDots := count

			for count < f.Start && sLine[count] == -1 {
				dots++
				count++
			}

			if dots >= f.Size {
				for j := 0; j < f.Size; j++ {
					sLine[startDots+j] = i
					sLine[f.Start+j] = -1
				}
				break
			}
		}
	}

	checkP2(sLine)
}

func checkP1(sLine []int) {
	sum := 0
	for i, num := range sLine {
		if num == -1 {
			continue
		}

		sum += num * i
	}

	fmt.Println("p1:", sum)
}

func p1(input string) {
	var sLine []int
	id := 0

	for i := 0; i < len(input); i++ {
		length, _ := strconv.Atoi(string(input[i]))

		if i%2 == 0 {
			for k := 0; k < length; k++ {
				sLine = append(sLine, id)
			}
			id++
		} else {
			for j := 0; j < length; j++ {
				sLine = append(sLine, -1)
			}
		}
	}

	for {
		freeIdx := -1
		for i, c := range sLine {
			if c == -1 {
				freeIdx = i
				break
			}
		}

		fileIdx := -1
		for i := len(sLine) - 1; i >= 0; i-- {
			if sLine[i] != -1 {
				fileIdx = i
				break
			}
		}

		if freeIdx == -1 || fileIdx == -1 || fileIdx < freeIdx {
			break
		}

		sLine[freeIdx], sLine[fileIdx] = sLine[fileIdx], -1
	}

	checkP1(sLine)
}
func main() {
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	p1(input)
	p2(input)
}
