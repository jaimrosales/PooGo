package main

//Los tipos teclarados existen de dos formas las declaraciones de alias y las definiciones de tipo

//Las declaraciones de alias: permiten crear identificadores a alias que hacen referencia a un tipo existente.
//Creacion :
//	type	myalias = tipoa hacer referencia

//LAs definiciones de tipo permiten definir un nuevo tipo basado en un tipo existente, solo extraera las propiedades del tipo base pero no hereda los metodos
//Creacion:
// type nuevoTipo	tipo base

import "fmt"

type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

//Declaracion de alias
type myAlias = course

//Definicion de tipo
type newCourse course

func main() {
	c := course{name: "Go"} //Tipoo normal
	c.Print()
	fmt.Printf("El tipo es: %T\n", c)

	d := myAlias{name: "Ve"} //Utilizando declaracion de tipo, d es de tipo course aunque dice myAlias
	d.Print()
	fmt.Printf("El tipo es: %T\n", d)

	e := newCourse{name: "Go"} //Se usa el tipo creado a partir de course
	// e.Print()                  No existe el metodo de print
	fmt.Printf("El tipo es: %T\n", e)

}
