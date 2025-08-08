package main

import "fmt"

type storager interface {
	Get() string
	Set(string)
}

type Person struct {
	name string
}

func NewPerson(name string) Person {
	return Person{name}
}

func (p Person) Get() string {
	return p.name
}

func (p *Person) Set(name string) { //Se usa un puntero para cambiar directamente el valor de p y no solo durante el metodo
	p.name = name
}

func exec(s storager, name string) { // de esta manera puede un metodo puede usar una interface de interface para poder trabajar con distintos tipos de interfaces
	s.Set(name)
	fmt.Println(s.Get())
}

// Lo de atras funciona sin embargo no es lo recomendable por go por tener una mescla de metodos que trabajan y no con punteros acontinuacion otro ejemplo utilizando solo punteros
// este tipo de practica mejora o facilita la implementacion de interfaces

type Human struct {
	denomination string
}

func NewHuman(denomination string) *Human {
	return &Human{denomination}
}

func (h *Human) Get() string {
	return h.denomination
}

func (h *Human) Set(denomination string) {
	h.denomination = denomination
}

func execHuman(s storager, denomination string) {
	s.Set(denomination)
	fmt.Print(s.Get())
}

func main() {
	p := NewPerson("Jaime")
	p.Set("Ricardo")
	fmt.Println(p.Get())
	exec(storager(&p), "jaime") //en este caso yo use casting pero no es necesario, el casting para indicar que p es de tipo storage, y directamente detecta que p tiene a set como puntero por lo que automaticamente anade el operador de puntero
	o := NewHuman("frederic")
	execHuman(o, "emiliano") // en este caso no es necesario senalar que es un puntero debido a que todos los metodos de Human (set and get) estan utilizando punteros y no una mescla de normales y punteros
}
