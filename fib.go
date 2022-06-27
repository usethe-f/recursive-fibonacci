package main

import (
	"errors"
	"fmt"
)

// Create a variable to represent the max Fibonacci index
var maxIndex uint64 = 29

// Create a mechanism to set maxIndex
func setMaxIndex(newValue uint64) {
	err := validateUint64(newValue)
	if err != nil {
		return
	}
	maxIndex = newValue
}

// Create a buffer to hold Fibonacci numbers sequentially
var seq = []uint64{1} // 1 is the base number

// Create a cache to hold fibonacci numbers that have been seen before
// so that recursion doesn't cause extra complexity
var cache = make(map[uint64]uint64)

// Ensure unsigned integer
func validateUint64(check uint64) error {
	var chk interface{}
	chk = check
	switch chk.(type) {
	case uint64:
		return nil
	default:
		return errors.New(fmt.Sprintf("expected uint64, got %T", check))
	}
}

// Run the Fibonacci recursor and return a function that returns a []uint64
// that represents the Fibonacci sequence to the {maxIndex}th index, and an
// error if one is present
func fib() func() ([]uint64, error) {
	// TODO: refactor to allow for negative Fibonacci sequences
	err := validateUint64(maxIndex)

	if err == nil {
		fibRecursor(maxIndex)
	}

	return func() ([]uint64, error) { return seq, nil }
}

// fibRecursor returns a single Fibonacci number recursively, while
// adding each intermediary number to a buffer
func fibRecursor(num uint64) uint64 {
	// If the recursion reaches full depth, return the number
	if num <= 1 {
		return num
	}

	// If the cache has already seen this number, return it
	if res, ok := cache[num]; ok {
		return res
	}

	var newNum uint64 = 0

	// If there is a cache miss, calculate the value, cache it, and add it to the buffer
	newNum = fibRecursor(num-1) + fibRecursor(num-2)
	cache[num] = newNum
	seq = append(seq, newNum)

	// Return the next Fibonacci number
	return newNum
}

func main() {
	f := fib()
	res, err := f()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
