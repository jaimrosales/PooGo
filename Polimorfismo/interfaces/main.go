package main

//Es un conjunto de firma de metodos

import "fmt"

//para crear una interface se hace con la palabra type seguida del nombre del metodo a usar dentro de la interface+er y despues la palabra interface{}
// dentro de la interface se especifica el conjunto de junta de metodos

//cualquier tipo valido que implemente el metodo greet estara satisfaciendo la interface, sin necesidad de la palabra implement como en otros lenguajes
// Un tipo puede satisfacer multiples interfaces

type Greeter interface { // la interface hara referencia al metodo de greet
	Greet()
}
type Byer interface {
	Bye()
}

//la estructura implementa el metodo
type Person struct {
	Name string
}

func (p Person) Greet() { //EL metodo Greet pertenece a persona e implementa la interface Greeter
	fmt.Printf("Hola soy %s", p.Name)
}
func (p Person) Bye() { //Segunda interface usada por el tipo Person
	fmt.Printf("Adios soy %s", p.Name)
}

type Text string //Se crea un tipo definido, es decir tiene base string pero permite anadir los metodos

func (t Text) Greet() { //Se crea el metodo Greet del tipo definido texto
	fmt.Printf("Hola soy %s", t)
}
func (t Text) Bye() { //Segunda interface usada por el tipo TExt
	fmt.Printf("adios soy %s", t)
}

// Los valores de las interfaces se pueden considerar como una tupla de un valor y un tipo concreto
func GreetAll(gs ...Greeter) {
	for _, g := range gs {
		g.Greet()
		fmt.Printf("\t Valor: %v, Tipo: %T\n", g, g)
	}
}
func ByeAll(bs ...Byer) { //Metodo para imprimir la interface byer
	for _, b := range bs {
		b.Bye()
		fmt.Printf("\t Valor: %v, Tipo: %T\n", b, b)
	}
}

//Dado a que ByeAll y GreetAll hacen codigo repetido para evitar esto se puede hacer un solo metodo que haga ambos
// Para solucionar esto se pueden hace un tipo embebido de interfaces, pero la interfaces deben de ser disjuntas
type ByeGreeter interface {
	Greeter
	Byer
}

func InterfaceAll(ia ...ByeGreeter) {
	for _, ex := range ia {
		ex.Bye()
		ex.Greet()
		fmt.Printf("\t Valor: %v, Tipo: %T\n", ex, ex)
	}
}

func main() {
	var g Greeter = Person{Name: "Jaime "}
	var h Greeter = Text("hannia\n") //Devido a que Text es un tipo definido y no una estructura para senalar el string a usar es necesario hacer un casting como en el ejemplo
	p := Person{Name: "Ricardo "}
	var t Text = "Emiliano"
	g.Greet()
	h.Greet()
	GreetAll(p, t) //imprimira los valores y el tipo de cada uno de sus valores
	ByeAll(p, t)   //imprimira los valores y el tipo de cada uno de sus valores
	InterfaceAll(p, t)
}
