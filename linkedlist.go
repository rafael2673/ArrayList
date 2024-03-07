package main

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (list *LinkedList) Add(value int)                         {}
func (list *LinkedList) AddOnIndex(value int, index int) error {}
func (list *LinkedList) RemoveLast()                           {}
func (list *LinkedList) RemoveOnIndex(index int) error         {}
func (list *LinkedList) Get(index int) (int, error)            {}
func (list *LinkedList) Update(value int, index int) error     {}
func (list *LinkedList) Size() int                             {}
