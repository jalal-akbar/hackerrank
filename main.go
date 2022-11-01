package main

import (
	"fmt"

	"github.com/jalal-akbar/hackerrank/warmup"
)

func main() {
	// 1. Simple Array Sum
	nums := []int{3, 10}
	sas := warmup.SimpleArraySum(nums)
	fmt.Println("Simple Array Sum : ", sas)
	// 2. Compare The Triplets
	a := []int32{4, 3, 1}
	b := []int32{2, 1, 3}
	ctt := warmup.CompareTriplets(a, b)
	fmt.Println("Compare The Triplet", ctt)
	// 3. A Very Big Sum
	ar := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	avbs := warmup.AVeryBigSum(ar)
	fmt.Println("A Very Big Sum :", avbs)
	// Diagonal Difference
	var arr = [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 8},
	}
	dd := warmup.DiagonalDifference(arr)
	fmt.Println("Diagonal Differemce", dd)
	// Plus Minus
	pm := []int32{1, 1, 0, -1, -1}
	warmup.PlusMinus(pm)

}
