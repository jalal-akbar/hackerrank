package easy

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func CompareTriplets(a []int32, b []int32) []int32 {
	var aPoint int32 = 0
	var bPoint int32 = 0

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			aPoint++
		} else if a[i] < b[i] {
			bPoint++
		} else {
			aPoint = aPoint + 0
			bPoint = bPoint + 0
		}
	}

	return []int32{aPoint, bPoint}

}
