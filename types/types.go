package main

import (
	"fmt"
	"time"
)

// Crear tipo de dato
type Animal struct {
	name string
	edad int
	raza string
}

func main() {
	var ok bool = true
	var suma int
	var division float32 = 7.0 / 3.0
	var name string = "Cristhian"
	var lastName string = "Apolo"
	lastName = "Apolo Cevallos"
	// Infiere el tipo de dato
	pais := "Ecuador"
	// Error
	// pais = 2
	const year = 2022

	var dog Animal = Animal{
		name: "Firulais",
		edad: 4,
		raza: "Shitzu",
	}

	var cat Animal = createAnimal("Misifu ", 6, "Egipcio")

	fmt.Println(sumaFunc())

	fmt.Println("Name " + name + " " + lastName + ". De: " + pais)
	time.Sleep(time.Second * 3)
	fmt.Println("5+5 =", suma)
	fmt.Println("division", division)
	fmt.Println("year", year)
	fmt.Println("Perro", dog)
	fmt.Println("Gato", cat.name)

	fmt.Println(ok && false)
	fmt.Println(ok || false)
	fmt.Println(!ok)
}

func createAnimal(n string, e int, r string) Animal {
	var animal = Animal{
		n,
		e,
		r,
	}
	return animal
}

func sumaFunc() (n1 int, n2 int) {
	n1 = 5
	n2 = 5
	return
}
