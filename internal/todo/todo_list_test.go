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

		item := todo.Item{
			Content: "Item-1",
		}
		index := l.Add(item)

		items := l.ListItems()

		assert.Equal(t, item.Content, items[index.N].Content)
	})
}
