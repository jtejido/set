package set

import "sync"

var (
	sss *safeSet
	_   Set = sss
)

type safeSet struct {
	s *unsafeSet
	sync.RWMutex
}

func newSafeSet() safeSet {
	return safeSet{s: newUnsafeSet()}
}

func (s *safeSet) Add(i ...Equality) {
	s.Lock()
	s.s.Add(i...)
	s.Unlock()
}

func (s *safeSet) Contains(i ...Equality) bool {
	s.RLock()
	ret := s.s.Contains(i...)
	s.RUnlock()
	return ret
}

func (s *safeSet) Clear() {
	s.Lock()
	s.s = newUnsafeSet()
	s.Unlock()
}

func (s *safeSet) Remove(i Equality) {
	s.Lock()

	for k := range s.s.index {
		if i.Equals(k) {
			delete(s.s.index, k)
			break
		}
	}

	s.Unlock()
}

func (s *safeSet) Len() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.s.index)
}

func (s *safeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		s.RLock()
		for key := range s.s.index {
			ch <- key
		}
		close(ch)
		s.RUnlock()
	}()

	return ch
}

func (s *safeSet) Equal(other Set) bool {
	o := other.(*safeSet)

	s.RLock()
	o.RLock()

	ret := s.s.Equal(o.s)
	s.RUnlock()
	o.RUnlock()
	return ret
}

func (s *safeSet) Clone() Set {
	s.RLock()

	unsafeClone := s.s.Clone().(*unsafeSet)
	ret := &safeSet{s: unsafeClone}
	s.RUnlock()
	return ret
}

func (s *safeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, s.Len())
	s.RLock()
	for key := range s.s.index {
		keys = append(keys, key)
	}
	s.RUnlock()
	return keys
}

func (s *safeSet) RemoveFrom(other Set) {
	o := other.(*safeSet)
	s.RLock()
	o.RLock()
	defer s.RUnlock()
	defer o.RUnlock()

	s.s.RemoveFrom(o.s)
}

func (s *safeSet) AddFrom(other Set) {
	o := other.(*safeSet)
	s.RLock()
	o.RLock()
	defer s.RUnlock()
	defer o.RUnlock()

	s.s.AddFrom(o.s)
}

func (s *safeSet) RetainFrom(other Set) {
	o := other.(*safeSet)

	s.RLock()
	o.RLock()
	defer s.RUnlock()
	defer o.RUnlock()

	s.s.RetainFrom(o.s)
}
