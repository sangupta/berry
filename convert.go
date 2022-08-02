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
	"errors"
	"fmt"
	"strconv"
)

//
// Convert the given value represented as `interface{}`
// to boolean. Returns `false` if value is `nil`. Returns
// `true` if the value is numeric and not-zero, or if the
// value is `string` and represents a `bool`, or if the
// value is non-primitive.
//
func ConvertToBool(value interface{}) (bool, error) {
	if value == nil {
		return false, nil
	}

	switch v := value.(type) {
	case int, int8, int16, int32, int64:
		return v != 0, nil

	case uint, uint8, uint16, uint32, uint64:
		return v != 0, nil

	case float32, float64:
		return v != 0, nil

	case string:
		return strconv.ParseBool(v)

	case bool:
		return v, nil
	}

	// any non-nil struct is considered true
	return false, nil
}

//
// Convert the given value represented as `interface{}`
// to its string representation.
//
func ConvertToString(value interface{}) string {
	if value == nil {
		return ""
	}

	return fmt.Sprintf("%v", value)
}

//
// Convert the given value represented as `interface{}`
// to a `uint64` value. Returns the `defaultValue` is the value
// is `nil`. Returns an `error` if the value cannot be
// cast to a number. For non-primitive values, the `defaultValue`
// is returned.
//
func ConvertToUint64(value interface{}, defaultValue uint64) (uint64, error) {
	if value == nil {
		return defaultValue, nil
	}

	switch v := value.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		c, err := v.(uint64)
		if err {
			return c, errors.New("Unable to convert value to uint64")
		}

	case string:
		return strconv.ParseUint(v, 10, 64)
	}

	// any non-nil struct is considered true
	return defaultValue, nil
}

//
// Convert the given value represented as `interface{}`
// to a `int64` value. Returns the `defaultValue` is the value
// is `nil`. Returns an `error` if the value cannot be
// cast to a number. For non-primitive values, the `defaultValue`
// is returned.
//
func ConvertToInt64(value interface{}, defaultValue int64) (int64, error) {
	if value == nil {
		return defaultValue, nil
	}

	switch v := value.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		c, ok := v.(int64)
		if !ok {
			return c, errors.New("Unable to convert value to int64")
		}

	case string:
		return strconv.ParseInt(v, 10, 64)
	}

	// any non-nil struct is considered true
	return defaultValue, nil
}

//
// Convert the given value represented as `interface{}`
// to a `float64` value. Returns the `defaultValue` is the value
// is `nil`. Returns an `error` if the value cannot be
// cast to a number. For non-primitive values, the `defaultValue`
// is returned.
//
func ConvertToFloat64(value interface{}, defaultValue float64) (float64, error) {
	if value == nil {
		return defaultValue, nil
	}

	switch v := value.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		c, ok := v.(float64)
		if !ok {
			return c, errors.New("Unable to convert value to float64")
		}

	case string:
		return strconv.ParseFloat(v, 64)
	}

	// any non-nil struct is considered true
	return defaultValue, nil
}
