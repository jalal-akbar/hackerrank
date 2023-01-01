package warmup

import "fmt"

func StairCase(n int32) {
	var i int32
	var j int32
	var k int32
	// for i = 0; i < n; i++ { // 0 1
	// 	for j = 0; j <= i; j++ { // 0
	// 		fmt.Print("#")
	// 	}
	// 	fmt.Println("*")
	// }
	// for i = 0; i < n; i++ { // 0 1
	// 	for j = i; j < n; j++ { // 0
	// 		fmt.Print("*")
	// 	}
	// 	for k = n; k > n-i-1; k-- {
	// 		fmt.Print("#")
	// 	}
	// 	fmt.Println()
	// }
	for i = 0; i < n; i++ {
		for j = 0; j < n-i-1; j++ {
			fmt.Print("*")
		}
		for k = n; k > n-1-i; k-- {
			fmt.Print("#")
		}
		fmt.Println("")
	}
}

// i0=******#
// i1,
// i2,
// i3,
// i4,
// i5,
