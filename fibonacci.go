package piscine

// Write a recursive function that returns the value at the position index in the fibonacci sequence.
// The first value is at index 0.
// The sequence starts this way: 0, 1, 1, 2, 3 etc...
// A negative index will return -1.
// for is forbidden for this exercise.

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index <= 1 {
		return index
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
