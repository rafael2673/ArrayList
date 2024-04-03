package list

import (
	"errors"
)

type NodeDoubleLinked struct {
	value int
	prev  *NodeDoubleLinked
	next  *NodeDoubleLinked
}

type DoubleLinkedList struct {
	head *NodeDoubleLinked
	tail *NodeDoubleLinked
	size int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (list *DoubleLinkedList) Add(value int) {
	newNode := &NodeDoubleLinked{value: value, prev: nil, next: nil}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
	list.size++
}

func (list *DoubleLinkedList) AddOnIndex(value int, index int) error {
	if index <= list.size && index > -1 {
		newNode := &NodeDoubleLinked{value: value, prev: nil, next: nil}
		if list.head == nil {
			list.head = newNode
			list.tail = newNode
		} else if index == 0 {
			newNode.next = list.head
			list.head.prev = newNode
			list.head = newNode
		} else if index == list.size {
			newNode.prev = list.tail
			list.tail.next = newNode
			list.tail = newNode
		} else {
			current := list.head

			for i := 0; i < index-1; i++ {
				current = current.next
			}

			newNode.prev = current
			newNode.next = current.next
			current.next.prev = newNode
			current.next = newNode
		}
		list.size++
		return nil
	} else {
		return errors.New("index is not accessible")
	}
}

func (list *DoubleLinkedList) RemoveLast() {
	if list.tail == nil {
		return
	}
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
	} else {
		list.tail = list.tail.prev
		list.tail.next = nil
	}
	list.size--
}

func (list *DoubleLinkedList) RemoveOnIndex(index int) error {
	if list.size == 0 {
		return errors.New("list is empty")
	}
	if index <= list.size && index > -1 {
		if list.size == 1 {
			list.head = nil
			list.tail = nil
		} else if index == 0 {
			list.head = list.head.next
			list.head.prev = nil
		} else if index == list.size-1 {
			list.tail = list.tail.prev
			list.tail.next = nil
		} else {
			current := list.head

			for i := 0; i < index; i++ {
				current = current.next
			}

			current.prev.next = current.next
			current.next.prev = current.prev
		}

		list.size--
		return nil
	} else {
		return errors.New("index is not accessible")
	}
}

func (list *DoubleLinkedList) Get(index int) (int, error) {
	if list.size == 0 {
		return -1, errors.New("list is empty")
	}
	if index <= list.size && index > -1 {
		current := list.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		return current.value, nil
	}
	return -1, errors.New("index is not accessible")
}

func (list *DoubleLinkedList) Update(value int, index int) error {
	if list.size == 0 {
		return errors.New("list is empty")
	}
	if index <= list.size && index > -1 {
		current := list.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		current.value = value
	}

	return nil
}

func (list *DoubleLinkedList) Size() int {
	return list.size
}
