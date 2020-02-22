package inventory

// Item : Basic struct for Inventory Item
type Item struct {
	ID   int
	Name string
}

// Inventory : splice of Inventory Items
type Inventory []Item

// GetValByID : Get Inventory Item Name by ID
func (i Inventory) GetValByID(id int) string {
	for _, item := range i {
		if item.ID == id {
			return item.Name
		}
	}
	return "Item Not Found"
}

// GetValByID : Get Inventory Item ID by Name
func (i Inventory) GetIDByVal(val string) int {
	for _, item := range i {
		if item.Name == val {
			return item.ID
		}
	}
	return 0
}

// GetVals : Get splice of Inventory Item Names
func (i Inventory) GetVals() []string {
	var v []string
	for _, item := range i {
		v = append(v, item.Name)
	}
	return v
}
