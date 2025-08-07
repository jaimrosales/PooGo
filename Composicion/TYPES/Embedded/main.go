package main

//embeber tipos dentro de structuras
//es la capacidad de que un tipo tenga otros tipos dentro de sus campos, funcionan como herencia sin ser herencia, permitiendo crar algo similar a subclases
//pero esto permite que un tipo pueda tener diferentes tipos embebidos

//En Go no existe sobre escritura de metodos lo que existe es la ocultacion de metodos, La forma en que funciona es que go verifica primero si el metodo o campo existe
//primeramente dentro de la estructura principal, y posterior buscara en la estructura embebida nivel por nivel,
// es decir buscara por capas, pasando y revisando a la siguiente capa si existe dicho metodo, en caso de existir ignorara si existe algun metodo llamado igual
//en alguna structura embebida, para ingresar a los metodos ocultos tendra que usarse la forma de . para ingresar, variable.claseEmbebida.metodo

//Cuando se tienen varias structura embebidas dentro de otra, y dos o mas comparten el nombre de algun metodo, no se podra hacer el ocultamiento,
//De hecho se generara un conflicto indicando que la referencia es ambigua, para solucionar el conflicto o no se hace referencia(basicamente hacer como que no existe, casi como eliminarlo del codiggo)
//O bien indicar precisamente a que estructura pertenece: variable.claseEmbebida1.metodo o variable.claseEmbebida2.metodo

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

type Human struct {
	Age      uint
	Children uint
}

func NewHuman(age, children uint) Human {
	return Human{age, children}
}

type Employee struct { // se embebe la estructura persona dentro de la estructura empleado
	Person // al poner el nombre de la structura, se embebe la estructura persona dentro de empleado con todos los campos y metodos, sin embargo el receptr sera el tipo interno cuando se manden a llamar los metodos
	Human  //segunda estructura embebida al empleado
	Salary float64
}

// Constructor que usa el constructor de Newperson, de esta manera el empleado tiene totalmente embebido la estructura persona, como tipo herencia pero mas sencillo compuesto
func NewEmploye(name string, age uint, chindren uint, salary float64) Employee {
	return Employee{NewPerson(name, age), NewHuman(age, chindren), salary}
}

func (e Employee) Payroll() {
	fmt.Println(e.Salary * 0.90)
}
func (p Employee) Greet() { //metodo de clase superior, no es sobre escritura, es ocultamiento
	fmt.Println("Hello desde empleado")
}

//
func main() {
	// se crea el objeto de nuevo empleado
	e := NewEmploye("jaime", 30, 2, 1000000)

	//e aunque es de tipo persona puede hacer uso de los campos y metodos de la clase persona adicional a los metodos de la clase empleado
	//Se podria crear una clase tipo estudiante por ejemplo con diferente metodo la cual tenga embebida la clase persona pero no la clase empleado, tendria nombre y edad pero no salario
	//esto abre paso a las subclases
	fmt.Println(e.Name, e.Human.Age, e.Person.Age) // para ingresar al campo age al existir en dos estructuras distintas, y no sobre la misma escalera de capas, es necesario especificar la estructura como si fuera un campo y senalar la edad, caso contrario marcara antiguedad, cabe resaltar que go marca accesible desde el principio pero al escribirse marcara error
	e.Greet()                                      //llama el metodo de la estructura empleado
	e.Person.Greet()                               // metodo greet de persona, si no existiera el greet de empleado entraria directamente al greet de persona al no existir uno en la estructura de empleado
	e.Payroll()
}
