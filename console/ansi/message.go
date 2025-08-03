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
