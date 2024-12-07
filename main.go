package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	day2()
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

		var ascending bool
		safe := true

		for i, value := range fields {
			valueNumber, _ := strconv.ParseInt(value, 10, 64)

			if i == 0 {
				nextValue, _ := strconv.ParseInt(fields[i+1], 10, 64)
				ascending = valueNumber < nextValue
			}

			if i != 0 {
				previousValue, _ := strconv.ParseInt(fields[i-1], 10, 64)
				step := math.Abs(float64(valueNumber) - float64(previousValue))

				isAscending := value > fields[i-1]
				if (isAscending != ascending) || step < 1 || step > 3 {
					safe = false
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
