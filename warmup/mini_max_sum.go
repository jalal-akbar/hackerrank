package warmup

func MiniMaxSum(arr []int32) {
	var res int32 = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < arr[i+1] {
			res = res + arr[i]
		}
	}

}
