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

type MockTestingT struct {
	failed bool
}

func (mt *MockTestingT) Fail() {
	mt.failed = true
}
func (mt *MockTestingT) FailNow() {
	mt.failed = true
}
func (mt *MockTestingT) Name() string {
	return "Mock"
}
func (mt *MockTestingT) Failed() (failed bool) {
	failed = mt.failed
	mt.failed = false
	return failed
}
func (mt *MockTestingT) Passed() (passed bool) {
	passed = !mt.failed
	mt.failed = false
	return passed
}

func Test_FailIf(t *testing.T) {
	mt := &MockTestingT{}

	droids := 1234
	if !FailIf(mt, droids != 0, "These are not the droids you are looking for") || mt.Passed() {
		t.FailNow()
	}

	err := errors.New("This is really bad")
	if !FatalIfError(mt, err) || mt.Passed() {
		t.FailNow()
	}
}

func Test_FailIfError(t *testing.T) {
	mt := &MockTestingT{}

	if FailIfError(mt, nil, "No error here") || mt.Failed() {
		t.FailNow()
	}

	t.Run("subtest", func(t *testing.T) {
		err := errors.New("error here")
		if !FailIfError(mt, err, "Not good") || mt.Passed() {
			t.FailNow()
		}
	})
}
