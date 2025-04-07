package util

func When[T any](condition bool, tv, fv T) T {
	if condition {
		return tv
	}
	return fv
}
