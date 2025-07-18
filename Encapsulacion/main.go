package main

// en go no se tiene las palabras reservadas publicas o privadas, sin embargo se tienen los identificadores exportado o no exportado que senalan que metodos
// si se podran exportar del paquete y que son de uso interno del paquete

// en go no se tienen los constructures ya que para inicializar objetos basta con llamar la estructura y declarar las variables
// hay algunos casos en los que se quiere realizar validaciones, o se requiera un constructuctor, se puede declarar una funcion New tipo exportada
// la cual servira como parametro inicial y retorne una estructura inicializada

//En go no se tienen  los metodos setters and getters, ya que no son necesarios debido a que al cambiar el identificador de exportado/noexportado se pueden asignar y obtener valores
//en caso de que se quieran utilizar setter and getters, es necesario poner la propiedad como no extortada y crear un setter que pertenesca a la clase, y que reciva como argumento el mismo tipo de dato de la variable a editar
//posterior a eso para cambiar el valor, es necesario que el setter sea exportado
// En el caso del Get no es necesario poner el prefijo Get basta con nombrar al metodo de la misma manera que la propiedad pero con mayuscula
import (
	"course/course"
	"fmt"
)

func main() {
	Go := course.New("Go desde cero", 0, false)
	Go.SetUserID([]uint{12, 56, 89})
	Go.SetClasses(map[uint]string{
		1: "introduccion",
		2: "estructuras",
		3: "maps",
	})

	Go.SetName("OOP con GO")
	fmt.Println(Go.Name())
	Go.PrintClasses()
	fmt.Printf("%+v", Go)

}
