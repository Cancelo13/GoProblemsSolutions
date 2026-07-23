package cards

func FavoriteCards() []int {
	return []int{2, 6, 9}
}

func GetItem(items []int, index int) int {
	if index < 0 || index >= len(items) {
		return -1
	}
	return items[index]
}

func SetItem(items []int, index int, value int) []int {
    if index < 0 || index >= len(items){
        return append(items, value)
    }
	items[index] = value
	return items
}

func PrependItems(items []int, values ...int) []int {
	if len(values) == 0 {
		return items
	}
	return append(values, items...)
}

func RemoveItem(items []int, index int) []int {
	if index < 0 || index >= len(items) {
		return items
	}
	var newItems []int
	for i := 0; i < len(items); i++ {
		if i == index {
			continue
		} else {
			newItems = append(newItems, items[i])
		}
	}
	return newItems
}

