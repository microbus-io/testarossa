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

import "testing"

type Asserter struct {
	t TestingT
}

func For(t *testing.T) *Asserter {
	return &Asserter{t: t}
}

// Error fails the test if err is nil.
func (a *Asserter) Error(err error, args ...any) bool {
	return Error(a.t, err, args...)
}

// NoError fails the test if err is not nil.
func (a *Asserter) NoError(err error, args ...any) bool {
	return NoError(a.t, err, args...)
}

// Equal fails the test if the two values are not equal.
func (a *Asserter) Equal(expected any, actual any, args ...any) bool {
	return Equal(a.t, expected, actual, args...)
}

// NotEqual fails the test if the two values are equal.
func (a *Asserter) NotEqual(expected any, actual any, args ...any) bool {
	return NotEqual(a.t, expected, actual, args...)
}

// Zero fails the test if the value is not 0.
func (a *Asserter) Zero(actual any, args ...any) bool {
	return Zero(a.t, actual, args...)
}

// NotZero fails the test if the value is the 0 value of its type.
func (a *Asserter) NotZero(actual any, args ...any) bool {
	return NotZero(a.t, actual, args...)
}

// True fails the test if the condition is false.
func (a *Asserter) True(condition bool, args ...any) bool {
	return True(a.t, condition, args...)
}

// False fails the test if the condition is true.
func (a *Asserter) False(condition bool, args ...any) bool {
	return False(a.t, condition, args...)
}

// Contains fails the test if the string does not contain a substring.
func (a *Asserter) Contains(whole any, sub any, args ...any) bool {
	return Contains(a.t, whole, sub, args...)
}

// NotContains fails the test if the string contain a substring.
func (a *Asserter) NotContains(whole any, sub any, args ...any) bool {
	return NotContains(a.t, whole, sub, args...)
}

// Len fails the test if the length of the string, slice, array, map or chan does not match the expected len.
func (a *Asserter) Len(obj any, length int, args ...any) bool {
	return Len(a.t, obj, length, args...)
}

// Nil fails the test if the object is not nil.
func (a *Asserter) Nil(obj any, args ...any) bool {
	return Nil(a.t, obj, args...)
}

// NotNil fails the test if the object is nil.
func (a *Asserter) NotNil(obj any, args ...any) bool {
	return NotNil(a.t, obj, args...)
}
