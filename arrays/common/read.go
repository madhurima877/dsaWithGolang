package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadData(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var testCases [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")

		var arr []int

		for _, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				return nil, err
			}

			arr = append(arr, num)
		}

		testCases = append(testCases, arr)
	}

	return testCases, scanner.Err()
}
