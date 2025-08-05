package course

import "fmt"

//un identificador exportado sera aquel en el que a nivel de paquete empiecen con letra mayuscula
type course struct {
	name    string
	price   float64
	isFree  bool
	userID  []uint
	classes map[uint]string
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
		name:   name,
		price:  price,
		isFree: isFree,
	}

}

//metodo no exportado
/*func (c *course) changePrice(price float64) { // se usa el operador de puntero para senalar que se esta apuntando hacia una direccion y no solo se esta haciendo uso del dato
	c.price = price
}*/

//metodo exportado
func (c *course) PrintClasses() { // el metodo tiene un argumento receptor llamado c de estructura course, el metodo pertenece a la estructura
	text := "las clases son: "
	for _, class := range c.classes {
		text += class + ", "

	}
	fmt.Println(text[:len(text)-2])
}

//Setters and Getters
func (c *course) SetName(name string) { c.name = name }
func (c *course) Name() string {
	return c.name
}

func (c *course) SetPrice(price float64) { c.price = price }
func (c *course) Price() float64 {
	return c.price
}

func (c *course) SetIsFree(isFree bool) { c.isFree = isFree }
func (c *course) IsFree() bool {
	return c.isFree
}
func (c *course) SetUserID(userID []uint) { c.userID = userID }
func (c *course) UserID() []uint {
	return c.userID
}
func (c *course) SetClasses(classes map[uint]string) {
	c.classes = classes
}
