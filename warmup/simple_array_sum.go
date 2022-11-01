package warmup

func SimpleArraySum(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res = res + nums[i]
	}
	return res
}
