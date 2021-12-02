package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var hPos, depth, aim int

	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(command[1])

		switch command[0] {
		case "forward":
			hPos += n
			depth += n * aim
		case "down":
			aim += n
		case "up":
			aim -= n
		}
	}

	fmt.Println("Final Position * Final Depth:", hPos*depth)
}
