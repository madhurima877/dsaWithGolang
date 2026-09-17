package main

import (
	"dsaWithGolang/arrays/common"
	"fmt"
)

func main() {
	testCases, err := common.ReadData("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, arr := range testCases {
		result := FirstMissingPositive(arr)
		fmt.Println("Input:", arr)
		fmt.Println("Result:", result)
	}
}
