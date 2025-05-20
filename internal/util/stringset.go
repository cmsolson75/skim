package util

type StringSet map[string]struct{}

func NewSet(items []string) StringSet {
	s := make(StringSet)
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

func (s StringSet) Has(item string) bool {
	_, ok := s[item]
	return ok
}
