package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var gamma, epsilon int

	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	firstLine := scanner.Text()
	oneCountArr := make([]int, len(firstLine)) // zero count per index
	totalLines := 1

	for scanner.Scan() {
		line := scanner.Text()
		totalLines++

		for i := 0; i < len(line); i++ {
			if line[i] == '1' {
				oneCountArr[i]++
			}
		}
	}

	multiplier := len(firstLine) - 1

	// check if zeroes are the majority of the column
	for _, v := range oneCountArr {
		if v >= totalLines/2 {
			gamma += int(math.Pow(2, float64(multiplier)))
		} else {
			epsilon += int(math.Pow(2, float64(multiplier)))
		}

		multiplier--
	}

	fmt.Println("Power Consumption:", gamma*epsilon)
}
