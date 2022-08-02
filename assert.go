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

import "reflect"

//
// Check if value is a primitve number or not. Returns
// `false` if value is `nil`.
//
func IsNumber(value interface{}) bool {
	if value == nil {
		return false
	}

	switch v := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		// sangupta: this is pretty stupid of the compiler to throw
		// an unused variable error, if `v` is not used here.
		if v != 0 {
			return true
		}

		return true

	}

	return false
}

//
// Check if value represents a slice or not. Returns
// `false` if value is `nil`.
//
func IsSlice(value interface{}) bool {
	if value == nil {
		return false
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice:
		return true
	}

	return false
}

//
// Check if value represents a `map` or not. Returns
// `false` if value is `nil`.
//
func IsMap(value interface{}) bool {
	if value == nil {
		return false
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Map:
		return true
	}

	return false
}
