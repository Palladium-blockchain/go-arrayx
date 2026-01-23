package arrayx

func Map[TIn any, TOut any](ts []TIn, f func(TIn) TOut) []TOut {
	res := make([]TOut, len(ts))
	for i, t := range ts {
		res[i] = f(t)
	}
	return res
}

func FindIf[T any](ts []T, f func(T) bool) *T {
	for _, t := range ts {
		if f(t) {
			return &t
		}
	}
	return nil
}

func Filter[T any](ts []T, f func(T) bool) []T {
	res := make([]T, 0)
	for _, t := range ts {
		if f(t) {
			res = append(res, t)
		}
	}
	return res
}

func Flatten[TIn, TOut any](ts []TIn, f func(TIn) []TOut) []TOut {
	res := make([]TOut, 0)
	for _, t := range ts {
		res = append(res, f(t)...)
	}
	return res
}

func In[T comparable](ts []T, item T) bool {
	for _, t := range ts {
		if item == t {
			return true
		}
	}
	return false
}

func ContainsAll[T comparable](have, required []T) bool {
	set := make(map[T]struct{}, len(have))

	for _, v := range have {
		set[v] = struct{}{}
	}

	for _, r := range required {
		if _, ok := set[r]; !ok {
			return false
		}
	}

	return true
}
