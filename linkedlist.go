package list

import "errors"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (list *LinkedList) Add(value int) {
	newNode := &Node{val: value, next: nil}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head

		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}

	list.size++
}

func (list *LinkedList) AddOnIndex(value int, index int) error {
	if index >= 0 && index <= list.size {
		newNode := &Node{val: value, next: nil}
		if index == 0 {
			newNode.next = list.head
			list.head = newNode
		} else {
			current := list.head
			for i := 1; current.next != nil && i < index-1; i++ {
				current = current.next
			}

			newNode.next = current.next
			current.next = newNode
		}
		list.size++
		return nil
	} else {
		return errors.New("index not accessible")
	}
}

func (list *LinkedList) RemoveLast() {}
func (list *LinkedList) RemoveOnIndex(index int) error {
	if index >= 0 && index <= list.size {
		return nil
	} else {
		return errors.New("index not accessible")
	}
}
func (list *LinkedList) Get(index int) (int, error) {
	if index >= 0 && index <= list.size {
		return index, nil
	} else {
		return -1, errors.New("index not accessible")
	}
}
func (list *LinkedList) Update(value int, index int) error {
	if index >= 0 && index <= list.size {
		return nil
	} else {
		return errors.New("index not accessible")
	}
}
func (list *LinkedList) Size() int {
	return list.size
}
