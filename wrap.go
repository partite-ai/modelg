package modelg

import "reflect"

type Unwrappable interface {
	Unwrap(target any) bool
}

func Unwrap[T any](src any) (T, bool) {
	var target T
	if uw, ok := src.(Unwrappable); ok {
		if uw.Unwrap(&target) {
			return target, true
		}
	}

	return target, false
}

func unwrapOrDelegrate(src any, target any) bool {
	tv := reflect.ValueOf(target)
	if tv.Kind() != reflect.Ptr || tv.IsNil() {
		return false
	}
	tv = tv.Elem()

	if reflect.TypeOf(src).AssignableTo(tv.Type()) {
		tv.Set(reflect.ValueOf(src))
		return true
	}
	if uw, ok := src.(Unwrappable); ok {
		return uw.Unwrap(target)
	}
	return false
}
