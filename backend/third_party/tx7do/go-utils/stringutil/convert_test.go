package stringutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeToString(t *testing.T) {
	assert.Equal(t, "true", BoolToString(true))
	assert.Equal(t, "12", IntToString(12))
	assert.Equal(t, "8", Int8ToString(8))
	assert.Equal(t, "16", Int16ToString(16))
	assert.Equal(t, "123", Int32ToString(123))
	assert.Equal(t, "456", Int64ToString(456))
	assert.Equal(t, "12", UintToString(12))
	assert.Equal(t, "8", Uint8ToString(8))
	assert.Equal(t, "16", Uint16ToString(16))
	assert.Equal(t, "123", Uint32ToString(123))
	assert.Equal(t, "456", Uint64ToString(456))
	assert.Equal(t, "1.25", Float32ToString(1.25))
	assert.Equal(t, "2.5", Float64ToString(2.5))
	assert.Equal(t, Float64ToString(3.1415), DoubleToString(3.1415))
}

func TestStringToType(t *testing.T) {
	b, err := StringToBool(" true ")
	assert.NoError(t, err)
	assert.True(t, b)

	i, err := StringToInt(" 10 ")
	assert.NoError(t, err)
	assert.Equal(t, 10, i)

	i8, err := StringToInt8(" 11 ")
	assert.NoError(t, err)
	assert.Equal(t, int8(11), i8)

	i16, err := StringToInt16(" 12 ")
	assert.NoError(t, err)
	assert.Equal(t, int16(12), i16)

	i32, err := StringToInt32(" 20 ")
	assert.NoError(t, err)
	assert.Equal(t, int32(20), i32)

	i64, err := StringToInt64(" 30 ")
	assert.NoError(t, err)
	assert.Equal(t, int64(30), i64)

	u, err := StringToUint(" 40 ")
	assert.NoError(t, err)
	assert.Equal(t, uint(40), u)

	u8, err := StringToUint8(" 41 ")
	assert.NoError(t, err)
	assert.Equal(t, uint8(41), u8)

	u16, err := StringToUint16(" 42 ")
	assert.NoError(t, err)
	assert.Equal(t, uint16(42), u16)

	u32, err := StringToUint32(" 50 ")
	assert.NoError(t, err)
	assert.Equal(t, uint32(50), u32)

	u64, err := StringToUint64(" 60 ")
	assert.NoError(t, err)
	assert.Equal(t, uint64(60), u64)

	f32, err := StringToFloat32(" 1.25 ")
	assert.NoError(t, err)
	assert.InDelta(t, float32(1.25), f32, 1e-6)

	f64, err := StringToFloat64(" 2.5 ")
	assert.NoError(t, err)
	assert.InDelta(t, float64(2.5), f64, 1e-9)

	d, err := StringToDouble(" 3.5 ")
	assert.NoError(t, err)
	assert.InDelta(t, float64(3.5), d, 1e-9)

	iBase, err := StringToIntBase("ff", 16)
	assert.NoError(t, err)
	assert.Equal(t, 255, iBase)

	i8Base, err := StringToInt8Base("7f", 16)
	assert.NoError(t, err)
	assert.Equal(t, int8(127), i8Base)

	i16Base, err := StringToInt16Base("7fff", 16)
	assert.NoError(t, err)
	assert.Equal(t, int16(32767), i16Base)

	i32Base, err := StringToInt32Base("177", 8)
	assert.NoError(t, err)
	assert.Equal(t, int32(127), i32Base)

	i64Base, err := StringToInt64Base("0x7f", 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(127), i64Base)

	uBase, err := StringToUintBase("1111", 2)
	assert.NoError(t, err)
	assert.Equal(t, uint(15), uBase)

	u8Base, err := StringToUint8Base("ff", 16)
	assert.NoError(t, err)
	assert.Equal(t, uint8(255), u8Base)

	u16Base, err := StringToUint16Base("ffff", 16)
	assert.NoError(t, err)
	assert.Equal(t, uint16(65535), u16Base)

	u32Base, err := StringToUint32Base("377", 8)
	assert.NoError(t, err)
	assert.Equal(t, uint32(255), u32Base)

	u64Base, err := StringToUint64Base("0xff", 0)
	assert.NoError(t, err)
	assert.Equal(t, uint64(255), u64Base)
}

func TestStringToTypeErrors(t *testing.T) {
	_, err := StringToBool("not-bool")
	assert.Error(t, err)

	_, err = StringToInt32("abc")
	assert.Error(t, err)

	_, err = StringToInt8("128")
	assert.Error(t, err)

	_, err = StringToFloat64("abc")
	assert.Error(t, err)

	_, err = StringToUint32("-1")
	assert.Error(t, err)

	_, err = StringToUint8("256")
	assert.Error(t, err)

	_, err = StringToIntBase("g", 16)
	assert.Error(t, err)

	_, err = StringToInt8Base("80", 16)
	assert.Error(t, err)

	_, err = StringToUintBase("2", 2)
	assert.Error(t, err)

	_, err = StringToUint8Base("100", 16)
	assert.Error(t, err)
}

func TestStringToTypeOrDefault(t *testing.T) {
	assert.True(t, StringToBoolOrDefault("true", false))
	assert.Equal(t, 10, StringToIntOrDefault("10", 1))
	assert.Equal(t, int8(11), StringToInt8OrDefault("11", 1))
	assert.Equal(t, int16(12), StringToInt16OrDefault("12", 1))
	assert.Equal(t, int32(20), StringToInt32OrDefault("20", 2))
	assert.Equal(t, int64(30), StringToInt64OrDefault("30", 3))
	assert.Equal(t, uint(40), StringToUintOrDefault("40", 4))
	assert.Equal(t, uint8(41), StringToUint8OrDefault("41", 4))
	assert.Equal(t, uint16(42), StringToUint16OrDefault("42", 4))
	assert.Equal(t, uint32(50), StringToUint32OrDefault("50", 5))
	assert.Equal(t, uint64(60), StringToUint64OrDefault("60", 6))
	assert.InDelta(t, float32(1.25), StringToFloat32OrDefault("1.25", 0), 1e-6)
	assert.InDelta(t, float64(2.5), StringToFloat64OrDefault("2.5", 0), 1e-9)
	assert.InDelta(t, float64(3.5), StringToDoubleOrDefault("3.5", 0), 1e-9)

	assert.Equal(t, true, StringToBoolOrDefault("bad", true))
	assert.Equal(t, 11, StringToIntOrDefault("bad", 11))
	assert.Equal(t, int8(12), StringToInt8OrDefault("bad", 12))
	assert.Equal(t, int16(13), StringToInt16OrDefault("bad", 13))
	assert.Equal(t, int32(22), StringToInt32OrDefault("bad", 22))
	assert.Equal(t, int64(33), StringToInt64OrDefault("bad", 33))
	assert.Equal(t, uint(44), StringToUintOrDefault("bad", 44))
	assert.Equal(t, uint8(45), StringToUint8OrDefault("bad", 45))
	assert.Equal(t, uint16(46), StringToUint16OrDefault("bad", 46))
	assert.Equal(t, uint32(55), StringToUint32OrDefault("bad", 55))
	assert.Equal(t, uint64(66), StringToUint64OrDefault("bad", 66))
	assert.Equal(t, float32(1.5), StringToFloat32OrDefault("bad", 1.5))
	assert.Equal(t, float64(2.5), StringToFloat64OrDefault("bad", 2.5))
	assert.Equal(t, float64(3.5), StringToDoubleOrDefault("bad", 3.5))

	assert.Equal(t, 255, StringToIntBaseOrDefault("ff", 16, 1))
	assert.Equal(t, 1, StringToIntBaseOrDefault("gg", 16, 1))
	assert.Equal(t, int8(127), StringToInt8BaseOrDefault("7f", 16, 1))
	assert.Equal(t, int8(1), StringToInt8BaseOrDefault("80", 16, 1))
	assert.Equal(t, int16(32767), StringToInt16BaseOrDefault("7fff", 16, 1))
	assert.Equal(t, int16(1), StringToInt16BaseOrDefault("8000", 16, 1))
	assert.Equal(t, int32(127), StringToInt32BaseOrDefault("177", 8, 1))
	assert.Equal(t, int32(1), StringToInt32BaseOrDefault("178", 8, 1))
	assert.Equal(t, int64(127), StringToInt64BaseOrDefault("0x7f", 0, 1))
	assert.Equal(t, int64(1), StringToInt64BaseOrDefault("0xzz", 0, 1))
	assert.Equal(t, uint(15), StringToUintBaseOrDefault("1111", 2, 1))
	assert.Equal(t, uint(1), StringToUintBaseOrDefault("2222", 2, 1))
	assert.Equal(t, uint8(255), StringToUint8BaseOrDefault("ff", 16, 1))
	assert.Equal(t, uint8(1), StringToUint8BaseOrDefault("100", 16, 1))
	assert.Equal(t, uint16(65535), StringToUint16BaseOrDefault("ffff", 16, 1))
	assert.Equal(t, uint16(1), StringToUint16BaseOrDefault("10000", 16, 1))
	assert.Equal(t, uint32(255), StringToUint32BaseOrDefault("377", 8, 1))
	assert.Equal(t, uint32(1), StringToUint32BaseOrDefault("478", 8, 1))
	assert.Equal(t, uint64(255), StringToUint64BaseOrDefault("0xff", 0, 1))
	assert.Equal(t, uint64(1), StringToUint64BaseOrDefault("0xzz", 0, 1))
}

func TestPtrToStringPtr(t *testing.T) {
	b := true
	assert.Equal(t, "true", *BoolPtrToStringPtr(&b))

	i := 7
	assert.Equal(t, "7", *IntPtrToStringPtr(&i))

	i8 := int8(8)
	assert.Equal(t, "8", *Int8PtrToStringPtr(&i8))

	i16 := int16(16)
	assert.Equal(t, "16", *Int16PtrToStringPtr(&i16))

	i32 := int32(32)
	assert.Equal(t, "32", *Int32PtrToStringPtr(&i32))

	i64 := int64(64)
	assert.Equal(t, "64", *Int64PtrToStringPtr(&i64))

	u := uint(9)
	assert.Equal(t, "9", *UintPtrToStringPtr(&u))

	u8 := uint8(10)
	assert.Equal(t, "10", *Uint8PtrToStringPtr(&u8))

	u16 := uint16(17)
	assert.Equal(t, "17", *Uint16PtrToStringPtr(&u16))

	u32 := uint32(33)
	assert.Equal(t, "33", *Uint32PtrToStringPtr(&u32))

	u64 := uint64(65)
	assert.Equal(t, "65", *Uint64PtrToStringPtr(&u64))

	f32 := float32(1.25)
	assert.Equal(t, "1.25", *Float32PtrToStringPtr(&f32))

	f64 := float64(2.5)
	assert.Equal(t, "2.5", *Float64PtrToStringPtr(&f64))
	assert.Equal(t, "2.5", *DoublePtrToStringPtr(&f64))

	assert.Nil(t, BoolPtrToStringPtr(nil))
	assert.Nil(t, IntPtrToStringPtr(nil))
	assert.Nil(t, Int8PtrToStringPtr(nil))
	assert.Nil(t, Int16PtrToStringPtr(nil))
	assert.Nil(t, Int32PtrToStringPtr(nil))
	assert.Nil(t, Int64PtrToStringPtr(nil))
	assert.Nil(t, UintPtrToStringPtr(nil))
	assert.Nil(t, Uint8PtrToStringPtr(nil))
	assert.Nil(t, Uint16PtrToStringPtr(nil))
	assert.Nil(t, Uint32PtrToStringPtr(nil))
	assert.Nil(t, Uint64PtrToStringPtr(nil))
	assert.Nil(t, Float32PtrToStringPtr(nil))
	assert.Nil(t, Float64PtrToStringPtr(nil))
	assert.Nil(t, DoublePtrToStringPtr(nil))
}

func TestStringPtrToTypePtr(t *testing.T) {
	sBool := "true"
	assert.Equal(t, true, *StringPtrToBoolPtr(&sBool))

	sInt := "7"
	assert.Equal(t, 7, *StringPtrToIntPtr(&sInt))

	sInt8 := "8"
	assert.Equal(t, int8(8), *StringPtrToInt8Ptr(&sInt8))

	sInt16 := "16"
	assert.Equal(t, int16(16), *StringPtrToInt16Ptr(&sInt16))

	sInt32 := "32"
	assert.Equal(t, int32(32), *StringPtrToInt32Ptr(&sInt32))

	sInt64 := "64"
	assert.Equal(t, int64(64), *StringPtrToInt64Ptr(&sInt64))

	sUint := "9"
	assert.Equal(t, uint(9), *StringPtrToUintPtr(&sUint))

	sUint8 := "10"
	assert.Equal(t, uint8(10), *StringPtrToUint8Ptr(&sUint8))

	sUint16 := "17"
	assert.Equal(t, uint16(17), *StringPtrToUint16Ptr(&sUint16))

	sUint32 := "33"
	assert.Equal(t, uint32(33), *StringPtrToUint32Ptr(&sUint32))

	sUint64 := "65"
	assert.Equal(t, uint64(65), *StringPtrToUint64Ptr(&sUint64))

	sFloat32 := "1.25"
	assert.InDelta(t, float32(1.25), *StringPtrToFloat32Ptr(&sFloat32), 1e-6)

	sFloat64 := "2.5"
	assert.InDelta(t, float64(2.5), *StringPtrToFloat64Ptr(&sFloat64), 1e-9)
	assert.InDelta(t, float64(2.5), *StringPtrToDoublePtr(&sFloat64), 1e-9)

	assert.Nil(t, StringPtrToBoolPtr(nil))
	assert.Nil(t, StringPtrToIntPtr(nil))
	assert.Nil(t, StringPtrToInt8Ptr(nil))
	assert.Nil(t, StringPtrToInt16Ptr(nil))
	assert.Nil(t, StringPtrToInt32Ptr(nil))
	assert.Nil(t, StringPtrToInt64Ptr(nil))
	assert.Nil(t, StringPtrToUintPtr(nil))
	assert.Nil(t, StringPtrToUint8Ptr(nil))
	assert.Nil(t, StringPtrToUint16Ptr(nil))
	assert.Nil(t, StringPtrToUint32Ptr(nil))
	assert.Nil(t, StringPtrToUint64Ptr(nil))
	assert.Nil(t, StringPtrToFloat32Ptr(nil))
	assert.Nil(t, StringPtrToFloat64Ptr(nil))
	assert.Nil(t, StringPtrToDoublePtr(nil))

	badBool := "bad"
	badInt8 := "128"
	badUint8 := "256"
	badFloat := "x"
	badHex := "gg"
	assert.Nil(t, StringPtrToBoolPtr(&badBool))
	assert.Nil(t, StringPtrToInt8Ptr(&badInt8))
	assert.Nil(t, StringPtrToUint8Ptr(&badUint8))
	assert.Nil(t, StringPtrToFloat64Ptr(&badFloat))
	assert.Nil(t, StringPtrToIntPtrBase(&badHex, 16))
	assert.Nil(t, StringPtrToUintPtrBase(&badHex, 16))
	assert.Nil(t, StringPtrToInt8PtrBase(&badHex, 16))
	assert.Nil(t, StringPtrToInt16PtrBase(&badHex, 16))
	assert.Nil(t, StringPtrToInt32PtrBase(&badHex, 16))
	assert.Nil(t, StringPtrToInt64PtrBase(&badHex, 16))
	assert.Nil(t, StringPtrToUint8PtrBase(&badHex, 16))
	assert.Nil(t, StringPtrToUint16PtrBase(&badHex, 16))
	assert.Nil(t, StringPtrToUint32PtrBase(&badHex, 16))
	assert.Nil(t, StringPtrToUint64PtrBase(&badHex, 16))

	hex := "ff"
	hex7f := "7f"
	bin := "1111"
	oct := "177"
	assert.Equal(t, 255, *StringPtrToIntPtrBase(&hex, 16))
	assert.Equal(t, uint(15), *StringPtrToUintPtrBase(&bin, 2))
	assert.Equal(t, int8(127), *StringPtrToInt8PtrBase(&hex7f, 16))
	assert.Equal(t, int16(255), *StringPtrToInt16PtrBase(&hex, 16))
	assert.Equal(t, int32(127), *StringPtrToInt32PtrBase(&oct, 8))
	assert.Equal(t, int64(127), *StringPtrToInt64PtrBase(&oct, 8))
	assert.Equal(t, uint8(255), *StringPtrToUint8PtrBase(&hex, 16))
	assert.Equal(t, uint16(255), *StringPtrToUint16PtrBase(&hex, 16))
	assert.Equal(t, uint32(127), *StringPtrToUint32PtrBase(&oct, 8))
	assert.Equal(t, uint64(127), *StringPtrToUint64PtrBase(&oct, 8))
	assert.Nil(t, StringPtrToIntPtrBase(nil, 16))
	assert.Nil(t, StringPtrToUintPtrBase(nil, 16))
	assert.Nil(t, StringPtrToInt8PtrBase(nil, 16))
	assert.Nil(t, StringPtrToInt16PtrBase(nil, 16))
	assert.Nil(t, StringPtrToInt32PtrBase(nil, 16))
	assert.Nil(t, StringPtrToInt64PtrBase(nil, 16))
	assert.Nil(t, StringPtrToUint8PtrBase(nil, 16))
	assert.Nil(t, StringPtrToUint16PtrBase(nil, 16))
	assert.Nil(t, StringPtrToUint32PtrBase(nil, 16))
	assert.Nil(t, StringPtrToUint64PtrBase(nil, 16))
}
