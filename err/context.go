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

package err

import (
	"encoding/json"
)

// Context defines an interface for error handling with additional context.
// It allows access to the error itself, extra details, and different
// representations such as string or JSON.
type Context interface {
	// Error returns the original error.
	Error() error
	// More returns additional context about the error.
	More() string
	// Nil reports whether the error is nil.
	Nil() bool
	// String returns a formatted string with the error and additional context.
	String() string
	// Marshal returns a JSON representation of the error and context.
	Marshal() ([]byte, error)
	// SetMore sets a "more" context value.
	SetMore(string)
	// SetError sets a error to the underlying error field.
	SetError(error)
}

type implementation struct {
	err  error
	more string
}

// NewContext creates a new Context with the given error and additional context.
func NewContext(err error, more string) Context {
	return &implementation{err, more}
}

func (i *implementation) Error() error {
	return i.err
}

func (i *implementation) More() string {
	return i.more
}

func (i *implementation) Nil() bool {
	return i.err == nil
}

func (i *implementation) String() string {
	return i.err.Error() + ": " + i.more
}

type jsonPayload struct {
	Error string `json:"error"`
	More  string `json:"more"`
}

func (i *implementation) Marshal() ([]byte, error) {
	data := jsonPayload{
		Error: i.err.Error(),
		More:  i.more,
	}
	return json.Marshal(data)
}

func (i *implementation) SetMore(value string) {
	i.more = value
}

func (i *implementation) SetError(value error) {
	i.err = value
}
