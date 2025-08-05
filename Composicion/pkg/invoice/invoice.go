package invoice

import (
	"Composicion/pkg/customer"
	"Composicion/pkg/invoiceitem"
)

// Invoice is the struct of a invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer  //relacion 1 a 1
	items   []invoiceitem.Item //relacion 1 a muchos, al crear un slice de tipo item
}

// New return a new Invoice
func New(country, city string, client customer.Customer, items []invoiceitem.Item) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}

}

// SetTotal is the setter of invoice.total
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}
