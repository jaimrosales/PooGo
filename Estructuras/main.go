package main

import "fmt"

//Las estructuras son colecciones de campos con las caracteristicas del elemento a abstraer

//Type <Name> Struct{campos}
type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserID  []uint
	Classes map[uint]string
}

//Para crear las instancias se asigna el objeto de instancia a una variable
func main() {
	Go := Course{ //Se instancia la estructura
		Name:   "Go desde Cero",
		Price:  12.34,
		IsFree: false,
		UserID: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "introduccion",
			2: "estructuras",
			3: "maps",
		},
	}
	//Otra forma de declarar las estructuras es sin necesidad de especificar a que caracteristica se hace referencia pero es necesario ingrasar
	//todas las caracteristicas y en el mismo orden, en caso de que te saltes una o no ingreses una, o el tipo cambie marcara error

	Python := Course{ //Se instancia la estructura
		"Python desde Cero",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "introduccion",
			2: "estructuras",
			3: "maps",
		},
	}

	//Tambien se puede agregar caracteristicas especificas sin necesidad de agregar todas, se les agregara el valor cero a las caracteristicas no instanciadas
	css := Course{
		Name:   "css desde cero",
		IsFree: true,
	}

	//Se puede instanciar la estructura sin declarar nada y posteriormente agregarlo con el operador "."
	Js := Course{}
	Js.Name = "Curso JS"
	Js.UserID = []uint{12, 67}

	fmt.Println(Go.Name) //Se usa la estructura
	fmt.Println(Python.Name)
	fmt.Printf("%+v\n", css)
	fmt.Printf("%+v", Js)
}
