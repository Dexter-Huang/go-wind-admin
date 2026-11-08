package stringutil

import (
	"fmt"
	"math/rand/v2"
	"regexp"
	"strconv"
	"testing"
)

// ****************************** TESTS ********************************************

func TestRandomSeed(t *testing.T) {

	// count, start, end, letters, numbers := 5, 0, 0, true, true
	random := rand.New(rand.NewPCG(10, 0))
	out := "UOGFM"

	// Test 1: Simulating RandomAlphaNumeric(count int)
	if x, _ := RandomSeed(5, 0, 0, true, true, nil, random); x != out {
		t.Errorf("RandomSeed(%v, %v, %v, %v, %v, %v, %v) = %v, want %v", 5, 0, 0, true, true, nil, random, x, out)
	}

	// Test 2: Simulating RandomAlphabetic(count int)
	out = "OUiKn"

	if x, _ := RandomSeed(5, 0, 0, true, false, nil, random); x != out {
		t.Errorf("RandomSeed(%v, %v, %v, %v, %v, %v, %v) = %v, want %v", 5, 0, 0, true, false, nil, random, x, out)
	}

	// Test 3: Simulating RandomNumeric(count int)
	out = "23353"

	if x, _ := RandomSeed(5, 0, 0, false, true, nil, random); x != out {
		t.Errorf("RandomSeed(%v, %v, %v, %v, %v, %v, %v) = %v, want %v", 5, 0, 0, false, true, nil, random, x, out)
	}

	// Test 4: Simulating RandomAscii(count int)
	out = "AEQz9"

	if x, _ := RandomSeed(5, 32, 127, false, false, nil, random); x != out {
		t.Errorf("RandomSeed(%v, %v, %v, %v, %v, %v, %v) = %v, want %v", 5, 32, 127, false, false, nil, random, x, out)
	}

	// Test 5: Simulating RandomSeed(...) with custom chars
	chars := []rune{'1', '2', '3', 'a', 'b', 'c'}
	out = "ab3ab"

	if x, _ := RandomSeed(5, 0, 0, false, false, chars, random); x != out {
		t.Errorf("RandomSeed(%v, %v, %v, %v, %v, %v, %v) = %v, want %v", 5, 0, 0, false, false, chars, random, x, out)
	}

}

// ****************************** EXAMPLES ********************************************

func ExampleRandomSeed() {

	var seed int64 = 10 // If you change this seed #, the random sequence below will change
	random := rand.New(rand.NewPCG(uint64(seed), 0))
	chars := []rune{'1', '2', '3', 'a', 'b', 'c'}

	rand1, _ := RandomSeed(5, 0, 0, true, true, nil, random)      // RandomAlphaNumeric (Alphabets and numbers possible)
	rand2, _ := RandomSeed(5, 0, 0, true, false, nil, random)     // RandomAlphabetic (Only alphabets)
	rand3, _ := RandomSeed(5, 0, 0, false, true, nil, random)     // RandomNumeric (Only numbers)
	rand4, _ := RandomSeed(5, 32, 127, false, false, nil, random) // RandomAscii (Alphabets, numbers, and other ASCII chars)
	rand5, _ := RandomSeed(5, 0, 0, true, true, chars, random)    // RandomSeed with custom characters

	fmt.Println(rand1)
	fmt.Println(rand2)
	fmt.Println(rand3)
	fmt.Println(rand4)
	fmt.Println(rand5)
	// Output:
	// UOGFM
	// OUiKn
	// 23353
	// AEQz9
	// ab3ab
}

func TestRandomAlphaNumeric(t *testing.T) {
	for i := 0; i < 16; i++ {
		out, _ := RandomAlphaNumeric(20)
		fmt.Println(out)
	}
}

func TestRandAlphaNumeric_FuzzOnlyNumeric(t *testing.T) {

	// Testing for a reported regression in which some versions produced
	// a predictably small set of chars.
	iters := 1000
	charlen := 0
	for i := 0; i < 16; i++ {
		numOnly := 0
		charlen++
		for i := 0; i < iters; i++ {
			out, err := RandomAlphaNumeric(charlen)
			if err != nil {
				t.Fatal("func failed to produce a random thinger")
			}
			if _, err := strconv.Atoi(out); err == nil {
				numOnly++
			}

			m, err := regexp.MatchString("^[0-9a-zA-Z]+$", out)
			if err != nil {
				t.Fatal(err)
			}
			if !m {
				t.Fatal("Character is not alphanum")
			}
		}

		if numOnly == iters {
			t.Fatalf("Got %d numeric-only random sequences", numOnly)
		}
	}

}

// ====================== Wrapper function tests ======================

func TestRandomNonAlphaNumeric(t *testing.T) {
	for i := 0; i < 100; i++ {
		out, err := RandomNonAlphaNumeric(20)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// RandomNonAlphaNumeric uses letters=false, numbers=false which means no filter,
		// so surrogates may cause the rune count to differ from requested count.
		// Just verify the output is non-empty and no error occurred.
		if len(out) == 0 {
			t.Fatal("expected non-empty output")
		}
	}
}

func TestRandomAscii(t *testing.T) {
	for i := 0; i < 100; i++ {
		out, err := RandomAscii(50)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(out) != 50 {
			t.Fatalf("expected length 50, got %d", len(out))
		}
		for _, ch := range out {
			if ch < 32 || ch > 126 {
				t.Fatalf("character %d out of ASCII printable range [32, 126]", ch)
			}
		}
	}
}

func TestRandomNumeric(t *testing.T) {
	for i := 0; i < 100; i++ {
		out, err := RandomNumeric(30)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(out) != 30 {
			t.Fatalf("expected length 30, got %d", len(out))
		}
		m, err := regexp.MatchString("^[0-9]+$", out)
		if err != nil {
			t.Fatal(err)
		}
		if !m {
			t.Fatalf("expected all digits, got %q", out)
		}
	}
}

func TestRandomAlphabetic(t *testing.T) {
	for i := 0; i < 100; i++ {
		out, err := RandomAlphabetic(30)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(out) != 30 {
			t.Fatalf("expected length 30, got %d", len(out))
		}
		m, err := regexp.MatchString("^[a-zA-Z]+$", out)
		if err != nil {
			t.Fatal(err)
		}
		if !m {
			t.Fatalf("expected all alphabetic, got %q", out)
		}
	}
}

func TestRandomAlphaNumericCustom(t *testing.T) {
	// letters=true, numbers=true => same as RandomAlphaNumeric
	for i := 0; i < 50; i++ {
		out, err := RandomAlphaNumericCustom(20, true, true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(out) != 20 {
			t.Fatalf("expected length 20, got %d", len(out))
		}
		m, _ := regexp.MatchString("^[0-9a-zA-Z]+$", out)
		if !m {
			t.Fatalf("expected alphanumeric, got %q", out)
		}
	}

	// letters=true, numbers=false => same as RandomAlphabetic
	for i := 0; i < 50; i++ {
		out, err := RandomAlphaNumericCustom(20, true, false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		m, _ := regexp.MatchString("^[a-zA-Z]+$", out)
		if !m {
			t.Fatalf("expected alphabetic only, got %q", out)
		}
	}

	// letters=false, numbers=true => same as RandomNumeric
	for i := 0; i < 50; i++ {
		out, err := RandomAlphaNumericCustom(20, false, true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		m, _ := regexp.MatchString("^[0-9]+$", out)
		if !m {
			t.Fatalf("expected numeric only, got %q", out)
		}
	}
}

func TestRandom(t *testing.T) {
	// Test with custom chars via variadic parameter
	chars := []rune{'a', 'b', 'c', '1', '2', '3'}
	for i := 0; i < 50; i++ {
		out, err := Random(10, 0, 0, false, false, chars...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(out) != 10 {
			t.Fatalf("expected length 10, got %d", len(out))
		}
		m, _ := regexp.MatchString("^[abc123]+$", out)
		if !m {
			t.Fatalf("expected only [abc123], got %q", out)
		}
	}

	// Test without custom chars
	out, err := Random(5, 0, 0, true, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out) != 5 {
		t.Fatalf("expected length 5, got %d", len(out))
	}
}

// ====================== Error path tests ======================

func TestRandomSeed_CountZero(t *testing.T) {
	random := rand.New(rand.NewPCG(1, 0))
	out, err := RandomSeed(0, 0, 0, true, true, nil, random)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "" {
		t.Fatalf("expected empty string, got %q", out)
	}
}

func TestRandomSeed_CountNegative(t *testing.T) {
	random := rand.New(rand.NewPCG(1, 0))
	_, err := RandomSeed(-1, 0, 0, true, true, nil, random)
	if err == nil {
		t.Fatal("expected error for negative count, got nil")
	}
	expected := "randomstringutils illegal argument: Requested random string length -1 is less than 0"
	if err.Error() != expected {
		t.Fatalf("expected error %q, got %q", expected, err.Error())
	}
}

func TestRandomSeed_EmptyChars(t *testing.T) {
	random := rand.New(rand.NewPCG(1, 0))
	chars := []rune{}
	_, err := RandomSeed(5, 0, 0, true, true, chars, random)
	if err == nil {
		t.Fatal("expected error for empty chars, got nil")
	}
	expected := "randomstringutils illegal argument: The chars array must not be empty"
	if err.Error() != expected {
		t.Fatalf("expected error %q, got %q", expected, err.Error())
	}
}

func TestRandomSeed_EndLessThanOrEqualStart(t *testing.T) {
	random := rand.New(rand.NewPCG(1, 0))

	// end == start
	_, err := RandomSeed(5, 10, 10, false, false, nil, random)
	if err == nil {
		t.Fatal("expected error for end <= start, got nil")
	}
	expected := "randomstringutils illegal argument: Parameter end (10) must be greater than start (10)"
	if err.Error() != expected {
		t.Fatalf("expected error %q, got %q", expected, err.Error())
	}

	// end < start
	_, err = RandomSeed(5, 20, 10, false, false, nil, random)
	if err == nil {
		t.Fatal("expected error for end < start, got nil")
	}
	expected2 := "randomstringutils illegal argument: Parameter end (10) must be greater than start (20)"
	if err.Error() != expected2 {
		t.Fatalf("expected error %q, got %q", expected2, err.Error())
	}
}

func TestRandomSeed_EndGreaterThanCharsLength(t *testing.T) {
	random := rand.New(rand.NewPCG(1, 0))
	chars := []rune{'a', 'b', 'c'}

	_, err := RandomSeed(5, 0, 10, false, false, chars, random)
	if err == nil {
		t.Fatal("expected error for end > len(chars), got nil")
	}
	expected := "randomstringutils illegal argument: Parameter end (10) cannot be greater than len(chars) (3)"
	if err.Error() != expected {
		t.Fatalf("expected error %q, got %q", expected, err.Error())
	}
}

func TestRandomSeed_WithCustomCharsAndStartEnd(t *testing.T) {
	random := rand.New(rand.NewPCG(42, 0))
	chars := []rune{'a', 'b', 'c', 'd', 'e', 'f'}

	// Use start=1, end=4 => should only use chars[1..3] = 'b','c','d'
	out, err := RandomSeed(20, 1, 4, false, false, chars, random)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out) != 20 {
		t.Fatalf("expected length 20, got %d", len(out))
	}
	m, _ := regexp.MatchString("^[bcd]+$", out)
	if !m {
		t.Fatalf("expected only [bcd], got %q", out)
	}
}
