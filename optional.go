package optional

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

func (o Option[T]) GetOr(fallback T) T {
	if o.value != nil {
		return *o.value
	} else {
		return fallback
	}
}

func (o Option[T]) ToPtr() *T {
	return o.value
}
