package option

// Option represent optional value
type Option[T any] interface {
	// true if Option[T] is None
	IsEmpty() bool

	// if Option[T] is None return parameter default
	// else return the value
	GetOrElse(d T) T

	// if Option[T] is None return parameter other
	// else return self
	OrElse(other Option[T]) Option[T]

	// return the value
	// if Option[T] is None, then panic
	Get() T
}

// option is a implementation of Option
type option[T any] struct {
	ref *T
}

// None return Option[T] which without value
func None[T any]() Option[T] {
	return &option[T]{}
}

// Some return Option[T] which with value
func Some[T any](value T) Option[T] {
	return &option[T]{
		ref: &value,
	}
}

// true if Option[T] is None
func (o *option[T]) IsEmpty() bool {
	return o.ref == nil
}

// if Option[T] is None return parameter default
// else return the value
func (o *option[T]) GetOrElse(d T) T {
	if o.IsEmpty() {
		return d
	}
	return *o.ref
}

// if Option[T] is None return parameter other
// else return self
func (o *option[T]) OrElse(other Option[T]) Option[T] {
	if o.IsEmpty() {
		return other
	}
	return o
}

// return the value
// if Option[T] is None, then panic
func (o *option[T]) Get() T {
	return *o.ref
}
