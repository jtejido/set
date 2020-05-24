package set

import "sync"

var (
	ss *safeOrderedSet
	_  Set = ss
)

type safeOrderedSet struct {
	s *unsafeOrderedSet
	sync.RWMutex
}

func newSafeOrderedSet() safeOrderedSet {
	return safeOrderedSet{s: newUnsafeOrderedSet()}
}

func (s *safeOrderedSet) Add(i ...Equality) {
	s.Lock()
	s.s.Add(i...)
	s.Unlock()
}

func (s *safeOrderedSet) Contains(i ...Equality) bool {
	s.RLock()
	ret := s.s.Contains(i...)
	s.RUnlock()
	return ret
}

func (s *safeOrderedSet) Clear() {
	s.Lock()
	s.s = newUnsafeOrderedSet()
	s.Unlock()
}

func (s *safeOrderedSet) Remove(i Equality) {
	s.Lock()

	for k, index := range s.s.index {
		if i.Equals(k) {
			if _, found := s.s.m[index]; !found {
				return
			}

			delete(s.s.m, index)
			// Remove the key
			for i := range s.s.keys {
				if s.s.keys[i] == index {
					s.s.keys = append(s.s.keys[:i], s.s.keys[i+1:]...)
					break
				}
			}

			delete(s.s.index, k)
			break
		}
	}

	s.Unlock()
}

func (s *safeOrderedSet) Len() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.s.m)
}

func (s *safeOrderedSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		s.RLock()
		for _, key := range s.s.keys {
			ch <- s.s.m[key]
		}
		close(ch)
		s.RUnlock()
	}()

	return ch
}

func (s *safeOrderedSet) Equal(other Set) bool {
	o := other.(*safeOrderedSet)

	s.RLock()
	o.RLock()

	ret := s.s.Equal(o.s)
	s.RUnlock()
	o.RUnlock()
	return ret
}

func (s *safeOrderedSet) Clone() Set {
	s.RLock()

	unsafeClone := s.s.Clone().(*unsafeOrderedSet)
	ret := &safeOrderedSet{s: unsafeClone}
	s.RUnlock()
	return ret
}

func (s *safeOrderedSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, s.Len())
	s.RLock()
	for _, key := range s.s.keys {
		keys = append(keys, s.s.m[key])
	}
	s.RUnlock()
	return keys
}

func (s *safeOrderedSet) RemoveFrom(other Set) {
	o := other.(*safeOrderedSet)
	s.RLock()
	o.RLock()
	defer s.RUnlock()
	defer o.RUnlock()

	s.s.RemoveFrom(o.s)
}

func (s *safeOrderedSet) AddFrom(other Set) {
	o := other.(*safeOrderedSet)
	s.RLock()
	o.RLock()
	defer s.RUnlock()
	defer o.RUnlock()

	s.s.AddFrom(o.s)

}

func (s *safeOrderedSet) RetainFrom(other Set) {
	o := other.(*safeOrderedSet)

	s.RLock()
	o.RLock()
	defer s.RUnlock()
	defer o.RUnlock()

	s.s.RetainFrom(o.s)
}
