/**
 * berry - Utility functions for Go.
 *
 * MIT License.
 * Copyright (c) 2022, Sandeep Gupta.
 * https://github.com/sangupta/berry
 *
 * Use of this source code is governed by a MIT style license
 * that can be found in LICENSE file in the code repository:
 */

package berry

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToBool(t *testing.T) {
	var f uintptr
	b, err := ConvertToBool(f)
	assert.False(t, b)
	assert.NoError(t, err)

	// int
	b, err = ConvertToBool(int(2))
	assert.True(t, b)
	assert.NoError(t, err)
	b, err = ConvertToBool(int(0))
	assert.False(t, b)
	assert.NoError(t, err)

	// int8
	b, err = ConvertToBool(int8(2))
	assert.True(t, b)
	assert.NoError(t, err)
	b, err = ConvertToBool(int8(0))
	assert.False(t, b)
	assert.NoError(t, err)

	// int16
	b, err = ConvertToBool(int16(2))
	assert.True(t, b)
	assert.NoError(t, err)
	b, err = ConvertToBool(int16(0))
	assert.False(t, b)
	assert.NoError(t, err)

	// int32
	b, err = ConvertToBool(int32(2))
	assert.True(t, b)
	assert.NoError(t, err)
	b, err = ConvertToBool(int32(0))
	assert.False(t, b)
	assert.NoError(t, err)

	// int64
	b, err = ConvertToBool(int64(2))
	assert.True(t, b)
	assert.NoError(t, err)
	b, err = ConvertToBool(int64(0))
	assert.False(t, b)
	assert.NoError(t, err)

	// uint
	b, err = ConvertToBool(uint(2))
	assert.True(t, b)
	assert.NoError(t, err)
	b, err = ConvertToBool(uint(0))
	assert.False(t, b)
	assert.NoError(t, err)

	// uint8
	b, err = ConvertToBool(uint8(2))
	assert.True(t, b)
	assert.NoError(t, err)
	b, err = ConvertToBool(uint8(0))
	assert.False(t, b)
	assert.NoError(t, err)

	// uint16
	b, err = ConvertToBool(uint16(2))
	assert.True(t, b)
	assert.NoError(t, err)
	b, err = ConvertToBool(uint16(0))
	assert.False(t, b)
	assert.NoError(t, err)

	// uint32
	b, err = ConvertToBool(uint32(2))
	assert.True(t, b)
	assert.NoError(t, err)
	b, err = ConvertToBool(uint32(0))
	assert.False(t, b)
	assert.NoError(t, err)

	// uint64
	b, err = ConvertToBool(uint64(2))
	assert.True(t, b)
	assert.NoError(t, err)
	b, err = ConvertToBool(uint64(0))
	assert.False(t, b)
	assert.NoError(t, err)

	// float32
	b, err = ConvertToBool(float32(2))
	assert.True(t, b)
	assert.NoError(t, err)
	b, err = ConvertToBool(float32(0))
	assert.False(t, b)
	assert.NoError(t, err)

	// float64
	b, err = ConvertToBool(float64(2))
	assert.True(t, b)
	assert.NoError(t, err)
	b, err = ConvertToBool(float64(0))
	assert.False(t, b)
	assert.NoError(t, err)

	// string
	b, err = ConvertToBool("true")
	assert.True(t, b)
	assert.NoError(t, err)
	b, err = ConvertToBool("TRUE")
	assert.True(t, b)
	assert.NoError(t, err)
	b, err = ConvertToBool("TRue")
	assert.True(t, b)
	assert.NoError(t, err)
	b, err = ConvertToBool("")
	assert.False(t, b)
	assert.NoError(t, err)

	// bool
	b, err = ConvertToBool(true)
	assert.True(t, b)
	assert.NoError(t, err)
	b, err = ConvertToBool(false)
	assert.False(t, b)
	assert.NoError(t, err)
}

func TestConvertToString(t *testing.T) {
	assert.Equal(t, "2", ConvertToString(2))
}

func TestConvertToUint64(t *testing.T) {
	value, err := ConvertToUint64(int(32), 0)
	assert.Equal(t, uint64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToUint64(int8(32), 0)
	assert.Equal(t, uint64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToUint64(int16(32), 0)
	assert.Equal(t, uint64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToUint64(int32(32), 0)
	assert.Equal(t, uint64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToUint64(int64(32), 0)
	assert.Equal(t, uint64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToUint64(uint(32), 0)
	assert.Equal(t, uint64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToUint64(uint8(32), 0)
	assert.Equal(t, uint64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToUint64(uint16(32), 0)
	assert.Equal(t, uint64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToUint64(uint32(32), 0)
	assert.Equal(t, uint64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToUint64(uint64(32), 0)
	assert.Equal(t, uint64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToUint64(float32(32.2), 0)
	assert.Equal(t, uint64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToUint64(float32(32.8), 0)
	assert.Equal(t, uint64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToUint64(float64(32.2), 0)
	assert.Equal(t, uint64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToUint64(float64(32.8), 0)
	assert.Equal(t, uint64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToUint64("32", 0)
	assert.Equal(t, uint64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToUint64("32.2", 0)
	assert.Error(t, err)

	value, err = ConvertToUint64("32.8", 0)
	assert.Error(t, err)

	value, err = ConvertToUint64(strings.Builder{}, 0)
	assert.Equal(t, uint64(0), value)
	assert.NoError(t, err)
}

func TestConvertToInt64(t *testing.T) {
	value, err := ConvertToInt64(int(32), 0)
	assert.Equal(t, int64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToInt64(int8(32), 0)
	assert.Equal(t, int64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToInt64(int16(32), 0)
	assert.Equal(t, int64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToInt64(int32(32), 0)
	assert.Equal(t, int64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToInt64(int64(32), 0)
	assert.Equal(t, int64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToInt64(uint(32), 0)
	assert.Equal(t, int64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToInt64(uint8(32), 0)
	assert.Equal(t, int64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToInt64(uint16(32), 0)
	assert.Equal(t, int64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToInt64(uint32(32), 0)
	assert.Equal(t, int64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToInt64(uint64(32), 0)
	assert.Equal(t, int64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToInt64(float32(32.2), 0)
	assert.Equal(t, int64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToInt64(float32(32.8), 0)
	assert.Equal(t, int64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToInt64(float64(32.2), 0)
	assert.Equal(t, int64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToInt64(float64(32.8), 0)
	assert.Equal(t, int64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToInt64("32", 0)
	assert.Equal(t, int64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToInt64("32.2", 0)
	assert.Error(t, err)

	value, err = ConvertToInt64("32.8", 0)
	assert.Error(t, err)

	value, err = ConvertToInt64(strings.Builder{}, 0)
	assert.Equal(t, int64(0), value)
	assert.NoError(t, err)
}

func TestConvertToFloat64(t *testing.T) {
	value, err := ConvertToFloat64(int(32), 0)
	assert.Equal(t, float64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToFloat64(int8(32), 0)
	assert.Equal(t, float64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToFloat64(int16(32), 0)
	assert.Equal(t, float64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToFloat64(int32(32), 0)
	assert.Equal(t, float64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToFloat64(int64(32), 0)
	assert.Equal(t, float64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToFloat64(uint(32), 0)
	assert.Equal(t, float64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToFloat64(uint8(32), 0)
	assert.Equal(t, float64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToFloat64(uint16(32), 0)
	assert.Equal(t, float64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToFloat64(uint32(32), 0)
	assert.Equal(t, float64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToFloat64(uint64(32), 0)
	assert.Equal(t, float64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToFloat64(float32(32.2), 0)
	assert.InDelta(t, float64(32.2), value, 0.000001)
	assert.NoError(t, err)

	value, err = ConvertToFloat64(float32(32.8), 0)
	assert.InDelta(t, float64(32.8), value, 0.000001)
	assert.NoError(t, err)

	value, err = ConvertToFloat64(float64(32.2), 0)
	assert.InDelta(t, float64(32.2), value, 0.000001)
	assert.NoError(t, err)

	value, err = ConvertToFloat64(float64(32.8), 0)
	assert.InDelta(t, float64(32.8), value, 0.000001)
	assert.NoError(t, err)

	value, err = ConvertToFloat64("32", 0)
	assert.Equal(t, float64(32), value)
	assert.NoError(t, err)

	value, err = ConvertToFloat64("32.2", 0)
	assert.InDelta(t, float64(32.2), value, 0.000001)
	assert.NoError(t, err)

	value, err = ConvertToFloat64("32.8", 0)
	assert.InDelta(t, float64(32.8), value, 0.000001)
	assert.NoError(t, err)

	value, err = ConvertToFloat64(strings.Builder{}, 0)
	assert.Equal(t, float64(0), value)
	assert.NoError(t, err)
}
