package list

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

// to do
func (list *DoubleLinkedList) AddOnIndex(value int, index int) error {
	return nil
}

// to do
func (list *DoubleLinkedList) RemoveLast() {}

// to do
func (list *DoubleLinkedList) RemoveOnIndex(index int) error {
	return nil
}

// to do
func (list *DoubleLinkedList) Get(index int) (int, error) {
	return -1, nil
}

// to do
func (list *DoubleLinkedList) Update(value int, index int) error {
	return nil
}

// to do
func (list *DoubleLinkedList) Size() int {
	return list.size
}
