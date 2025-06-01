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

func Test_Equality(t *testing.T) {
	mt := &MockTestingT{}

	Equal(mt, 1, 0, "You are not the %d", 1)
	mt.Failed(t)

	Equal(mt, 1, 1, "All things equal")
	mt.Passed(t)

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

	Contains(mt, "hello world", "hello")
	mt.Passed(t)
	Contains(mt, "hello world", "goodbye")
	mt.Failed(t)

	NotContains(mt, "hello world", "hello")
	mt.Failed(t)
	NotContains(mt, "hello world", "goodbye")
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
	Contains(mt, []int{1, 2, 3}, "4")
	mt.Failed(t)

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

	Contains(mt, map[string]string{"x": "X", "y": "Y"}, "x")
	mt.Passed(t)
	Contains(mt, map[string]string{"x": "X", "y": "Y"}, "z")
	mt.Failed(t)

	NotContains(mt, map[string]string{"x": "X", "y": "Y"}, "x")
	mt.Failed(t)
	NotContains(mt, map[string]string{"x": "X", "y": "Y"}, "z")
	mt.Passed(t)

	NotContains(mt, map[string]string{"x": "X", "y": "Y"}, 1)
	mt.Passed(t)
}
