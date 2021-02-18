package main

import "fmt"

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	arrayIdx := 0
	seqIdx := 0
	for arrayIdx < len(array) && seqIdx < len(sequence) {
		if array[arrayIdx] == sequence[seqIdx] {
			seqIdx += 1
		}
		arrayIdx += 1
	}
	return seqIdx == len(sequence)
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}
	ok := IsValidSubsequence(array, sequence)
	if ok {
		fmt.Println("true")
		return
	}
	fmt.Println("false")
}
