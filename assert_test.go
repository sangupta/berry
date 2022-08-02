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

func TestIsNumber(t *testing.T) {
	assert.False(t, IsNumber(nil))

	// positive
	assert.True(t, IsNumber(int(0)))
	assert.True(t, IsNumber(int(12)))

	assert.True(t, IsNumber(int8(0)))
	assert.True(t, IsNumber(int8(12)))

	assert.True(t, IsNumber(int16(0)))
	assert.True(t, IsNumber(int16(12)))

	assert.True(t, IsNumber(int32(0)))
	assert.True(t, IsNumber(int32(12)))

	assert.True(t, IsNumber(int64(0)))
	assert.True(t, IsNumber(int64(12)))

	assert.True(t, IsNumber(uint(0)))
	assert.True(t, IsNumber(uint(12)))

	assert.True(t, IsNumber(uint8(0)))
	assert.True(t, IsNumber(uint8(12)))

	assert.True(t, IsNumber(uint16(0)))
	assert.True(t, IsNumber(uint16(12)))

	assert.True(t, IsNumber(uint32(0)))
	assert.True(t, IsNumber(uint32(12)))

	assert.True(t, IsNumber(uint64(0)))
	assert.True(t, IsNumber(uint64(12)))

	assert.True(t, IsNumber(float32(0)))
	assert.True(t, IsNumber(float32(12)))

	assert.True(t, IsNumber(float64(0)))
	assert.True(t, IsNumber(float64(12)))

	// negative
	var xi13 string = "12"
	assert.False(t, IsNumber(xi13))
}

func TestIsSlice(t *testing.T) {
	assert.True(t, IsSlice(make([]int, 1)))
	assert.True(t, IsSlice(make([]int8, 1)))
	assert.True(t, IsSlice(make([]int16, 1)))
	assert.True(t, IsSlice(make([]int32, 1)))
	assert.True(t, IsSlice(make([]int64, 1)))
	assert.True(t, IsSlice(make([]uint, 1)))
	assert.True(t, IsSlice(make([]uint8, 1)))
	assert.True(t, IsSlice(make([]uint16, 1)))
	assert.True(t, IsSlice(make([]uint32, 1)))
	assert.True(t, IsSlice(make([]uint64, 1)))
	assert.True(t, IsSlice(make([]float32, 1)))
	assert.True(t, IsSlice(make([]float64, 1)))
	assert.True(t, IsSlice(make([]string, 1)))

	assert.False(t, IsSlice(nil))
	assert.False(t, IsSlice("123"))
	assert.False(t, IsSlice(64.34))
	assert.False(t, IsSlice(strings.Builder{}))
}

func TestIsMap(t *testing.T) {
	assert.True(t, IsMap(make(map[string]int, 1)))
	assert.True(t, IsMap(make(map[string]int8, 1)))
	assert.True(t, IsMap(make(map[string]int16, 1)))
	assert.True(t, IsMap(make(map[string]int32, 1)))
	assert.True(t, IsMap(make(map[string]int64, 1)))
	assert.True(t, IsMap(make(map[string]uint, 1)))
	assert.True(t, IsMap(make(map[string]uint8, 1)))
	assert.True(t, IsMap(make(map[string]uint16, 1)))
	assert.True(t, IsMap(make(map[string]uint32, 1)))
	assert.True(t, IsMap(make(map[string]uint64, 1)))
	assert.True(t, IsMap(make(map[string]float32, 1)))
	assert.True(t, IsMap(make(map[string]float64, 1)))
	assert.True(t, IsMap(make(map[string]string, 1)))

	assert.False(t, IsMap(nil))
	assert.False(t, IsMap("123"))
	assert.False(t, IsMap(64.34))
	assert.False(t, IsMap(strings.Builder{}))
	assert.False(t, IsMap(make([]int, 1)))
	assert.False(t, IsMap(make([]int8, 1)))
	assert.False(t, IsMap(make([]int16, 1)))
	assert.False(t, IsMap(make([]int32, 1)))
	assert.False(t, IsMap(make([]int64, 1)))
	assert.False(t, IsMap(make([]uint, 1)))
	assert.False(t, IsMap(make([]uint8, 1)))
	assert.False(t, IsMap(make([]uint16, 1)))
	assert.False(t, IsMap(make([]uint32, 1)))
	assert.False(t, IsMap(make([]uint64, 1)))
	assert.False(t, IsMap(make([]float32, 1)))
	assert.False(t, IsMap(make([]float64, 1)))
	assert.False(t, IsMap(make([]string, 1)))
}
