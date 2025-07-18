package mapn

type Fixed[K, V comparable] struct {
	err      error
	fixedLen int
	m        *Dynamic[K, V]
}

func NewFixed[K, V comparable](fixedLen int, m map[K]V) *Fixed[K, V] {
	d := NewDynamic(map[K]V{})
	if d.Error() != nil {
		return &Fixed[K, V]{d.Error(), fixedLen, nil}
	}
	if fixedLen < 0 {
		return &Fixed[K, V]{ErrNegativeFixedLength, fixedLen, nil}
	}
	if fixedLen < len(m) {
		return &Fixed[K, V]{ErrFixedLengthOutOfBounds, fixedLen, nil}
	}
	if len(m) > fixedLen {
		return &Fixed[K, V]{ErrOutOfBounds, fixedLen, nil}
	}
	d.Grow(fixedLen)
	for key, value := range m {
		d.Add(key, value)
	}
	return &Fixed[K, V]{nil, fixedLen, d}
}

func (f *Fixed[K, V]) Length() int {
	return f.m.Length()
}

func (f *Fixed[K, V]) Error() error {
	return f.err
}

func (f *Fixed[K, V]) Exists(reqkey K) bool {
	if ok := f.errorsOk(); !ok {
		return false
	}
	return f.m.Exists(reqkey)
}

func (f *Fixed[K, V]) Add(reqkey K, reqvalue V) (k K, v V) {
	if ok := f.errorsOk(); !ok {
		return
	}
	res1, res2 := f.m.Add(reqkey, reqvalue)
	if f.m.Error() != nil {
		f.err = f.m.Error()
	}
	k = res1
	v = res2
	return
}

func (f *Fixed[K, V]) Set(reqkey K, reqvalue V) (k K, v V) {
	if ok := f.errorsOk(); !ok {
		return
	}
	res1, res2 := f.m.Set(reqkey, reqvalue)
	if f.m.Error() != nil {
		f.err = f.m.Error()
	}
	k = res1
	v = res2
	return
}

func (f *Fixed[K, V]) Keys() []K {
	if ok := f.errorsOk(); !ok {
		return []K{}
	}
	keys := f.m.Keys()
	if f.m.Error() != nil {
		f.err = f.m.Error()
		return []K{}
	}
	return keys
}

func (f *Fixed[K, V]) Values() []V {
	if ok := f.errorsOk(); !ok {
		return []V{}
	}
	values := f.m.Values()
	if f.m.Error() != nil {
		f.err = f.m.Error()
		return []V{}
	}
	return values
}

func (f *Fixed[K, V]) Delete(reqkey K) (k K) {
	if ok := f.errorsOk(); !ok {
		return
	}
	res1 := f.m.Delete(reqkey)
	if f.m.Error() != nil {
		f.err = f.m.Error()
		return
	}
	k = res1
	return
}

func (f *Fixed[K, V]) Get(reqkey K) (k K, v V) {
	if ok := f.errorsOk(); !ok {
		return
	}
	res1, res2 := f.m.Get(reqkey)
	if f.m.Error() != nil {
		f.err = f.m.Error()
		return
	}
	k = res1
	v = res2
	return
}

func (f *Fixed[K, V]) String() string {
	if ok := f.errorsOk(); !ok {
		return ""
	}
	res := f.m.String()
	if f.m.Error() != nil {
		f.err = f.m.Error()
		return ""
	}
	return res
}

func (f *Fixed[K, V]) outOfBounds() bool {
	return f.m.Length() > f.fixedLen
}

func (f *Fixed[K, V]) errorsOk() bool {
	if f.err != nil {
		return false
	}
	if f.m.Error() != nil {
		f.err = f.m.Error()
		return false
	}
	if f.outOfBounds() {
		f.err = ErrOutOfBounds
		return false
	}
	return true
}
