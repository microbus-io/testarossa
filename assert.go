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
	"fmt"
	"reflect"
	"strings"
)

// Error fails the test if err is nil.
func Error(t TestingT, err error, msgAndArgs ...any) bool {
	isNil := err == nil || reflect.ValueOf(err).IsNil()
	return !FailIf(
		t,
		isNil,
		"Expected error",
		formatMsgAndArgs(msgAndArgs...),
	)
}

// ErrorContains fails the test if the err is nil or if it does not contain the substring.
func ErrorContains(t TestingT, err error, substr string, msgAndArgs ...any) bool {
	isNil := err == nil || reflect.ValueOf(err).IsNil()
	return !FailIf(
		t,
		isNil || !strings.Contains(err.Error(), substr),
		fmt.Sprintf("Expected error to contain '%s'", substr),
		formatMsgAndArgs(msgAndArgs...),
	)
}

// NoError fails the test if err is not nil.
func NoError(t TestingT, err error, msgAndArgs ...any) bool {
	isNil := err == nil || reflect.ValueOf(err).IsNil()
	return !FailIf(
		t,
		!isNil,
		"Expected no error",
		err,
		formatMsgAndArgs(msgAndArgs...),
	)
}

// Equal fails the test if the two values are not equal.
func Equal(t TestingT, expected any, actual any, msgAndArgs ...any) bool {
	return !FailIf(
		t,
		!reflect.DeepEqual(expected, actual),
		fmt.Sprintf("Expected %v, actual %v", expected, actual),
		formatMsgAndArgs(msgAndArgs...),
	)
}

// NotEqual fails the test if the two values are equal.
func NotEqual(t TestingT, expected any, actual any, msgAndArgs ...any) bool {
	return !FailIf(
		t,
		reflect.DeepEqual(expected, actual),
		fmt.Sprintf("Expected actual to differ from %v", expected),
		formatMsgAndArgs(msgAndArgs...),
	)
}

// Zero fails the test if the value is not 0.
func Zero[T int8 | int16 | int32 | int64 | int | float32 | float64](t TestingT, actual T, msgAndArgs ...any) bool {
	return !FailIf(
		t,
		actual != T(0),
		fmt.Sprintf("Expected 0, actual %v", actual),
		formatMsgAndArgs(msgAndArgs...),
	)
}

// NotZero fails the test if the value is 0.
func NotZero[T int8 | int16 | int32 | int64 | int | float32 | float64](t TestingT, actual T, msgAndArgs ...any) bool {
	return !FailIf(
		t,
		actual == T(0),
		fmt.Sprintf("Expected 0, actual %v", actual),
		formatMsgAndArgs(msgAndArgs...),
	)
}

// True fails the test if the condition is false.
func True(t TestingT, condition bool, msgAndArgs ...any) bool {
	return !FailIf(
		t,
		!condition,
		"Expected condition to be true",
		formatMsgAndArgs(msgAndArgs...),
	)
}

// False fails the test if the condition is true.
func False(t TestingT, condition bool, msgAndArgs ...any) bool {
	return !FailIf(
		t,
		condition,
		"Expected condition to be true",
		formatMsgAndArgs(msgAndArgs...),
	)
}

// Contains fails the test if the string does not contain a substring.
func Contains(t TestingT, s string, substr string, msgAndArgs ...any) bool {
	return !FailIf(
		t,
		!strings.Contains(s, substr),
		fmt.Sprintf("Expected '%v' to contain '%v'", s, substr),
		formatMsgAndArgs(msgAndArgs...),
	)
}

// NotContains fails the test if the string contain a substring.
func NotContains(t TestingT, s string, substr string, msgAndArgs ...any) bool {
	return !FailIf(
		t,
		strings.Contains(s, substr),
		fmt.Sprintf("Expected '%v' not to contain '%v'", s, substr),
		formatMsgAndArgs(msgAndArgs...),
	)
}

// SliceContains fails the test if the slice does not contain the item.
func SliceContains[T any](t TestingT, slice []T, contains T, msgAndArgs ...any) bool {
	found := false
	for i := 0; i < len(slice) && !found; i++ {
		found = reflect.DeepEqual(slice[i], contains)
	}
	return !FailIf(
		t,
		!found,
		fmt.Sprintf("'%v' not found in %v", contains, slice),
		formatMsgAndArgs(msgAndArgs...),
	)
}

// SliceNotContains fails the test if the slice contains the item.
func SliceNotContains[T any](t TestingT, slice []T, contains T, msgAndArgs ...any) bool {
	found := false
	for i := 0; i < len(slice) && !found; i++ {
		found = reflect.DeepEqual(slice[i], contains)
	}
	return !FailIf(
		t,
		found,
		fmt.Sprintf("'%v' found in %v", contains, slice),
		formatMsgAndArgs(msgAndArgs...),
	)
}

// SliceLen fails the test if the length of the slice does not match the expected len.
func SliceLen[T any](t TestingT, slice []T, length int, msgAndArgs ...any) bool {
	return !FailIf(
		t,
		len(slice) != length,
		fmt.Sprintf("Expected length %d, actual %d", length, len(slice)),
		formatMsgAndArgs(msgAndArgs...),
	)
}

// StrLen fails the test if the length of the string does not match the expected len.
func StrLen(t TestingT, s string, length int, msgAndArgs ...any) bool {
	return !FailIf(
		t,
		len(s) != length,
		fmt.Sprintf("Expected '%s' to be of length %d, actual %d", s, length, len(s)),
		formatMsgAndArgs(msgAndArgs...),
	)
}

// SliceEqual fails the test if the two values are not equal.
func SliceEqual[T comparable](t TestingT, expected []T, actual []T, msgAndArgs ...any) bool {
	eq := len(expected) == len(actual)
	for i := 0; i < len(expected) && eq; i++ {
		eq = actual[i] == expected[i]
	}
	return !FailIf(
		t,
		!eq,
		fmt.Sprintf("Expected %v, actual %v", expected, actual),
		formatMsgAndArgs(msgAndArgs...),
	)
}

// SliceNotEqual fails the test if the two values are equal.
func SliceNotEqual[T comparable](t TestingT, expected []T, actual []T, msgAndArgs ...any) bool {
	eq := len(expected) == len(actual)
	for i := 0; i < len(expected) && eq; i++ {
		eq = actual[i] == expected[i]
	}
	return !FailIf(
		t,
		eq,
		fmt.Sprintf("Expected actual to differ from %v", expected),
		formatMsgAndArgs(msgAndArgs...),
	)
}

// Nil fails the test if the object is not nil.
func Nil(t TestingT, obj any, msgAndArgs ...any) bool {
	isNil := obj == nil || reflect.ValueOf(obj).IsNil()
	return !FailIf(
		t,
		!isNil,
		fmt.Sprintf("Expected nil, actual %v", obj),
		formatMsgAndArgs(msgAndArgs...),
	)
}

// NotNil fails the test if the object is nil.
func NotNil(t TestingT, obj any, msgAndArgs ...any) bool {
	isNil := obj == nil || reflect.ValueOf(obj).IsNil()
	return !FailIf(
		t,
		isNil,
		"Expected object not to be nil",
		formatMsgAndArgs(msgAndArgs...),
	)
}

func formatMsgAndArgs(msgAndArgs ...any) string {
	if len(msgAndArgs) == 0 {
		return ""
	}
	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		if msgAsStr, ok := msg.(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msg)
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}
