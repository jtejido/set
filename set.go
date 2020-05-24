package set

type Set interface {
	Add(i ...Equality)
	Len() int
	Clear()
	Contains(i ...Equality) bool
	Equal(other Set) bool
	Iter() <-chan interface{}
	Remove(i Equality)
	ToSlice() []interface{}
	RemoveFrom(other Set)
	AddFrom(other Set)
	RetainFrom(other Set)
	Clone() Set
}

type Equality interface {
	Equals(interface{}) bool
}

func NewOrderedSet(s ...Equality) Set {
	set := newSafeOrderedSet()
	for _, item := range s {
		set.Add(item)
	}
	return &set
}

func NewOrderedSetFromSlice(s []Equality) Set {
	a := NewOrderedSet(s...)
	return a
}

func NewUnsafeOrderedSet() Set {
	return newUnsafeOrderedSet()
}

func NewUnsafeOrderedSetFromSlice(s []Equality) Set {
	a := NewUnsafeOrderedSet()
	for _, item := range s {
		a.Add(item)
	}
	return a
}

func NewSet(s ...Equality) Set {
	set := newSafeSet()
	for _, item := range s {
		set.Add(item)
	}
	return &set
}

func NewSetFromSlice(s []Equality) Set {
	a := NewSet(s...)
	return a
}

func NewUnsafeSet() Set {
	return newUnsafeSet()
}

func NewUnsafeSetFromSlice(s []Equality) Set {
	a := NewUnsafeSet()
	for _, item := range s {
		a.Add(item)
	}
	return a
}
