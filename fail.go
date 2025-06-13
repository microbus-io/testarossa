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
	var sb strings.Builder
	i := 0
	for i < len(args) {
		val := ""
		if str, ok := args[i].(string); ok {
			pctCount := strings.Count(str, "%") - 2*strings.Count(str, "%%")
			pctCount = min(pctCount, len(args)-i-1)
			val = fmt.Sprintf(str, args[i+1:i+1+pctCount]...)
			i += pctCount + 1
		} else {
			val = fmt.Sprintf("%+v", args[i])
			i++
		}
		if val == "" {
			continue
		}
		val = strings.ReplaceAll(val, "\n", "\n    ")
		sb.WriteString("    ")
		sb.WriteString(val)
		sb.WriteString("\n")
	}
	if len(args) == 0 {
		sb.WriteString("\n")
	}
	fmt.Printf("--- FAIL: %s\n%s%s", t.Name(), stackTrace(), sb.String())
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

func stackTrace() (stackTrace string) {
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
		if strings.HasPrefix(funcName, "testarossa.") && !strings.HasPrefix(funcName, "testarossa.Test") {
			continue
		}
		if funcName == "testing.tRunner" {
			break
		}
		stackTrace = fmt.Sprintf("    %s:%d\n%s", file, line, stackTrace)
		if strings.Contains(funcName, ".Test") || strings.Contains(funcName, ".Benchmark") {
			break
		}
	}
	return stackTrace
}
