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

package ansi

type message struct {
	textColor       Color
	backgroundColor Color
	value           string
}

func (m *message) String() string {
	return m.value
}

func (m *message) BgColor() Color {
	return m.backgroundColor
}

func (m *message) FgColor() Color {
	return m.textColor
}

func (m *message) SetBgColor(c Color) Base {
	m.backgroundColor = c
	m.value = ApplyBoth(m.value, m.textColor, c)
	return m
}

func (m *message) SetFgColor(c Color) Base {
	m.textColor = c
	m.value = ApplyBoth(m.value, c, m.backgroundColor)
	return m
}
