package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	// 1. Creating slices
	var a []int                         // nil slice
	b := []int{1, 2, 3}                 // literal
	c := make([]int, 3)                // [0 0 0]
	d := make([]int, 3, 5)             // len=3, cap=5
	fmt.Println("a (nil):", a)
	fmt.Println("b (literal):", b)
	fmt.Println("c (make):", c)
	fmt.Println("d (len 3, cap 5):", d)

	// 2. Appending elements
	b = append(b, 4, 5)
	fmt.Println("After append:", b)

	// 3. Copying slices
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println("Copied slice:", dst)

	// 4. Slicing from an array
	arr := [5]int{100, 200, 300, 400, 500}
	sliced := arr[1:4]
	fmt.Println("Sliced array [1:4]:", sliced)

	// 5. Sorting slices
	nums := []int{33, 11, 44, 22}
	sort.Ints(nums)
	fmt.Println("Sorted (asc):", nums)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println("Sorted (desc):", nums)

	// 6. Length and capacity
	fmt.Println("len(nums):", len(nums))
	fmt.Println("cap(nums):", cap(nums))

	// 7. Remove element at index 2
	nums = append(nums[:2], nums[3:]...)
	fmt.Println("After removing index 2:", nums)

	// 8. Nil and empty check
	if a == nil {
		fmt.Println("Slice a is nil")
	}
	if len(a) == 0 {
		fmt.Println("Slice a is empty")
	}

	// 9. Iterating over slice
	fmt.Print("Iterating over dst: ")
	for i, v := range dst {
		fmt.Printf("[%d:%d] ", i, v)
	}
	fmt.Println()

	// 10. Comparing slices (Go 1.18+)
	eq := slices.Equal([]int{1, 2, 3}, []int{1, 2, 3})
	fmt.Println("Slices equal:", eq)

	// Bonus: Custom sorting
	custom := []int{7, 3, 9, 1}
	sort.Slice(custom, func(i, j int) bool {
		return custom[i] > custom[j] // descending
	})
	fmt.Println("Custom sorted (desc):", custom)
}
