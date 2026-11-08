package stringutil

import (
	"strconv"
	"strings"
)

// BoolToString converts bool to string.
func BoolToString(v bool) string {
	return strconv.FormatBool(v)
}

// IntToString converts int to string.
func IntToString(v int) string {
	return strconv.FormatInt(int64(v), 10)
}

// Int8ToString converts int8 to string.
func Int8ToString(v int8) string {
	return strconv.FormatInt(int64(v), 10)
}

// Int16ToString converts int16 to string.
func Int16ToString(v int16) string {
	return strconv.FormatInt(int64(v), 10)
}

// Int32ToString converts int32 to string.
func Int32ToString(v int32) string {
	return strconv.FormatInt(int64(v), 10)
}

// Int64ToString converts int64 to string.
func Int64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

// UintToString converts uint to string.
func UintToString(v uint) string {
	return strconv.FormatUint(uint64(v), 10)
}

// Uint8ToString converts uint8 to string.
func Uint8ToString(v uint8) string {
	return strconv.FormatUint(uint64(v), 10)
}

// Uint16ToString converts uint16 to string.
func Uint16ToString(v uint16) string {
	return strconv.FormatUint(uint64(v), 10)
}

// Uint32ToString converts uint32 to string.
func Uint32ToString(v uint32) string {
	return strconv.FormatUint(uint64(v), 10)
}

// Uint64ToString converts uint64 to string.
func Uint64ToString(v uint64) string {
	return strconv.FormatUint(v, 10)
}

// Float32ToString converts float32 to string.
func Float32ToString(v float32) string {
	return strconv.FormatFloat(float64(v), 'f', -1, 32)
}

// Float64ToString converts float64 to string.
func Float64ToString(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

// DoubleToString is an alias for Float64ToString.
func DoubleToString(v float64) string {
	return Float64ToString(v)
}

// StringToBool converts string to bool.
func StringToBool(s string) (bool, error) {
	return strconv.ParseBool(strings.TrimSpace(s))
}

// StringToInt converts string to int.
func StringToInt(s string) (int, error) {
	v, err := strconv.ParseInt(strings.TrimSpace(s), 10, 0)
	return int(v), err
}

// StringToIntBase converts string to int using the given base.
// base can be 2 to 36, or 0 to infer base from string prefix.
func StringToIntBase(s string, base int) (int, error) {
	v, err := strconv.ParseInt(strings.TrimSpace(s), base, 0)
	return int(v), err
}

// StringToInt8 converts string to int8.
func StringToInt8(s string) (int8, error) {
	v, err := strconv.ParseInt(strings.TrimSpace(s), 10, 8)
	return int8(v), err
}

// StringToInt8Base converts string to int8 using the given base.
func StringToInt8Base(s string, base int) (int8, error) {
	v, err := strconv.ParseInt(strings.TrimSpace(s), base, 8)
	return int8(v), err
}

// StringToInt16 converts string to int16.
func StringToInt16(s string) (int16, error) {
	v, err := strconv.ParseInt(strings.TrimSpace(s), 10, 16)
	return int16(v), err
}

// StringToInt16Base converts string to int16 using the given base.
func StringToInt16Base(s string, base int) (int16, error) {
	v, err := strconv.ParseInt(strings.TrimSpace(s), base, 16)
	return int16(v), err
}

// StringToInt32 converts string to int32.
func StringToInt32(s string) (int32, error) {
	v, err := strconv.ParseInt(strings.TrimSpace(s), 10, 32)
	return int32(v), err
}

// StringToInt32Base converts string to int32 using the given base.
func StringToInt32Base(s string, base int) (int32, error) {
	v, err := strconv.ParseInt(strings.TrimSpace(s), base, 32)
	return int32(v), err
}

// StringToInt64 converts string to int64.
func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(strings.TrimSpace(s), 10, 64)
}

// StringToInt64Base converts string to int64 using the given base.
func StringToInt64Base(s string, base int) (int64, error) {
	return strconv.ParseInt(strings.TrimSpace(s), base, 64)
}

// StringToUint converts string to uint.
func StringToUint(s string) (uint, error) {
	v, err := strconv.ParseUint(strings.TrimSpace(s), 10, 0)
	return uint(v), err
}

// StringToUintBase converts string to uint using the given base.
// base can be 2 to 36, or 0 to infer base from string prefix.
func StringToUintBase(s string, base int) (uint, error) {
	v, err := strconv.ParseUint(strings.TrimSpace(s), base, 0)
	return uint(v), err
}

// StringToUint8 converts string to uint8.
func StringToUint8(s string) (uint8, error) {
	v, err := strconv.ParseUint(strings.TrimSpace(s), 10, 8)
	return uint8(v), err
}

// StringToUint8Base converts string to uint8 using the given base.
func StringToUint8Base(s string, base int) (uint8, error) {
	v, err := strconv.ParseUint(strings.TrimSpace(s), base, 8)
	return uint8(v), err
}

// StringToUint16 converts string to uint16.
func StringToUint16(s string) (uint16, error) {
	v, err := strconv.ParseUint(strings.TrimSpace(s), 10, 16)
	return uint16(v), err
}

// StringToUint16Base converts string to uint16 using the given base.
func StringToUint16Base(s string, base int) (uint16, error) {
	v, err := strconv.ParseUint(strings.TrimSpace(s), base, 16)
	return uint16(v), err
}

// StringToUint32 converts string to uint32.
func StringToUint32(s string) (uint32, error) {
	v, err := strconv.ParseUint(strings.TrimSpace(s), 10, 32)
	return uint32(v), err
}

// StringToUint32Base converts string to uint32 using the given base.
func StringToUint32Base(s string, base int) (uint32, error) {
	v, err := strconv.ParseUint(strings.TrimSpace(s), base, 32)
	return uint32(v), err
}

// StringToUint64 converts string to uint64.
func StringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(strings.TrimSpace(s), 10, 64)
}

// StringToUint64Base converts string to uint64 using the given base.
func StringToUint64Base(s string, base int) (uint64, error) {
	return strconv.ParseUint(strings.TrimSpace(s), base, 64)
}

// StringToFloat32 converts string to float32.
func StringToFloat32(s string) (float32, error) {
	v, err := strconv.ParseFloat(strings.TrimSpace(s), 32)
	return float32(v), err
}

// StringToFloat64 converts string to float64.
func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(strings.TrimSpace(s), 64)
}

// StringToDouble is an alias for StringToFloat64.
func StringToDouble(s string) (float64, error) {
	return StringToFloat64(s)
}

// StringToBoolOrDefault converts string to bool and falls back to defaultValue on parse error.
func StringToBoolOrDefault(s string, defaultValue bool) bool {
	v, err := StringToBool(s)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToIntOrDefault converts string to int and falls back to defaultValue on parse error.
func StringToIntOrDefault(s string, defaultValue int) int {
	v, err := StringToInt(s)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToIntBaseOrDefault converts string to int with base and falls back to defaultValue on parse error.
func StringToIntBaseOrDefault(s string, base int, defaultValue int) int {
	v, err := StringToIntBase(s, base)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToInt8OrDefault converts string to int8 and falls back to defaultValue on parse error.
func StringToInt8OrDefault(s string, defaultValue int8) int8 {
	v, err := StringToInt8(s)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToInt8BaseOrDefault converts string to int8 with base and falls back to defaultValue on parse error.
func StringToInt8BaseOrDefault(s string, base int, defaultValue int8) int8 {
	v, err := StringToInt8Base(s, base)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToInt16OrDefault converts string to int16 and falls back to defaultValue on parse error.
func StringToInt16OrDefault(s string, defaultValue int16) int16 {
	v, err := StringToInt16(s)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToInt16BaseOrDefault converts string to int16 with base and falls back to defaultValue on parse error.
func StringToInt16BaseOrDefault(s string, base int, defaultValue int16) int16 {
	v, err := StringToInt16Base(s, base)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToInt32OrDefault converts string to int32 and falls back to defaultValue on parse error.
func StringToInt32OrDefault(s string, defaultValue int32) int32 {
	v, err := StringToInt32(s)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToInt32BaseOrDefault converts string to int32 with base and falls back to defaultValue on parse error.
func StringToInt32BaseOrDefault(s string, base int, defaultValue int32) int32 {
	v, err := StringToInt32Base(s, base)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToInt64OrDefault converts string to int64 and falls back to defaultValue on parse error.
func StringToInt64OrDefault(s string, defaultValue int64) int64 {
	v, err := StringToInt64(s)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToInt64BaseOrDefault converts string to int64 with base and falls back to defaultValue on parse error.
func StringToInt64BaseOrDefault(s string, base int, defaultValue int64) int64 {
	v, err := StringToInt64Base(s, base)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToUintOrDefault converts string to uint and falls back to defaultValue on parse error.
func StringToUintOrDefault(s string, defaultValue uint) uint {
	v, err := StringToUint(s)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToUintBaseOrDefault converts string to uint with base and falls back to defaultValue on parse error.
func StringToUintBaseOrDefault(s string, base int, defaultValue uint) uint {
	v, err := StringToUintBase(s, base)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToUint8OrDefault converts string to uint8 and falls back to defaultValue on parse error.
func StringToUint8OrDefault(s string, defaultValue uint8) uint8 {
	v, err := StringToUint8(s)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToUint8BaseOrDefault converts string to uint8 with base and falls back to defaultValue on parse error.
func StringToUint8BaseOrDefault(s string, base int, defaultValue uint8) uint8 {
	v, err := StringToUint8Base(s, base)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToUint16OrDefault converts string to uint16 and falls back to defaultValue on parse error.
func StringToUint16OrDefault(s string, defaultValue uint16) uint16 {
	v, err := StringToUint16(s)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToUint16BaseOrDefault converts string to uint16 with base and falls back to defaultValue on parse error.
func StringToUint16BaseOrDefault(s string, base int, defaultValue uint16) uint16 {
	v, err := StringToUint16Base(s, base)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToUint32OrDefault converts string to uint32 and falls back to defaultValue on parse error.
func StringToUint32OrDefault(s string, defaultValue uint32) uint32 {
	v, err := StringToUint32(s)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToUint32BaseOrDefault converts string to uint32 with base and falls back to defaultValue on parse error.
func StringToUint32BaseOrDefault(s string, base int, defaultValue uint32) uint32 {
	v, err := StringToUint32Base(s, base)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToUint64OrDefault converts string to uint64 and falls back to defaultValue on parse error.
func StringToUint64OrDefault(s string, defaultValue uint64) uint64 {
	v, err := StringToUint64(s)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToUint64BaseOrDefault converts string to uint64 with base and falls back to defaultValue on parse error.
func StringToUint64BaseOrDefault(s string, base int, defaultValue uint64) uint64 {
	v, err := StringToUint64Base(s, base)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToFloat32OrDefault converts string to float32 and falls back to defaultValue on parse error.
func StringToFloat32OrDefault(s string, defaultValue float32) float32 {
	v, err := StringToFloat32(s)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToFloat64OrDefault converts string to float64 and falls back to defaultValue on parse error.
func StringToFloat64OrDefault(s string, defaultValue float64) float64 {
	v, err := StringToFloat64(s)
	if err != nil {
		return defaultValue
	}

	return v
}

// StringToDoubleOrDefault converts string to float64 and falls back to defaultValue on parse error.
func StringToDoubleOrDefault(s string, defaultValue float64) float64 {
	v, err := StringToDouble(s)
	if err != nil {
		return defaultValue
	}

	return v
}

func toStringPtr[T any](src *T, fn func(T) string) *string {
	if src == nil {
		return nil
	}

	v := fn(*src)
	return &v
}

func fromStringPtr[T any](src *string, fn func(string) (T, error)) *T {
	if src == nil {
		return nil
	}

	v, err := fn(*src)
	if err != nil {
		return nil
	}

	return &v
}

func fromStringPtrWithBase[T any](src *string, base int, fn func(string, int) (T, error)) *T {
	if src == nil {
		return nil
	}

	v, err := fn(*src, base)
	if err != nil {
		return nil
	}

	return &v
}

// BoolPtrToStringPtr converts *bool to *string. Returns nil when src is nil.
func BoolPtrToStringPtr(src *bool) *string { return toStringPtr(src, BoolToString) }

// IntPtrToStringPtr converts *int to *string. Returns nil when src is nil.
func IntPtrToStringPtr(src *int) *string { return toStringPtr(src, IntToString) }

// Int8PtrToStringPtr converts *int8 to *string. Returns nil when src is nil.
func Int8PtrToStringPtr(src *int8) *string { return toStringPtr(src, Int8ToString) }

// Int16PtrToStringPtr converts *int16 to *string. Returns nil when src is nil.
func Int16PtrToStringPtr(src *int16) *string { return toStringPtr(src, Int16ToString) }

// Int32PtrToStringPtr converts *int32 to *string. Returns nil when src is nil.
func Int32PtrToStringPtr(src *int32) *string { return toStringPtr(src, Int32ToString) }

// Int64PtrToStringPtr converts *int64 to *string. Returns nil when src is nil.
func Int64PtrToStringPtr(src *int64) *string { return toStringPtr(src, Int64ToString) }

// UintPtrToStringPtr converts *uint to *string. Returns nil when src is nil.
func UintPtrToStringPtr(src *uint) *string { return toStringPtr(src, UintToString) }

// Uint8PtrToStringPtr converts *uint8 to *string. Returns nil when src is nil.
func Uint8PtrToStringPtr(src *uint8) *string { return toStringPtr(src, Uint8ToString) }

// Uint16PtrToStringPtr converts *uint16 to *string. Returns nil when src is nil.
func Uint16PtrToStringPtr(src *uint16) *string { return toStringPtr(src, Uint16ToString) }

// Uint32PtrToStringPtr converts *uint32 to *string. Returns nil when src is nil.
func Uint32PtrToStringPtr(src *uint32) *string { return toStringPtr(src, Uint32ToString) }

// Uint64PtrToStringPtr converts *uint64 to *string. Returns nil when src is nil.
func Uint64PtrToStringPtr(src *uint64) *string { return toStringPtr(src, Uint64ToString) }

// Float32PtrToStringPtr converts *float32 to *string. Returns nil when src is nil.
func Float32PtrToStringPtr(src *float32) *string { return toStringPtr(src, Float32ToString) }

// Float64PtrToStringPtr converts *float64 to *string. Returns nil when src is nil.
func Float64PtrToStringPtr(src *float64) *string { return toStringPtr(src, Float64ToString) }

// DoublePtrToStringPtr converts *float64 to *string. Returns nil when src is nil.
func DoublePtrToStringPtr(src *float64) *string { return toStringPtr(src, DoubleToString) }

// StringPtrToBoolPtr converts *string to *bool. Returns nil when src is nil or parse fails.
func StringPtrToBoolPtr(src *string) *bool { return fromStringPtr(src, StringToBool) }

// StringPtrToIntPtr converts *string to *int. Returns nil when src is nil or parse fails.
func StringPtrToIntPtr(src *string) *int { return fromStringPtr(src, StringToInt) }

// StringPtrToIntPtrBase converts *string to *int with base. Returns nil when src is nil or parse fails.
func StringPtrToIntPtrBase(src *string, base int) *int {
	return fromStringPtrWithBase(src, base, StringToIntBase)
}

// StringPtrToInt8Ptr converts *string to *int8. Returns nil when src is nil or parse fails.
func StringPtrToInt8Ptr(src *string) *int8 { return fromStringPtr(src, StringToInt8) }

// StringPtrToInt8PtrBase converts *string to *int8 with base. Returns nil when src is nil or parse fails.
func StringPtrToInt8PtrBase(src *string, base int) *int8 {
	return fromStringPtrWithBase(src, base, StringToInt8Base)
}

// StringPtrToInt16Ptr converts *string to *int16. Returns nil when src is nil or parse fails.
func StringPtrToInt16Ptr(src *string) *int16 { return fromStringPtr(src, StringToInt16) }

// StringPtrToInt16PtrBase converts *string to *int16 with base. Returns nil when src is nil or parse fails.
func StringPtrToInt16PtrBase(src *string, base int) *int16 {
	return fromStringPtrWithBase(src, base, StringToInt16Base)
}

// StringPtrToInt32Ptr converts *string to *int32. Returns nil when src is nil or parse fails.
func StringPtrToInt32Ptr(src *string) *int32 { return fromStringPtr(src, StringToInt32) }

// StringPtrToInt32PtrBase converts *string to *int32 with base. Returns nil when src is nil or parse fails.
func StringPtrToInt32PtrBase(src *string, base int) *int32 {
	return fromStringPtrWithBase(src, base, StringToInt32Base)
}

// StringPtrToInt64Ptr converts *string to *int64. Returns nil when src is nil or parse fails.
func StringPtrToInt64Ptr(src *string) *int64 { return fromStringPtr(src, StringToInt64) }

// StringPtrToInt64PtrBase converts *string to *int64 with base. Returns nil when src is nil or parse fails.
func StringPtrToInt64PtrBase(src *string, base int) *int64 {
	return fromStringPtrWithBase(src, base, StringToInt64Base)
}

// StringPtrToUintPtr converts *string to *uint. Returns nil when src is nil or parse fails.
func StringPtrToUintPtr(src *string) *uint { return fromStringPtr(src, StringToUint) }

// StringPtrToUintPtrBase converts *string to *uint with base. Returns nil when src is nil or parse fails.
func StringPtrToUintPtrBase(src *string, base int) *uint {
	return fromStringPtrWithBase(src, base, StringToUintBase)
}

// StringPtrToUint8Ptr converts *string to *uint8. Returns nil when src is nil or parse fails.
func StringPtrToUint8Ptr(src *string) *uint8 { return fromStringPtr(src, StringToUint8) }

// StringPtrToUint8PtrBase converts *string to *uint8 with base. Returns nil when src is nil or parse fails.
func StringPtrToUint8PtrBase(src *string, base int) *uint8 {
	return fromStringPtrWithBase(src, base, StringToUint8Base)
}

// StringPtrToUint16Ptr converts *string to *uint16. Returns nil when src is nil or parse fails.
func StringPtrToUint16Ptr(src *string) *uint16 { return fromStringPtr(src, StringToUint16) }

// StringPtrToUint16PtrBase converts *string to *uint16 with base. Returns nil when src is nil or parse fails.
func StringPtrToUint16PtrBase(src *string, base int) *uint16 {
	return fromStringPtrWithBase(src, base, StringToUint16Base)
}

// StringPtrToUint32Ptr converts *string to *uint32. Returns nil when src is nil or parse fails.
func StringPtrToUint32Ptr(src *string) *uint32 { return fromStringPtr(src, StringToUint32) }

// StringPtrToUint32PtrBase converts *string to *uint32 with base. Returns nil when src is nil or parse fails.
func StringPtrToUint32PtrBase(src *string, base int) *uint32 {
	return fromStringPtrWithBase(src, base, StringToUint32Base)
}

// StringPtrToUint64Ptr converts *string to *uint64. Returns nil when src is nil or parse fails.
func StringPtrToUint64Ptr(src *string) *uint64 { return fromStringPtr(src, StringToUint64) }

// StringPtrToUint64PtrBase converts *string to *uint64 with base. Returns nil when src is nil or parse fails.
func StringPtrToUint64PtrBase(src *string, base int) *uint64 {
	return fromStringPtrWithBase(src, base, StringToUint64Base)
}

// StringPtrToFloat32Ptr converts *string to *float32. Returns nil when src is nil or parse fails.
func StringPtrToFloat32Ptr(src *string) *float32 { return fromStringPtr(src, StringToFloat32) }

// StringPtrToFloat64Ptr converts *string to *float64. Returns nil when src is nil or parse fails.
func StringPtrToFloat64Ptr(src *string) *float64 { return fromStringPtr(src, StringToFloat64) }

// StringPtrToDoublePtr converts *string to *float64. Returns nil when src is nil or parse fails.
func StringPtrToDoublePtr(src *string) *float64 { return fromStringPtr(src, StringToDouble) }
