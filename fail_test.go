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
	"errors"
	"testing"
)

func TestMe(t *testing.T) {
	Equal(t, 1, 0, "You are not the one")
	err := errors.New("This is bad")
	NoError(t, err)

	droids := 1234
	FailIf(t, droids != 0, "These are not the droids you are looking for")
	err = errors.New("This is really bad")
	FatalIfError(t, err)
}
