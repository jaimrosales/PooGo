package main

// en go no se tiene las palabras reservadas publicas o privadas, sin embargo se tienen los identificadores exportado o no exportado que senalan que metodos
// si se podran exportar del paquete y que son de uso interno del paquete

// en go no se tienen los constructures ya que para inicializar objetos basta con llamar la estructura y declarar las variables
// hay algunos casos en los que se quiere realizar validaciones, o se requiera un constructuctor, se puede declarar una funcion New tipo exportada
// la cual servira como parametro inicial y retorne una estructura inicializada

import "course/course"

func main() {
	Go := course.New("Go desde cero", 0, false)
	Go.UserID = []uint{12, 56, 89}
	Go.Classes = map[uint]string{
		1: "introduccion",
		2: "estructuras",
		3: "maps",
	}

	Go.PrintClasses()

}
