package main

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
