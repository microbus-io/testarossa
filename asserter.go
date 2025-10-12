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

type Asserter struct {
	t TestingT
}

func For(t TestingT) *Asserter {
	return &Asserter{t: t}
}

// Error fails the test if err is nil.
func (tt *Asserter) Error(err error, args ...any) bool {
	return Error(tt.t, err, args...)
}

// NoError fails the test if err is not nil.
func (tt *Asserter) NoError(err error, args ...any) bool {
	return NoError(tt.t, err, args...)
}

// Equal fails the test if the two values are not equal.
// Note: the expected value comes before the actual value in the argument list.
func (tt *Asserter) Equal(expected any, actual any, args ...any) bool {
	return Equal(tt.t, expected, actual, args...)
}

// NotEqual fails the test if the two values are equal.
// Note: the expected value comes before the actual value in the argument list.
func (tt *Asserter) NotEqual(expected any, actual any, args ...any) bool {
	return NotEqual(tt.t, expected, actual, args...)
}

// Zero fails the test if the value is not 0.
func (tt *Asserter) Zero(actual any, args ...any) bool {
	return Zero(tt.t, actual, args...)
}

// NotZero fails the test if the value is the 0 value of its type.
func (tt *Asserter) NotZero(actual any, args ...any) bool {
	return NotZero(tt.t, actual, args...)
}

// True fails the test if the condition is false.
func (tt *Asserter) True(condition bool, args ...any) bool {
	return True(tt.t, condition, args...)
}

// False fails the test if the condition is true.
func (tt *Asserter) False(condition bool, args ...any) bool {
	return False(tt.t, condition, args...)
}

// Contains fails the test if the string does not contain a substring.
func (tt *Asserter) Contains(whole any, sub any, args ...any) bool {
	return Contains(tt.t, whole, sub, args...)
}

// NotContains fails the test if the string contain a substring.
func (tt *Asserter) NotContains(whole any, sub any, args ...any) bool {
	return NotContains(tt.t, whole, sub, args...)
}

// Len fails the test if the length of the string, slice, array, map or chan does not match the expected len.
func (tt *Asserter) Len(obj any, length int, args ...any) bool {
	return Len(tt.t, obj, length, args...)
}

// Nil fails the test if the object is not nil.
func (tt *Asserter) Nil(obj any, args ...any) bool {
	return Nil(tt.t, obj, args...)
}

// NotNil fails the test if the object is nil.
func (tt *Asserter) NotNil(obj any, args ...any) bool {
	return NotNil(tt.t, obj, args...)
}

/*
Expect fails the test if any of the paired values are not equal.
Note: the expected value comes after the actual value in each pair.

	result, err := doSomething(p1, p2)
	tt.Expect(err, nil, result, 4321)
*/
func (tt *Asserter) Expect(actualExpectedPairs ...any) bool {
	return Expect(tt.t, actualExpectedPairs...)
}

/*
HTMLMatch fails the test if no HTML element matching the CSS selector query was found
to also match the regular expression by the inner text of any of its descendants.

Examples:

	tt.HTMLMatch(html, `TR > TD > A.expandable[href]`, "")
	tt.HTMLMatch(html, `DIV#main_panel`, `^Help$``)
	tt.HTMLMatch(html, `TR TD INPUT[name="x"]`, `[0-9]+``)
*/
func (tt *Asserter) HTMLMatch(htmlBody []byte, cssSelectorQuery string, innerTextRegExp string, args ...any) bool {
	return HTMLMatch(tt.t, htmlBody, cssSelectorQuery, innerTextRegExp, args...)
}

/*
HTMLNotMatch fails the test if at least one HTML element matching the CSS selector query was found
to also match the regular expression by the inner text of any of its descendants.

Examples:

	HTMLNotMatch(t, html, `TR > TD > A.expandable[href]`, "")
	HTMLNotMatch(t, html, `DIV#main_panel`, "^Help$")
	HTMLNotMatch(t, html, `TR TD INPUT[name="x"]`, `[0-9]+``)
*/
func (tt *Asserter) HTMLNotMatch(htmlBody []byte, cssSelectorQuery string, innerTextRegExp string, args ...any) bool {
	return HTMLNotMatch(tt.t, htmlBody, cssSelectorQuery, innerTextRegExp, args...)
}
