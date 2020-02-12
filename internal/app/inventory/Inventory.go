package inventory

type Item struct {
	ID int
	Name string
}

type Inventory []Item

func (i Inventory) GetValByID(id int) string {
	for _, item := range i {
		if item.ID == id {
			return item.Name
		}
	}
	return "Item Not Found"
}

func (i Inventory) GetIDByVal(val string) int {
	for _, item := range i {
		if item.Name == val {
			return item.ID
		}
	}
	return 0
}

func (i Inventory) GetVals() []string {
	var v []string
	for _, item := range i {
		v = append(v, item.Name)
	}
	return v
}