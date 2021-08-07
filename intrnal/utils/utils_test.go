package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldChunkSlice(t *testing.T) {
	assert := assert.New(t)

	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := Chunk(source, 2)
	assert.Len(result, 5)

	expected := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
		{9},
	}
	assert.Equal(expected, result)
}

func TestShouldFlipWithoutDuplicates(t *testing.T) {
	assert := assert.New(t)

	source := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
		"key4": 4,
		"key5": 5,
	}

	expected := map[int]string{
		1: "key1",
		2: "key2",
		3: "key3",
		4: "key4",
		5: "key5",
	}

	result := Flip(source)

	assert.Equal(expected, result)
}

func TestShouldFlipWithDuplicates(t *testing.T) {
	assert := assert.New(t)

	source := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
		"key4": 3,
		"key5": 3,
	}

	result := Flip(source)

	assert.Len(result, 3)

	keys := make([]int, 0)
	for k := range result {
		keys = append(keys, k)
	}

	assert.Contains(keys, 1)
	assert.Contains(keys, 2)
	assert.Contains(keys, 3)
}

func TestShouldDiffSlices(t *testing.T) {
	assert := assert.New(t)

	source := []int{1, 2, 3, 4, 5}
	comparable := []int{3, 4}

	expected := []int{1, 2, 5}

	result := Diff(source, comparable)

	assert.Equal(expected, result)
}

func TestShouldDiffSlicesWithSourceDuplicates(t *testing.T) {
	assert := assert.New(t)

	source := []int{1, 2, 3, 4, 5, 5}
	comparable := []int{3, 4}

	expected := []int{1, 2, 5, 5}

	result := Diff(source, comparable)

	assert.Equal(expected, result)
}

func TestShouldDiffSlicesWithComparableDuplicates(t *testing.T) {
	assert := assert.New(t)

	source := []int{1, 2, 3, 4, 5}
	comparable := []int{3, 4, 4}

	expected := []int{1, 2, 5}

	result := Diff(source, comparable)

	assert.Equal(expected, result)
}
