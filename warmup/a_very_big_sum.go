package warmup

/*
 * Complete the 'aVeryBigSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER_ARRAY ar as parameter.
 */

func AVeryBigSum(ar []int64) int64 {
	// Write your code here
	var res int64 = 0

	for i := 0; i < len(ar); i++ {
		if len(ar[:i]) < 10 || len(ar[:i]) > 1 {
			res = res + ar[i]
		}
	}
	return res
}

// 5
// 1000000001 1000000002 1000000003 1000000004 1000000005

//5000000015
