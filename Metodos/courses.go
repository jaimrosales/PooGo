package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserID  []uint
	Classes map[uint]string
}

func (c Course) Classesimpresas() { // el metodo tiene un argumento receptor llamado c de estructura course, el metodo pertenece a la estructura
	text := "las clases son: "
	for _, class := range c.Classes {
		text += class + ", "

	}
	fmt.Println(text[:len(text)-2])
}

func (c *Course) ChangePrice(price float64) { // se usa el operador de puntero para senalar que se esta apuntando hacia una direccion y no solo se esta haciendo uso del dato
	c.Price = price
}

//1.5
