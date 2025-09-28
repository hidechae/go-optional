package optional

import "errors"

var ErrGetFromNone = errors.New("get from none")

type Option[T any] struct {
	value []T
}

const idx = iota

func Some[T any](v T) Option[T] {
	return Option[T]{
		value: []T{v},
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		value: nil,
	}
}

func FromPtr[T any](v *T) Option[T] {
	if v == nil {
		return None[T]()
	}
	return Some(*v)
}

func (o Option[T]) IsSone() bool {
	return o.value != nil
}

func (o Option[T]) IsNone() bool {
	return o.value == nil
}

func (o Option[T]) Get() (T, error) {
	if o.IsNone() {
		var defaultValue T
		return defaultValue, ErrGetFromNone
	}
	return o.value[idx], nil
}

func (o Option[T]) GetOr(fallback T) T {
	if o.IsNone() {
		return fallback
	}
	return o.value[idx]
}

func (o Option[T]) ToPtr() *T {
	v, err := o.Get()
	if err != nil {
		return nil
	}
	return &v
}
