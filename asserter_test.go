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

func TestAsserter_All(t *testing.T) {
	mt := &MockTestingT{}
	tt := For(mt)
	if !tt.Error(errors.New("bad")) || mt.Failed() {
		t.FailNow()
	}
	if !tt.NoError(nil) || mt.Failed() {
		t.FailNow()
	}
	if !tt.Equal(1, 1) || mt.Failed() {
		t.FailNow()
	}
	if !tt.NotEqual(1, 0) || mt.Failed() {
		t.FailNow()
	}
	if !tt.Zero(0) || mt.Failed() {
		t.FailNow()
	}
	if !tt.NotZero(1) || mt.Failed() {
		t.FailNow()
	}
	if !tt.True(true) || mt.Failed() {
		t.FailNow()
	}
	if !tt.False(false) || mt.Failed() {
		t.FailNow()
	}
	if !tt.Contains([]int{1, 2, 3}, 1) || mt.Failed() {
		t.FailNow()
	}
	if !tt.NotContains([]int{1, 2, 3}, 0) || mt.Failed() {
		t.FailNow()
	}
	if !tt.Len([]int{1, 2, 3}, 3) || mt.Failed() {
		t.FailNow()
	}
	if !tt.Nil(nil) || mt.Failed() {
		t.FailNow()
	}
	if !tt.NotNil(1) || mt.Failed() {
		t.FailNow()
	}
	if !tt.Expect(1, 1, "x", "x") || mt.Failed() {
		t.FailNow()
	}
}

func TestAsserter_HTML(t *testing.T) {
	mt := &MockTestingT{}
	tt := For(mt)
	htmlBody := []byte(`<html><body><div class="banner" id="id123">Cool <b>Banner</b>!</div></body></html>`)

	if !tt.HTMLMatch(htmlBody, "B", "") || mt.Failed() {
		t.FailNow()
	}
	if !tt.HTMLMatch(htmlBody, "DIV.banner", "") || mt.Failed() {
		t.FailNow()
	}
	if !tt.HTMLMatch(htmlBody, "DIV#id123", "") || mt.Failed() {
		t.FailNow()
	}
	if !tt.HTMLMatch(htmlBody, "B", "Banner") || mt.Failed() {
		t.FailNow()
	}

	if !tt.HTMLMatch(htmlBody, "DIV", "Banner") || mt.Failed() {
		t.FailNow()
	}
	if tt.HTMLNotMatch(htmlBody, "DIV", "Banner") || mt.Passed() {
		t.FailNow()
	}

	if !tt.HTMLNotMatch(htmlBody, "DIV", "Title") || mt.Failed() {
		t.FailNow()
	}
	if tt.HTMLMatch(htmlBody, "DIV", "Title") || mt.Passed() {
		t.FailNow()
	}

	// Errors
	if tt.HTMLMatch(htmlBody, "DIV.#id123", "Banner") || mt.Passed() {
		t.FailNow()
	}
	if tt.HTMLMatch(htmlBody, "DIV", "[Banner") || mt.Passed() {
		t.FailNow()
	}
	if tt.HTMLMatch([]byte(`<html><</html>`), "DIV", "") || mt.Passed() {
		t.FailNow()
	}
}
