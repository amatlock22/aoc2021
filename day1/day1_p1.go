package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	incCount := 0

	scanner.Scan()

	prevNum, _ := strconv.Atoi(scanner.Text())

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())

		if num > prevNum {
			incCount++
		}

		prevNum = num
	}

	fmt.Println("Increase Count:", incCount)
}
