package orders

import "github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"

func OrderItemListToMap(items []*entity.OrderItem) map[string]int {

	order := make(map[string]int)

	for _, item := range items {
		order[item.MenuItemID] = item.Quantity
	}

	return order
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
