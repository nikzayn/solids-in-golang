package cart

import (
	"testing"

	"github.com/nikzayn/solids-in-golang/SRP/item"
	"github.com/stretchr/testify/assert"
)

func TestAddItemsInCart(t *testing.T) {
	t.Run("Add items in cart", func(t *testing.T) {
		var c Cart
		addItem1 := item.Item{Name: "Bread", Price: 44.4}
		addItem2 := item.Item{Name: "Eggs", Price: 34.4}
		addItem3 := item.Item{Name: "Milk", Price: 50.9}
		addItem4 := item.Item{Name: "Cheese", Price: 66.4}

		c.AddItem(addItem1)
		c.AddItem(addItem2)
		c.AddItem(addItem3)
		c.AddItem(addItem4)

		assert.Equal(t, c.items, c.items)
	})
}

func TestRemoveItemsInCart(t *testing.T) {
	t.Run("Add items in cart", func(t *testing.T) {
		var c Cart
		addItem1 := item.Item{Name: "Bread", Price: 44.4}
		addItem2 := item.Item{Name: "Eggs", Price: 34.4}
		addItem3 := item.Item{Name: "Milk", Price: 50.9}
		addItem4 := item.Item{Name: "Cheese", Price: 66.4}

		c.AddItem(addItem1)
		c.AddItem(addItem2)
		c.AddItem(addItem3)
		c.AddItem(addItem4)

		c.RemoveItem(addItem1)
		c.RemoveItem(addItem2)
		c.RemoveItem(addItem3)
		c.RemoveItem(addItem4)

		assert.Equal(t, c.items, c.items)
	})
}

func BenchmarkAddItemsInCart(b *testing.B) {
	var c Cart
	addItem := item.Item{Name: "Bread", Price: 44.4}

	for i := 0; i < b.N; i++ {
		c.AddItem(addItem)
	}
}

func BenchmarkRemoveItemsInCart(b *testing.B) {
	var c Cart
	addItem := item.Item{Name: "Bread", Price: 44.4}

	for i := 0; i < b.N; i++ {
		c.RemoveItem(addItem)
	}
}
