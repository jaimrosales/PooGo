package invoiceitem

//item contains of information of a invoice item
type Item struct {
	id      uint
	product string
	value   float64
}

//New returns a new item el constructor
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

//value getter of item.value
func (i Item) Value() float64 {
	return i.value

}
