package main

import (
	"fmt"
)

func modifyValue(ptr *int) {
	*ptr = 100
}

type Person struct {
	Name string
	Age  int
}

func modifyPerson(ptr *Person) {
	ptr.Age = 23
	ptr.Name = "Jack Parra"
}

func main() {
	var x int = 10
	var ptr *int = &x // el ptr apunta el valor de x
	fmt.Println(x)    // output: 10
	fmt.Println(ptr)  // output: 0xc0000140a8
	fmt.Println(*ptr) // output: 10

	fmt.Println("\n\tGESTION DE MEMORIA CON PUNTEROS")
	fmt.Println(ptr)  // output: 0xc0000160c0
	fmt.Println(*ptr) // output: 0
	*ptr = 10
	fmt.Println(*ptr) // output: 10
	ptr = nil

	fmt.Println("\n\tARGUMENTOS DE FUNCIONES")
	var y int = 42
	fmt.Println("Before modification:", y)
	modifyValue(&y) //para apuntar al puntero
	fmt.Println("After modification:", y)

	fmt.Println("\n\tPUNTEROS Y ESTRUCTURAS DE DATOS")
	p := Person{Name: "Javier", Age: 46}
	fmt.Println("Before modification:", p)
	modifyPerson(&p)
	fmt.Println("After modification:", p)
}
