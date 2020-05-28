# This is a Golang implementation of a set (Ordered and Unordered).
[![Build Status](https://travis-ci.com/jtejido/set.svg?branch=master)](https://travis-ci.com/jtejido/set)
[![codecov](https://codecov.io/gh/jtejido/set/branch/master/graph/badge.svg)](https://codecov.io/gh/jtejido/set)
[![MIT license](https://img.shields.io/badge/license-MIT-blue)](https://opensource.org/licenses/MIT)

## Set API

```golang
type Set interface {
	Add(i ...interface{})
	Len() int
	Clear()
	Contains(i ...interface{}) bool
	Equal(other Set) bool
	Iter() <-chan interface{}
	Remove(i interface{})
	ToSlice() []interface{}
	AddFrom(other Set)  // in-place Union with other set
	RetainFrom(other Set) // in-place Intersect with other set
	RemoveFrom(other Set) // in-place Difference with other set
	Clone() Set
}
```