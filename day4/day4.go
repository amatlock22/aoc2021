package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	drawnNumsStr := strings.Split(scanner.Text(), ",")
	drawnNums := make([]int, len(drawnNumsStr))

	for i, s := range drawnNumsStr {
		drawnNums[i], _ = strconv.Atoi(s)
	}

	scanner.Scan() // throw out first empty line

	counter := 0
	boardNum := 0

	boards := make([][]int, 1)
	boards[0] = make([]int, 25)

	for scanner.Scan() {
		numsStr := strings.Fields(scanner.Text())

		// end of board
		if len(numsStr) == 0 {
			boardNum++
			boards = append(boards, make([]int, 25))
			continue
		}

		for i, n := range numsStr {
			num, _ := strconv.Atoi(n)
			boards[boardNum][counter+i] = num
		}

		counter += 5
		if counter == 25 {
			counter = 0
		}
	}

	// making duplicate so there are fresh board copies in each part
	uneditedBoards := make([][]int, len(boards))
	for i := range boards {
		uneditedBoards[i] = make([]int, len(boards[i]))
		copy(uneditedBoards[i], boards[i])
	}

	fmt.Println("Part 1:", part1(drawnNums, boards))
	fmt.Println("Part 2:", part2(drawnNums, uneditedBoards))
}

func part1(drawnNums []int, boards [][]int) int {
	for _, num := range drawnNums {
		for _, board := range boards {
			markBoard(board, num)
			if checkWin(board) {
				return getScore(board) * num
			}
		}
	}

	return 0
}

func part2(drawnNums []int, boards [][]int) int {
	lastDeletedScore := 0

	for _, num := range drawnNums {
		if len(boards) == 0 {
			break
		}

		k := 0

		for _, board := range boards {
			markBoard(board, num)
			// keep boards that didn't win
			if !checkWin(board) {
				boards[k] = board
				k++
			} else {
				lastDeletedScore = getScore(board) * num
			}
		}

		boards = boards[:k]
	}

	return lastDeletedScore
}

func markBoard(board []int, num int) {
	for i, n := range board {
		if n == num {
			board[i] = -1
			break
		}
	}
}

func checkWin(board []int) bool {
	for i := 0; i < 5; i++ {
		j := i * 5

		// row check
		if board[j] == -1 && board[j+1] == -1 && board[j+2] == -1 && board[j+3] == -1 && board[j+4] == -1 {
			return true
		}

		// column check
		if board[i] == -1 && board[i+5] == -1 && board[i+10] == -1 && board[i+15] == -1 && board[i+20] == -1 {
			return true
		}
	}

	return false
}

func getScore(board []int) int {
	score := 0

	for _, n := range board {
		if n != -1 {
			score += n
		}
	}

	return score
}
