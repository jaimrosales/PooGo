package main

import "fmt"

// ya se utilizo el poliforfismo con las interfaces greet o getter
// utilizando tipos diferentes
type PayMethod interface {
	Pay()
}

type Paypal struct{}
type Cash struct{}
type CreditCard struct{}

func (p Paypal) Pay() {
	fmt.Println("Pagado con Paypal")
}
func (p Cash) Pay() {
	fmt.Println("Pagado con Cash")
}
func (p CreditCard) Pay() {
	fmt.Println("Pagado con Tarjeta de credito")
}

func FactoryPay(method uint) PayMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		return nil
	}
}

func main() {
	var met uint
	fmt.Println("digite un numero de pago")
	fmt.Println("\t1:PayPal\t2:Cash\t3:CreditCard")
	_, err := fmt.Scanln(&met)
	if err != nil {
		panic("debe digitar un metodo valido")
	}
	if met > 3 {
		panic("debe digitar un metodo valido")
	}
	metodo := FactoryPay(met)
	metodo.Pay()
}
