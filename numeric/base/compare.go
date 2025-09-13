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

package base

// Compare is the result of a comparison between two values.
type Compare int

const (
	// CompareLess indicates that the first value is less than the second.
	CompareLess Compare = -1
	// CompareEqual indicates that the two values are equal.
	CompareEqual Compare = 0
	// CompareGreater indicates that the first value is greater than the second.
	CompareGreater Compare = 1
)
