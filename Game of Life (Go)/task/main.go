package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	lines    int
	columns  int
	universe [][]rune
)

func main() {
	// write your code here
	//_, err := fmt.Scan(&lines)
	//if err != nil {
	//	log.Fatal(err)
	//}
	lines = 110
	columns = 350
	universe = createUniverse(lines, columns)
	printUniverse(universe, 1)
	for i := 0; i < 1000; i++ {
		universe = nextGeneration(universe)
		printUniverse(universe, i+1)
	}
}

func calculateAlive(matrix [][]rune) int {
	alive := 0
	for i := range matrix {
		for s := range matrix[i] {
			if matrix[i][s] == 'O' {
				alive++
			}
		}
	}
	return alive
}

func calculateNeighbours(matrix [3][3]rune) int {
	neighbours := 0
	for i := range matrix {
		for s := range matrix[i] {
			if matrix[i][s] == 'O' {
				neighbours++
			}
		}
	}
	return neighbours
}

func nextGeneration(pgu [][]rune) [][]rune {
	ngu := createUniverse(lines, columns)
	for i := range pgu {
		for s := range pgu[i] {
			if pgu[i][s] == 'O' {
				neighbours := calculateNeighbours([3][3]rune{
					{pgu[indexSubstr(i, 1, lines)][indexSubstr(s, 1, lines)], pgu[indexSubstr(i, 1, lines)][s], pgu[indexSubstr(i, 1, lines)][indexSum(s, 1, lines)]},
					{pgu[i][indexSubstr(s, 1, lines)], ' ', pgu[i][indexSum(s, 1, lines)]},
					{pgu[indexSum(i, 1, lines)][indexSubstr(s, 1, lines)], pgu[indexSum(i, 1, lines)][s], pgu[indexSum(i, 1, lines)][indexSum(s, 1, lines)]},
				})
				if neighbours == 2 || neighbours == 3 {
					ngu[i][s] = 'O'
				} else {
					ngu[i][s] = ' '

				}
			} else {
				neighbours := calculateNeighbours([3][3]rune{
					{pgu[indexSubstr(i, 1, lines)][indexSubstr(s, 1, lines)], pgu[indexSubstr(i, 1, lines)][s], pgu[indexSubstr(i, 1, lines)][indexSum(s, 1, lines)]},
					{pgu[i][indexSubstr(s, 1, lines)], ' ', pgu[i][indexSum(s, 1, lines)]},
					{pgu[indexSum(i, 1, lines)][indexSubstr(s, 1, lines)], pgu[indexSum(i, 1, lines)][s], pgu[indexSum(i, 1, lines)][indexSum(s, 1, lines)]},
				})
				if neighbours == 3 {
					ngu[i][s] = 'O'
				} else {
					ngu[i][s] = ' '
				}
			}
		}
	}
	return ngu
}

func indexSum(index, addition, length int) (resultIndex int) {
	if index+addition > length-1 {
		resultIndex = index + addition - length
	} else {
		resultIndex = index + addition
	}
	return
}

func indexSubstr(index, subtrahend, length int) (resultIndex int) {
	if index-subtrahend < 0 {
		resultIndex = length + index - subtrahend
	} else {
		resultIndex = index - subtrahend
	}
	return
}

func createUniverse(lines, columns int) [][]rune {
	var un [][]rune
	for line := 0; line < lines; line++ {
		un = append(un, []rune{})
		for s := 0; s < columns; s++ {
			cell := ' '
			if rand.Intn(2) == 1 {
				cell = 'O'
			}
			un[line] = append(un[line], cell)
		}
	}
	return un
}

func printUniverse(universeToPrint [][]rune, generation int) {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("Generation #%v\n", generation)
	fmt.Printf("Alive: %v\n", calculateAlive(universeToPrint))
	for line := range universeToPrint {
		fmt.Printf("%v\n", string(universeToPrint[line]))
	}
	time.Sleep(33 * time.Millisecond)
}
