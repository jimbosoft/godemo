package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlusMinus(t *testing.T) {
	arr := []int32{-1, -20, 32, 0, 67}

	fmt.Printf("Hello")
	pos, neg, zero, err := plusMinus(arr)
	require.NoError(t, err)
	require.Equal(t, "0.400000", pos, "positive incorrect")
	require.Equal(t, "0.400000", neg, "negative incorrect")
	require.Equal(t, "0.200000", zero, "zero incorrect")
}
