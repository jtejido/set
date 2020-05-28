package set

var (
	s unsafeSet
	_ Set = s
)

type unsafeSet map[interface{}]struct{}

func newUnsafeSet() unsafeSet {
	return make(map[interface{}]struct{})
}

func (s unsafeSet) Add(i ...interface{}) {
	for _, item := range i {
		if _, found := s[item]; found {
			continue
		}

		s[item] = struct{}{}
	}
}

func (s unsafeSet) Contains(i ...interface{}) bool {
	for _, item := range i {
		if _, found := s[item]; !found {
			return false
		}
	}
	return true
}

func (s unsafeSet) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s unsafeSet) Remove(i interface{}) {
	if _, found := s[i]; !found {
		return
	}

	delete(s, i)
}

func (s unsafeSet) Len() int {
	return len(s)
}

func (s unsafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for key := range s {
			ch <- key
		}

		close(ch)
	}()

	return ch
}

func (s unsafeSet) Equal(other Set) bool {
	_ = other.(unsafeSet)

	if s.Len() != other.Len() {
		return false
	}
	for elem := range s {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (s unsafeSet) Clone() Set {
	clonedSet := newUnsafeSet()
	for key := range s {
		if key != nil {
			clonedSet.Add(key)
		}
	}
	return clonedSet
}

func (s unsafeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, s.Len())
	for key := range s {
		keys = append(keys, key)
	}

	return keys
}

func (s unsafeSet) RemoveFrom(other Set) {
	for elem := range s {
		if other.Contains(elem) {
			s.Remove(elem)
		}
	}
}

func (s unsafeSet) AddFrom(other Set) {
	o := other.(unsafeSet)

	for elem := range o {
		if !s.Contains(elem) {
			s.Add(elem)
		}
	}
}

func (s unsafeSet) RetainFrom(other Set) {
	for elem := range s {
		if !other.Contains(elem) {
			s.Remove(elem)
		}
	}
}
