package main

//1.3
// Los metodos en Go son funciones con un argumento receptor, el argumento receptor se especifica entre parentesis entre func y el nombre de la funcion
// en el se especifica el tipo de dato y el nombre del argumento receptor,
//1.4
//Los metodos en Go se declaran por fuera de la estructura,esto para poder tener metodos para otros tipos de datos distintos a estructuras, ademas permite la
//creacion de metodos en distinto archivo siempre y cuando pertenescan al mismo paquete

import "fmt"

/* La estructura como tal tambien se puede mover a otro archivo mientras esten dentro del mismo paquete se compartiran
type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserID  []uint
	Classes map[uint]string
}*/

//funcion normal
func PrintClasses(c Course) {
	text := "las clases son: "
	for _, class := range c.Classes {
		text += class + ", "

	}
	fmt.Println(text[:len(text)-2])
}

//metodo
func (c Course) ImprimirClase() { // el metodo tiene un argumento receptor llamado c de estructura course, el metodo pertenece a la estructura
	text := "las clases son: "
	for _, class := range c.Classes {
		text += class + ", "

	}
	fmt.Println(text[:len(text)-2])
}

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
	Ir := &Course{ //Se instancia la Ir tipo puntero de curso
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

	//PrintClasses(Go)
	Go.ImprimirClase()
	Go.Classesimpresas()
	(&Go).ChangePrice(67.12) //se usa el operador de direccion para poder enviar como tal la direccion de la estructura y no solo enviar una copia de los datos
	(*Ir).ImprimirClase()

	//Si se quiere trabajar con la estructura se utiliza un receptor de tipo puntero, yu si se quiere nada mas con trabajar con los datos sin cambiarlos
	//se utilizan receptores normales
	//Como tip cuando se trbajan con interfaces es de que si algun metodo es de tipo puntero configurar todos los metodos de la estructura como tipo puntero
	//aunque no se necesite actualizar un valor, esto debido a que las interfaces diferencian entre si un metodo es puntero o no

	fmt.Println(Go.Price)

}
