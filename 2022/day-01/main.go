package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var path = "./input.txt"

	content := ReadInput(path)

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

	biggesCalorie := MaxSlice(sumsOfCalories)

	fmt.Println("The Elf carrying the most Calories has a total of", biggesCalorie)
}

func MaxSlice(slice []uint) uint {
	var max, temp uint

	for _, element := range slice {
		if element > temp {
			temp = element
			max = temp
		}
	}

	return max
}

func ReadInput(path string) []string {
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
