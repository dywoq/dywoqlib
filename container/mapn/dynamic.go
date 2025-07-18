package mapn

import "github.com/dywoq/dywoqlib/container/slice"

type Dynamic[K, V comparable] struct {
	err error
	m   map[K]V
}

func NewDynamic[K, V comparable](m map[K]V) *Dynamic[K, V] {
	return &Dynamic[K, V]{nil, m}
}

func (d *Dynamic[K, V]) Length() int {
	return len(d.m)
}

func (d *Dynamic[K, V]) Error() error {
	return d.err
}

func (d *Dynamic[K, V]) Grow(i int) {
	if d.err != nil {
		return
	}
	d.m = make(map[K]V, i)
}

func (d *Dynamic[K, V]) Exists(reqkey K) (exists bool) {
	if d.err != nil {
		return false
	}
	_, exists = d.m[reqkey]
	return
}

func (d *Dynamic[K, V]) Add(reqkey K, reqvalue V) (k K, v V) {
	if d.err != nil {
		return
	}
	if !d.Exists(reqkey) {
		d.m[reqkey] = reqvalue
		k = reqkey
		v = reqvalue
		return
	}
	d.err = ErrKeyAlreadyExist
	return
}

func (d *Dynamic[K, V]) Set(reqkey K, reqvalue V) (k K, v V) {
	if d.err != nil {
		return
	}
	if d.Exists(reqkey) {
		d.m[reqkey] = reqvalue
		k = reqkey
		v = reqvalue
		return
	}
	d.err = ErrKeyNotFound
	return
}

func (d *Dynamic[K, V]) Keys() []K {
	if d.err != nil {
		return []K{}
	}
	keys := slice.NewFixed[K](len(d.m))
	// we don't use keys.Grow() here because we don't need to preallocate the slice -
	// slice.NewFixed already does that
	// or because there's no method Grow() in slice.Fixed xD
	for key := range d.m {
		keys.Append(key)
	}
	if keys.Error() != nil {
		d.err = keys.Error()
		return []K{}
	}
	return keys.Native()
}

func (d *Dynamic[K, V]) Values() []V {
	if d.err != nil {
		return []V{}
	}
	values := slice.NewFixed[V](len(d.m))
	// there's same situation as described above
	for _, value := range d.m {
		values.Append(value)
	}
	if values.Error() != nil {
		d.err = values.Error()
		return []V{}
	}
	return values.Native()
}

func (d *Dynamic[K, V]) Delete(reqkey K) (k K) {
	if d.err != nil {
		return
	}
	if d.Exists(reqkey) {
		delete(d.m, reqkey)
		k = reqkey
		return
	}
	d.err = ErrKeyNotFound
	return
}

func (d *Dynamic[K, V]) Get(reqkey K) (k K, v V) {
	if d.err != nil {
		return
	}
	if d.Exists(reqkey) {
		k = reqkey
		v = d.m[reqkey]
		return
	}
	d.err = ErrKeyNotFound
	return
}
