package main

//embeber tipos dentro de structuras
import "fmt"

type Person struct {
	Name string
	Age  uint
}

func NewPerson(name string, age uint) Person {
	return Person{name, age}
}
func (p Person) Greet() {
	fmt.Println("Hello")
}

type Employee struct { // se embebe la estructura persona dentro de la estructura empleado
	Person // al poner el nombre de la structura, se embebe la estructura persona dentro de empleado con todos los campos y metodos, sin embargo el receptr sera el tipo interno cuando se manden a llamar los metodos
	Salary float64
}

// Constructor que usa el constructor de Newperson, de esta manera el empleado tiene totalmente embebido la estructura persona, como tipo herencia pero mas sencillo compuesto
func NewEmploye(name string, age uint, salary float64) Employee {
	return Employee{NewPerson(name, age), salary}
}

func (e Employee) Payroll() {
	fmt.Println(e.Salary * 0.90)
}

//
func main() {
	// se crea el objeto de nuevo empleado
	e := NewEmploye("jaime", 30, 1000000)

	//e aunque es de tipo persona puede hacer uso de los campos y metodos de la clase persona adicional a los metodos de la clase empleado
	//Se podria crear una clase tipo estudiante por ejemplo con diferente metodo la cual tenga embebida la clase persona pero no la clase empleado, tendria nombre y edad pero no salario
	//esto abre paso a las subclases
	fmt.Println(e.Name, e.Age)
	e.Greet()
	e.Payroll()
}
