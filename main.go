package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use: go run main.go <archivo_transcripcion>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := strings.TrimSuffix(inputFile, ".txt") + "_cleaned.txt"

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer output.Close()

	scanner := bufio.NewScanner(file)
	timePattern := regexp.MustCompile(`\d{2}:\d{2}:\d{2}\.\d{3} --> \d{2}:\d{2}:\d{2}\.\d{3}.*`)
	tagPattern := regexp.MustCompile(`<.*?>`)
	lastLine := ""

	for scanner.Scan() {
		line := scanner.Text()
		if timePattern.MatchString(line) {
			continue
		}
		cleanedLine := tagPattern.ReplaceAllString(line, "")
		cleanedLine = strings.TrimSpace(cleanedLine)
		if cleanedLine != "" && cleanedLine != lastLine {
			if _, err := output.WriteString(cleanedLine + " "); err != nil {
				fmt.Printf("Error writing output file: %v\n", err)
				os.Exit(1)
			}
			lastLine = cleanedLine
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("File cleaned created: %s\n", outputFile)
}

