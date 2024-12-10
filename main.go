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
)

func main() {
	day3()
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
