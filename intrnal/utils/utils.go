package utils

import "fmt"

func Chunk(slice []int, chunkSize int) [][]int {
	var chunks [][]int
	sliceLen := len(slice)
	for i := 0; i < sliceLen; i += chunkSize {
		end := i + chunkSize

		if end > sliceLen {
			end = sliceLen
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

// Notice: If a value has several occurrences,
// the random key will be used as its value, and all others will be lost.
// Notice: The outputting key order may not match the incoming key order
func Flip(source map[string]int) map[int]string {
	flippedSource := make(map[int]string)
	for k, v := range source {
		fmt.Println(v, k)

		flippedSource[v] = k
	}

	return flippedSource
}

func Diff(source []int, comparable []int) []int {
	hasMap := make(map[int]int, len(source))
	for i, v := range comparable {
		hasMap[v] = i
	}

	diff := make([]int, 0)
	for _, v := range source {
		if _, ok := hasMap[v]; ok {
			continue
		}

		diff = append(diff, v)
	}

	return diff
}
