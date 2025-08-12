package main

// una interfas nula no tiene valor, ni un tipo concreto
//esto ocurre cuando existe la interface pero no existe ningun metodo de los de la interface
// en el ejemplo X(), si se intenta utilizar X() al no existir ningun tipo que utilice X(), marcara un error

//Empty interface o interfas vacia
//Es aquella interface que no tiene ningun metodo para el caso empty, Usalmente es util para manejar valores de tipos desconocidos
// es decir cuando no sabemos el tipo existente
//Cualquier tipo existente dentro del paquete implementara la empty interface
import "fmt"

type exampler interface {
	X()
}

func wrapper(i interface{}) {
	fmt.Printf("valor: %v, Tipe:%T", i, i)

}

func main() {
	//var e exampler
	//e.X()

	wrapper(12)
	wrapper(1.2)
	wrapper(false)
	wrapper("fieiha")
}
