/*
Copyright 2024-2025 Microbus LLC and various contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testarossa

import (
	"reflect"
	"slices"
	"strings"
)

// Error fails the test if err is nil.
func Error(t TestingT, err error, args ...any) bool {
	isNil := err == nil
	if len(args) == 0 {
		args = []any{"Expected error"}
	}
	return !FailIf(
		t,
		isNil,
		args...,
	)
}

// ErrorContains fails the test if the err is nil or if it does not contain the substring.
//
// Deprecated: Use Contains
func ErrorContains(t TestingT, err error, substr string, args ...any) bool {
	return Contains(t, err, substr, args...)
}

// NoError fails the test if err is not nil.
func NoError(t TestingT, err error, args ...any) bool {
	isNil := err == nil
	if len(args) == 0 {
		args = []any{"Expected no error", err}
	}
	return !FailIf(
		t,
		!isNil,
		args...,
	)
}

// Equal fails the test if the two values are not equal.
func Equal(t TestingT, expected any, actual any, args ...any) bool {
	if len(args) == 0 {
		if reflect.TypeOf(actual) != reflect.TypeOf(expected) {
			args = []any{"Expected type %v, actual type %v", reflect.TypeOf(expected), reflect.TypeOf(actual)}
		} else {
			args = []any{"Expected '%v', actual '%v'", expected, actual}
		}
	}
	return !FailIf(
		t,
		!reflect.DeepEqual(expected, actual),
		args...,
	)
}

// NotEqual fails the test if the two values are equal.
func NotEqual(t TestingT, expected any, actual any, args ...any) bool {
	if len(args) == 0 {
		args = []any{"Expected not to equal '%v'", expected}
	}
	return !FailIf(
		t,
		reflect.DeepEqual(expected, actual),
		args...,
	)
}

// Zero fails the test if the value is not the 0 value of its type.
func Zero(t TestingT, actual any, args ...any) bool {
	if len(args) == 0 {
		args = []any{"Expected zero, actual '%v'", actual}
	}
	return !FailIf(
		t,
		!reflect.ValueOf(actual).IsZero(),
		args...,
	)
}

// NotZero fails the test if the value is the 0 value of its type.
func NotZero(t TestingT, actual any, args ...any) bool {
	if len(args) == 0 {
		args = []any{"Expected not to be zero, actual '%v'", actual}
	}
	return !FailIf(
		t,
		reflect.ValueOf(actual).IsZero(),
		args...,
	)
}

// True fails the test if the condition is false.
func True(t TestingT, condition bool, args ...any) bool {
	if len(args) == 0 {
		args = []any{"Expected condition to be true"}
	}
	return !FailIf(
		t,
		!condition,
		args...,
	)
}

// False fails the test if the condition is true.
func False(t TestingT, condition bool, args ...any) bool {
	if len(args) == 0 {
		args = []any{"Expected condition to be false"}
	}
	return !FailIf(
		t,
		condition,
		args...,
	)
}

// Contains fails the test if a string or error don't contain a substring,
// or if a slice doesn't contain an element
// or if a map doesn't contain a key.
func Contains(t TestingT, whole any, sub any, args ...any) bool {
	if isNil(whole) {
		if len(args) == 0 {
			args = []any{"Nil does not contain '%v'", sub}
		}
		return !FailIf(
			t,
			true,
			args...,
		)
	}
	if err, ok := whole.(error); ok {
		whole = err.Error()
	}
	if len(args) == 0 {
		args = []any{"Expected '%v' to contain '%v'", whole, sub}
	}
	wholeValue := reflect.ValueOf(whole)
	subValue := reflect.ValueOf(sub)
	if wholeValue.Type().Kind() == reflect.String && subValue.Type().Kind() == reflect.String {
		return !FailIf(
			t,
			!strings.Contains(wholeValue.String(), subValue.String()),
			args...,
		)
	}
	if wholeValue.Type().Kind() == reflect.Slice || wholeValue.Type().Kind() == reflect.Array {
		arr := make([]any, wholeValue.Len())
		for i := range wholeValue.Len() {
			arr = append(arr, wholeValue.Index(i).Interface())
		}
		return !FailIf(
			t,
			!slices.Contains(arr, sub),
			args...,
		)
	}
	if wholeValue.Type().Kind() == reflect.Map {
		arr := make([]any, wholeValue.Len())
		for _, key := range wholeValue.MapKeys() {
			arr = append(arr, key.Interface())
		}
		return !FailIf(
			t,
			!slices.Contains(arr, sub),
			args...,
		)
	}
	if len(args) == 0 {
		args = []any{"Type %v doesn't support containment", wholeValue.Type()}
	}
	return !FailIf(t, true, args...)
}

// NotContains fails the test if a string or error contain a substring,
// or if a slice contains an element
// or if a map contains a key.
func NotContains(t TestingT, whole any, sub any, args ...any) bool {
	if isNil(whole) {
		return false
	}
	if err, ok := whole.(error); ok {
		whole = err.Error()
	}
	if len(args) == 0 {
		args = []any{"Expected '%v' not to contain '%v'", whole, sub}
	}
	wholeValue := reflect.ValueOf(whole)
	subValue := reflect.ValueOf(sub)
	if wholeValue.Type().Kind() == reflect.String && subValue.Type().Kind() == reflect.String {
		return !FailIf(
			t,
			strings.Contains(wholeValue.String(), subValue.String()),
			args...,
		)
	}
	if wholeValue.Type().Kind() == reflect.Slice || wholeValue.Type().Kind() == reflect.Array {
		arr := make([]any, wholeValue.Len())
		for i := range wholeValue.Len() {
			arr = append(arr, wholeValue.Index(i).Interface())
		}
		return !FailIf(
			t,
			slices.Contains(arr, sub),
			args...,
		)
	}
	if wholeValue.Type().Kind() == reflect.Map {
		arr := make([]any, wholeValue.Len())
		for _, key := range wholeValue.MapKeys() {
			arr = append(arr, key.Interface())
		}
		return !FailIf(
			t,
			slices.Contains(arr, sub),
			args...,
		)
	}
	if len(args) == 0 {
		args = []any{"Type %v doesn't support containment", wholeValue.Type()}
	}
	return !FailIf(t, true, args...)
}

// SliceContains fails the test if the slice does not contain the item.
//
// Deprecated: Use Contains
func SliceContains[T comparable](t TestingT, slice []T, contains T, args ...any) bool {
	return Contains(t, slice, contains, args...)
}

// SliceNotContains fails the test if the slice contains the item.
//
// Deprecated: Use NotContains
func SliceNotContains[T comparable](t TestingT, slice []T, contains T, args ...any) bool {
	return NotContains(t, slice, contains, args...)
}

// SliceLen fails the test if the length of the slice does not match the expected len.
//
// Deprecated: Use Len
func SliceLen[T any](t TestingT, slice []T, length int, args ...any) bool {
	return Len(t, slice, length, args...)
}

// StrLen fails the test if the length of the string does not match the expected len.
//
// Deprecated: Use Len
func StrLen(t TestingT, s string, length int, args ...any) bool {
	return Len(t, s, length, args...)
}

// MapLen fails the test if the length of the map does not match the expected len.
//
// Deprecated: Use Len
func MapLen[K comparable, V any](t TestingT, m map[K]V, length int, args ...any) bool {
	return Len(t, m, length, args...)
}

// Len fails the test if the length of the string, slice, array, map or chan does not match the expected len.
func Len(t TestingT, obj any, length int, args ...any) bool {
	objType := reflect.TypeOf(obj)
	if FailIf(t, objType.Kind() != reflect.Slice && objType.Kind() != reflect.Array && objType.Kind() != reflect.Map &&
		objType.Kind() != reflect.String && objType.Kind() != reflect.Chan, "%v doesn't have a length") {
		return true
	}
	actualLength := reflect.ValueOf(obj).Len()
	if len(args) == 0 {
		args = []any{"Expected length %d, actual %d", length, actualLength}
	}
	return !FailIf(
		t,
		actualLength != length,
		args...,
	)
}

// SliceEqual fails the test if the two values are not equal.
//
// Deprecated: Use Equal
func SliceEqual[T comparable](t TestingT, expected []T, actual []T, args ...any) bool {
	return Equal(t, expected, actual, args...)
}

// SliceNotEqual fails the test if the two values are equal.
//
// Deprecated: Use NotEqual
func SliceNotEqual[T comparable](t TestingT, expected []T, actual []T, args ...any) bool {
	return NotEqual(t, expected, actual, args...)
}

// isNil checks for nil of an interface.
func isNil(obj any) bool {
	defer func() { recover() }()
	return obj == nil || reflect.ValueOf(obj).IsNil()
}

// Nil fails the test if the object is not nil.
func Nil(t TestingT, obj any, args ...any) bool {
	isNil := isNil(obj)
	if len(args) == 0 {
		args = []any{"Expected nil, actual '%v'", obj}
	}
	return !FailIf(
		t,
		!isNil,
		args...,
	)
}

// NotNil fails the test if the object is nil.
func NotNil(t TestingT, obj any, args ...any) bool {
	isNil := isNil(obj)
	if len(args) == 0 {
		args = []any{"Expected not to be nil"}
	}
	return !FailIf(
		t,
		isNil,
		args...,
	)
}
