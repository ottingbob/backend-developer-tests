package fizzbuzz

import (
	"testing"
)

const (
	fizz     = "Fizz"
	buzz     = "Buzz"
	fizzBuzz = "FizzBuzz"
)

type TestResults struct {
	TestCase         *testing.T
	RunNum           int64
	FizzNum          int64
	BuzzNum          int64
	CountedNums      int64
	ExpectedNums     int64
	CountedFizz      int64
	ExpectedFizz     int64
	CountedBuzz      int64
	ExpectedBuzz     int64
	CountedFizzBuzz  int64
	ExpectedFizzBuzz int64
}

func calculateExpectedCounts(runNum, fizzNum, buzzNum int64) (int64, int64, int64, int64) {
	var expectedFizzBuzz, expectedFizz, expectedBuzz, expectedNums int64 = 0, 0, 0, 0
	if fizzNum != 0 && buzzNum != 0 {
		expectedFizzBuzz = runNum / (fizzNum * buzzNum)
	}
	if expectedFizzBuzz < 0 {
		expectedFizzBuzz *= -1
	}

	if fizzNum != 0 {
		if fizzNum < 0 {
			fizzNum *= -1
		}
		expectedFizz = (runNum / fizzNum) - expectedFizzBuzz
	}
	if expectedFizz < 0 {
		expectedFizz *= -1
	}

	if buzzNum != 0 {
		if buzzNum < 0 {
			buzzNum *= -1
		}
		expectedBuzz = (runNum / buzzNum) - expectedFizzBuzz
	}
	if expectedBuzz < 0 {
		expectedBuzz *= -1
	}

	expectedNums = runNum - (expectedFizzBuzz + expectedFizz + expectedBuzz)

	return expectedFizzBuzz, expectedFizz, expectedBuzz, expectedNums
}

func calculateCounts(results []string) (int64, int64, int64, int64) {
	var countedFizz, countedBuzz, countedFizzBuzz, countedNums int64 = 0, 0, 0, 0

	for _, val := range results {
		switch val {
		case fizz:
			countedFizz++
		case buzz:
			countedBuzz++
		case fizzBuzz:
			countedFizzBuzz++
		default:
			countedNums++
		}
	}

	return countedFizz, countedBuzz, countedFizzBuzz, countedNums
}

func assertCounts(tr *TestResults) {
	if tr.CountedNums != tr.ExpectedNums {
		tr.TestCase.Fatalf(`FizzBuzz(%d, %d, %d). Received [%d] counted numbers and expected [%d]`, tr.RunNum, tr.FizzNum, tr.BuzzNum, tr.CountedNums, tr.ExpectedNums)
	} else if tr.CountedFizz != tr.ExpectedFizz {
		tr.TestCase.Fatalf(`FizzBuzz(%d, %d, %d). Received [%d] counted Fizz and expected [%d]`, tr.RunNum, tr.FizzNum, tr.BuzzNum, tr.CountedFizz, tr.ExpectedFizz)
	} else if tr.CountedBuzz != tr.ExpectedBuzz {
		tr.TestCase.Fatalf(`FizzBuzz(%d, %d, %d). Received [%d] counted Buzz and expected [%d]`, tr.RunNum, tr.FizzNum, tr.BuzzNum, tr.CountedBuzz, tr.ExpectedBuzz)
	} else if tr.CountedFizzBuzz != tr.ExpectedFizzBuzz {
		tr.TestCase.Fatalf(`FizzBuzz(%d, %d, %d). Received [%d] counted FizzBuzz and expected [%d]`, tr.RunNum, tr.FizzNum, tr.BuzzNum, tr.CountedFizzBuzz, tr.ExpectedFizzBuzz)
	}
}

func TestFizzBuzz_SmallInputs(t *testing.T) {
	var runNum, fizzNum, buzzNum int64 = 16, 3, 5
	expectedFizzBuzz, expectedFizz, expectedBuzz, expectedNums := calculateExpectedCounts(runNum, fizzNum, buzzNum)
	t.Logf("Expected Fizz: %d, Expected Buzz: %d, Expected FB: %d, Expected Nums: %d", expectedFizz, expectedBuzz, expectedFizzBuzz, expectedNums)

	results := FizzBuzz(runNum, fizzNum, buzzNum)

	countedFizz, countedBuzz, countedFizzBuzz, countedNums := calculateCounts(results)
	tr := &TestResults{
		TestCase:         t,
		RunNum:           runNum,
		FizzNum:          fizzNum,
		BuzzNum:          buzzNum,
		CountedNums:      countedNums,
		ExpectedNums:     expectedNums,
		CountedFizz:      countedFizz,
		ExpectedFizz:     expectedFizz,
		CountedBuzz:      countedBuzz,
		ExpectedBuzz:     expectedBuzz,
		CountedFizzBuzz:  countedFizzBuzz,
		ExpectedFizzBuzz: expectedFizzBuzz,
	}
	assertCounts(tr)
}

func TestFizzBuzz_LargeInputs(t *testing.T) {
	var runNum, fizzNum, buzzNum int64 = 1024, 13, 37
	expectedFizzBuzz, expectedFizz, expectedBuzz, expectedNums := calculateExpectedCounts(runNum, fizzNum, buzzNum)
	t.Logf("Expected Fizz: %d, Expected Buzz: %d, Expected FB: %d, Expected Nums: %d", expectedFizz, expectedBuzz, expectedFizzBuzz, expectedNums)

	results := FizzBuzz(runNum, fizzNum, buzzNum)

	countedFizz, countedBuzz, countedFizzBuzz, countedNums := calculateCounts(results)
	tr := &TestResults{
		TestCase:         t,
		RunNum:           runNum,
		FizzNum:          fizzNum,
		BuzzNum:          buzzNum,
		CountedNums:      countedNums,
		ExpectedNums:     expectedNums,
		CountedFizz:      countedFizz,
		ExpectedFizz:     expectedFizz,
		CountedBuzz:      countedBuzz,
		ExpectedBuzz:     expectedBuzz,
		CountedFizzBuzz:  countedFizzBuzz,
		ExpectedFizzBuzz: expectedFizzBuzz,
	}
	assertCounts(tr)
}

func TestFizzBuzz_LargerInputs(t *testing.T) {
	var runNum, fizzNum, buzzNum int64 = 331024, 23, 137
	expectedFizzBuzz, expectedFizz, expectedBuzz, expectedNums := calculateExpectedCounts(runNum, fizzNum, buzzNum)
	t.Logf("Expected Fizz: %d, Expected Buzz: %d, Expected FB: %d, Expected Nums: %d", expectedFizz, expectedBuzz, expectedFizzBuzz, expectedNums)

	results := FizzBuzz(runNum, fizzNum, buzzNum)

	countedFizz, countedBuzz, countedFizzBuzz, countedNums := calculateCounts(results)
	tr := &TestResults{
		TestCase:         t,
		RunNum:           runNum,
		FizzNum:          fizzNum,
		BuzzNum:          buzzNum,
		CountedNums:      countedNums,
		ExpectedNums:     expectedNums,
		CountedFizz:      countedFizz,
		ExpectedFizz:     expectedFizz,
		CountedBuzz:      countedBuzz,
		ExpectedBuzz:     expectedBuzz,
		CountedFizzBuzz:  countedFizzBuzz,
		ExpectedFizzBuzz: expectedFizzBuzz,
	}
	assertCounts(tr)
}

func TestFizzBuzz_ZeroRuns(t *testing.T) {
	var runNum, fizzNum, buzzNum int64 = 0, 12, 137
	expectedFizzBuzz, expectedFizz, expectedBuzz, expectedNums := calculateExpectedCounts(runNum, fizzNum, buzzNum)
	t.Logf("Expected Fizz: %d, Expected Buzz: %d, Expected FB: %d, Expected Nums: %d", expectedFizz, expectedBuzz, expectedFizzBuzz, expectedNums)

	results := FizzBuzz(runNum, fizzNum, buzzNum)

	countedFizz, countedBuzz, countedFizzBuzz, countedNums := calculateCounts(results)
	tr := &TestResults{
		TestCase:         t,
		RunNum:           runNum,
		FizzNum:          fizzNum,
		BuzzNum:          buzzNum,
		CountedNums:      countedNums,
		ExpectedNums:     expectedNums,
		CountedFizz:      countedFizz,
		ExpectedFizz:     expectedFizz,
		CountedBuzz:      countedBuzz,
		ExpectedBuzz:     expectedBuzz,
		CountedFizzBuzz:  countedFizzBuzz,
		ExpectedFizzBuzz: expectedFizzBuzz,
	}
	assertCounts(tr)
}

func TestFizzBuzz_ZeroFizz(t *testing.T) {
	var runNum, fizzNum, buzzNum int64 = 1000, 0, 13
	expectedFizzBuzz, expectedFizz, expectedBuzz, expectedNums := calculateExpectedCounts(runNum, fizzNum, buzzNum)
	t.Logf("Expected Fizz: %d, Expected Buzz: %d, Expected FB: %d, Expected Nums: %d", expectedFizz, expectedBuzz, expectedFizzBuzz, expectedNums)

	results := FizzBuzz(runNum, fizzNum, buzzNum)

	countedFizz, countedBuzz, countedFizzBuzz, countedNums := calculateCounts(results)
	tr := &TestResults{
		TestCase:         t,
		RunNum:           runNum,
		FizzNum:          fizzNum,
		BuzzNum:          buzzNum,
		CountedNums:      countedNums,
		ExpectedNums:     expectedNums,
		CountedFizz:      countedFizz,
		ExpectedFizz:     expectedFizz,
		CountedBuzz:      countedBuzz,
		ExpectedBuzz:     expectedBuzz,
		CountedFizzBuzz:  countedFizzBuzz,
		ExpectedFizzBuzz: expectedFizzBuzz,
	}
	assertCounts(tr)
}

func TestFizzBuzz_ZeroBuzz(t *testing.T) {
	var runNum, fizzNum, buzzNum int64 = 1000, 20, 0
	expectedFizzBuzz, expectedFizz, expectedBuzz, expectedNums := calculateExpectedCounts(runNum, fizzNum, buzzNum)
	t.Logf("Expected Fizz: %d, Expected Buzz: %d, Expected FB: %d, Expected Nums: %d", expectedFizz, expectedBuzz, expectedFizzBuzz, expectedNums)

	results := FizzBuzz(runNum, fizzNum, buzzNum)

	countedFizz, countedBuzz, countedFizzBuzz, countedNums := calculateCounts(results)
	tr := &TestResults{
		TestCase:         t,
		RunNum:           runNum,
		FizzNum:          fizzNum,
		BuzzNum:          buzzNum,
		CountedNums:      countedNums,
		ExpectedNums:     expectedNums,
		CountedFizz:      countedFizz,
		ExpectedFizz:     expectedFizz,
		CountedBuzz:      countedBuzz,
		ExpectedBuzz:     expectedBuzz,
		CountedFizzBuzz:  countedFizzBuzz,
		ExpectedFizzBuzz: expectedFizzBuzz,
	}
	assertCounts(tr)
}

func TestFizzBuzz_Zero_Fizz_Zero_Buzz(t *testing.T) {
	var runNum, fizzNum, buzzNum int64 = 1000, 0, 0
	expectedFizzBuzz, expectedFizz, expectedBuzz, expectedNums := calculateExpectedCounts(runNum, fizzNum, buzzNum)
	t.Logf("Expected Fizz: %d, Expected Buzz: %d, Expected FB: %d, Expected Nums: %d", expectedFizz, expectedBuzz, expectedFizzBuzz, expectedNums)

	results := FizzBuzz(runNum, fizzNum, buzzNum)

	countedFizz, countedBuzz, countedFizzBuzz, countedNums := calculateCounts(results)
	tr := &TestResults{
		TestCase:         t,
		RunNum:           runNum,
		FizzNum:          fizzNum,
		BuzzNum:          buzzNum,
		CountedNums:      countedNums,
		ExpectedNums:     expectedNums,
		CountedFizz:      countedFizz,
		ExpectedFizz:     expectedFizz,
		CountedBuzz:      countedBuzz,
		ExpectedBuzz:     expectedBuzz,
		CountedFizzBuzz:  countedFizzBuzz,
		ExpectedFizzBuzz: expectedFizzBuzz,
	}
	assertCounts(tr)
}

func TestFizzBuzz_Fizz_Out_Of_Range(t *testing.T) {
	var runNum, fizzNum, buzzNum int64 = 100, 101, 3
	expectedFizzBuzz, expectedFizz, expectedBuzz, expectedNums := calculateExpectedCounts(runNum, fizzNum, buzzNum)
	t.Logf("Expected Fizz: %d, Expected Buzz: %d, Expected FB: %d, Expected Nums: %d", expectedFizz, expectedBuzz, expectedFizzBuzz, expectedNums)

	results := FizzBuzz(runNum, fizzNum, buzzNum)

	countedFizz, countedBuzz, countedFizzBuzz, countedNums := calculateCounts(results)
	tr := &TestResults{
		TestCase:         t,
		RunNum:           runNum,
		FizzNum:          fizzNum,
		BuzzNum:          buzzNum,
		CountedNums:      countedNums,
		ExpectedNums:     expectedNums,
		CountedFizz:      countedFizz,
		ExpectedFizz:     expectedFizz,
		CountedBuzz:      countedBuzz,
		ExpectedBuzz:     expectedBuzz,
		CountedFizzBuzz:  countedFizzBuzz,
		ExpectedFizzBuzz: expectedFizzBuzz,
	}
	assertCounts(tr)
}

func TestFizzBuzz_Buzz_Out_Of_Range(t *testing.T) {
	var runNum, fizzNum, buzzNum int64 = 100, 11, 300
	expectedFizzBuzz, expectedFizz, expectedBuzz, expectedNums := calculateExpectedCounts(runNum, fizzNum, buzzNum)
	t.Logf("Expected Fizz: %d, Expected Buzz: %d, Expected FB: %d, Expected Nums: %d", expectedFizz, expectedBuzz, expectedFizzBuzz, expectedNums)

	results := FizzBuzz(runNum, fizzNum, buzzNum)

	countedFizz, countedBuzz, countedFizzBuzz, countedNums := calculateCounts(results)
	tr := &TestResults{
		TestCase:         t,
		RunNum:           runNum,
		FizzNum:          fizzNum,
		BuzzNum:          buzzNum,
		CountedNums:      countedNums,
		ExpectedNums:     expectedNums,
		CountedFizz:      countedFizz,
		ExpectedFizz:     expectedFizz,
		CountedBuzz:      countedBuzz,
		ExpectedBuzz:     expectedBuzz,
		CountedFizzBuzz:  countedFizzBuzz,
		ExpectedFizzBuzz: expectedFizzBuzz,
	}
	assertCounts(tr)
}

func TestFizzBuzz_Fizz_And_Buzz_Out_Of_Range(t *testing.T) {
	var runNum, fizzNum, buzzNum int64 = 100, 111, 300
	expectedFizzBuzz, expectedFizz, expectedBuzz, expectedNums := calculateExpectedCounts(runNum, fizzNum, buzzNum)
	t.Logf("Expected Fizz: %d, Expected Buzz: %d, Expected FB: %d, Expected Nums: %d", expectedFizz, expectedBuzz, expectedFizzBuzz, expectedNums)

	results := FizzBuzz(runNum, fizzNum, buzzNum)

	countedFizz, countedBuzz, countedFizzBuzz, countedNums := calculateCounts(results)
	tr := &TestResults{
		TestCase:         t,
		RunNum:           runNum,
		FizzNum:          fizzNum,
		BuzzNum:          buzzNum,
		CountedNums:      countedNums,
		ExpectedNums:     expectedNums,
		CountedFizz:      countedFizz,
		ExpectedFizz:     expectedFizz,
		CountedBuzz:      countedBuzz,
		ExpectedBuzz:     expectedBuzz,
		CountedFizzBuzz:  countedFizzBuzz,
		ExpectedFizzBuzz: expectedFizzBuzz,
	}
	assertCounts(tr)
}

func TestFizzBuzz_Negative_Runs(t *testing.T) {
	var runNum, fizzNum, buzzNum int64 = -100, 11, 3
	var expectedFizzBuzz, expectedFizz, expectedBuzz, expectedNums int64 = 0, 0, 0, 0
	t.Logf("Expected Fizz: %d, Expected Buzz: %d, Expected FB: %d, Expected Nums: %d", expectedFizz, expectedBuzz, expectedFizzBuzz, expectedNums)

	results := FizzBuzz(runNum, fizzNum, buzzNum)

	countedFizz, countedBuzz, countedFizzBuzz, countedNums := calculateCounts(results)
	tr := &TestResults{
		TestCase:         t,
		RunNum:           runNum,
		FizzNum:          fizzNum,
		BuzzNum:          buzzNum,
		CountedNums:      countedNums,
		ExpectedNums:     expectedNums,
		CountedFizz:      countedFizz,
		ExpectedFizz:     expectedFizz,
		CountedBuzz:      countedBuzz,
		ExpectedBuzz:     expectedBuzz,
		CountedFizzBuzz:  countedFizzBuzz,
		ExpectedFizzBuzz: expectedFizzBuzz,
	}
	assertCounts(tr)
}

func TestFizzBuzz_Negative_Fizz(t *testing.T) {
	var runNum, fizzNum, buzzNum int64 = 16, -3, 5
	expectedFizzBuzz, expectedFizz, expectedBuzz, expectedNums := calculateExpectedCounts(runNum, fizzNum, buzzNum)
	t.Logf("Expected Fizz: %d, Expected Buzz: %d, Expected FB: %d, Expected Nums: %d", expectedFizz, expectedBuzz, expectedFizzBuzz, expectedNums)

	results := FizzBuzz(runNum, fizzNum, buzzNum)

	countedFizz, countedBuzz, countedFizzBuzz, countedNums := calculateCounts(results)
	tr := &TestResults{
		TestCase:         t,
		RunNum:           runNum,
		FizzNum:          fizzNum,
		BuzzNum:          buzzNum,
		CountedNums:      countedNums,
		ExpectedNums:     expectedNums,
		CountedFizz:      countedFizz,
		ExpectedFizz:     expectedFizz,
		CountedBuzz:      countedBuzz,
		ExpectedBuzz:     expectedBuzz,
		CountedFizzBuzz:  countedFizzBuzz,
		ExpectedFizzBuzz: expectedFizzBuzz,
	}
	assertCounts(tr)
}

func TestFizzBuzz_Negative_Buzz(t *testing.T) {
	var runNum, fizzNum, buzzNum int64 = 102, 7, -5
	expectedFizzBuzz, expectedFizz, expectedBuzz, expectedNums := calculateExpectedCounts(runNum, fizzNum, buzzNum)
	t.Logf("Expected Fizz: %d, Expected Buzz: %d, Expected FB: %d, Expected Nums: %d", expectedFizz, expectedBuzz, expectedFizzBuzz, expectedNums)

	results := FizzBuzz(runNum, fizzNum, buzzNum)

	countedFizz, countedBuzz, countedFizzBuzz, countedNums := calculateCounts(results)
	tr := &TestResults{
		TestCase:         t,
		RunNum:           runNum,
		FizzNum:          fizzNum,
		BuzzNum:          buzzNum,
		CountedNums:      countedNums,
		ExpectedNums:     expectedNums,
		CountedFizz:      countedFizz,
		ExpectedFizz:     expectedFizz,
		CountedBuzz:      countedBuzz,
		ExpectedBuzz:     expectedBuzz,
		CountedFizzBuzz:  countedFizzBuzz,
		ExpectedFizzBuzz: expectedFizzBuzz,
	}
	assertCounts(tr)
}

func TestFizzBuzz_Negative_Fizz_And_Buzz(t *testing.T) {
	var runNum, fizzNum, buzzNum int64 = 1020, -3, -5
	expectedFizzBuzz, expectedFizz, expectedBuzz, expectedNums := calculateExpectedCounts(runNum, fizzNum, buzzNum)
	t.Logf("Expected Fizz: %d, Expected Buzz: %d, Expected FB: %d, Expected Nums: %d", expectedFizz, expectedBuzz, expectedFizzBuzz, expectedNums)

	results := FizzBuzz(runNum, fizzNum, buzzNum)

	countedFizz, countedBuzz, countedFizzBuzz, countedNums := calculateCounts(results)
	tr := &TestResults{
		TestCase:         t,
		RunNum:           runNum,
		FizzNum:          fizzNum,
		BuzzNum:          buzzNum,
		CountedNums:      countedNums,
		ExpectedNums:     expectedNums,
		CountedFizz:      countedFizz,
		ExpectedFizz:     expectedFizz,
		CountedBuzz:      countedBuzz,
		ExpectedBuzz:     expectedBuzz,
		CountedFizzBuzz:  countedFizzBuzz,
		ExpectedFizzBuzz: expectedFizzBuzz,
	}
	assertCounts(tr)
}
