package todo

import (
	"testing"

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
	assert.Equal(t, false, true)
}
