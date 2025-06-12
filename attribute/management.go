// Copyright 2025 dywoq
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package attribute

import (
	"fmt"
	"log"
	"runtime"
)

type management struct{}

func (m management) output(message string, mode Mode) {
	switch mode {
	case SoftMode:
		fmt.Println(message)
	case StrictMode:
		log.SetFlags(0)
		log.Fatalln(message)
	}
}

func (m management) functionName(skip int) string {
	zero := "unknown" // the variable is to be returned if error happens
	ret, _, _, ok := runtime.Caller(skip)
	if !ok {
		return zero
	}
	fn := runtime.FuncForPC(ret)
	if fn == nil {
		return zero
	}
	return fn.Name()
}

func (m management) targetNumberSkip() int {
	return 2
}

func (m management) sourceNumberSkip() int {
	return 3
}
