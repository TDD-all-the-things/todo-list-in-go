package todo_test

import (
	"testing"

	"github.com/TDD-all-the-things/todo-list-in-go/internal/todo"
	"github.com/stretchr/testify/assert"
)

// $todo add <item>
// 1.<item>
// Item <itemIndex> added.

// $todo done <itemIndex>
// Item <itemIndex> done.

// $todo list
// 1. <item1>
// 2. <item2>
// Total: 2 items.

// $todo list --all
// 1. <item1>
// 2. <item2>
// 3. [Done] <item3>
// Total: 3 items, 1 item done

func TestTodoList(t *testing.T) {

	t.Run("add + list", func(t *testing.T) {
		l := todo.NewList()

		item1 := todo.Item{Content: "Item-1"}
		item2 := todo.Item{Content: "Item-2"}
		item3 := todo.Item{Content: "Item-3"}

		index1 := l.Add(item1)
		index2 := l.Add(item2)
		index3 := l.Add(item3)

		items := l.ListItems()

		assert.Equal(t, item1.Content, items[index1.N].Content)
		assert.Equal(t, item2.Content, items[index2.N].Content)
		assert.Equal(t, item3.Content, items[index3.N].Content)
	})
}
