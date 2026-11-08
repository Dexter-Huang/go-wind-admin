package stringutil

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"unicode"
)

// CryptoRandomNonAlphaNumeric generates a cryptographically secure random string of the given length
// containing characters that are not filtered by letter or digit checks.
// Equivalent to CryptoRandomAlphaNumericCustom(count, false, false).
func CryptoRandomNonAlphaNumeric(count int) (string, error) {
	return CryptoRandomAlphaNumericCustom(count, false, false)
}

// CryptoRandomAscii generates a cryptographically secure random string consisting of
// printable ASCII characters (code points 32 through 126).
func CryptoRandomAscii(count int) (string, error) {
	return CryptoRandom(count, 32, 127, false, false)
}

// CryptoRandomNumeric generates a cryptographically secure random string consisting only of
// digit characters (0-9).
func CryptoRandomNumeric(count int) (string, error) {
	return CryptoRandom(count, 0, 0, false, true)
}

// CryptoRandomAlphabetic generates a cryptographically secure random string consisting only of
// alphabetic characters (a-z, A-Z).
func CryptoRandomAlphabetic(count int) (string, error) {
	return CryptoRandom(count, 0, 0, true, false)
}

// CryptoRandomAlphaNumeric generates a cryptographically secure random string consisting of
// alphabetic characters (a-z, A-Z) and digits (0-9).
func CryptoRandomAlphaNumeric(count int) (string, error) {
	return CryptoRandom(count, 0, 0, true, true)
}

// CryptoRandomAlphaNumericCustom generates a cryptographically secure random string with
// customizable character filters.
// Set letters to true to allow alphabetic characters, numbers to true to allow digits.
// If both are false, no character filtering is applied.
func CryptoRandomAlphaNumericCustom(count int, letters bool, numbers bool) (string, error) {
	return CryptoRandom(count, 0, 0, letters, numbers)
}

// CryptoRandom generates a cryptographically secure random string of the specified length
// using crypto/rand as the entropy source.
//
// Parameters:
//   - count:   the length of the random string to generate
//   - start:   the start of the character range (inclusive); ignored when 0 and chars is nil
//   - end:     the end of the character range (exclusive); ignored when 0 and chars is nil
//   - letters: whether to allow only letter characters
//   - numbers: whether to allow only digit characters
//   - chars:   optional custom character set (variadic); when provided, start/end index into this slice
//
// When start and end are both 0 and no chars are provided, the character range defaults to
// [0, MaxInt32) if neither letters nor numbers is true, or [' ', 'z'+1) otherwise.
//
// Returns an error if count < 0, chars is non-nil but empty, end <= start (when start/end are
// explicitly provided), or end > len(chars) when chars is provided.
func CryptoRandom(count int, start int, end int, letters bool, numbers bool, chars ...rune) (string, error) {
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
			ch = rune(getCryptoRandomInt(gap) + int64(start))
		} else {
			ch = chars[getCryptoRandomInt(gap)+int64(start)]
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
					buffer[count] = rune(55296 + getCryptoRandomInt(128))
				}
			} else if ch >= 55296 && ch <= 56191 { // High surrogates range (Partial)
				if count == 0 {
					count++
				} else {
					// Insert low surrogate
					buffer[count] = rune(56320 + getCryptoRandomInt(128))
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

// getCryptoRandomInt returns a cryptographically secure random integer in [0, count).
func getCryptoRandomInt(count int) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(count)))
	if err != nil {
		panic(err)
	}
	return nBig.Int64()
}
