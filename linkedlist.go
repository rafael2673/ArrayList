package list

import (
	"errors"
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
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
			for i := 1; i < index; i++ {
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

func (list *LinkedList) RemoveLast() {
	if list.head == nil {
		return
	} else if list.head.next == nil {
		list.head = nil
	} else {
		current := list.head
		for current.next.next != nil {
			current = current.next
		}
		current.next = nil
	}
	list.size--

}

func (list *LinkedList) RemoveOnIndex(index int) error {
	if index >= 0 && index <= list.size {
		if index == 0 {
			list.head = list.head.next
		} else {
			current := list.head

			for i := 1; i < index; i++ {
				current = current.next
			}

			current.next = current.next.next
		}
		list.size--
		return nil
	} else {
		return errors.New("index not accessible")
	}
}

func (list *LinkedList) Get(index int) (int, error) {
	if index >= 0 && index <= list.size {
		current := list.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		return current.val, nil
	} else {
		return -1, errors.New("index not accessible")
	}
}
func (list *LinkedList) GetAll() {
	fmt.Println("Size: ", list.Size())
	current := list.head
	fmt.Println(current.val)
	for i := 0; i < list.size-1; i++ {
		current = current.next
		fmt.Println(current.val)
	}
}
func (list *LinkedList) Update(value int, index int) error {
	if index >= 0 && index <= list.size {
		current := list.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		current.val = value
		return nil
	} else {
		return errors.New("index not accessible")
	}
}
func (list *LinkedList) Size() int {
	return list.size
}

func (list *LinkedList) Reverse() {
	newList := &LinkedList{}

	current := list.head

	for current != nil {
		newList.AddOnIndex(current.val, 0)
		current = current.next
	}

	list.head = newList.head
}
