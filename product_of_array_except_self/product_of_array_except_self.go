package productexceptself

// ProductExceptSelf returns a slice where res[i] is the product of all values
// in nums except nums[i].
//
// Constraints:
// - No division
// - O(n) time
// - Use an extra slice for result
//
// Expected algorithm: Prefix/Suffix Products
// Target concepts: prefix accumulation, reverse pass
func ProductExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	length := len(nums)
	results := make([]int, length)

	for i := range nums {
		results[i] = 1
	}

	left := 1
	right := 1

	for i := range nums {
		results[i] *= left
		left *= nums[i]

		j := length - 1 - i
		results[j] *= right
		right *= nums[j]

	}

	return results
}
