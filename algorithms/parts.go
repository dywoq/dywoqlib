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

package algorithms

// Parts break k into n parts, including n as the last point.
func Parts(n, k int) []float64 {
	result := make([]float64, k)
	step := float64(n) / float64(k+1)
	for i := 1; i <= k; i++ {
		result[i-1] = step * float64(i)
	}
	return result
}
