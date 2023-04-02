package pricing

import (
	"testing"

	"github.com/nikzayn/solids-in-golang/SRP/item"
	"github.com/stretchr/testify/assert"
)

func TestPricing(t *testing.T) {
	var result []item.Item
	var totalPrice float64
	item_1 := item.Item{Name: "Bread", Price: 44.4}
	item_2 := item.Item{Name: "Eggs", Price: 34.4}
	item_3 := item.Item{Name: "Milk", Price: 50.9}
	item_4 := item.Item{Name: "Cheese", Price: 66.4}

	result = append(result, item_1, item_2, item_3, item_4)

	for _, price := range result {
		totalPrice += price.Price
	}

	assert.Equal(t, CalculateTotalPrice(result), totalPrice)
}

func BenchmarkPricing(b *testing.B) {
	var result []item.Item
	item := item.Item{Name: "Bread", Price: 44.4}
	result = append(result, item)

	for i := 0; i < b.N; i++ {
		CalculateTotalPrice(result)
	}
}
