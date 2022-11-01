package warmup

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */
func DiagonalDifference(arr [][]int32) int32 {
	// 1
	// var leftToRight int32 = 0
	// var rightToLeft int32 = 0
	// for i := 0; i < len(arr); i++ {
	// 	for j := i; j <= i; j++ {
	// 		leftToRight = leftToRight + arr[i][j]
	// 	}
	// 	for k := len(arr) - 1 - i; k >= len(arr)-1-i; k-- {
	// 		rightToLeft = rightToLeft + arr[i][k]
	// 	}
	// }
	// if rightToLeft > leftToRight {
	// 	return rightToLeft - leftToRight
	// }
	// return leftToRight - rightToLeft
	// 2
	var i, j, k, l int
	var leftToRight int32
	var rightToLeft int32
	for i = 0; i < len(arr); i++ {
		for j = i; j <= i; j++ {
			leftToRight = leftToRight + arr[i][j]
		}
	}
	for k = 0; k < len(arr); k++ { // 0 1
		for l = len(arr) - 1 - k; l >= len(arr)-1-k; l-- { // 2
			rightToLeft = rightToLeft + arr[k][l]
		}
	}
	if rightToLeft > leftToRight {
		return rightToLeft - leftToRight
	}
	return leftToRight - rightToLeft
}
