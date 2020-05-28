package set

type Set interface {
	Add(i ...interface{})
	Len() int
	Clear()
	Contains(i ...interface{}) bool
	Equal(other Set) bool
	Iter() <-chan interface{}
	Remove(i interface{})
	ToSlice() []interface{}
	RemoveFrom(other Set)
	AddFrom(other Set)
	RetainFrom(other Set)
	Clone() Set
}

func NewOrderedSet(s ...interface{}) Set {
	set := newSafeOrderedSet()
	for _, item := range s {
		set.Add(item)
	}
	return &set
}

func NewOrderedSetFromSlice(s []interface{}) Set {
	a := NewOrderedSet(s...)
	return a
}

func NewUnsafeOrderedSet() Set {
	return newUnsafeOrderedSet()
}

func NewUnsafeOrderedSetFromSlice(s []interface{}) Set {
	a := NewUnsafeOrderedSet()
	for _, item := range s {
		a.Add(item)
	}
	return a
}

func NewSet(s ...interface{}) Set {
	set := newSafeSet()
	for _, item := range s {
		set.Add(item)
	}
	return set
}

func NewSetFromSlice(s []interface{}) Set {
	a := NewSet(s...)
	return a
}

func NewUnsafeSet() Set {
	return newUnsafeSet()
}

func NewUnsafeSetFromSlice(s []interface{}) Set {
	a := NewUnsafeSet()
	for _, item := range s {
		a.Add(item)
	}
	return a
}
