package set

var (
	s *unsafeSet
	_ Set = s
)

type unsafeSet struct {
	index map[interface{}]struct{}
}

func newUnsafeSet() *unsafeSet {
	return &unsafeSet{index: make(map[interface{}]struct{})}
}

func (s *unsafeSet) Add(i ...Equality) {
	for _, item := range i {
		var caught bool
		for k := range s.index {
			if item.Equals(k) {
				caught = true
				break
			}
		}

		if caught {
			continue
		}

		s.index[item] = struct{}{}
	}
}

func (s *unsafeSet) Contains(i ...Equality) bool {
	for _, val := range i {
		if val != nil {
			if v, ok := val.(Equality); ok {
				for k := range s.index {
					if v.Equals(k) {
						return true
					}
				}
			}
		}
	}
	return false
}

func (s *unsafeSet) Clear() {
	*s = *newUnsafeSet()
}

func (s *unsafeSet) Remove(i Equality) {
	for k := range s.index {
		if i.Equals(k) {
			delete(s.index, k)
			break
		}
	}
}

func (s *unsafeSet) Len() int {
	return len(s.index)
}

func (s *unsafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for key := range s.index {
			ch <- key
		}

		close(ch)
	}()

	return ch
}

func (s *unsafeSet) Equal(other Set) bool {
	_ = other.(*unsafeSet)

	if s.Len() != other.Len() {
		return false
	}
	for elem := range s.index {
		if elem != nil {
			if v, ok := elem.(Equality); ok {
				if !other.Contains(v) {
					return false
				}
			}
		}
	}
	return true
}

func (s *unsafeSet) Clone() Set {
	clonedSet := newUnsafeSet()
	for key := range s.index {
		if key != nil {
			clonedSet.Add(key.(Equality))
		}
	}
	return clonedSet
}

func (s *unsafeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, s.Len())
	for key := range s.index {
		keys = append(keys, key)
	}

	return keys
}

func (s *unsafeSet) RemoveFrom(other Set) {
	for elem := range s.index {
		if elem != nil {
			if v, ok := elem.(Equality); ok {
				if other.Contains(v) {
					s.Remove(v)
				}
			}
		}
	}
}

func (s *unsafeSet) AddFrom(other Set) {
	o := other.(*unsafeSet)

	for elem := range o.index {
		if elem != nil {
			if v, ok := elem.(Equality); ok {
				if !s.Contains(v) {
					s.Add(v)
				}
			}
		}
	}
}

func (s *unsafeSet) RetainFrom(other Set) {
	for elem := range s.index {
		if elem != nil {
			if v, ok := elem.(Equality); ok {
				if !other.Contains(v) {
					s.Remove(v)
				}
			}
		}
	}
}
