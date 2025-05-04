/*
source:
https://leetcode.com/problems/range-sum-query-immutable/description/
*/
package main

type NumArray struct {
	cache []int
}

func cacheSum(this *NumArray, nums []int) {
	this.cache = make([]int, len(nums))
	this.cache[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		this.cache[i] = this.cache[i-1] + nums[i]
	}
}

func Constructor(nums []int) NumArray {
	obj := NumArray{}
	cacheSum(&obj, nums)
	return obj
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.cache[right]
	}
	return this.cache[right] - this.cache[left-1]
}
