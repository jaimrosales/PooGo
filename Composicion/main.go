package main

import (
	"Composicion/pkg/customer"
	"Composicion/pkg/invoice"
	"Composicion/pkg/invoiceitem"
	"fmt"
)

func main() {
	i := invoice.New(
		"mexico",
		"ixtlanDelRio",
		customer.New("lupe", "5 mayo 14", "3241084565"),
		/*[]invoiceitem.Item{
			invoiceitem.New(1,"mantenimiento preventivo basico",5000),
			invoiceitem.New(2, "lavado de inyectores con ultrasonido",1000),
			invoiceitem.New(3, "revision de suspencion",200),
		},*/ // Este se puede sustituir por un constructor de NewItems, metodo que permite crear varios item en conjunto
		invoiceitem.NewItems(
			invoiceitem.New(1, "mantenimiento preventivo basico", 5000),
			invoiceitem.New(2, "lavado de inyectores con ultrasonido", 1000),
			invoiceitem.New(3, "revision de suspencion", 200),
		),
	)
	i.SetTotal()

	fmt.Printf("%v", i)
}
