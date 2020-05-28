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

func (s *safeOrderedSet) Add(i ...interface{}) {
	s.Lock()
	s.s.Add(i...)
	s.Unlock()
}

func (s *safeOrderedSet) Contains(i ...interface{}) bool {
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

func (s *safeOrderedSet) Remove(i interface{}) {
	s.Lock()
	s.s.Remove(i)
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
	s.Lock()
	o.RLock()
	s.s.RemoveFrom(o.s)
	s.Unlock()
	o.RUnlock()
}

func (s *safeOrderedSet) AddFrom(other Set) {
	o := other.(*safeOrderedSet)
	s.Lock()
	o.RLock()
	s.s.AddFrom(o.s)
	s.Unlock()
	o.RUnlock()
}

func (s *safeOrderedSet) RetainFrom(other Set) {
	o := other.(*safeOrderedSet)
	s.Lock()
	o.RLock()
	s.s.RetainFrom(o.s)
	s.Unlock()
	o.RUnlock()
}
