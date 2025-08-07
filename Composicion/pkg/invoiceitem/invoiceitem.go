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
/*func (i Item) Value() float64 {
	return i.value

}*/ // El getter se puede eliminar ya que ahora el calculo del total se hace desde el tipo Items

type Items []Item

func NewItems(items ...Item) Items {
	var is Items
	for _, item := range items {
		is = append(is, item)
	}
	return is
}

func (is Items) Total() float64 {
	var total float64
	for _, item := range is {
		total += item.value
	}
	return total
}
