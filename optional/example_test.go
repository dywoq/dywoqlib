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

package optional_test

import (
	"errors"
	"fmt"

	"github.com/dywoq/dywoqlib/err"
	"github.com/dywoq/dywoqlib/optional"
)

func ExampleMaybe() {
	defer func() {
		if recovered := recover(); recovered != nil {
			fmt.Printf("recovered: %v\n", recovered)
		}
	}()

	opt := optional.Int(20)

	// present
	present := opt.Present()
	fmt.Printf("present: %v\n", present)

	// getting
	val, ok := opt.Get()
	if !ok {
		fmt.Printf("ok: %v\n", ok)
	}
	fmt.Printf("val: %v\n", val)

	// else
	opt = optional.Int()
	elseOpt := opt.Else(30)
	fmt.Printf("elseOpt: %v\n", elseOpt)

	// filtering
	opt = optional.Int(20)
	filtered := opt.Filter(func(i int) bool {
		return i < 30
	})
	filteredVal, ok := filtered.Get()
	if !ok {
		fmt.Printf("ok: %v\n", ok)
	}
	fmt.Printf("filteredVal: %v\n", filteredVal)

	// or
	opt = optional.Int()
	or := opt.Or(func() int {
		return 10 + 10
	})
	fmt.Printf("or: %v\n", or)

	// error context
	opt = optional.NoneContext[int](err.NewContext(
		errors.New("something went wrong"),
		"optional value is not present on purpose, this is its fate",
	))
	err := opt.Error()
	fmt.Printf("err: %v\n", err.String())

	// unwrapping
	opt = optional.Int()
	unwrapped := opt.Unwrap()
	fmt.Printf("unwrapped: %v\n", unwrapped)

	// Output:
	// present: true
	// val: 20
	// elseOpt: 30
	// filteredVal: 20
	// or: 20
	// err: something went wrong: optional value is not present on purpose, this is its fate
	// recovered: github.com/dywoq/dywoqlib/optional: not present
}
