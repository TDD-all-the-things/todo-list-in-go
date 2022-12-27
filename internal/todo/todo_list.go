package todo

type List interface {
	Add(item Item) ItemIndex
	Done(index ItemIndex) error
	ListItems() []Item
	ListAllItems() []Item
}

type Item struct {
}

type ItemIndex struct {
	N int
}

type list struct {
}

func NewList() List {
	return &list{}
}

func (l *list) Add(item Item) ItemIndex {
	return ItemIndex{}
}

func (l *list) Done(i ItemIndex) error {
	return nil
}

func (l *list) ListItems() []Item {
	return nil
}

func (l *list) ListAllItems() []Item {
	return nil
}
