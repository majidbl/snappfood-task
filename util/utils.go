package util

func In[T comparable](t T, arr []T) bool {
	for _, a := range arr {
		if a == t {
			return a == t
		}
	}

	return false
}
