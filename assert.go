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
	"bytes"
	"encoding"
	"fmt"
	"reflect"
	"strings"
)

// Error fails the test if err is nil.
func Error(t TestingT, err error, args ...any) bool {
	msgArgs := []any{"Expected error"}
	return !FailIf(
		t,
		err == nil,
		append(msgArgs, args...)...,
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
	msgArgs := []any{"Expected no error", err}
	return !FailIf(
		t,
		err != nil,
		append(msgArgs, args...)...,
	)
}

// Equal fails the test if the two values are not equal.
func Equal(t TestingT, expected any, actual any, args ...any) bool {
	var msgArgs []any
	if reflect.TypeOf(actual) != reflect.TypeOf(expected) {
		msgArgs = []any{"Expected type %v, actual type %v", reflect.TypeOf(expected), reflect.TypeOf(actual)}
	} else {
		msgArgs = []any{"Expected '%v', actual '%v'", v(expected), v(actual)}
	}
	return !FailIf(
		t,
		!reflect.DeepEqual(expected, actual),
		append(msgArgs, args...)...,
	)
}

// NotEqual fails the test if the two values are equal.
func NotEqual(t TestingT, expected any, actual any, args ...any) bool {
	msgArgs := []any{"Expected not to equal '%v'", v(expected)}
	return !FailIf(
		t,
		reflect.DeepEqual(expected, actual),
		append(msgArgs, args...)...,
	)
}

// Zero fails the test if the value is not the 0 value of its type.
// Nils are considered zero.
func Zero(t TestingT, actual any, args ...any) bool {
	msgArgs := []any{"Expected zero, actual '%v'", v(actual)}
	return !FailIf(
		t,
		!isNil(actual) && !reflect.ValueOf(actual).IsZero(),
		append(msgArgs, args...)...,
	)
}

// NotZero fails the test if the value is the 0 value of its type.
// Nils are considered zero.
func NotZero(t TestingT, actual any, args ...any) bool {
	msgArgs := []any{"Expected not to be zero, actual '%v'", v(actual)}
	return !FailIf(
		t,
		isNil(actual) || reflect.ValueOf(actual).IsZero(),
		append(msgArgs, args...)...,
	)
}

// True fails the test if the condition is false.
func True(t TestingT, condition bool, args ...any) bool {
	msgArgs := []any{"Expected condition to be true"}
	return !FailIf(
		t,
		!condition,
		append(msgArgs, args...)...,
	)
}

// False fails the test if the condition is true.
func False(t TestingT, condition bool, args ...any) bool {
	msgArgs := []any{"Expected condition to be false"}
	return !FailIf(
		t,
		condition,
		append(msgArgs, args...)...,
	)
}

// Contains fails the test if a string or error don't contain a substring,
// or if a byte slice doesn't contain a byte subslice,
// or if a slice doesn't contain an element,
// or if a map doesn't contain a key.
func Contains(t TestingT, whole any, sub any, args ...any) bool {
	if isNil(whole) {
		msgArgs := []any{"Nil does not contain '%v'", v(sub)}
		return !FailIf(
			t,
			true,
			append(msgArgs, args...)...,
		)
	}
	if err, ok := whole.(error); ok {
		whole = err.Error()
	}
	msgArgs := []any{"Expected '%v' to contain '%v'", v(whole), v(sub)}
	// Strings
	if w, ok := whole.(string); ok {
		if s, ok := sub.(string); ok {
			return !FailIf(
				t,
				!strings.Contains(w, s),
				append(msgArgs, args...)...,
			)
		}
	}
	// []byte
	if w, ok := whole.([]byte); ok {
		if s, ok := sub.([]byte); ok {
			return !FailIf(
				t,
				!bytes.Contains(w, s),
				append(msgArgs, args...)...,
			)
		}
	}
	wholeValue := reflect.ValueOf(whole)
	if wholeValue.Type().Kind() == reflect.Slice || wholeValue.Type().Kind() == reflect.Array {
		found := false
		for i := range wholeValue.Len() {
			if reflect.DeepEqual(wholeValue.Index(i).Interface(), sub) {
				found = true
				break
			}
		}
		return !FailIf(
			t,
			!found,
			append(msgArgs, args...)...,
		)
	}
	if wholeValue.Type().Kind() == reflect.Map {
		found := false
		for _, key := range wholeValue.MapKeys() {
			if reflect.DeepEqual(key.Interface(), sub) {
				found = true
				break
			}
		}
		return !FailIf(
			t,
			!found,
			append(msgArgs, args...)...,
		)
	}
	msgArgs = []any{"Type %v doesn't support containment", wholeValue.Type()}
	return !FailIf(
		t,
		true,
		append(msgArgs, args...)...,
	)
}

// NotContains fails the test if a string or error contain a substring,
// or if a byte slice contains a byte subslice,
// or if a slice contains an element,
// or if a map contains a key.
func NotContains(t TestingT, whole any, sub any, args ...any) bool {
	if isNil(whole) {
		return false
	}
	if err, ok := whole.(error); ok {
		whole = err.Error()
	}
	msgArgs := []any{"Expected '%v' not to contain '%v'", v(whole), v(sub)}
	// Strings
	if w, ok := whole.(string); ok {
		if s, ok := sub.(string); ok {
			return !FailIf(
				t,
				strings.Contains(w, s),
				append(msgArgs, args...)...,
			)
		}
	}
	// []byte
	if w, ok := whole.([]byte); ok {
		if s, ok := sub.([]byte); ok {
			return !FailIf(
				t,
				bytes.Contains(w, s),
				append(msgArgs, args...)...,
			)
		}
	}
	wholeValue := reflect.ValueOf(whole)
	if wholeValue.Type().Kind() == reflect.Slice || wholeValue.Type().Kind() == reflect.Array {
		found := false
		for i := range wholeValue.Len() {
			if reflect.DeepEqual(wholeValue.Index(i).Interface(), sub) {
				found = true
				break
			}
		}
		return !FailIf(
			t,
			found,
			append(msgArgs, args...)...,
		)
	}
	if wholeValue.Type().Kind() == reflect.Map {
		found := false
		for _, key := range wholeValue.MapKeys() {
			if reflect.DeepEqual(key.Interface(), sub) {
				found = true
				break
			}
		}
		return !FailIf(
			t,
			found,
			append(msgArgs, args...)...,
		)
	}
	msgArgs = []any{"Type %v doesn't support containment", wholeValue.Type()}
	return !FailIf(
		t,
		true,
		append(msgArgs, args...)...,
	)
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
	actualLength := 0
	if !isNil(obj) {
		objType := reflect.TypeOf(obj)
		hasLength := false ||
			objType.Kind() == reflect.Slice ||
			objType.Kind() == reflect.Array ||
			objType.Kind() == reflect.Map ||
			objType.Kind() == reflect.String ||
			objType.Kind() == reflect.Chan
		if FailIf(t, !hasLength, "Type %v doesn't have a length", objType) {
			return true
		}
		actualLength = reflect.ValueOf(obj).Len()
	}
	msgArgs := []any{"Expected length %d, actual %d", length, actualLength}
	return !FailIf(
		t,
		actualLength != length,
		append(msgArgs, args...)...,
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
	msgArgs := []any{"Expected nil, actual '%v'", v(obj)}
	return !FailIf(
		t,
		!isNil,
		append(msgArgs, args...)...,
	)
}

// NotNil fails the test if the object is nil.
func NotNil(t TestingT, obj any, args ...any) bool {
	isNil := isNil(obj)
	msgArgs := []any{"Expected not to be nil"}
	return !FailIf(
		t,
		isNil,
		append(msgArgs, args...)...,
	)
}

// v converts o to a string of no more than 1K in length.
func v(o any) string {
	truncate := func(s string) string {
		if len(s) <= 1024 {
			return s
		}
		rs := []rune(s)
		if len(rs) <= 1024 {
			return s
		}
		return string(append(rs[:1024], rune('\u2026')))
	}
	if s, ok := o.(string); ok {
		return truncate(s)
	}
	if tm, ok := o.(encoding.TextMarshaler); ok && !isNil(tm) {
		if txt, err := tm.MarshalText(); err == nil {
			return truncate(string(txt))
		}
	}
	if s, ok := o.(fmt.Stringer); ok && !isNil(s) {
		return truncate(s.String())
	}
	return truncate(fmt.Sprintf("%v", o))
}
