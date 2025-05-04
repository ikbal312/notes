/*
source
https://leetcode.com/problems/subarray-sum-equals-k
*/

package main

func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	m[0] = 1 // base case when cumulative start from 0
	cumulativeSum := 0
	kOccurrence := 0
	for _, num := range nums {
		cumulativeSum += num
		if count, exists := m[cumulativeSum-k]; exists {
			kOccurrence += count
		}
		m[cumulativeSum]++

	}
	return kOccurrence
}
