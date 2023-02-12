package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{"quarter_of_a_dozen": 3, "half_of_a_dozen": 6, "dozen": 12, "small_gross": 120, "gross": 144, "great_gross": 1728}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	additionalAmount, inUnits := units[unit]
	if inUnits == false {
		return false
	}

	bill[item] += additionalAmount

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	oldAmount, inBill := bill[item]
	removedAmount, inUnits := units[unit]
	if inBill == false {
		return false
	}
	if inUnits == false {
		return false
	}
	if bill[item] < removedAmount {
		return false
	} else if oldAmount == removedAmount {
		delete(bill, item)
	} else {
		bill[item] = oldAmount - removedAmount
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	oldAmount, inBill := bill[item]
	if inBill == false {
		return 0, false
	}
	return oldAmount, true
}
