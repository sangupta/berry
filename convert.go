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
	"fmt"
	"strconv"
	"strings"
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
	case int:
		data, _ := value.(int)
		return data != 0, nil

	case int8:
		data, _ := value.(int8)
		return data != 0, nil

	case int16:
		data, _ := value.(int16)
		return data != 0, nil

	case int32:
		data, _ := value.(int32)
		return data != 0, nil

	case int64:
		data, _ := value.(int64)
		return data != 0, nil

	case uint:
		data, _ := value.(uint)
		return data != 0, nil

	case uint8:
		data, _ := value.(uint8)
		return data != 0, nil

	case uint16:
		data, _ := value.(uint16)
		return data != 0, nil

	case uint32:
		data, _ := value.(uint32)
		return data != 0, nil

	case uint64:
		data, _ := value.(uint64)
		return data != 0, nil

	case float32:
		data, _ := value.(float32)
		return data != 0, nil

	case float64:
		data, _ := value.(float64)
		return data != 0, nil

	case string:
		v = strings.ToLower(strings.TrimSpace(v))
		switch v {
		case "1", "true", "t":
			return true, nil
		default:
			return false, nil
		}

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
	case int:
		data, _ := value.(int)
		return uint64(data), nil

	case int8:
		data, _ := value.(int8)
		return uint64(data), nil

	case int16:
		data, _ := value.(int16)
		return uint64(data), nil

	case int32:
		data, _ := value.(int32)
		return uint64(data), nil

	case int64:
		data, _ := value.(int64)
		return uint64(data), nil

	case uint:
		data, _ := value.(uint)
		return uint64(data), nil

	case uint8:
		data, _ := value.(uint8)
		return uint64(data), nil

	case uint16:
		data, _ := value.(uint16)
		return uint64(data), nil

	case uint32:
		data, _ := value.(uint32)
		return uint64(data), nil

	case uint64:
		data, _ := value.(uint64)
		return uint64(data), nil

	case float32:
		data, _ := value.(float32)
		return uint64(data), nil

	case float64:
		data, _ := value.(float64)
		return uint64(data), nil

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
	case int:
		data, _ := value.(int)
		return int64(data), nil

	case int8:
		data, _ := value.(int8)
		return int64(data), nil

	case int16:
		data, _ := value.(int16)
		return int64(data), nil

	case int32:
		data, _ := value.(int32)
		return int64(data), nil

	case int64:
		data, _ := value.(int64)
		return int64(data), nil

	case uint:
		data, _ := value.(uint)
		return int64(data), nil

	case uint8:
		data, _ := value.(uint8)
		return int64(data), nil

	case uint16:
		data, _ := value.(uint16)
		return int64(data), nil

	case uint32:
		data, _ := value.(uint32)
		return int64(data), nil

	case uint64:
		data, _ := value.(uint64)
		return int64(data), nil

	case float32:
		data, _ := value.(float32)
		return int64(data), nil

	case float64:
		data, _ := value.(float64)
		return int64(data), nil

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
	case int:
		data, _ := value.(int)
		return float64(data), nil

	case int8:
		data, _ := value.(int8)
		return float64(data), nil

	case int16:
		data, _ := value.(int16)
		return float64(data), nil

	case int32:
		data, _ := value.(int32)
		return float64(data), nil

	case int64:
		data, _ := value.(int64)
		return float64(data), nil

	case uint:
		data, _ := value.(uint)
		return float64(data), nil

	case uint8:
		data, _ := value.(uint8)
		return float64(data), nil

	case uint16:
		data, _ := value.(uint16)
		return float64(data), nil

	case uint32:
		data, _ := value.(uint32)
		return float64(data), nil

	case uint64:
		data, _ := value.(uint64)
		return float64(data), nil

	case float32:
		data, _ := value.(float32)
		return float64(data), nil

	case float64:
		data, _ := value.(float64)
		return float64(data), nil

	case string:
		return strconv.ParseFloat(v, 64)
	}

	// any non-nil struct is considered true
	return defaultValue, nil
}
