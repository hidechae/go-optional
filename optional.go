package optional

import "errors"

var ErrGetValueFromNone = errors.New("get value from none")

type Option[T any] struct {
	value *T
}

func Some[T any](v T) Option[T] {
	return Option[T]{
		value: &v,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		value: nil,
	}
}

func FromPtr[T any](v *T) Option[T] {
	return Option[T]{
		value: v,
	}
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
		return defaultValue, ErrGetValueFromNone
	} else {
		return *o.value, nil
	}
}

func (o Option[T]) GetOr(fallback T) T {
	if o.IsNone() {
		return fallback
	} else {
		return *o.value
	}
}

func (o Option[T]) ToPtr() *T {
	return o.value
}
