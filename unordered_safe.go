package set

import "sync"

var (
	sss safeSet
	_   Set = sss
)

var mu = sync.RWMutex{}

type safeSet unsafeSet

func newSafeSet() safeSet {
	return safeSet(newUnsafeSet())
}

func (s safeSet) Add(i ...interface{}) {
	mu.Lock()
	unsafeSet(s).Add(i...)
	mu.Unlock()
}

func (s safeSet) Contains(i ...interface{}) bool {
	mu.RLock()
	ret := unsafeSet(s).Contains(i...)
	mu.RUnlock()
	return ret
}

func (s safeSet) Clear() {
	mu.Lock()
	for k := range s {
		delete(s, k)
	}
	mu.Unlock()
}

func (s safeSet) Remove(i interface{}) {
	mu.Lock()
	if _, found := s[i]; !found {
		return
	}
	delete(s, i)
	mu.Unlock()
}

func (s safeSet) Len() int {
	mu.RLock()
	defer mu.RUnlock()
	return len(s)
}

func (s safeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for key := range s {
			ch <- key
		}
		close(ch)
	}()

	return ch
}

func (s safeSet) Equal(other Set) bool {
	o := other.(safeSet)

	mu.RLock()
	ret := unsafeSet(s).Equal(unsafeSet(o))
	mu.RUnlock()

	return ret
}

func (s safeSet) Clone() Set {
	mu.RLock()
	defer mu.RUnlock()
	unsafeClone := unsafeSet(s).Clone()
	ret := safeSet(unsafeClone.(unsafeSet))

	return ret
}

func (s safeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, s.Len())
	mu.RLock()
	for key := range s {
		keys = append(keys, key)
	}
	mu.RUnlock()
	return keys
}

func (s safeSet) RemoveFrom(other Set) {
	mu.Lock()
	o := other.(safeSet)
	unsafeSet(s).RemoveFrom(unsafeSet(o))
	mu.Unlock()
}

func (s safeSet) AddFrom(other Set) {
	mu.Lock()
	o := other.(safeSet)
	unsafeSet(s).AddFrom(unsafeSet(o))
	mu.Unlock()
}

func (s safeSet) RetainFrom(other Set) {
	mu.Lock()
	o := other.(safeSet)
	unsafeSet(s).RetainFrom(unsafeSet(o))
	mu.Unlock()
}
