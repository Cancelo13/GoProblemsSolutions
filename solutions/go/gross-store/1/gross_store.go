package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

func NewBill() map[string]int{
	return map[string]int{}
}

func AddItem(bill map[string]int, units map[string]int, item string, unit string) bool {
	if units[unit] == 0 {
		return false
	}
	bill[item] += units[unit]
	return true
}

func RemoveItem(bill map[string]int, units map[string]int, item string, unit string) bool {
	if bill[item] == 0 || units[unit] == 0 || bill[item]-units[unit] < 0 {
		return false
	}
	bill[item] -= units[unit]
	if bill[item] == 0 {
		delete(bill, item)
	}
	return true
}

func GetItem(bill map[string]int, item string) (int, bool) {
	if bill[item] == 0 {
		return 0, false
	}
	return bill[item], true
}
