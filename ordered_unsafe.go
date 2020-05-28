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

func (s *unsafeOrderedSet) Add(i ...interface{}) {
	for _, item := range i {
		if _, found := s.index[item]; found {
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

func (s *unsafeOrderedSet) Contains(i ...interface{}) bool {
	for _, item := range i {
		if _, found := s.index[item]; !found {
			return false
		}
	}
	return true
}

func (s *unsafeOrderedSet) Clear() {
	*s = *newUnsafeOrderedSet()
}

func (s *unsafeOrderedSet) Remove(i interface{}) {
	index, found := s.index[i]
	if !found {
		return
	}

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

	delete(s.index, i)
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
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (s *unsafeOrderedSet) Clone() Set {
	clonedSet := newUnsafeOrderedSet()
	for _, key := range s.keys {
		if s.m[key] != nil {
			clonedSet.Add(s.m[key])
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
		if other.Contains(elem) {
			s.Remove(elem)
		}
	}
}

func (s *unsafeOrderedSet) AddFrom(other Set) {
	o := other.(*unsafeOrderedSet)

	for elem := range o.index {
		if !s.Contains(elem) {
			s.Add(elem)
		}
	}
}

func (s *unsafeOrderedSet) RetainFrom(other Set) {
	for elem := range s.index {
		if !other.Contains(elem) {
			s.Remove(elem)
		}
	}
}
