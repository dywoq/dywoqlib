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

var event func()

// SetEvent sets custom event to attributes-functions.
// If you want to reset the event, you can use SetEvent(nil).
//
// Example without the custom event:
// 	func check() {
//		attribute.Deprecated() // or any attribute-function (attribute.Todo)
// 	}
//
// 	func main() {
//		check()
//	}
//
// Example with custom event:
// 	func check() {
//		attribute.SetEvent(func() { fmt.Println("custom event") })
//		attribute.Deprecated()
// 	}
//
// 	func main() {
//		check()
//	}
//
func SetEvent(value func()) {
	event = value
}
