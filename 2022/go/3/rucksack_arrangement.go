package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := "rearrangeFile.txt"
	rearrangeFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	sum := 0
	scanner := bufio.NewScanner(rearrangeFile)
	for scanner.Scan() {
		line := scanner.Text()
		firstSubString, secondSubString := splitStringInHalf(string(line))
		repeatedChar := findRepeatedChar(firstSubString, secondSubString)
		priority := getPriorityFromAscii(repeatedChar)
		sum += priority
	}
	fmt.Printf("Result: %d", sum)
}

func splitStringInHalf(line string) (string, string) {
	lineCharNumber := len(line)
	halfStringIndex := lineCharNumber / 2
	firstSubString := line[:halfStringIndex]
	secondSubString := line[halfStringIndex:]
	return string(firstSubString), string(secondSubString)
}

func findRepeatedChar(firsString string, secondString string) string {
	for i := 0; i < len(firsString); i++ {
		for j := 0; j < len(secondString); j++ {
			if firsString[i] == secondString[j] {
				return string(firsString[i])
			}
		}
	}
	return ""
}

func getPriorityFromAscii(char string) int {
	asciiValue := char[0]
	if asciiValue < 97 {
		return int(asciiValue) - 38
	}
	return int(asciiValue) - 96
}
