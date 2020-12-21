package fizzbuzz

import "strconv"

// FizzBuzz performs a FizzBuzz operation over a range of integers
//
// Given a range of integers:
// - Return "Fizz" if the integer is divisible by the `fizzAt` value.
// - Return "Buzz" if the integer is divisible by the `buzzAt` value.
// - Return "FizzBuzz" if the integer is divisible by both the `fizzAt` and
//   `buzzAt` values.
// - Return the original number if is is not divisible by either the `fizzAt` or
//   the `buzzAt` values.
func FizzBuzz(total, fizzAt, buzzAt int64) []string {
	if total < 0 {
		return make([]string, 0)
	}
	result := make([]string, total)

	for i := int64(1); i <= total; i++ {
		fizzDiv := isDivisible(i, fizzAt)
		buzzDiv := isDivisible(i, buzzAt)

		if fizzDiv {
			result[i-1] = "Fizz"
		}

		if buzzDiv {
			result[i-1] += "Buzz"
			continue
		}

		if !fizzDiv && !buzzDiv {
			result[i-1] = strconv.FormatInt(i, 10)
		}
	}

	return result
}

func isDivisible(i, number int64) bool {
	if number == 0 {
		return false
	}

	return i%number == 0
}
