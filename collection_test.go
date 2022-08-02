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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceMax(t *testing.T) {
	max, err := SliceMax([]int{2, 4, 8, 16, 32, 64, 54, 42, 12})
	assert.NoError(t, err)
	assert.Equal(t, int(64), max)

	max, err = SliceMax([]int{2})
	assert.NoError(t, err)
	assert.Equal(t, int(2), max)

	// negative, nil
	var foo []int
	max, err = SliceMax(foo)
	assert.Error(t, err)

	foo2 := []int{}
	max, err = SliceMax(foo2)
	assert.Error(t, err)
}

func TestSliceMin(t *testing.T) {
	max, err := SliceMin([]int{2, 4, 8, 16, 32, 64, 54, 42, 12, -2})
	assert.NoError(t, err)
	assert.Equal(t, int(-2), max)

	max, err = SliceMin([]int{2})
	assert.NoError(t, err)
	assert.Equal(t, int(2), max)

	// negative, nil
	var foo []int
	max, err = SliceMin(foo)
	assert.Error(t, err)

	foo2 := []int{}
	max, err = SliceMin(foo2)
	assert.Error(t, err)
}

func TestSliceContains(t *testing.T) {
	var foo []int
	assert.False(t, SliceContains(foo, 2))

	foo = []int{}
	assert.False(t, SliceContains(foo, 2))

	foo = []int{4}
	assert.False(t, SliceContains(foo, 2))

	foo = []int{4, 8}
	assert.False(t, SliceContains(foo, 2))

	foo = []int{4, 2, 8}
	assert.True(t, SliceContains(foo, 2))
}

func TestSliceIndex(t *testing.T) {
	var foo []int
	assert.Equal(t, -1, SliceIndex(foo, 2))

	foo = []int{}
	assert.Equal(t, -1, SliceIndex(foo, 2))

	foo = []int{4}
	assert.Equal(t, -1, SliceIndex(foo, 2))

	foo = []int{4, 8}
	assert.Equal(t, -1, SliceIndex(foo, 2))

	foo = []int{4, 2, 8}
	assert.Equal(t, 1, SliceIndex(foo, 2))
}

func TestSliceModify(t *testing.T) {
	var foo []int
	modifier := func(x int) int {
		return x * 2
	}

	assert.False(t, SliceModify(foo, modifier))

	foo = []int{2, 4, 6, 8}
	assert.True(t, SliceModify(foo, modifier))
	assert.Equal(t, []int{4, 8, 12, 16}, foo)
}

func TestAreSlicesEqual(t *testing.T) {
	var foo1 []int
	foo2 := []int{2, 4, 6, 8}

	// nil
	assert.False(t, AreSlicesEqual(foo1, foo2))
	assert.False(t, AreSlicesEqual(foo2, foo1))

	// diff length
	foo1 = []int{2, 4}
	assert.False(t, AreSlicesEqual(foo1, foo2))

	// same length
	foo2 = []int{4, 6}
	assert.False(t, AreSlicesEqual(foo1, foo2))

	foo2 = []int{2, 4}
	assert.True(t, AreSlicesEqual(foo1, foo2))
}

func TestSliceReverse(t *testing.T) {
	var foo []int
	assert.Nil(t, SliceReverse(foo))

	foo = []int{}
	assert.Equal(t, 0, len(SliceReverse(foo)))

	foo = []int{2}
	assert.Equal(t, []int{2}, SliceReverse(foo))

	foo = []int{2, 4, 6, 8}
	assert.Equal(t, []int{8, 6, 4, 2}, SliceReverse(foo))

	foo = []int{2, 4, 8}
	assert.Equal(t, []int{8, 4, 2}, SliceReverse(foo))
}
