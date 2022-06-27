package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"testing"
)

// Read the Fibonacci sequence to the 93rd index and return it as []uint
func populateFib(t *testing.T) ([93]uint, error) {
	fibFile, err := os.Open("fib.csv")
	if err != nil {
		return [93]uint{}, err
	}

	reader := csv.NewReader(fibFile)

	fibSequence, err := reader.ReadAll()
	if err != nil {
		return [93]uint{}, err
	}

	uintFibSequence := [93]uint{}

	for i := 0; i < len(fibSequence); i++ {
		fibNum, err := strconv.Atoi(fibSequence[0][i])
		if err != nil {
			t.Log(err)
			return [93]uint{}, err
		}

		uintFibSequence[i] = uint(fibNum)
	}
	err = fibFile.Close()
	if err != nil {
		return [93]uint{}, err
	}
	return uintFibSequence, nil
}

////////////////////////
// test cases         //
////////////////////////

// Fail if the length of the []uint returned by fib() with
// a maxIndex of i is different than that of Fibonacci[i]
func checkLen(t *testing.T, master *[93]uint, i int) error {
	if i == 0 {
		return nil
	}
	setMaxIndex(uint(i))
	f := fib()
	fibSlice, err := f()
	got := uint(fibSlice[i])
	want := uint(0)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if got != want {
		t.Logf("got %d, expected %d at index %d", got, want, i)
		t.Fail()
	}
	return nil
}

// Fail if the last index of the []uint returned by fib() with
// a maxIndex of i is not equal to Fibonacci[i]
func checkLastNum(t *testing.T, master *[93]uint, i int) error {
	if i == 0 {
		return nil
	}

	setMaxIndex(uint(i))
	f := fib()
	fibSlice, err := f()
	var got uint = master[i-1]
	var want uint = fibSlice[i-1]

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if got != want {
		t.Logf("got %d, expected %d", got, want)
		t.Fail()
	}

	return nil
}

func TestFib(t *testing.T) {
	fibFromCSV, err := populateFib(t)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	tests := []func(t *testing.T, seq *[93]uint, i int) error{
		checkLen,
		checkLastNum,
	}

	for i, tc := range tests {
		tc(t, &fibFromCSV, i)
	}
}
