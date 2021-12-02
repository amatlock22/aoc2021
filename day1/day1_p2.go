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
	num1, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	num2, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	num3, _ := strconv.Atoi(scanner.Text())

	for scanner.Scan() {
		num4, _ := strconv.Atoi(scanner.Text())

		if num2+num3+num4 > num1+num2+num3 {
			incCount++
		}

		num1, num2, num3 = num2, num3, num4
	}

	fmt.Println("Increase Count:", incCount)
}
