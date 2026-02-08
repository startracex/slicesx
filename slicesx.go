package slicesx

import "reflect"

// Map applies fn to each element of the slice and returns a new slice of results.
func Map[S ~[]E, E any, R any](raw S, fn func(E) R) []R {
	return MapIndex(raw, func(v E, _ int) R { return fn(v) })
}

// MapIndex applies fn to each element of the slice and returns a new slice of results.
func MapIndex[S ~[]E, E any, R any](raw S, fn func(E, int) R) []R {
	result := make([]R, 0, len(raw))
	for i, v := range raw {
		result = append(result, fn(v, i))
	}
	return result
}

// Filter returns a new slice containing only elements for which fn returns true.
func Filter[S ~[]E, E any](raw S, fn func(E) bool) S {
	return FilterIndex(raw, func(v E, _ int) bool { return fn(v) })
}

// FilterIndex returns a new slice containing only elements for which fn returns true.
func FilterIndex[S ~[]E, E any](raw S, fn func(E, int) bool) S {
	result := make(S, 0, len(raw))
	for i, v := range raw {
		if fn(v, i) {
			result = append(result, v)
		}
	}
	return result
}

// ForEach executes fn for each element in the slice.
func ForEach[S ~[]E, E any](raw S, fn func(E)) {
	ForEach2(raw, func(v E, _ int) { fn(v) })
}

// ForEach2 executes fn for each element in the slice.
func ForEach2[S ~[]E, E any](raw S, fn func(E, int)) {
	for i, v := range raw {
		fn(v, i)
	}
}

// Some returns true if any element satisfies fn.
func Some[S ~[]E, E any](raw S, fn func(E) bool) bool {
	return SomeIndex(raw, func(v E, _ int) bool { return fn(v) })
}

// SomeIndex returns true if any element satisfies fn.
func SomeIndex[S ~[]E, E any](raw S, fn func(E, int) bool) bool {
	for i, v := range raw {
		if fn(v, i) {
			return true
		}
	}
	return false
}

// Every returns true if all elements satisfy fn.
func Every[S ~[]E, E any](raw S, fn func(E) bool) bool {
	return EveryIndex(raw, func(v E, _ int) bool { return fn(v) })
}

// EveryIndex returns true if all elements satisfy fn.
func EveryIndex[S ~[]E, E any](raw S, fn func(E, int) bool) bool {
	for i, v := range raw {
		if !fn(v, i) {
			return false
		}
	}
	return true
}

// Reduce left-folds the slice.
func Reduce[S ~[]E, E any, R any](raw S, fn func(R, E) R, initial R) R {
	for _, v := range raw {
		initial = fn(initial, v)
	}
	return initial
}

// ReduceIndex left-folds the slice.
func ReduceIndex[S ~[]E, E any, R any](raw S, fn func(R, E, int) R, initial R) R {
	for i, v := range raw {
		initial = fn(initial, v, i)
	}
	return initial
}

// ReduceRight right-folds the slice.
func ReduceRight[S ~[]E, E any, R any](raw S, fn func(R, E) R, initial R) R {
	for i := len(raw) - 1; i >= 0; i-- {
		initial = fn(initial, raw[i])
	}
	return initial
}

// ReduceRightIndex right-folds the slice.
func ReduceRightIndex[S ~[]E, E any, R any](raw S, fn func(R, E, int) R, initial R) R {
	for i := len(raw) - 1; i >= 0; i-- {
		initial = fn(initial, raw[i], i)
	}
	return initial
}

// Unshift prepends values to the slice.
func Unshift[S ~[]E, E any](raw S, value ...E) S {
	result := make(S, 0, len(value)+len(raw))
	result = append(result, value...)
	result = append(result, raw...)
	return result
}

// Push appends values to the slice.
func Push[S ~[]E, E any](raw S, value ...E) S {
	return append(raw, value...)
}

// Shift removes the first element.
func Shift[S ~[]E, E any](raw S) S {
	if len(raw) == 0 {
		return S{}
	}
	result := make(S, 0, len(raw)-1)
	return append(result, raw[1:]...)
}

// Pop removes the last element.
func Pop[S ~[]E, E any](raw S) S {
	if len(raw) == 0 {
		return S{}
	}
	result := make(S, 0, len(raw)-1)
	return append(result, raw[:len(raw)-1]...)
}

// Reverse returns a reversed copy of the slice.
func Reverse[S ~[]E, E any](raw S) S {
	n := len(raw)
	result := make(S, n)
	for i := 0; i < n; i++ {
		result[i] = raw[n-1-i]
	}
	return result
}

// Flat flattens a slice-of-slices by one level.
func Flat[S ~[]T, T ~[]E, E any](raw S) T {
	total := 0
	for _, inner := range raw {
		total += len(inner)
	}
	result := make(T, 0, total)

	for _, inner := range raw {
		result = append(result, inner...)
	}
	return result
}

// FlatMap maps each element to a slice and flattens one level.
func FlatMap[S ~[]E, E any, T any](raw S, fn func(E) []T) []T {
	return FlatMapIndex(raw, func(v E, _ int) []T { return fn(v) })
}

// FlatMapIndex maps each element to a slice and flattens one level.
func FlatMapIndex[S ~[]E, E any, T any](raw S, fn func(E, int) []T) []T {
	result := make([]T, 0, len(raw))
	for i, v := range raw {
		result = append(result, fn(v, i)...)
	}
	return result
}

// Splice removes deleteCount elements at start and inserts values.
func Splice[S ~[]E, E any](raw S, start int, deleteCount int, values ...E) S {
	if start < 0 {
		start = 0
	}
	if start > len(raw) {
		start = len(raw)
	}

	end := start + deleteCount
	if end > len(raw) {
		end = len(raw)
	}

	result := make(S, 0, len(raw)-deleteCount+len(values))
	result = append(result, raw[:start]...)
	result = append(result, values...)
	result = append(result, raw[end:]...)
	return result
}

// IsSlice checks if value is a slice.
func IsSlice(value any) bool {
	if value == nil {
		return false
	}
	return reflect.TypeOf(value).Kind() == reflect.Slice
}
