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

// Run the Fibonacci recursor and return a function that returns a string of uint64
// that represents the Fibonacci sequence to the {maxIndex}th index
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

// Rationale and problem-solving
//
// As this was meant to be a whiteboardish assessment, I would like to use this
// section to walk you through the process that I used.
//
// When I see 'Fibonacci', I grin with the same giddiness that I did the first
// time I finally grokked the concept of recursion. Recursion makes a relatively
// quick job out of the Fibonacci series. However, recursive calls greatly increase
// overhead. I decided that I wanted to show off my recursion chops, and my optimization
// chops, to boot.
//
// This solution crawls a binary tree to calculate a number at Fibonacci[n]
// by adding the sums of the tree nodes, working backwards from the greatest Fibonacci index
// to the smallest. As nodes split out left and right, there will be repeated calculations.
// To reduce this overhead, a cache is employed. Each node is checked against the cache. If
// the number is already cached, no calculation is required - the cached value is returned.
// In the event of a cache miss, the new Fibonacci number is calculated and cached.
//
// Now, I had a pretty nice way to tell me what the /last/ number in the Fibonacci sequence was.
// However, the requirements state that I must /expand/ the sequence up to n, not calculate
// the final number. I employed a buffer in the form of a slice of uint64 to solve this. For every
// cache miss, a new number is pushed to the buffer. At the end of recursion, a function is
// returned that returns the buffer.
//
// The last change that I made was to change all of the int datatypes to uint64 and catch negs.
// Since the logic driving this calculator would have to be altered to return negative Fibonacci
// sequences, this is a modicum of protection. The real solution would be to refactor the code
// to gracefully handle negative Fibonacci sequences, but that was not a requirement.
//
// Finally, to make sure everything was wired properly, I wrote tests.
