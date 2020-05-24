package set

var (
	us *unsafeOrderedSet
	_  Set = us
)

type unsafeOrderedSet struct {
	currentIndex int
	m            map[int]interface{}
	index        map[interface{}]int
	keys         []int
}

func newUnsafeOrderedSet() *unsafeOrderedSet {
	return &unsafeOrderedSet{currentIndex: 0, m: make(map[int]interface{}), keys: make([]int, 0), index: make(map[interface{}]int)}
}

func (s *unsafeOrderedSet) Add(i ...Equality) {
	for _, item := range i {
		var caught bool
		for k, _ := range s.index {
			if item.Equals(k) {
				caught = true
				break
			}
		}

		if caught {
			continue
		}

		if _, ok := s.m[s.currentIndex]; !ok {
			s.keys = append(s.keys, s.currentIndex)
		}

		s.m[s.currentIndex] = item
		s.index[item] = s.currentIndex
		s.currentIndex++
	}
}

func (s *unsafeOrderedSet) Contains(i ...Equality) bool {
	for _, val := range i {
		for k, _ := range s.index {
			if val != nil {
				if v, ok := val.(Equality); ok {
					if v.Equals(k) {
						return true
					}
				}
			}
		}
	}
	return false
}

func (s *unsafeOrderedSet) Clear() {
	*s = *newUnsafeOrderedSet()
}

func (s *unsafeOrderedSet) Remove(i Equality) {
	for k, index := range s.index {
		if i.Equals(k) {
			if _, found := s.m[index]; !found {
				return
			}

			delete(s.m, index)
			for i := range s.keys {
				if s.keys[i] == index {
					s.keys = append(s.keys[:i], s.keys[i+1:]...)
					break
				}
			}

			delete(s.index, k)
			break
		}
	}
}

func (s *unsafeOrderedSet) Len() int {
	return len(s.m)
}

func (s *unsafeOrderedSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for _, key := range s.keys {
			ch <- s.m[key]
		}

		close(ch)
	}()

	return ch
}

func (s *unsafeOrderedSet) Equal(other Set) bool {
	_ = other.(*unsafeOrderedSet)

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

func (s *unsafeOrderedSet) Clone() Set {
	clonedSet := newUnsafeOrderedSet()
	for _, key := range s.keys {
		if s.m[key] != nil {
			clonedSet.Add(s.m[key].(Equality))
		}
	}
	return clonedSet
}

func (s *unsafeOrderedSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, s.Len())
	for _, key := range s.keys {
		keys = append(keys, s.m[key])
	}

	return keys
}

func (s *unsafeOrderedSet) RemoveFrom(other Set) {
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

func (s *unsafeOrderedSet) AddFrom(other Set) {
	o := other.(*unsafeOrderedSet)

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

func (s *unsafeOrderedSet) RetainFrom(other Set) {
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
