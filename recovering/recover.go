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

package recovering

import (
	"log"
)

// Recover catches the encountered panic and returns it. Unlike builtin recover(),
// the function outputs the log message about panic.
// The message of caught panic would look like this:
// recovering.Recover(): caught the panic: <caught panic() value>
func Recover() (r any) {
	if r = recover(); r != nil {
		log.Printf("recovering.Recover(): caught the panic: %v", r)
	}
	return
}
