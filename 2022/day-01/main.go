package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var path = "./input.txt"

	content := readInput(path)

	var sumsOfCalories []uint
	var elfTotalCalories uint

	for _, calorie := range content {
		calorieInt, err := strconv.Atoi(calorie)

		elfTotalCalories += uint(calorieInt)

		// If there is an empty space
		if err != nil {
			sumsOfCalories = append(sumsOfCalories, elfTotalCalories)
			elfTotalCalories = 0
			continue
		}
	}
  // Add the calories for the last elf in the list (the file)
  sumsOfCalories = append(sumsOfCalories, elfTotalCalories)

  // Part 1
  // Sort the slice from max to min
  sort.Slice(sumsOfCalories, func(i, j int) bool {
    return sumsOfCalories[i] > sumsOfCalories[j]
  })

  biggestCalorie := sumsOfCalories[0]
	fmt.Println("The Elf carrying the most Calories has a total of", biggestCalorie)

  // Part 2
  topThrees := sumsOfCalories[:3]
  var total uint

  for _, element := range topThrees {
    total += element
  }

  fmt.Println("The total of calories carried by the top 3 Elves is", total)
}

func readInput(path string) []string {
	// Open the file
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var lines []string

	// Read the file line per line
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
