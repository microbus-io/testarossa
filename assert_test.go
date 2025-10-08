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
	"time"
)

func Test_Equality(t *testing.T) {
	mt := &MockTestingT{}

	if Equal(mt, 1, 0, "You are not the %d", 1) || mt.Passed() {
		t.FailNow()
	}

	if !Equal(mt, 1, 1, "All things equal") || mt.Failed() {
		t.FailNow()
	}

	if Equal(mt, "foo", 1) || mt.Passed() {
		t.FailNow()
	}
	if !NotEqual(mt, "foo", 1) || mt.Failed() {
		t.FailNow()
	}

	if !Equal(mt, []int{1, 2, 3}, []int{1, 2, 3}) || mt.Failed() {
		t.FailNow()
	}
	if NotEqual(mt, []int{1, 2, 3}, []int{1, 2, 3}) || mt.Passed() {
		t.FailNow()
	}

	if Equal(mt, []int{3, 2, 1}, []int{1, 2, 3}) || mt.Passed() {
		t.FailNow()
	}
	if !NotEqual(mt, []int{3, 2, 1}, []int{1, 2, 3}) || mt.Failed() {
		t.FailNow()
	}

	if !Equal(mt, map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "b": 2, "c": 3}) || mt.Failed() {
		t.FailNow()
	}
	if NotEqual(mt, map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "b": 2, "c": 3}) || mt.Passed() {
		t.FailNow()
	}

	var ptr *struct{}
	if !Equal(mt, nil, ptr) || mt.Failed() {
		t.FailNow()
	}
	if NotEqual(mt, nil, ptr) || mt.Passed() {
		t.FailNow()
	}

	ptr = &struct{}{}
	if Equal(mt, nil, ptr) || mt.Passed() {
		t.FailNow()
	}
	if !NotEqual(mt, nil, ptr) || mt.Failed() {
		t.FailNow()
	}

	var arr []int
	if !Equal(mt, nil, arr) || mt.Failed() {
		t.FailNow()
	}
	if NotEqual(mt, nil, arr) || mt.Passed() {
		t.FailNow()
	}

	arr = []int{1, 2, 3}
	if Equal(mt, nil, arr) || mt.Passed() {
		t.FailNow()
	}
	if !NotEqual(mt, nil, arr) || t.Failed() {
		t.FailNow()
	}

	var err error
	if !Equal(mt, nil, err) || mt.Failed() {
		t.FailNow()
	}
	if NotEqual(mt, nil, err) || mt.Passed() {
		t.FailNow()
	}

	err = errors.New("failed")
	if Equal(mt, nil, err) || mt.Passed() {
		t.FailNow()
	}
	if !NotEqual(mt, nil, err) || mt.Failed() {
		t.FailNow()
	}
}

func Test_Errors(t *testing.T) {
	mt := &MockTestingT{}

	var err error
	if !NoError(mt, err) || mt.Failed() {
		t.FailNow()
	}
	if Error(mt, err) || mt.Passed() {
		t.FailNow()
	}

	err = errors.New("This is bad")
	if NoError(mt, err) || mt.Passed() {
		t.FailNow()
	}
	if !Error(mt, err) || mt.Failed() {
		t.FailNow()
	}

	if !ErrorContains(mt, err, "bad") || mt.Failed() {
		t.FailNow()
	}
	if ErrorContains(mt, err, "really") || mt.Passed() {
		t.FailNow()
	}
}

func Test_Contains(t *testing.T) {
	mt := &MockTestingT{}

	// string
	if !Contains(mt, "hello world", "hello") || mt.Failed() {
		t.FailNow()
	}
	if Contains(mt, "hello world", "goodbye") || mt.Passed() {
		t.FailNow()
	}

	if NotContains(mt, "hello world", "hello") || mt.Passed() {
		t.FailNow()
	}
	if !NotContains(mt, "hello world", "goodbye") || mt.Failed() {
		t.FailNow()
	}

	// []byte
	if !Contains(mt, []byte("ABC"), []byte("AB")) || mt.Failed() {
		t.FailNow()
	}
	if Contains(mt, []byte("ABC"), []byte("X")) || mt.Passed() {
		t.FailNow()
	}
	if Contains(mt, []byte("ABC"), []byte("ABCD")) || mt.Passed() {
		t.FailNow()
	}
	if Contains(mt, []byte{}, []byte("X")) || mt.Passed() {
		t.FailNow()
	}

	// string and []byte
	if !Contains(mt, "hello world", []byte("hello")) || mt.Failed() {
		t.FailNow()
	}
	if !NotContains(mt, "hello world", []byte("foo")) || mt.Failed() {
		t.FailNow()
	}
	if !Contains(mt, []byte("hello world"), "hello") || mt.Failed() {
		t.FailNow()
	}
	if !NotContains(mt, []byte("hello world"), "foo") || mt.Failed() {
		t.FailNow()
	}

	// slice
	if NotContains(mt, []byte("ABC"), []byte("AB")) || mt.Passed() {
		t.FailNow()
	}
	if !NotContains(mt, []byte("ABC"), []byte("X")) || mt.Failed() {
		t.FailNow()
	}
	if !NotContains(mt, []byte("ABC"), []byte("ABCD")) || mt.Failed() {
		t.FailNow()
	}
	if !NotContains(mt, []byte{}, []byte("X")) || mt.Failed() {
		t.FailNow()
	}

	if !Contains(mt, []int{1, 2, 3}, 3) || mt.Failed() {
		t.FailNow()
	}
	if Contains(mt, []int{1, 2, 3}, 4) || mt.Passed() {
		t.FailNow()
	}

	if NotContains(mt, []int{1, 2, 3}, 3) || mt.Passed() {
		t.FailNow()
	}
	if !NotContains(mt, []int{1, 2, 3}, 4) || mt.Failed() {
		t.FailNow()
	}

	if Contains(mt, []int{1, 2, 3}, "3") || mt.Passed() {
		t.FailNow()
	}
	if Contains(mt, []int{1, 2, 3}, nil) || mt.Passed() {
		t.FailNow()
	}
	if !NotContains(mt, []int{1, 2, 3}, "3") || mt.Failed() {
		t.FailNow()
	}
	if !NotContains(mt, []int{1, 2, 3}, nil) || mt.Failed() {
		t.FailNow()
	}

	type E struct {
		X int
	}

	if !Contains(mt, []E{{1}, {2}, {3}}, E{3}) || mt.Failed() {
		t.FailNow()
	}
	if Contains(mt, []E{{1}, {2}, {3}}, E{4}) || mt.Passed() {
		t.FailNow()
	}

	if NotContains(mt, []E{{1}, {2}, {3}}, E{3}) || mt.Passed() {
		t.FailNow()
	}
	if !NotContains(mt, []E{{1}, {2}, {3}}, E{4}) || mt.Failed() {
		t.FailNow()
	}

	if Contains(mt, []E{{1}, {2}, {3}}, 1) || mt.Passed() {
		t.FailNow()
	}
	if Contains(mt, []E{{1}, {2}, {3}}, nil) || mt.Passed() {
		t.FailNow()
	}
	if !NotContains(mt, []E{{1}, {2}, {3}}, 1) || mt.Failed() {
		t.FailNow()
	}
	if !NotContains(mt, []E{{1}, {2}, {3}}, nil) || mt.Failed() {
		t.FailNow()
	}

	// map
	if !Contains(mt, map[string]string{"x": "X", "y": "Y"}, "x") || mt.Failed() {
		t.FailNow()
	}
	if Contains(mt, map[string]string{"x": "X", "y": "Y"}, "z") || mt.Passed() {
		t.FailNow()
	}

	if NotContains(mt, map[string]string{"x": "X", "y": "Y"}, "x") || mt.Passed() {
		t.FailNow()
	}
	if !NotContains(mt, map[string]string{"x": "X", "y": "Y"}, "z") || mt.Failed() {
		t.FailNow()
	}

	if Contains(mt, map[string]string{"x": "X", "y": "Y"}, 1) || mt.Passed() {
		t.FailNow()
	}
	if Contains(mt, map[string]string{"x": "X", "y": "Y"}, nil) || mt.Passed() {
		t.FailNow()
	}
	if !NotContains(mt, map[string]string{"x": "X", "y": "Y"}, 1) || mt.Failed() {
		t.FailNow()
	}
	if !NotContains(mt, map[string]string{"x": "X", "y": "Y"}, nil) || mt.Failed() {
		t.FailNow()
	}

	// error
	err := errors.New("This is bad")
	if !Contains(mt, err, "bad") || mt.Failed() {
		t.FailNow()
	}
	if Contains(mt, err, "really") || mt.Passed() {
		t.FailNow()
	}
	if NotContains(mt, err, "bad") || mt.Passed() {
		t.FailNow()
	}
	if !NotContains(mt, err, "really") || mt.Failed() {
		t.FailNow()
	}

	err = nil
	if Contains(mt, err, "bad") || mt.Passed() {
		t.FailNow()
	}

	// nil container
	if Contains(mt, nil, 1) || mt.Passed() {
		t.FailNow()
	}
	if !NotContains(mt, nil, 1) || mt.Failed() {
		t.FailNow()
	}

	// Unsupported type
	if Contains(mt, 1, 1) || mt.Passed() {
		t.FailNow()
	}
	if !NotContains(mt, 1, 1) || mt.Failed() {
		t.FailNow()
	}
}

func Test_Len(t *testing.T) {
	mt := &MockTestingT{}

	if !Len(mt, nil, 0) || mt.Failed() {
		t.FailNow()
	}

	if !Len(mt, "foo", 3) || mt.Failed() {
		t.FailNow()
	}

	if !Len(mt, []int{1, 2, 3}, 3) || mt.Failed() {
		t.FailNow()
	}
	if !Len(mt, []int{}, 0) || mt.Failed() {
		t.FailNow()
	}

	if !Len(mt, map[int]int{1: 10, 2: 20, 3: 30}, 3) || mt.Failed() {
		t.FailNow()
	}
	if !Len(mt, map[int]int{}, 0) || mt.Failed() {
		t.FailNow()
	}

	ch := make(chan bool, 3)
	ch <- true
	ch <- true
	ch <- true
	if !Len(mt, ch, 3) || mt.Failed() {
		t.FailNow()
	}

	if Len(mt, 1, 1) || mt.Passed() {
		t.FailNow()
	}
	if Len(mt, false, 1) || mt.Passed() {
		t.FailNow()
	}
}

func Test_Nil(t *testing.T) {
	mt := &MockTestingT{}

	i := 123
	for _, x := range []any{1, true, &i, struct{}{}, "foo"} {
		if !NotNil(mt, x) || mt.Failed() {
			t.FailNow()
		}
		if Nil(mt, x) || mt.Passed() {
			t.FailNow()
		}
	}
	if !Nil(mt, nil) || mt.Failed() {
		t.FailNow()
	}
	if NotNil(mt, nil) || mt.Passed() {
		t.FailNow()
	}
}

func Test_TrueFalse(t *testing.T) {
	mt := &MockTestingT{}

	if !True(mt, true) || mt.Failed() {
		t.FailNow()
	}
	if True(mt, false) || mt.Passed() {
		t.FailNow()
	}

	if !False(mt, false) || mt.Failed() {
		t.FailNow()
	}
	if False(mt, true) || mt.Passed() {
		t.FailNow()
	}
}

func Test_Zero(t *testing.T) {
	mt := &MockTestingT{}

	for _, z := range []any{0, "", false, time.Time{}, struct{}{}, nil} {
		if !Zero(mt, z) || mt.Failed() {
			t.FailNow()
		}
		if NotZero(mt, z) || mt.Passed() {
			t.FailNow()
		}
	}
	i := 123
	for _, nz := range []any{1, "x", true, time.Now(), struct{ x int }{5}, &i} {
		if !NotZero(mt, nz) || mt.Failed() {
			t.FailNow()
		}
		if Zero(mt, nz) || mt.Passed() {
			t.FailNow()
		}
	}
}

func Test_SliceLen(t *testing.T) {
	mt := &MockTestingT{}

	fullSlice := []int{1, 2, 3}
	if !SliceLen(mt, fullSlice, 3) || mt.Failed() {
		t.FailNow()
	}
	if SliceLen(mt, fullSlice, 2) || mt.Passed() {
		t.FailNow()
	}

	emptySlice := []int{}
	if !SliceLen(mt, emptySlice, 0) || mt.Failed() {
		t.FailNow()
	}
	if SliceLen(mt, emptySlice, 2) || mt.Passed() {
		t.FailNow()
	}

	var nilSlice []int
	if !SliceLen(mt, nilSlice, 0) || mt.Failed() {
		t.FailNow()
	}
	if SliceLen(mt, nilSlice, 2) || mt.Passed() {
		t.FailNow()
	}
}

func Test_MapLen(t *testing.T) {
	mt := &MockTestingT{}

	fullMap := map[int]int{1: 1, 2: 2, 3: 3}
	if !MapLen(mt, fullMap, 3) || mt.Failed() {
		t.FailNow()
	}
	if MapLen(mt, fullMap, 2) || mt.Passed() {
		t.FailNow()
	}

	emptyMap := map[int]int{}
	if !MapLen(mt, emptyMap, 0) || mt.Failed() {
		t.FailNow()
	}
	if MapLen(mt, emptyMap, 2) || mt.Passed() {
		t.FailNow()
	}

	var nilMap map[int]int
	if !MapLen(mt, nilMap, 0) || mt.Failed() {
		t.FailNow()
	}
	if MapLen(mt, nilMap, 2) || mt.Passed() {
		t.FailNow()
	}
}

func Test_StrLen(t *testing.T) {
	mt := &MockTestingT{}

	if !StrLen(mt, "abc", 3) || mt.Failed() {
		t.FailNow()
	}
	if StrLen(mt, "abc", 2) || mt.Passed() {
		t.FailNow()
	}

	if !StrLen(mt, "", 0) || mt.Failed() {
		t.FailNow()
	}
	if StrLen(mt, "", 2) || mt.Passed() {
		t.FailNow()
	}
}

func Test_SliceContains(t *testing.T) {
	mt := &MockTestingT{}

	s := []int{1, 2, 3}
	if !SliceContains(mt, s, 1) || mt.Failed() {
		t.FailNow()
	}
	if SliceContains(mt, s, 0) || mt.Passed() {
		t.FailNow()
	}
	if SliceNotContains(mt, s, 1) || mt.Passed() {
		t.FailNow()
	}
	if !SliceNotContains(mt, s, 0) || mt.Failed() {
		t.FailNow()
	}

	s = []int{}
	if SliceContains(mt, s, 0) || mt.Passed() {
		t.FailNow()
	}
	if !SliceNotContains(mt, s, 0) || mt.Failed() {
		t.FailNow()
	}

	s = nil
	if SliceContains(mt, s, 0) || mt.Passed() {
		t.FailNow()
	}
	if !SliceNotContains(mt, s, 0) || mt.Failed() {
		t.FailNow()
	}
}

func Test_SliceQual(t *testing.T) {
	mt := &MockTestingT{}

	if !SliceEqual(mt, []int{1, 2, 3}, []int{1, 2, 3}) || mt.Failed() {
		t.FailNow()
	}
	if SliceEqual(mt, []int{1, 2}, []int{1, 2, 3}) || mt.Passed() {
		t.FailNow()
	}
	if SliceEqual(mt, []int{}, []int{1, 2, 3}) || mt.Passed() {
		t.FailNow()
	}
	if SliceEqual(mt, nil, []int{1, 2, 3}) || mt.Passed() {
		t.FailNow()
	}
	if SliceEqual(mt, []int{1, 2, 3}, nil) || mt.Passed() {
		t.FailNow()
	}

	if SliceNotEqual(mt, []int{1, 2, 3}, []int{1, 2, 3}) || mt.Passed() {
		t.FailNow()
	}
	if !SliceNotEqual(mt, []int{1, 2}, []int{1, 2, 3}) || mt.Failed() {
		t.FailNow()
	}
	if !SliceNotEqual(mt, []int{}, []int{1, 2, 3}) || mt.Failed() {
		t.FailNow()
	}
	if !SliceNotEqual(mt, nil, []int{1, 2, 3}) || mt.Failed() {
		t.FailNow()
	}
	if !SliceNotEqual(mt, []int{1, 2, 3}, nil) || mt.Failed() {
		t.FailNow()
	}
}

type stringer struct {
	x string
}

func (s *stringer) String() string {
	return "Stringer" + s.x
}

func Test_Stringer(t *testing.T) {
	mt := &MockTestingT{}
	if !NotNil(mt, &stringer{}) || mt.Failed() {
		t.FailNow()
	}
	if Nil(mt, &stringer{}) || mt.Passed() {
		t.FailNow()
	}
	if !Nil(mt, (*stringer)(nil)) || mt.Failed() {
		t.FailNow()
	}

	if !Equal(mt, "Stringer", v(&stringer{})) || mt.Failed() {
		t.FailNow()
	}
	if Equal(mt, "Stringer", v((*stringer)(nil))) || mt.Passed() {
		t.FailNow()
	}
	if !Equal(mt, "<nil>", v((*stringer)(nil))) || mt.Failed() {
		t.FailNow()
	}
}

type textMarshaler struct {
	x string
}

func (tm *textMarshaler) MarshalText() (text []byte, err error) {
	return []byte("TextMarshaler" + tm.x), nil
}

func Test_TextMarshaler(t *testing.T) {
	mt := &MockTestingT{}
	if !NotNil(mt, &textMarshaler{}) || mt.Failed() {
		t.FailNow()
	}
	if Nil(mt, &textMarshaler{}) || mt.Passed() {
		t.FailNow()
	}
	if !Nil(mt, (*textMarshaler)(nil)) || mt.Failed() {
		t.FailNow()
	}

	if !Equal(mt, "TextMarshaler", v(&textMarshaler{})) || mt.Failed() {
		t.FailNow()
	}
	if Equal(mt, "TextMarshaler", v((*textMarshaler)(nil))) || mt.Passed() {
		t.FailNow()
	}
	if !Equal(mt, "<nil>", v((*textMarshaler)(nil))) || mt.Failed() {
		t.FailNow()
	}
}

func Test_Expect(t *testing.T) {
	mt := &MockTestingT{}

	var err error
	x := 1
	s := "hello"
	if !Expect(mt, err, nil, x, 1, s, "hello") || mt.Failed() {
		t.FailNow()
	}
	if Expect(mt, err, nil, x, 2, s, "hello") || mt.Passed() {
		t.FailNow()
	}
	if Expect(mt, err, nil, x, 1, s, "world") || mt.Passed() {
		t.FailNow()
	}

	err = errors.New("failed")
	x = 0
	s = ""
	if Expect(mt, err, nil, x, 1, s, "hello") || mt.Passed() {
		t.FailNow()
	}
}
