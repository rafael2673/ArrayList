package main

type List interface {
	Add(value int)
	AddOnIndex(value int, index int) error
	RemoveLast()
	RemoveOnIndex(index int) error
	Get(index int) (int, error)
	Update(value int, index int) error
	Size() int
}
