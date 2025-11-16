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

// X rock
// Y paper
// Z scissor
func calculateScore(strategyFile *os.File) int {
	score := 0
	pointsOfResponses := map[string]int{
		"X":1,
		"Y":2,
		"Z":3,
	}
	pointsOfOponent:= map[string]int{
		"A":-1,
		"B":-2,
		"C":-3,
	}
	scanner := bufio.NewScanner(strategyFile)
	for scanner.Scan() {
		scoreOfDuel := 0
		line  := scanner.Text()
		oponent := string(line[0])
		response:= string(line[2])

		scoreOfResponse:= pointsOfResponses[response]
		scoreOfOponent:= pointsOfOponent[oponent]
		result := scoreOfResponse + scoreOfOponent
		if (result == 0) {
			scoreOfDuel += scoreOfResponse + 3	
		}
		if (result == 1 || result == -2) {
			scoreOfDuel += scoreOfResponse + 6	
		}
		if (result == -1 || result == 2) {
			scoreOfDuel += scoreOfResponse
		}
		score += scoreOfDuel
	}
	return score 
}

