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
func (mt *MockTestingT) Failed(t *testing.T) {
	if !mt.failed {
		t.Fatal()
	}
	mt.failed = false
}
func (mt *MockTestingT) Passed(t *testing.T) {
	if mt.failed {
		t.Fatal()
	}
	mt.failed = false
}

func Test_FailIf(t *testing.T) {
	mt := &MockTestingT{}

	droids := 1234
	FailIf(mt, droids != 0, "These are not the droids you are looking for")
	mt.Failed(t)

	err := errors.New("This is really bad")
	FatalIfError(mt, err)
	mt.Failed(t)
}

func Test_FailIfError(t *testing.T) {
	mt := &MockTestingT{}

	FailIfError(mt, nil, "No error here")
	mt.Passed(t)

	t.Run("subtest", func(t *testing.T) {
		err := errors.New("error here")
		FailIfError(mt, err, "Not good")
		mt.Failed(t)
	})
}
