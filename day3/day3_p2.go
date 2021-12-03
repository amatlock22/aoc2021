package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var oxygenList []string

	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		oxygenList = append(oxygenList, line)
	}

	co2List := make([]string, len(oxygenList))
	copy(co2List, oxygenList)
	oxygen, co2 := getOxygenAndCo2(oxygenList, co2List, 0)

	fmt.Println("Life Support Rating:", oxygen*co2)
}

func getOxygenAndCo2(oxygenList, co2List []string, i int) (oxygen int, co2 int) {
	if len(oxygenList) == 1 && len(co2List) == 1 {
		convertedBinary, _ := strconv.ParseInt(oxygenList[0], 2, 64)
		oxygen = int(convertedBinary)

		convertedBinary, _ = strconv.ParseInt(co2List[0], 2, 64)
		co2 = int(convertedBinary)

		return
	}

	oxygenOneCount, co2OneCount := 0, 0
	oxygenZeroCount, co2ZeroCount := 0, 0
	mostCommon, leastCommon := '1', '0'

	if len(oxygenList) > 1 {
		for _, v := range oxygenList {
			if v[i] == '1' {
				oxygenOneCount++
			} else {
				oxygenZeroCount++
			}
		}

		if oxygenOneCount < oxygenZeroCount {
			mostCommon = '0'
		}

		k := 0

		for j := range oxygenList {
			if oxygenList[j][i] == byte(mostCommon) {
				oxygenList[k] = oxygenList[j]
				k++
			}
		}

		oxygenList = oxygenList[:k]
	}

	if len(co2List) > 1 {
		for _, v := range co2List {
			if v[i] == '1' {
				co2OneCount++
			} else {
				co2ZeroCount++
			}
		}

		if co2OneCount < co2ZeroCount {
			leastCommon = '1'
		}

		k := 0

		for j := range co2List {
			if co2List[j][i] == byte(leastCommon) {
				co2List[k] = co2List[j]
				k++
			}
		}

		co2List = co2List[:k]
	}

	i++

	return getOxygenAndCo2(oxygenList, co2List, i)
}
