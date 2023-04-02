package cart

import (
	items "github.com/nikzayn/solids-in-golang/SRP/item"
)

// Define Cart struct in which total number of items come
type Cart struct {
	items []items.Item
}

/*
Description - Add item in your shopping cart
  - @Input() - item -> Item struct
  - TC - O(1) | SC - O(1)
*/
func (c *Cart) AddItem(item items.Item) {
	c.items = append(c.items, item)
}

/*
Description - Remove item in your shopping cart
  - @Input() - item -> Item struct
  - TC - O(N) | SC - O(N)
*/
func (c *Cart) RemoveItem(item items.Item) {
	for i, v := range c.items {
		if v == item {
			c.items = append(c.items[:i], c.items[i+1:]...)
			break
		}
	}
}
