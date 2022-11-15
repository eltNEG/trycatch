package trycatch

type Result[T any, E any] struct {
	Err *E
	Val T
}

func (r *Result[T, E]) HasError() bool {
	return r.Err != nil
}

type ErrorHandler[V any, R any, E any] func(r *Result[R, E], v V, err error)

func Try[T any, R any, E any](r *Result[R, E], f ErrorHandler[T, R, E]) func(value T, err error) T {
	return func(value T, err error) T {
		if err == nil {
			return value
		}
		f(r, value, err)
		panic("")
	}
}

func Catch[U any, E any](r *Result[U, E]) {
	recover()
	return
}

func Init[T any, E any](r *Result[T, E]) *Result[T, E] {
	r = &Result[T, E]{}
	return r
}
