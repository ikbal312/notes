/*

source:
https://leetcode.com/problems/contiguous-array/description/

*/

package main

func findMaxLength(nums []int) int {
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	mxlen := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sum += -1
		} else {
			sum += 1
		}

		if v, ok := m[sum]; ok {
			mxlen = max(mxlen, i-v)
		} else {
			m[sum] = i
		}
	}
	return mxlen
}
