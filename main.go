package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	day4()
}

func day1() {
	file, err := os.Open("./day1part1.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var list1 []int64
	var list2 []int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		num1, err1 := strconv.ParseInt(fields[0], 10, 64)
		num2, err2 := strconv.ParseInt(fields[1], 10, 64)
		if err1 == nil && err2 == nil {
			list1 = append(list1, num1)
			list2 = append(list2, num2)
		}
	}

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	similarityCounts := make(map[int64]int64, 0)

	for _, value := range list1 {
		_, present := similarityCounts[value]
		if present {
			fmt.Println("Can skip")
		}
		for _, secondVal := range list2 {
			if value == secondVal {
				similarityCounts[value] += 1
			}
		}
	}

	fmt.Println(similarityCounts)

	var sum int64

	for _, value := range list1 {
		similarityScore := similarityCounts[value]
		sum += value * similarityScore
	}

	fmt.Println(sum)
}

func day2() {
	file, err := os.Open("./day2input.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	safeReports := 0
	var safeReportListing [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		var level []float64
		for _, field := range fields {
			conv, err := strconv.ParseFloat(field, 64)
			if err != nil {
				fmt.Println("Error parsing float")
			}
			level = append(level, conv)
		}

		safe := checkRules(level)

		if !safe {
			for i := range level {
				cleanedLevel := removeIndexFromSlice(level, i)
				cleanedSafe := checkRules(cleanedLevel)
				if cleanedSafe {
					safe = cleanedSafe
					break
				}
			}
		}

		if safe {
			safeReports += 1
			safeReportListing = append(safeReportListing, fields)
		}
	}

	fmt.Println(safeReports)
}

func day3() {
	file, err := os.Open("./day3input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var input string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input += line
	}

	enabled := true
	var enabledStrings []string
	remainder := input

	for running := true; running; {
		if enabled {
			if strings.Contains(remainder, "don't()") {
				enabledString, cutRemainder, _ := strings.Cut(remainder, "don't()")
				enabledStrings = append(enabledStrings, enabledString)
				remainder = cutRemainder
				enabled = !enabled
			} else {
				enabledString, _, _ := strings.Cut(remainder, "don't()")
				enabledStrings = append(enabledStrings, enabledString)
				running = false
			}
		} else {
			if strings.Contains(remainder, "do()") {
				_, cutRemainder, _ := strings.Cut(remainder, "do()")
				remainder = cutRemainder
				enabled = !enabled
			} else {
				running = false
			}
		}
	}

	input = strings.Join(enabledStrings, "")

	regexFinder, err := regexp.Compile(`mul\(\d*,\d*\)`)
	if err != nil {
		fmt.Println("error making regex")
	}
	matches := regexFinder.FindAllString(input, -1)

	var sum float64
	for _, match := range matches {
		tempClean := strings.Split(match, "(")[1]
		clean := strings.Split(tempClean, ")")[0]
		var values []float64
		for _, value := range strings.Split(clean, ",") {
			value, _ := strconv.ParseFloat(value, 64)
			values = append(values, value)
		}
		sum += values[0] * values[1]
	}
	fmt.Println(fmt.Sprintf("%d", int(sum)))
}

func day4() {
	file, err := os.Open("./day4input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	characters := []string{}

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	lineLength := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineLength = utf8.RuneCountInString(line)
		for _, value := range line {
			characters = append(characters, fmt.Sprintf("%c", value))
		}
		lineNumber++
	}

	grid := make([][]string, lineNumber)
	for i := range grid {
		grid[i] = make([]string, lineLength)
		for j := range grid[i] {
			characterIndex := (i * lineLength) + j
			grid[i][j] = characters[characterIndex]
		}
	}

	fmt.Println(lineLength, lineNumber)
	for _, row := range grid {
		fmt.Println(row)
	}

	var wordsFound float64

	for x, row := range grid {
		for y := range row {
			// for i := 1; i <= 9; i++ {
			// 	if checkForWord(x, y, "XMAS", i, grid) {
			// 		wordsFound += 1
			// 	}
			// }
			if checkForX(x, y, grid) {
				wordsFound += 1
				continue
			}
		}
	}

	fmt.Printf("FOUND: %d \n", int(wordsFound))
}

func checkForX(x int, y int, grid [][]string) bool {
	if x < 1 || x > len(grid[0])-2 || y < 1 || y > len(grid)-2 {
		return false
	}
	if grid[x][y] != "A" {
		return false
	}

	if (grid[x-1][y-1] == "M" && grid[x+1][y+1] == "S") || (grid[x-1][y-1] == "S" && grid[x+1][y+1] == "M") {
		if (grid[x+1][y-1] == "M" && grid[x-1][y+1] == "S") || (grid[x-1][y+1] == "M" && grid[x+1][y-1] == "S") {
			return true
		}
	} else if (grid[x+1][y+1] == "M" && grid[x-1][y-1] == "S") || (grid[x+1][y+1] == "S" && grid[x-1][y-1] == "M") {
		if (grid[x+1][y-1] == "M" && grid[x-1][y+1] == "S") || (grid[x-1][y+1] == "M" && grid[x+1][y-1] == "S") {
			return true
		}
	}

	return false
}

func checkForWord(x int, y int, remaining string, direction int, grid [][]string) bool {
	// 1 2 3
	// 4   6
	// 7 8 9

	var nextCharacter string

	nextCharacter = firstCharacter(remaining)

	if grid[x][y] == nextCharacter {
		_, nextRemaining, _ := strings.Cut(remaining, nextCharacter)

		if nextRemaining == "" {
			return true
		}
		newX := x
		newY := y

		switch direction {
		case 1:
			newX -= 1
			newY -= 1
		case 2:
			newY -= 1
		case 3:
			newX += 1
			newY -= 1
		case 4:
			newX -= 1
		case 6:
			newX += 1
		case 7:
			newX -= 1
			newY += 1
		case 8:
			newY += 1
		case 9:
			newX += 1
			newY += 1
		}

		if newX < 0 || newX > len(grid[0])-1 {
			return false
		}
		if newY < 0 || newY > len(grid)-1 {
			return false
		}

		// check direction recursing here.
		if checkForWord(newX, newY, nextRemaining, direction, grid) {
			return true
		}
	}

	return false
}

func firstCharacter(str string) string {
	v := []rune(str)
	return string(v[:1])
}

func contains(slice []int, item int) bool {
	for _, elem := range slice {
		if elem == item {
			return true
		}
	}
	return false
}

func removeIndexFromSlice(slice []float64, index int) []float64 {
	result := append([]float64{}, slice[:index]...)
	result = append(result, slice[index+1:]...)
	return result
}

func checkRules(level []float64) bool {
	var ascending bool
	// only ascending or descending
	// diff between adjacent values is between 1 and 3
	ascending = level[0] < level[1]

	for i, value := range level {
		if i > 0 {
			if ascending {
				if level[i-1] > value {
					fmt.Println("Descending on an ascending row")
					return false
				}
			} else {
				if level[i-1] < value {
					fmt.Println("Ascending on an descending row")
					return false
				}
			}

			stepSize := math.Abs(level[i-1] - value)
			if stepSize < 1 || stepSize > 3 {
				fmt.Printf("Wrong step size of %d\n", int(stepSize))
				return false
			}
		}
	}

	return true
}
