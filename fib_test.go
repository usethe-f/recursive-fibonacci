package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"testing"
)

// Read the Fibonacci sequence to the 93rd index and return it as []uint64
func populateFib(t *testing.T) ([93]uint64, error) {
	fibFile, err := os.Open("fib.csv")
	if err != nil {
		return [93]uint64{}, err
	}

	reader := csv.NewReader(fibFile)

	fibSequence, err := reader.ReadAll()
	if err != nil {
		return [93]uint64{}, err
	}

	uint64FibSequence := [93]uint64{}

	for i := 0; i < len(fibSequence); i++ {
		fibNum, err := strconv.Atoi(fibSequence[0][i])
		if err != nil {
			t.Log(err)
			return [93]uint64{}, err
		}

		uint64FibSequence[i] = uint64(fibNum)
	}
	err = fibFile.Close()
	if err != nil {
		return [93]uint64{}, err
	}
	return uint64FibSequence, nil
}

////////////////////////
// test cases         //
////////////////////////

// Fail if the length of the []uint64 returned by fib() with
// a maxIndex of i is different than that of Fibonacci[i]
func checkLen(t *testing.T, master *[93]uint64, i int) error {
	if i == 0 {
		return nil
	}
	setMaxIndex(uint64(i))
	f := fib()
	fibSlice, err := f()
	got := uint64(fibSlice[i])
	want := uint64(0)

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

// Fail if the last index of the []uint64 returned by fib() with
// a maxIndex of i is not equal to Fibonacci[i]
func checkLastNum(t *testing.T, master *[93]uint64, i int) error {
	if i == 0 {
		return nil
	}

	setMaxIndex(uint64(i))
	f := fib()
	fibSlice, err := f()
	var got uint64 = master[i-1]
	var want uint64 = fibSlice[i-1]

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

	tests := []func(t *testing.T, seq *[93]uint64, i int) error{
		checkLen,
		checkLastNum,
	}

	for i, tc := range tests {
		tc(t, &fibFromCSV, i)
	}
}
