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

func Benchmark_Equal(b *testing.B) {
	mt := &MockTestingT{}
	for b.Loop() {
		Equal(mt, 1, 1)
	}
}

func Benchmark_EqualSlice(b *testing.B) {
	mt := &MockTestingT{}
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4, 5}
	for b.Loop() {
		Equal(mt, slice1, slice2)
	}
}

func Benchmark_ContainsString(b *testing.B) {
	mt := &MockTestingT{}
	for b.Loop() {
		Contains(mt, "hello world", "world")
	}
}

func Benchmark_ContainsSlice(b *testing.B) {
	mt := &MockTestingT{}
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for b.Loop() {
		Contains(mt, slice, 5)
	}
}

func Benchmark_NoError(b *testing.B) {
	mt := &MockTestingT{}
	for b.Loop() {
		NoError(mt, nil)
	}
}

func Benchmark_Error(b *testing.B) {
	mt := &MockTestingT{}
	err := errors.New("test error")
	for b.Loop() {
		Error(mt, err)
	}
}

func Benchmark_HTMLMatch(b *testing.B) {
	mt := &MockTestingT{}
	htmlBody := []byte(`<html><body><div class="banner">Hello World</div></body></html>`)
	for b.Loop() {
		HTMLMatch(mt, htmlBody, "DIV.banner", "World")
	}
}

func Benchmark_Asserter(b *testing.B) {
	mt := &MockTestingT{}
	tt := For(mt)
	for b.Loop() {
		tt.Equal(1, 1)
		tt.True(true)
		tt.NotNil(1)
	}
}

func Benchmark_FailIf(b *testing.B) {
	mt := &MockTestingT{}
	for b.Loop() {
		FailIf(mt, false, "should not fail")
	}
}

func Benchmark_Len(b *testing.B) {
	mt := &MockTestingT{}
	slice := []int{1, 2, 3, 4, 5}
	for b.Loop() {
		Len(mt, slice, 5)
	}
}
