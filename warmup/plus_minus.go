package warmup

import "fmt"

func PlusMinus(arr []int32) {
	// Write your code here
	l := len(arr)
	pos := 0
	neg := 0
	zer := 0

	for i := 0; i < l; i++ {
		if arr[i] > 0 {
			pos = pos + 1
		} else if arr[i] < 0 {
			neg = neg + 1
		} else {
			zer = zer + 1
		}
	}
	positive := float32(pos) / float32(l)
	negative := float32(neg) / float32(l)
	zero := float32(zer) / float32(l)
	fmt.Printf("%.6f\n", positive)
	fmt.Printf("%.6f\n", negative)
	fmt.Printf("%.6f\n", zero)
}
