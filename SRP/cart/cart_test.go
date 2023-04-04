package cart

import (
	"testing"

	"github.com/nikzayn/solids-in-golang/SRP/item"
)

func TestCart(t *testing.T) {
	t.Run("Add items in cart", func(t *testing.T) {
		var c *Cart
		addItem := item.Item{Name: "Bread", Price: 44.0}
		got := c.AddItem(addItem)
		want := item.Item{Name: "Bread", Price: 44.0}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
