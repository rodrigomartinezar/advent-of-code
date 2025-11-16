package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func main() {
	filePath := "rockPaperScissor_input.txt"
	strategyFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	score := calculateScore(strategyFile)
	fmt.Printf("Score: %d\n", score)
}
// A rock
// B paper
// C scissor

// X lose
// Y draw
// Z win 
func calculateScore(strategyFile *os.File) int {
	score := 0
	pointsOfResult:= map[string]int{
		"X":0,
		"Y":3,
		"Z":6,
	}
	selectionMapForResult := make(map[string]map[string]int)
	selectionMapForResult["A"] = map[string]int{
		"X": 3,
		"Y": 1,
		"Z": 2,
	}
	selectionMapForResult["B"] = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	selectionMapForResult["C"] = map[string]int{
		"X": 2,
		"Y": 3,
		"Z": 1,
	}
	scanner := bufio.NewScanner(strategyFile)
	for scanner.Scan() {
		scoreOfDuel := 0
		line  := scanner.Text()
		oponent := string(line[0])
		expectedResult:= string(line[2])

		pointOfResult := pointsOfResult[expectedResult]
		pointOfSelection := selectionMapForResult[oponent][expectedResult]
		scoreOfDuel += pointOfResult + pointOfSelection
		score += scoreOfDuel
	}
	return score 
}

