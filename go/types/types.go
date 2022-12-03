package types

// StringSet provides a general purpose string set.
type StringSet map[string]struct{}

// NewStringSet returns a new empty StringSet ready for use.
func NewStringSet() StringSet {
	return make(StringSet)
}

func NewStringSetFromList(ss []string) StringSet {
	result := make(StringSet)
	for _, s := range ss {
		result.Add(s)
	}

	return result
}

// Add adds the given val to the set.
func (s StringSet) Add(val string) {
	if s == nil {
		return
	}

	s[val] = struct{}{}
}

// Contains checks if val exists in the set.
func (s StringSet) Contains(val string) bool {
	if s == nil {
		return false
	}

	_, ok := s[val]

	return ok
}

// Remove deletes val from the set.
func (s StringSet) Remove(val string) {
	if s == nil {
		return
	}

	delete(s, val)
}

// Len returns the number of items in the set.
func (s StringSet) Len() int {
	return len(s)
}

// Items returns a list of all items in the set.
//
// The result items are unique.
func (s StringSet) Items() []string {
	result := make([]string, 0, len(s))
	for k := range s {
		result = append(result, k)
	}

	return result
}

// IsEmpty returns true of s is empty.
func (s StringSet) IsEmpty() bool {
	return s.Len() == 0
}
