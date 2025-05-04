
// maximum sum of subArray
// time complexity: O(n^2)
// space complexity: O(1)
func maxSumSubArrayBruteForce(nums []int) int {
	maxSum := nums[0]
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

// max sum of subArray kadanes algorithm
// time complexity: O(n)
// space complexity: O(1)
func maxSumSubArrayKadane(nums []int) int {
	maxSums := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}
