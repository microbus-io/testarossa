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
	"fmt"
	"runtime"
	"strings"
)

// TestingT is an interface that both *testing.T and *testing.B implement.
type TestingT interface {
	Fail()
	FailNow()
	Name() string
}

// FailIf fails the test if the condition is met.
// If returns back the result of evaluating the condition.
func FailIf(t TestingT, condition bool, args ...any) bool {
	if !condition {
		return false
	}
	filePath, lineNum := atSourceFileLine()
	var sb strings.Builder
	if len(args) > 1 {
		if format, ok := args[0].(string); ok && strings.Contains(format, "%") {
			v := fmt.Sprintf(format, args[1:]...)
			v = strings.ReplaceAll(v, "\n", "\n    ")
			sb.WriteString("    ")
			sb.WriteString(v)
			sb.WriteString("\n")
		}
	}
	if sb.Len() == 0 {
		for a := range args {
			v := fmt.Sprintf("%+v", args[a])
			if v == "" {
				continue
			}
			v = strings.ReplaceAll(v, "\n", "\n    ")
			sb.WriteString("    ")
			sb.WriteString(v)
			sb.WriteString("\n")
		}
	}
	if lineNum == 0 {
		fmt.Printf("--- FAIL: %s\n%s", t.Name(), sb.String())
	} else {
		fmt.Printf("--- FAIL: %s\n    %s:%d\n%s", t.Name(), filePath, lineNum, sb.String())
	}
	t.Fail()
	return true
}

// FailIfError is a shortcut to FailIf(t, err != nil, append([]any{err}, args...)...) .
func FailIfError(t TestingT, err error, args ...any) bool {
	return FailIf(t, err != nil, append([]any{err}, args...)...)
}

// FatalIf fails the test and stops further execution if the condition is met.
func FatalIf(t TestingT, condition bool, args ...any) {
	if condition {
		FailIf(t, condition, args...)
		t.FailNow()
	}
}

// FatalIfError is a shortcut to FatalIf(t, err != nil, append([]any{err}, args...)...) .
func FatalIfError(t TestingT, err error, args ...any) {
	FatalIf(t, err != nil, append([]any{err}, args...)...)
}

func atSourceFileLine() (filePath string, lineNum int) {
	for lvl := 2; true; lvl++ {
		pc, file, line, ok := runtime.Caller(lvl)
		if !ok {
			break
		}
		funcName := runtime.FuncForPC(pc).Name()
		p := strings.LastIndex(funcName, "/")
		if p > 0 {
			funcName = funcName[p+1:]
		}
		if strings.HasPrefix(funcName, "testarossa.") {
			continue
		}
		if lineNum == 0 {
			filePath = file
			lineNum = line
		}

		p = strings.Index(funcName, ".")
		if p > 0 {
			funcName = funcName[p+1:]
		}
		if strings.HasPrefix(funcName, "Test") || strings.HasPrefix(funcName, "Benchmark") {
			return file, line
		}
	}
	return filePath, lineNum
}
