package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	filePath := "caloriesTopElves.txt"
	caloriesFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	//	mostCaloriesElf, mostCalories := getElfWithMostCalories(caloriesFile)
	//	fmt.Printf("Most calories elf: %d  with %d calories \n", mostCaloriesElf, mostCalories)
	totalCalories := getTopNElves(caloriesFile, 3)
	fmt.Printf("Total calories: %d", totalCalories)
}

func getElfWithMostCalories(caloriesFile *os.File) (int, int) {
	scanner := bufio.NewScanner(caloriesFile)
	mostCalories := 0
	mostCaloriesElf := 1
	currentElfCalories := 0
	currentElf := 1
	for scanner.Scan() {
		currentCalories, _ := strconv.Atoi(scanner.Text())
		currentElfCalories += currentCalories
		if currentElfCalories > mostCalories {
			mostCaloriesElf = currentElf
			mostCalories = currentElfCalories
		}
		if currentCalories == 0 {
			currentCalories = 0
			currentElf += 1
			currentElfCalories = 0
		}
	}
	return mostCaloriesElf, mostCalories
}

func indexOf(element int, data []int) int {
	for index, elementInSlice:= range data {
		if elementInSlice == element {
			return index
		}
	}
	return -1
}

func getTopNElves(caloriesFile *os.File, numberOfElves int) int {
	scanner := bufio.NewScanner(caloriesFile)
	var topCalories []int
	currentElfCalories := 0
	for scanner.Scan() {
		currentCalories, _ := strconv.Atoi(scanner.Text())
		currentElfCalories += currentCalories
		if currentCalories == 0 {
			if len(topCalories) < numberOfElves {
				topCalories = append(topCalories, currentCalories)
			}

			minimumCalories := slices.Min(topCalories)
			if minimumCalories < currentElfCalories {
				indexOfMin := indexOf(minimumCalories, topCalories)
				topCalories = slices.Delete(topCalories, indexOfMin, indexOfMin+1)
				topCalories = append(topCalories, currentElfCalories)
			}
			currentCalories = 0
			currentElfCalories = 0
		}
	}
	total := 0
	for _, value := range topCalories {
		total += value
	}
	return total
}
