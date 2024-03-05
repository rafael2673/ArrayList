package main

import "errors"

type ArrayList struct {
	values   []int
	inserted int // Esse daqui é a quantidade de valores da lista
}

func (list *ArrayList) Add(value int) {
	if list.inserted >= len(list.values) {
	}
	list.values[list.inserted] = value
	list.inserted++
}

func (list *ArrayList) AddOnIndex(value int, index int) error {
	if index >= 0 && index <= list.inserted {
		if list.inserted >= len(list.values) {
		}
		for i := list.inserted; i < index; i-- {
			list.values[i] = list.values[i-1]
		}
		list.values[index] = value
		list.inserted++
		return nil
	} else {
		return errors.New("index not acessible")
	}
}

func (list *ArrayList) RemoveLast() {

}

func (list *ArrayList) RemoveOnIndex(index int) error {
	return errors.New("index not acessible")
}

func (list *ArrayList) Get(index int) (int, error) {
	return 0, errors.New("index not acessible")
}

func (list *ArrayList) Update(value int, index int) error {
	return errors.New("index not acessible")
}

func (list *ArrayList) Size() int {
	return 0
}
