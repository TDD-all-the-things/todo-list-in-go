package todo

type List interface {
	Add(item Item) ItemIndex
	Done(index ItemIndex) error
	ListItems() []Item
	ListAllItems() []Item
}

type Item struct {
	Content string
}

type ItemIndex struct {
	N int
}

type list struct {
	items []Item
}

func NewList() List {
	return &list{items: make([]Item, 0)}
}

func (l *list) Add(item Item) ItemIndex {
	l.items = append(l.items, item)
	return ItemIndex{N: len(l.items) - 1}
}

func (l *list) Done(i ItemIndex) error {
	return nil
}

func (l *list) ListItems() []Item {
	return append([]Item{}, l.items...)
}

func (l *list) ListAllItems() []Item {
	return nil
}
