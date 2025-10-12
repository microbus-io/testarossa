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
	"errors"
	"testing"
)

func TestTableDriven_Equal(t *testing.T) {
	testCases := []struct {
		name     string
		expected any
		actual   any
		want     bool
	}{
		{"equal ints", 1, 1, true},
		{"unequal ints", 1, 2, false},
		{"equal strings", "hello", "hello", true},
		{"unequal strings", "hello", "world", false},
		{"equal slices", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"unequal slices", []int{1, 2, 3}, []int{3, 2, 1}, false},
		{"nil equals nil", nil, nil, true},
		{"nil pointer equals nil", (*int)(nil), nil, true},
		{"zero slice equals nil", []int(nil), nil, true},
		{"empty slice not equals nil slice", []int{}, []int(nil), false},
		{"different types", "1", 1, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mt := &MockTestingT{}
			result := Equal(mt, tc.expected, tc.actual)
			if result != tc.want {
				t.Fatalf("Expected %v", tc.want)
			}
		})
	}
}

func TestTableDriven_Contains(t *testing.T) {
	testCases := []struct {
		name  string
		whole any
		sub   any
		want  bool
	}{
		{"string contains substring", "hello world", "world", true},
		{"string does not contain substring", "hello world", "goodbye", false},
		{"slice contains element", []int{1, 2, 3}, 2, true},
		{"slice does not contain element", []int{1, 2, 3}, 4, false},
		{"map contains key", map[string]int{"a": 1, "b": 2}, "a", true},
		{"map does not contain key", map[string]int{"a": 1, "b": 2}, "c", false},
		{"error contains substring", errors.New("bad error"), "bad", true},
		{"error does not contain substring", errors.New("bad error"), "good", false},
		{"byte slice contains byte slice", []byte("hello"), []byte("ell"), true},
		{"byte slice does not contain byte slice", []byte("hello"), []byte("xyz"), false},
		{"array contains element", [3]int{1, 2, 3}, 2, true},
		{"array does not contain element", [3]int{1, 2, 3}, 4, false},
		{"string contains byte slice", "hello", []byte("ell"), true},
		{"byte slice contains string", []byte("hello"), "ell", true},
		{"nil container", nil, 1, false},
		{"unsupported type", 123, 1, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mt := &MockTestingT{}
			result := Contains(mt, tc.whole, tc.sub)
			if result != tc.want {
				t.Fatalf("Expected %v", tc.want)
			}
		})
	}
}

func TestTableDriven_Len(t *testing.T) {
	testCases := []struct {
		name   string
		obj    any
		length int
		want   bool
	}{
		{"nil has length 0", nil, 0, true},
		{"empty string", "", 0, true},
		{"string with content", "hello", 5, true},
		{"empty slice", []int{}, 0, true},
		{"slice with elements", []int{1, 2, 3}, 3, true},
		{"empty map", map[string]int{}, 0, true},
		{"map with elements", map[string]int{"a": 1, "b": 2}, 2, true},
		{"array", [5]int{1, 2, 3, 4, 5}, 5, true},
		{"wrong length for string", "hello", 3, false},
		{"wrong length for slice", []int{1, 2, 3}, 5, false},
		{"unsupported type int", 123, 1, false},
		{"unsupported type bool", true, 1, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mt := &MockTestingT{}
			result := Len(mt, tc.obj, tc.length)
			if result != tc.want {
				t.Fatalf("Expected %v", tc.want)
			}
		})
	}
}

func TestTableDriven_Nil(t *testing.T) {
	i := 123
	testCases := []struct {
		name      string
		obj       any
		shouldNil bool
	}{
		{"nil is nil", nil, true},
		{"nil pointer is nil", (*int)(nil), true},
		{"nil slice is nil", []int(nil), true},
		{"nil map is nil", map[string]int(nil), true},
		{"nil error is nil", error(nil), true},
		{"integer is not nil", 123, false},
		{"string is not nil", "hello", false},
		{"bool is not nil", true, false},
		{"struct is not nil", struct{}{}, false},
		{"pointer to int is not nil", &i, false},
		{"empty string is not nil", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mt := &MockTestingT{}
			result := Nil(mt, tc.obj)
			if result != tc.shouldNil {
				t.Fatalf("Expected %v", tc.shouldNil)
			}
		})
	}
}

func TestTableDriven_Zero(t *testing.T) {
	testCases := []struct {
		name       string
		obj        any
		shouldZero bool
	}{
		{"zero int", 0, true},
		{"non-zero int", 1, false},
		{"empty string", "", true},
		{"non-empty string", "x", false},
		{"false bool", false, true},
		{"true bool", true, false},
		{"nil", nil, true},
		{"empty struct", struct{}{}, true},
		{"non-empty struct", struct{ x int }{5}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mt := &MockTestingT{}
			result := Zero(mt, tc.obj)
			if result != tc.shouldZero {
				t.Fatalf("Expected %v", tc.shouldZero)
			}
		})
	}
}

func TestTableDriven_HTMLMatch(t *testing.T) {
	testCases := []struct {
		name     string
		html     []byte
		selector string
		regex    string
		want     bool
	}{
		{"element exists", []byte(`<html><div>test</div></html>`), "DIV", "", true},
		{"element does not exist", []byte(`<html><span>test</span></html>`), "DIV", "", false},
		{"element with class", []byte(`<html><div class="main">test</div></html>`), "DIV.main", "", true},
		{"element with id", []byte(`<html><div id="header">test</div></html>`), "DIV#header", "", true},
		{"text matches", []byte(`<html><div>hello world</div></html>`), "DIV", "world", true},
		{"text does not match", []byte(`<html><div>hello world</div></html>`), "DIV", "goodbye", false},
		{"regex matches", []byte(`<html><div>test123</div></html>`), "DIV", `\d+`, true},
		{"regex does not match", []byte(`<html><div>test</div></html>`), "DIV", `\d+`, false},
		{"nested text", []byte(`<html><div><b>bold</b></div></html>`), "DIV", "bold", true},
		{"invalid selector", []byte(`<html><div>test</div></html>`), "DIV.", "", false},
		{"invalid regex", []byte(`<html><div>test</div></html>`), "DIV", "[unclosed", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mt := &MockTestingT{}
			result := HTMLMatch(mt, tc.html, tc.selector, tc.regex)
			if result != tc.want {
				t.Fatalf("Expected %v", tc.want)
			}
		})
	}
}
