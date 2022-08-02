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

import "errors"

//
// Defines the primitive types that define the
// `Number` interface.
//
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

//
// Find the minimum value from the slice of numbers. Returns
// an `error` if the slice is nil, or is empty.
//
func SliceMax[T Number](slice []T) (T, error) {
	if slice == nil {
		return 0, errors.New("Cannot find max on a nil slice")
	}

	length := len(slice)
	if length == 0 {
		return 0, errors.New("Cannot find min on an empty slice")
	}

	if length == 1 {
		return slice[0], nil
	}

	max := slice[0]
	for index := 1; index < length; index++ {
		value := slice[index]
		if max < value {
			max = value
		}
	}

	return max, nil
}

//
// Find the minimum value from the slice of numbers. Returns
// an `error` if the slice is nil, or is empty.
//
func SliceMin[T Number](slice []T) (T, error) {
	if slice == nil {
		return 0, errors.New("Cannot find max on a nil slice")
	}

	length := len(slice)
	if length == 0 {
		return 0, errors.New("Cannot find min on an empty slice")
	}

	if length == 1 {
		return slice[0], nil
	}

	min := slice[0]
	for index := 1; index < length; index++ {
		value := slice[index]
		if min > value {
			min = value
		}
	}

	return min, nil
}

//
// Check if the slice contains the given element or not. Returns
// `false` if the slice is `nil`, or empty or the element cannot
// be found.
//
func SliceContains[T comparable](slice []T, element T) bool {
	if slice == nil {
		return false
	}

	length := len(slice)
	if length == 0 {
		return false
	}

	for _, val := range slice {
		if val == element {
			return true
		}
	}

	return false
}

//
// Find the index of the given element in the slice. The
// functions returns `-1` if the slice is `nil` or empty
// or the element cannot be found.
//
func SliceIndex[T comparable](slice []T, element T) int {
	if slice == nil {
		return -1
	}

	length := len(slice)
	if length == 0 {
		return -1
	}

	for index, val := range slice {
		if val == element {
			return index
		}
	}

	return -1
}

//
// Modify the slice using a modifier function. Each element of
// the slice is sent to the modifier, and the returned value is
// updated in the slice at same index. The method has no effect
// if the slice is `nil`.
//
func SliceModify[T comparable](slice []T, modifier func(item T) T) {
	if slice == nil {
		return
	}

	for index := 0; index < len(slice); index++ {
		val := slice[index]
		updated := modifier(val)
		slice[index] = updated
	}
}

//
// Check if 2 slices are equal or not. The slices are considered
// equal if and only if there are of same length, and elements at
// the same index in slices are equal. If any of the slice is `nil`
// will return a `false`.
//
func AreSliceEqual[T comparable](slice1 []T, slice2 []T) bool {
	if slice1 == nil || slice2 == nil {
		return false
	}

	if len(slice1) != len(slice2) {
		return false
	}

	for index, val := range slice1 {
		if slice2[index] != val {
			return false
		}
	}

	return true
}
