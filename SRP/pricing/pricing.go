package pricing

import (
	items "github.com/nikzayn/solids-in-golang/SRP/item"
)

/*
Description - Calculate total price of items in your cart
  - @Input() - item -> Item struct
  - @Output() -> float64
  - TC - O(N) | SC - O(1)
*/
func CalculateTotalPrice(items []items.Item) float64 {
	var total float64
	for _, item := range items {
		total += item.Price
	}
	return total
}
