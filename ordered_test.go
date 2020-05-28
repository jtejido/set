package set

import "testing"

func makeOrderedSet(ints []int) Set {
	set := NewOrderedSet()
	for _, i := range ints {
		set.Add(i)
	}
	return set
}

func makeUnsafeOrderedSet(ints []int) Set {
	set := NewUnsafeOrderedSet()
	for _, i := range ints {
		set.Add(i)
	}
	return set
}

func TestNewOrderedSet(t *testing.T) {
	a := NewOrderedSet()
	if a.Len() != 0 {
		t.Error("NewOrderedSet should start out as an empty set")
	}

	assertEqual(NewOrderedSetFromSlice([]interface{}{}), NewOrderedSet(), t)
	assertEqual(NewOrderedSetFromSlice([]interface{}{1}), NewOrderedSet(1), t)
	assertEqual(NewOrderedSetFromSlice([]interface{}{1, 2}), NewOrderedSet(2, 1), t)
	assertEqual(NewOrderedSetFromSlice([]interface{}{"a"}), NewOrderedSet("a"), t)
	assertEqual(NewOrderedSetFromSlice([]interface{}{"a", "b"}), NewOrderedSet("b", "a"), t)
}

func TestNewUnsafeOrderedSet(t *testing.T) {
	a := NewUnsafeOrderedSet()

	if a.Len() != 0 {
		t.Error("NewSet should start out as an empty set")
	}
}

func TestAddOrderedSet(t *testing.T) {
	a := makeOrderedSet([]int{1, 2, 3})

	if a.Len() != 3 {
		t.Error("AddSet does not have a size of 3 even though 3 items were added to a new set")
	}
}

func TestAddUnsafeOrderedSet(t *testing.T) {
	a := makeUnsafeOrderedSet([]int{1, 2, 3})

	if a.Len() != 3 {
		t.Error("AddSet does not have a size of 3 even though 3 items were added to a new set")
	}
}

func TestAddOrderedSetNoDuplicate(t *testing.T) {
	a := makeOrderedSet([]int{7, 5, 3, 7})

	if a.Len() != 3 {
		t.Error("AddSetNoDuplicate set should have 3 elements since 7 is a duplicate")
	}

	if !(a.Contains(7) && a.Contains(5) && a.Contains(3)) {
		t.Error("AddSetNoDuplicate set should have a 7, 5, and 3 in it.")
	}
}

func TestAddUnsafeOrderedSetNoDuplicate(t *testing.T) {
	a := makeUnsafeOrderedSet([]int{7, 5, 3, 7})

	if a.Len() != 3 {
		t.Error("AddSetNoDuplicate set should have 3 elements since 7 is a duplicate")
	}

	if !(a.Contains(7) && a.Contains(5) && a.Contains(3)) {
		t.Error("AddSetNoDuplicate set should have a 7, 5, and 3 in it.")
	}
}

func TestRemoveOrderedSet(t *testing.T) {
	a := makeOrderedSet([]int{6, 3, 1})

	a.Remove(3)

	if a.Len() != 2 {
		t.Error("RemoveSet should only have 2 items in the set")
	}

	if !(a.Contains(6) && a.Contains(1)) {
		t.Error("RemoveSet should have only items 6 and 1 in the set")
	}

	a.Remove(6)
	a.Remove(1)

	if a.Len() != 0 {
		t.Error("RemoveSet should be an empty set after removing 6 and 1")
	}
}

func TestRemoveUnsafeOrderedSet(t *testing.T) {
	a := makeUnsafeOrderedSet([]int{6, 3, 1})

	a.Remove(3)

	if a.Len() != 2 {
		t.Error("RemoveSet should only have 2 items in the set")
	}

	if !(a.Contains(6) && a.Contains(1)) {
		t.Error("RemoveSet should have only items 6 and 1 in the set")
	}

	a.Remove(6)
	a.Remove(1)

	if a.Len() != 0 {
		t.Error("RemoveSet should be an empty set after removing 6 and 1")
	}
}

func TestContainsOrderedSet(t *testing.T) {
	a := NewOrderedSet()

	a.Add(71)

	if !a.Contains(71) {
		t.Error("ContainsSet should contain 71")
	}

	a.Remove(71)

	if a.Contains(71) {
		t.Error("ContainsSet should not contain 71")
	}

	a.Add(13)
	a.Add(7)
	a.Add(1)

	if !(a.Contains(13) && a.Contains(7) && a.Contains(1)) {
		t.Error("ContainsSet should contain 13, 7, 1")
	}
}

func TestContainsUnsafeOrderedSet(t *testing.T) {
	a := NewUnsafeOrderedSet()

	a.Add(71)

	if !a.Contains(71) {
		t.Error("ContainsSet should contain 71")
	}

	a.Remove(71)

	if a.Contains(71) {
		t.Error("ContainsSet should not contain 71")
	}

	a.Add(13)
	a.Add(7)
	a.Add(1)

	if !(a.Contains(13) && a.Contains(7) && a.Contains(1)) {
		t.Error("ContainsSet should contain 13, 7, 1")
	}
}

func TestContainsMultipleOrderedSet(t *testing.T) {
	a := makeOrderedSet([]int{8, 6, 7, 5, 3, 0, 9})

	if !a.Contains(8, 6, 7, 5, 3, 0, 9) {
		t.Error("ContainsAll should contain Jenny's phone number")
	}

	if a.Contains(8, 6, 11, 5, 3, 0, 9) {
		t.Error("ContainsAll should not have all of these numbers")
	}
}

func TestContainsMultipleUnsafeOrderedSet(t *testing.T) {
	a := makeUnsafeOrderedSet([]int{8, 6, 7, 5, 3, 0, 9})

	if !a.Contains(8, 6, 7, 5, 3, 0, 9) {
		t.Error("ContainsAll should contain Jenny's phone number")
	}

	if a.Contains(8, 6, 11, 5, 3, 0, 9) {
		t.Error("ContainsAll should not have all of these numbers")
	}
}

func TestClearOrderedSet(t *testing.T) {
	a := makeOrderedSet([]int{2, 5, 9, 10})

	a.Clear()

	if a.Len() != 0 {
		t.Error("ClearSet should be an empty set")
	}
}

func TestClearUnsafeOrderedSet(t *testing.T) {
	a := makeUnsafeOrderedSet([]int{2, 5, 9, 10})

	a.Clear()

	if a.Len() != 0 {
		t.Error("ClearSet should be an empty set")
	}
}

func TestLenOrderedSet(t *testing.T) {
	a := NewOrderedSet()

	if a.Len() != 0 {
		t.Error("set should be an empty set")
	}

	a.Add(1)

	if a.Len() != 1 {
		t.Error("set should have a size of 1")
	}

	a.Remove(1)

	if a.Len() != 0 {
		t.Error("set should be an empty set")
	}

	a.Add(9)

	if a.Len() != 1 {
		t.Error("set should have a size of 1")
	}

	a.Clear()

	if a.Len() != 0 {
		t.Error("set should have a size of 1")
	}
}

func TestLenUnsafeOrderedSet(t *testing.T) {
	a := NewUnsafeOrderedSet()

	if a.Len() != 0 {
		t.Error("set should be an empty set")
	}

	a.Add(1)

	if a.Len() != 1 {
		t.Error("set should have a size of 1")
	}

	a.Remove(1)

	if a.Len() != 0 {
		t.Error("set should be an empty set")
	}

	a.Add(9)

	if a.Len() != 1 {
		t.Error("set should have a size of 1")
	}

	a.Clear()

	if a.Len() != 0 {
		t.Error("set should have a size of 1")
	}
}

func TestOrderedSetUnion(t *testing.T) {
	a := NewOrderedSet()

	b := NewOrderedSet()
	b.Add(1)
	b.Add(2)
	b.Add(3)
	b.Add(4)
	b.Add(5)

	a.AddFrom(b)

	if a.Len() != 5 {
		t.Error("set a is unioned with an empty set and therefore should have 5 elements in it")
	}

	d := NewOrderedSet()
	d.Add(10)
	d.Add(14)
	d.Add(0)

	a.AddFrom(d)
	if a.Len() != 8 {
		t.Error("set e should should have 8 elements in it after being unioned with set d")
	}

	f := NewOrderedSet()
	f.Add(14)
	f.Add(3)

	f.AddFrom(a)
	if f.Len() != 8 {
		t.Error("set f should still have 8 elements in it after being unioned with set f that has duplicates")
	}
}

func TestUnsafeOrderedSetUnion(t *testing.T) {
	a := NewUnsafeOrderedSet()

	b := NewUnsafeOrderedSet()
	b.Add(1)
	b.Add(2)
	b.Add(3)
	b.Add(4)
	b.Add(5)

	a.AddFrom(b)

	if a.Len() != 5 {
		t.Error("set a is unioned with an empty set and therefore should have 5 elements in it")
	}

	d := NewUnsafeOrderedSet()
	d.Add(10)
	d.Add(14)
	d.Add(0)

	a.AddFrom(d)
	if a.Len() != 8 {
		t.Error("set a should should have 8 elements in it after being unioned with set d")
	}

	f := NewUnsafeOrderedSet()
	f.Add(14)
	f.Add(3)

	f.AddFrom(a)
	if f.Len() != 8 {
		t.Error("set f should still have 8 elements in it after being unioned with set f that has duplicates")
	}
}

func TestUnsafeOrderedSetAddAsSlice(t *testing.T) {
	a := NewUnsafeOrderedSet()
	b := NewUnsafeOrderedSet()
	b.Add(1)
	b.Add(2)
	b.Add(3)
	b.Add(4)
	b.Add(5)

	a.Add(b.ToSlice()...)

	if a.Len() != 5 {
		t.Error("set b is added with from an empty set a and therefore should have 5 elements in it")
	}

	d := NewUnsafeOrderedSet()
	d.Add(10)
	d.Add(14)
	d.Add(0)

	a.Add(d.ToSlice()...)
	if a.Len() != 8 {
		t.Error("set a should should have 8 elements in it after being added with with items from set d")
	}

	e := NewUnsafeOrderedSet()
	e.Add(14)
	e.Add(3)

	e.Add(a.ToSlice()...)
	if e.Len() != 8 {
		t.Error("set e should still have 8 elements in it after being added with items from set a that has duplicates")
	}
}

func TestOrderedSetIntersect(t *testing.T) {
	a := NewOrderedSet()
	a.Add(1)
	a.Add(3)
	a.Add(5)

	b := NewOrderedSet()
	a.Add(2)
	a.Add(4)
	a.Add(6)

	a.RetainFrom(b)

	if a.Len() != 0 {
		t.Error("set a should be the empty set because there is no common items to intersect")
	}

	a.Add(10)
	b.Add(10)

	a.RetainFrom(b)

	if !(a.Len() == 1 && a.Contains(10)) {
		t.Error("set a should have a size of 1 and contain the item 10")
	}
}

func TestUnsafeOrderedSetIntersect(t *testing.T) {
	a := NewUnsafeOrderedSet()
	a.Add(1)
	a.Add(3)
	a.Add(5)

	b := NewUnsafeOrderedSet()
	a.Add(2)
	a.Add(4)
	a.Add(6)

	a.RetainFrom(b)

	if a.Len() != 0 {
		t.Error("set a should be the empty set because there is no common items to intersect")
	}

	a.Add(10)
	b.Add(10)

	a.RetainFrom(b)

	if !(a.Len() == 1 && a.Contains(10)) {
		t.Error("set a should have a size of 1 and contain the item 10")
	}
}

func TestOrderedSetDifference(t *testing.T) {
	a := NewOrderedSet()
	a.Add(1)
	a.Add(2)
	a.Add(3)

	b := NewOrderedSet()
	b.Add(1)
	b.Add(3)
	b.Add(4)
	b.Add(5)
	b.Add(6)
	b.Add(99)

	a.RemoveFrom(b)

	if !(a.Len() == 1 && a.Contains(2)) {
		t.Error("the difference of set a to b is the set of 1 item: 2")
	}
}

func TestUnsafeOrderedSetDifference(t *testing.T) {
	a := NewUnsafeOrderedSet()
	a.Add(1)
	a.Add(2)
	a.Add(3)

	b := NewUnsafeOrderedSet()
	b.Add(1)
	b.Add(3)
	b.Add(4)
	b.Add(5)
	b.Add(6)
	b.Add(99)

	a.RemoveFrom(b)

	if !(a.Len() == 1 && a.Contains(2)) {
		t.Error("the difference of set a to b is the set of 1 item: 2")
	}
}

func TestOrderedSetEqual(t *testing.T) {
	a := NewOrderedSet()
	b := NewOrderedSet()

	if !a.Equal(b) {
		t.Error("Both a and b are empty sets, and should be equal")
	}

	a.Add(10)

	if a.Equal(b) {
		t.Error("a should not be equal to b because b is empty and a has item 1 in it")
	}

	b.Add(10)

	if !a.Equal(b) {
		t.Error("a is now equal again to b because both have the item 10 in them")
	}

	b.Add(8)
	b.Add(3)
	b.Add(47)

	if a.Equal(b) {
		t.Error("b has 3 more elements in it so therefore should not be equal to a")
	}

	a.Add(8)
	a.Add(3)
	a.Add(47)

	if !a.Equal(b) {
		t.Error("a and b should be equal with the same number of elements")
	}
}

func TestUnsafeOrderedSetEqual(t *testing.T) {
	a := NewUnsafeOrderedSet()
	b := NewUnsafeOrderedSet()

	if !a.Equal(b) {
		t.Error("Both a and b are empty sets, and should be equal")
	}

	a.Add(10)

	if a.Equal(b) {
		t.Error("a should not be equal to b because b is empty and a has item 1 in it")
	}

	b.Add(10)

	if !a.Equal(b) {
		t.Error("a is now equal again to b because both have the item 10 in them")
	}

	b.Add(8)
	b.Add(3)
	b.Add(47)

	if a.Equal(b) {
		t.Error("b has 3 more elements in it so therefore should not be equal to a")
	}

	a.Add(8)
	a.Add(3)
	a.Add(47)

	if !a.Equal(b) {
		t.Error("a and b should be equal with the same number of elements")
	}
}

func TestOrderedSetClone(t *testing.T) {
	a := NewOrderedSet()
	a.Add(1)
	a.Add(2)

	b := a.Clone()

	if !a.Equal(b) {
		t.Error("Clones should be equal")
	}

	a.Add(3)
	if a.Equal(b) {
		t.Error("a contains one more element, they should not be equal")
	}

	c := a.Clone()
	c.Remove(1)

	if a.Equal(c) {
		t.Error("C contains one element less, they should not be equal")
	}
}

func TestUnsafeOrderedSetClone(t *testing.T) {
	a := NewUnsafeOrderedSet()
	a.Add(1)
	a.Add(2)

	b := a.Clone()

	if !a.Equal(b) {
		t.Error("Clones should be equal")
	}

	a.Add(3)
	if a.Equal(b) {
		t.Error("a contains one more element, they should not be equal")
	}

	c := a.Clone()
	c.Remove(1)

	if a.Equal(c) {
		t.Error("C contains one element less, they should not be equal")
	}
}

func TestOrderedSetIter(t *testing.T) {
	a := NewOrderedSet()

	a.Add("Z")
	a.Add("Y")
	a.Add("X")
	a.Add("W")

	b := NewOrderedSet()
	for val := range a.Iter() {
		b.Add(val)
	}

	if !a.Equal(b) {
		t.Error("The sets are not equal after iterating (Iter) through the first set")
	}
}

func TestUnsafeOrderedSetIter(t *testing.T) {
	a := NewUnsafeOrderedSet()

	a.Add("Z")
	a.Add("Y")
	a.Add("X")
	a.Add("W")

	b := NewUnsafeOrderedSet()
	for val := range a.Iter() {
		b.Add(val)
	}

	if !a.Equal(b) {
		t.Error("The sets are not equal after iterating (Iter) through the first set")
	}
}

func TestEmptyOrderedSetProperties(t *testing.T) {
	empty := NewOrderedSet()

	a := NewOrderedSet()
	a.Add(1)
	a.Add("foo")
	a.Add("bar")

	b := NewOrderedSet()
	b.Add("one")
	b.Add("two")
	b.Add(3)
	b.Add(4)

	a.AddFrom(empty)
	if !a.Equal(a) {
		t.Error("The union of any set with the empty set is supposed to be equal to itself")
	}

	a.RetainFrom(empty)
	if !a.Equal(empty) {
		t.Error("The intesection of any set with the empty set is supposed to be the empty set")
	}

	if empty.Len() != 0 {
		t.Error("Len of the empty set is supposed to be zero")
	}

}

func TestToSliceUnthreadsafeOrdered(t *testing.T) {
	s := makeUnsafeOrderedSet([]int{1, 2, 3})
	setAsSlice := s.ToSlice()
	if len(setAsSlice) != s.Len() {
		t.Errorf("Set length is incorrect: %v", len(setAsSlice))
	}

	for _, i := range setAsSlice {
		if !s.Contains(i) {
			t.Errorf("Set is missing element: %v", i)
		}
	}
}
