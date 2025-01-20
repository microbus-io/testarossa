/*
Copyright 2024 Microbus LLC and various contributors

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
func ErrorContains(t TestingT, err error, substr string, args ...any) bool {
	isNil := err == nil
	if len(args) == 0 {
		if isNil {
			args = []any{"Expected error to contain '%v'", substr}
		} else {
			args = []any{"Expected error '%v' to contain '%v'", err, substr}
		}
	}
	return !FailIf(
		t,
		isNil || !strings.Contains(err.Error(), substr),
		args...,
	)
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
		args = []any{"Expected %v, actual %v", expected, actual}
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
		args = []any{"Expected actual not to equal %v", expected}
	}
	return !FailIf(
		t,
		reflect.DeepEqual(expected, actual),
		args...,
	)
}

// Zero fails the test if the value is not 0.
func Zero[T int8 | int16 | int32 | int64 | int | float32 | float64](t TestingT, actual T, args ...any) bool {
	if len(args) == 0 {
		args = []any{"Expected 0, actual %v", actual}
	}
	return !FailIf(
		t,
		actual != T(0),
		args...,
	)
}

// NotZero fails the test if the value is 0.
func NotZero[T int8 | int16 | int32 | int64 | int | float32 | float64](t TestingT, actual T, args ...any) bool {
	if len(args) == 0 {
		args = []any{"Expected not 0"}
	}
	return !FailIf(
		t,
		actual == T(0),
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

// Contains fails the test if the string does not contain a substring.
func Contains(t TestingT, s string, substr string, args ...any) bool {
	if len(args) == 0 {
		args = []any{"Expected '%v' to contain '%v'", s, substr}
	}
	return !FailIf(
		t,
		!strings.Contains(s, substr),
		args...,
	)
}

// NotContains fails the test if the string contain a substring.
func NotContains(t TestingT, s string, substr string, args ...any) bool {
	if len(args) == 0 {
		args = []any{"Expected '%v' not to contain '%v'", s, substr}
	}
	return !FailIf(
		t,
		strings.Contains(s, substr),
		args...,
	)
}

// SliceContains fails the test if the slice does not contain the item.
func SliceContains[T any](t TestingT, slice []T, contains T, args ...any) bool {
	found := false
	for i := 0; i < len(slice) && !found; i++ {
		found = reflect.DeepEqual(slice[i], contains)
	}
	if len(args) == 0 {
		args = []any{"Expected to find '%v' in %v", contains, slice}
	}
	return !FailIf(
		t,
		!found,
		args...,
	)
}

// SliceNotContains fails the test if the slice contains the item.
func SliceNotContains[T any](t TestingT, slice []T, contains T, args ...any) bool {
	found := false
	for i := 0; i < len(slice) && !found; i++ {
		found = reflect.DeepEqual(slice[i], contains)
	}
	if len(args) == 0 {
		args = []any{"Expected not to find '%v' in %v", contains, slice}
	}
	return !FailIf(
		t,
		found,
		args...,
	)
}

// SliceLen fails the test if the length of the slice does not match the expected len.
func SliceLen[T any](t TestingT, slice []T, length int, args ...any) bool {
	if len(args) == 0 {
		args = []any{"Expected length %d, actual %d", length, len(slice)}
	}
	return !FailIf(
		t,
		len(slice) != length,
		args...,
	)
}

// StrLen fails the test if the length of the string does not match the expected len.
func StrLen(t TestingT, s string, length int, args ...any) bool {
	if len(args) == 0 {
		args = []any{"Expected '%v' to be of length %d, actual %d", s, length, len(s)}
	}
	return !FailIf(
		t,
		len(s) != length,
		args...,
	)
}

// SliceEqual fails the test if the two values are not equal.
func SliceEqual[T comparable](t TestingT, expected []T, actual []T, args ...any) bool {
	eq := len(expected) == len(actual)
	for i := 0; i < len(expected) && eq; i++ {
		eq = actual[i] == expected[i]
	}
	if len(args) == 0 {
		args = []any{"Expected %v, actual %v", expected, actual}
	}
	return !FailIf(
		t,
		!eq,
		args...,
	)
}

// SliceNotEqual fails the test if the two values are equal.
func SliceNotEqual[T comparable](t TestingT, expected []T, actual []T, args ...any) bool {
	eq := len(expected) == len(actual)
	for i := 0; i < len(expected) && eq; i++ {
		eq = actual[i] == expected[i]
	}
	if len(args) == 0 {
		args = []any{"Expected actual to not equal %v", expected}
	}
	return !FailIf(
		t,
		eq,
		args...,
	)
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
		args = []any{"Expected nil, actual %v", obj}
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
