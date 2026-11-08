package stringutil

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"
	"unicode"
)

// RANDOM is a package-level pseudo-random number generator seeded with the current unix nano timestamp.
var RANDOM = rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))

// RandomNonAlphaNumeric generates a random string of the given length containing
// characters that are not filtered by letter or digit checks.
// Equivalent to RandomAlphaNumericCustom(count, false, false).
func RandomNonAlphaNumeric(count int) (string, error) {
	return RandomAlphaNumericCustom(count, false, false)
}

// RandomAscii generates a random string consisting of printable ASCII characters
// (code points 32 through 126).
func RandomAscii(count int) (string, error) {
	return Random(count, 32, 127, false, false)
}

// RandomNumeric generates a random string consisting only of digit characters (0-9).
func RandomNumeric(count int) (string, error) {
	return Random(count, 0, 0, false, true)
}

// RandomAlphabetic generates a random string consisting only of alphabetic characters (a-z, A-Z).
func RandomAlphabetic(count int) (string, error) {
	return Random(count, 0, 0, true, false)
}

// RandomAlphaNumeric generates a random string consisting of alphabetic characters (a-z, A-Z)
// and digits (0-9).
func RandomAlphaNumeric(count int) (string, error) {
	return Random(count, 0, 0, true, true)
}

// RandomAlphaNumericCustom generates a random string with customizable character filters.
// Set letters to true to allow alphabetic characters, numbers to true to allow digits.
// If both are false, no character filtering is applied.
func RandomAlphaNumericCustom(count int, letters bool, numbers bool) (string, error) {
	return Random(count, 0, 0, letters, numbers)
}

// Random generates a random string of the specified length.
//
// Parameters:
//   - count:  the length of the random string to generate
//   - start:  the start of the character range (inclusive); ignored when 0 and chars is nil
//   - end:    the end of the character range (exclusive); ignored when 0 and chars is nil
//   - letters:  whether to allow only letter characters
//   - numbers:  whether to allow only digit characters
//   - chars:    optional custom character set (variadic); when provided, start/end index into this slice
//
// When start and end are both 0 and no chars are provided, the character range defaults to
// [0, MaxInt32) if neither letters nor numbers is true, or [' ', 'z'+1) otherwise.
func Random(count int, start int, end int, letters bool, numbers bool, chars ...rune) (string, error) {
	return RandomSeed(count, start, end, letters, numbers, chars, RANDOM)
}

// RandomSeed generates a random string using a caller-supplied *rand.Rand, making the output
// fully deterministic and reproducible for a given seed.
//
// Parameters:
//   - count:   the length of the random string to generate; must be >= 0
//   - start:   the start of the character range (inclusive); ignored when 0 and chars is nil
//   - end:     the end of the character range (exclusive); ignored when 0 and chars is nil
//   - letters: whether to allow only letter characters
//   - numbers: whether to allow only digit characters
//   - chars:   optional custom character set; when non-nil, start/end index into this slice
//   - random:  the pseudo-random number generator to use
//
// Returns an error if count < 0, chars is non-nil but empty, end <= start (when start/end are
// explicitly provided), or end > len(chars) when chars is provided.
func RandomSeed(count int, start int, end int, letters bool, numbers bool, chars []rune, random *rand.Rand) (string, error) {
	if count == 0 {
		return "", nil
	} else if count < 0 {
		err := fmt.Errorf("randomstringutils illegal argument: Requested random string length %v is less than 0", count) // equiv to err := errors.New("...")
		return "", err
	}
	if chars != nil && len(chars) == 0 {
		err := fmt.Errorf("randomstringutils illegal argument: The chars array must not be empty")
		return "", err
	}

	if start == 0 && end == 0 {
		if chars != nil {
			end = len(chars)
		} else {
			if !letters && !numbers {
				end = math.MaxInt32
			} else {
				end = 'z' + 1
				start = ' '
			}
		}
	} else {
		if end <= start {
			err := fmt.Errorf("randomstringutils illegal argument: Parameter end (%v) must be greater than start (%v)", end, start)
			return "", err
		}

		if chars != nil && end > len(chars) {
			err := fmt.Errorf("randomstringutils illegal argument: Parameter end (%v) cannot be greater than len(chars) (%v)", end, len(chars))
			return "", err
		}
	}

	buffer := make([]rune, count)
	gap := end - start

	// high-surrogates range, (\uD800-\uDBFF) = 55296 - 56319
	//  low-surrogates range, (\uDC00-\uDFFF) = 56320 - 57343

	for count != 0 {
		count--
		var ch rune
		if chars == nil {
			ch = rune(random.IntN(gap) + start)
		} else {
			ch = chars[random.IntN(gap)+start]
		}

		if letters && unicode.IsLetter(ch) || numbers && unicode.IsDigit(ch) || !letters && !numbers {
			if ch >= 56320 && ch <= 57343 { // low surrogate range
				if count == 0 {
					count++
				} else {
					// Insert low surrogate
					buffer[count] = ch
					count--
					// Insert high surrogate
					buffer[count] = rune(55296 + random.IntN(128))
				}
			} else if ch >= 55296 && ch <= 56191 { // High surrogates range (Partial)
				if count == 0 {
					count++
				} else {
					// Insert low surrogate
					buffer[count] = rune(56320 + random.IntN(128))
					count--
					// Insert high surrogate
					buffer[count] = ch
				}
			} else if ch >= 56192 && ch <= 56319 {
				// private high surrogate, skip it
				count++
			} else {
				// not one of the surrogates*
				buffer[count] = ch
			}
		} else {
			count++
		}
	}
	return string(buffer), nil
}
