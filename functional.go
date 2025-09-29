package optional

func Map[T, V any](o Option[T], f func(T) V) Option[V] {
	v, err := o.Get()
	if err != nil {
		return None[V]()
	}
	return Some(f(v))
}
