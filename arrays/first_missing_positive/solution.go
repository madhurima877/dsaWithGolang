package main

func FirstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for nums[i] >= 1 &&
			nums[i] <= n &&
			nums[nums[i]-1] != nums[i] {

			val := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = val
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}