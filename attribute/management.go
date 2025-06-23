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
	"runtime"
)

type management struct{}

func (management) funcInfo(skip int) string {
	ret, _, _, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}
	f := runtime.FuncForPC(ret)
	if f == nil {
		return "unknown"
	}
	name := f.Name()
	return name
}

func (management) skipNums() (int, int) {
	// 2 - skip num of caller of an attributes
	// 3 - skip num of source of the warning
	return 2, 3
}
