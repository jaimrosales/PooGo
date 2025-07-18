package course

import "fmt"

//un identificador exportado sera aquel en el que a nivel de paquete empiecen con letra mayuscula
type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserID  []uint
	Classes map[uint]string
}

//funcion constructora
// esta retornara una estructura de puntero a curso para indicar que se apunta hacia un dato de memoria y no un dato de trabajo, y de igual manera se verifica
//que se podran usar los demas metodos de la estructura
//El nombre New es mas que suficiente para no ser redundante con NewCourse
func New(name string, price float64, isFree bool) *course {
	if price <= 0 {
		price = 30
	}

	return &course{
		Name:   name,
		Price:  price,
		IsFree: isFree,
	}

}

//metodo no exportado
func (c *course) changePrice(price float64) { // se usa el operador de puntero para senalar que se esta apuntando hacia una direccion y no solo se esta haciendo uso del dato
	c.Price = price
}

//metodo exportado
func (c *course) PrintClasses() { // el metodo tiene un argumento receptor llamado c de estructura course, el metodo pertenece a la estructura
	text := "las clases son: "
	for _, class := range c.Classes {
		text += class + ", "

	}
	fmt.Println(text[:len(text)-2])
}

//Minute 6
