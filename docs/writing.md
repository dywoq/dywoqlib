## Writing

### Writing benchmarks
1. First, you need to include package "github.com/dywoq/dywoqlib/internal/testing" as package "internal_testing.go":
```go
import internal_testing "github.com/dywoq/dywoqlib/internal/testing"
```
We need this package to set base parameters for benchmarks.
If Go standard package "testing" collides, name it as go_testing:
```go
import (
	go_testing "testing"
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
)
```
2. In your benchmarks use:
```go
func BenchmarkFoo(b *go_testing.B) {
	internal_testing.SetBase(b)
	// something...
}
```
1. Use `b.Loop()` instead of `for` statement:
```go
func BenchmarkFoo(b *go_testing.B) {
	internal_testing.SetBase(b)
	for b.Loop() {
		// something...
	}
}
```

### Panic
1. When you call standard Go `panic()`, you need to paste link of package you work in before the message.
2. Example:
```go
panic("github.com/dywoq/dywoqlib/foo/something: division by zero")
```

### Style
1. The indentation width must be 4.
2. Use `err.Context` if possible - this is useful for creating errors with more context. Only one restriction - **Don't use `err.Context` where standard Go `error` is already used to avoid breaking compatibility!**
3. Follow the standard Go naming - for example: `foo.Something`.
4. In the start of each Go file, you must include the copyright notice:
```go
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
```

### Naming branches
1. The naming format branches is:
```
<container>.<sub-packages if exist>.<the very short form of the issue, words should be split into .> - the issue id
```

2. The example:
```
container.unique.optimization - #32
```
