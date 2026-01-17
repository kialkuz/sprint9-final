package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElementsZeroSize(t *testing.T) {
	elements, err := generateRandomElements(0)

	require.Empty(t, elements)
	require.Error(t, err)
}

func TestGenerateRandomElementsNegativeSize(t *testing.T) {
	elements, err := generateRandomElements(-10)

	require.Empty(t, elements)
	require.Error(t, err)
}

func TestGenerateRandomElementsPositiveSize(t *testing.T) {
	elements, err := generateRandomElements(10)

	require.NotEmpty(t, elements)
	require.NoError(t, err)
}

func TestMaximumNoData(t *testing.T) {
	var elements []int

	max, err := maximum(elements)

	require.Zero(t, max)
	require.Error(t, err)
}

func TestMaximumSingleValue(t *testing.T) {
	elements := []int{1000}

	max, err := maximum(elements)

	require.Equal(t, max, 1000)
	require.NoError(t, err)
}

func TestMaximumMultipleValues(t *testing.T) {
	elements := []int{1000, 900, 1200, 1400}

	max, err := maximum(elements)

	require.Equal(t, max, 1400)
	require.NoError(t, err)
}

func TestMaxChunksNoData(t *testing.T) {
	var elements []int

	max, err := maximum(elements)

	require.Zero(t, max)
	require.Error(t, err)
}

func TestMaxChunksSingleValue(t *testing.T) {
	elements := []int{1000}

	max, err := maximum(elements)

	require.Equal(t, max, 1000)
	require.NoError(t, err)
}

func TestMaxChunksMultipleValues(t *testing.T) {
	elements := []int{1000, 900, 1200, 1400}

	max, err := maximum(elements)

	require.Equal(t, max, 1400)
	require.NoError(t, err)
}
