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

	Equal(mt, 1, 0, "You are not the %d", 1)
	mt.Failed(t)

	Equal(mt, 1, 1, "All things equal")
	mt.Passed(t)

	Equal(mt, "foo", 1)
	mt.Failed(t)
	NotEqual(mt, "foo", 1)
	mt.Passed(t)

	Equal(mt, []int{1, 2, 3}, []int{1, 2, 3})
	mt.Passed(t)
	NotEqual(mt, []int{1, 2, 3}, []int{1, 2, 3})
	mt.Failed(t)

	Equal(mt, []int{3, 2, 1}, []int{1, 2, 3})
	mt.Failed(t)
	NotEqual(mt, []int{3, 2, 1}, []int{1, 2, 3})
	mt.Passed(t)

	Equal(mt, map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "b": 2, "c": 3})
	mt.Passed(t)
	NotEqual(mt, map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "b": 2, "c": 3})
	mt.Failed(t)
}

func Test_Errors(t *testing.T) {
	mt := &MockTestingT{}

	var err error
	NoError(mt, err)
	mt.Passed(t)
	Error(mt, err)
	mt.Failed(t)

	err = errors.New("This is bad")
	NoError(mt, err)
	mt.Failed(t)
	Error(mt, err)
	mt.Passed(t)

	ErrorContains(mt, err, "bad")
	mt.Passed(t)
	ErrorContains(mt, err, "really")
	mt.Failed(t)
}

func Test_Contains(t *testing.T) {
	mt := &MockTestingT{}

	// string
	Contains(mt, "hello world", "hello")
	mt.Passed(t)
	Contains(mt, "hello world", "goodbye")
	mt.Failed(t)

	NotContains(mt, "hello world", "hello")
	mt.Failed(t)
	NotContains(mt, "hello world", "goodbye")
	mt.Passed(t)

	// []byte
	Contains(mt, []byte("ABC"), []byte("AB"))
	mt.Passed(t)
	Contains(mt, []byte("ABC"), []byte("X"))
	mt.Failed(t)
	Contains(mt, []byte("ABC"), []byte("ABCD"))
	mt.Failed(t)
	Contains(mt, []byte{}, []byte("X"))
	mt.Failed(t)

	// slice
	NotContains(mt, []byte("ABC"), []byte("AB"))
	mt.Failed(t)
	NotContains(mt, []byte("ABC"), []byte("X"))
	mt.Passed(t)
	NotContains(mt, []byte("ABC"), []byte("ABCD"))
	mt.Passed(t)
	NotContains(mt, []byte{}, []byte("X"))
	mt.Passed(t)

	Contains(mt, []int{1, 2, 3}, 3)
	mt.Passed(t)
	Contains(mt, []int{1, 2, 3}, 4)
	mt.Failed(t)

	NotContains(mt, []int{1, 2, 3}, 3)
	mt.Failed(t)
	NotContains(mt, []int{1, 2, 3}, 4)
	mt.Passed(t)

	Contains(mt, []int{1, 2, 3}, "3")
	mt.Failed(t)
	Contains(mt, []int{1, 2, 3}, nil)
	mt.Failed(t)
	NotContains(mt, []int{1, 2, 3}, "3")
	mt.Passed(t)
	NotContains(mt, []int{1, 2, 3}, nil)
	mt.Passed(t)

	type E struct {
		X int
	}

	Contains(mt, []E{{1}, {2}, {3}}, E{3})
	mt.Passed(t)
	Contains(mt, []E{{1}, {2}, {3}}, E{4})
	mt.Failed(t)

	NotContains(mt, []E{{1}, {2}, {3}}, E{3})
	mt.Failed(t)
	NotContains(mt, []E{{1}, {2}, {3}}, E{4})
	mt.Passed(t)

	Contains(mt, []E{{1}, {2}, {3}}, 1)
	mt.Failed(t)
	Contains(mt, []E{{1}, {2}, {3}}, nil)
	mt.Failed(t)
	NotContains(mt, []E{{1}, {2}, {3}}, 1)
	mt.Passed(t)
	NotContains(mt, []E{{1}, {2}, {3}}, nil)
	mt.Passed(t)

	// map
	Contains(mt, map[string]string{"x": "X", "y": "Y"}, "x")
	mt.Passed(t)
	Contains(mt, map[string]string{"x": "X", "y": "Y"}, "z")
	mt.Failed(t)

	NotContains(mt, map[string]string{"x": "X", "y": "Y"}, "x")
	mt.Failed(t)
	NotContains(mt, map[string]string{"x": "X", "y": "Y"}, "z")
	mt.Passed(t)

	Contains(mt, map[string]string{"x": "X", "y": "Y"}, 1)
	mt.Failed(t)
	Contains(mt, map[string]string{"x": "X", "y": "Y"}, nil)
	mt.Failed(t)
	NotContains(mt, map[string]string{"x": "X", "y": "Y"}, 1)
	mt.Passed(t)
	NotContains(mt, map[string]string{"x": "X", "y": "Y"}, nil)
	mt.Passed(t)

	// error
	err := errors.New("This is bad")
	Contains(mt, err, "bad")
	mt.Passed(t)
	Contains(mt, err, "really")
	mt.Failed(t)
	NotContains(mt, err, "bad")
	mt.Failed(t)
	NotContains(mt, err, "really")
	mt.Passed(t)

	err = nil
	Contains(mt, err, "bad")
	mt.Failed(t)

	// nil
	Contains(mt, nil, 1)
	mt.Failed(t)
	NotContains(mt, nil, 1)
	mt.Passed(t)

	// Unsupported type
	Contains(mt, 1, 1)
	mt.Failed(t)
	NotContains(mt, 1, 1)
	mt.Failed(t)
}

func Test_Len(t *testing.T) {
	mt := &MockTestingT{}

	Len(mt, nil, 0)
	mt.Passed(t)

	Len(mt, "foo", 3)
	mt.Passed(t)

	Len(mt, []int{1, 2, 3}, 3)
	mt.Passed(t)
	Len(mt, []int{}, 0)
	mt.Passed(t)

	Len(mt, map[int]int{1: 10, 2: 20, 3: 30}, 3)
	mt.Passed(t)
	Len(mt, map[int]int{}, 0)
	mt.Passed(t)

	ch := make(chan bool, 3)
	ch <- true
	ch <- true
	ch <- true
	Len(mt, ch, 3)
	mt.Passed(t)

	Len(mt, 1, 1)
	mt.Failed(t)
	Len(mt, false, 1)
	mt.Failed(t)
}

func Test_Nil(t *testing.T) {
	mt := &MockTestingT{}

	i := 123
	for _, x := range []any{1, true, &i, struct{}{}, "foo"} {
		NotNil(mt, x)
		mt.Passed(t)
		Nil(mt, x)
		mt.Failed(t)
	}
	Nil(mt, nil)
	mt.Passed(t)
	NotNil(mt, nil)
	mt.Failed(t)
}

func Test_TrueFalse(t *testing.T) {
	mt := &MockTestingT{}

	True(mt, true)
	mt.Passed(t)
	True(mt, false)
	mt.Failed(t)

	False(mt, false)
	mt.Passed(t)
	False(mt, true)
	mt.Failed(t)
}

func Test_Zero(t *testing.T) {
	mt := &MockTestingT{}

	for _, z := range []any{0, "", false, time.Time{}, struct{}{}, nil} {
		Zero(mt, z)
		mt.Passed(t)
		NotZero(mt, z)
		mt.Failed(t)
	}
	i := 123
	for _, nz := range []any{1, "x", true, time.Now(), struct{ x int }{5}, &i} {
		NotZero(mt, nz)
		mt.Passed(t)
		Zero(mt, nz)
		mt.Failed(t)
	}
}

func Test_SliceLen(t *testing.T) {
	mt := &MockTestingT{}

	fullSlice := []int{1, 2, 3}
	SliceLen(mt, fullSlice, 3)
	mt.Passed(t)
	SliceLen(mt, fullSlice, 2)
	mt.Failed(t)

	emptySlice := []int{}
	SliceLen(mt, emptySlice, 0)
	mt.Passed(t)
	SliceLen(mt, emptySlice, 2)
	mt.Failed(t)

	var nilSlice []int
	SliceLen(mt, nilSlice, 0)
	mt.Passed(t)
	SliceLen(mt, nilSlice, 2)
	mt.Failed(t)
}

func Test_MapLen(t *testing.T) {
	mt := &MockTestingT{}

	fullMap := map[int]int{1: 1, 2: 2, 3: 3}
	MapLen(mt, fullMap, 3)
	mt.Passed(t)
	MapLen(mt, fullMap, 2)
	mt.Failed(t)

	emptyMap := map[int]int{}
	MapLen(mt, emptyMap, 0)
	mt.Passed(t)
	MapLen(mt, emptyMap, 2)
	mt.Failed(t)

	var nilMap map[int]int
	MapLen(mt, nilMap, 0)
	mt.Passed(t)
	MapLen(mt, nilMap, 2)
	mt.Failed(t)
}

func Test_StrLen(t *testing.T) {
	mt := &MockTestingT{}

	StrLen(mt, "abc", 3)
	mt.Passed(t)
	StrLen(mt, "abc", 2)
	mt.Failed(t)

	StrLen(mt, "", 0)
	mt.Passed(t)
	StrLen(mt, "", 2)
	mt.Failed(t)
}

func Test_SliceContains(t *testing.T) {
	mt := &MockTestingT{}

	s := []int{1, 2, 3}
	SliceContains(mt, s, 1)
	mt.Passed(t)
	SliceContains(mt, s, 0)
	mt.Failed(t)
	SliceNotContains(mt, s, 1)
	mt.Failed(t)
	SliceNotContains(mt, s, 0)
	mt.Passed(t)

	s = []int{}
	SliceContains(mt, s, 0)
	mt.Failed(t)
	SliceNotContains(mt, s, 0)
	mt.Passed(t)

	s = nil
	SliceContains(mt, s, 0)
	mt.Failed(t)
	SliceNotContains(mt, s, 0)
	mt.Passed(t)
}

func Test_SliceQual(t *testing.T) {
	mt := &MockTestingT{}

	SliceEqual(mt, []int{1, 2, 3}, []int{1, 2, 3})
	mt.Passed(t)
	SliceEqual(mt, []int{1, 2}, []int{1, 2, 3})
	mt.Failed(t)
	SliceEqual(mt, []int{}, []int{1, 2, 3})
	mt.Failed(t)
	SliceEqual(mt, nil, []int{1, 2, 3})
	mt.Failed(t)
	SliceEqual(mt, []int{1, 2, 3}, nil)
	mt.Failed(t)

	SliceNotEqual(mt, []int{1, 2, 3}, []int{1, 2, 3})
	mt.Failed(t)
	SliceNotEqual(mt, []int{1, 2}, []int{1, 2, 3})
	mt.Passed(t)
	SliceNotEqual(mt, []int{}, []int{1, 2, 3})
	mt.Passed(t)
	SliceNotEqual(mt, nil, []int{1, 2, 3})
	mt.Passed(t)
	SliceNotEqual(mt, []int{1, 2, 3}, nil)
	mt.Passed(t)
}

type stringer struct {
	x string
}

func (s *stringer) String() string {
	return "Stringer" + s.x
}

func Test_Stringer(t *testing.T) {
	mt := &MockTestingT{}
	NotNil(mt, &stringer{})
	mt.Passed(t)
	Nil(mt, &stringer{})
	mt.Failed(t)
	Nil(mt, (*stringer)(nil))
	mt.Passed(t)

	Equal(mt, "Stringer", v(&stringer{}))
	mt.Passed(t)
	Equal(mt, "Stringer", v((*stringer)(nil)))
	mt.Failed(t)
	Equal(mt, "<nil>", v((*stringer)(nil)))
	mt.Passed(t)
}

type textMarshaler struct {
	x string
}

func (tm *textMarshaler) MarshalText() (text []byte, err error) {
	return []byte("TextMarshaler" + tm.x), nil
}

func Test_TextMarshaler(t *testing.T) {
	mt := &MockTestingT{}
	NotNil(mt, &textMarshaler{})
	mt.Passed(t)
	Nil(mt, &textMarshaler{})
	mt.Failed(t)
	Nil(mt, (*textMarshaler)(nil))
	mt.Passed(t)

	Equal(mt, "TextMarshaler", v(&textMarshaler{}))
	mt.Passed(t)
	Equal(mt, "TextMarshaler", v((*textMarshaler)(nil)))
	mt.Failed(t)
	Equal(mt, "<nil>", v((*textMarshaler)(nil)))
	mt.Passed(t)
}
