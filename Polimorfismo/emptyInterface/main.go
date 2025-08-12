package main

// una interfas nula no tiene valor, ni un tipo concreto
//esto ocurre cuando existe la interface pero no existe ningun metodo de los de la interface
// en el ejemplo X(), si se intenta utilizar X() al no existir ningun tipo que utilice X(), marcara un error

//Empty interface o interfas vacia
//Es aquella interface que no tiene ningun metodo para el caso empty, Usalmente es util para manejar valores de tipos desconocidos
// es decir cuando no sabemos el tipo existente
//Cualquier tipo existente dentro del paquete implementara la empty interface

// Las assertions o aserciones permiten validar el tipo concreto que tiene la interfas vacia,
// para ingresar las inserciones se escribe el objeto de interface i seguido de un punto i. seguido del tipo que se quiere validar entre parentesis
//i.(string) o i.(uint), las inserciones retornan el valor de la interface ademas de un bool valor, validacion := i.(uint)

// En caso de que tengamos que validar distintos tipos de tipos, ya que las assertions solo nos permiten validar un tipo a la vas para no tener millar de if ok
// se puede usar un tipe switch,
//La forma en que se usa es con un siwtch normal y en variable se pondra una assertion pero utilizando type de modo switch v := i.(type){} y se evalua los casos que se busquen validar

import (
	"fmt"
	"strings"
)

/*type exampler interface {
	X()
}*/

func wrapper(i interface{}) {
	fmt.Printf("valor: %v, Tipe:%T \n", i, i)
	//ejemplo de insercion
	v, ok := i.(string) //Se genera una insercion para validar si tipo string
	if ok {
		fmt.Printf("\t %s\n", strings.ToUpper(v))
	}

	//Ejemplo de type siwtch

	switch v := i.(type) {
	case string:
		fmt.Printf("\t %s\n", strings.ToUpper(v))
	case int:
		fmt.Println(v * 2)
	default:
		fmt.Printf("valor: %v, Tipe:%T \n", i, i)

	}

}

func main() {
	//var e exampler
	//e.X()

	wrapper(12)
	wrapper(1.2)
	wrapper("ricardo")
	wrapper(false)
	wrapper("jaime")

}
