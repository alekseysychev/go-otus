package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	Length       int
	FrontElement *ListItem
	BackElement  *ListItem
}

func (l *list) Len() int {
	return l.Length
}

func (l *list) Front() *ListItem {
	return l.FrontElement
}

func (l *list) Back() *ListItem {
	return l.BackElement
}

func (l *list) PushFront(v interface{}) *ListItem {
	// create element
	var e ListItem = ListItem{
		Value: v,
		Next:  l.FrontElement,
	}

	if l.Front() != nil {
		l.FrontElement.Prev = &e
	}

	// add element to first
	l.FrontElement = &e
	// add element to back
	if l.Back() == nil {
		l.BackElement = &e
	}
	// inc length
	l.Length++

	return &e
}
func (l *list) PushBack(v interface{}) *ListItem {
	// create element
	var e ListItem = ListItem{
		Value: v,
		Prev:  l.BackElement,
	}

	if l.Back() != nil {
		l.BackElement.Next = &e
	}

	// add element to first
	l.BackElement = &e
	// add element to back
	if l.Front() == nil {
		l.FrontElement = &e
	}
	// inc length
	l.Length++

	return &e
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		l.FrontElement = i.Next
	} else {
		i.Prev.Next = i.Next
	}
	if i.Next == nil {
		l.BackElement = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}
	l.Length--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	} else {
		i.Prev.Next = i.Next
	}
	if i.Next == nil {
		l.BackElement = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}
	l.Front().Prev = i
	i.Next = l.Front()
	l.FrontElement = i
}

func NewList() List {
	return new(list)
}
