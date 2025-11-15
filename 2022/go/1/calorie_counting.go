package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"strconv"
)

func main() {
	filePath := "calories.txt"
	caloriesFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	mostCaloriesElf, mostCalories := getElfWithMostCalories(caloriesFile)
	fmt.Printf("Most calories elf: %d  with %d calories \n", mostCaloriesElf, mostCalories)
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
		if (currentElfCalories > mostCalories) {
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
