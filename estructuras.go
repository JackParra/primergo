// package main

// import "fmt"

// // Definición básica de un struct
// type Usuario struct {
// 	ID     int
// 	Nombre string
// 	Email  string
// 	Activo bool
// }

// // Struct con tags (útil para JSON, validaciones, etc.)
// type Producto struct {
// 	ID     int     `json:"id" db:"product_id"`
// 	Nombre string  `json:"nombre" validate:"required"`
// 	Precio float64 `json:"precio" validate:"min=0"`
// }

// // Struct anidado
// type Empresa struct {
// 	Nombre    string
// 	Direccion Direccion
// 	Empleados []Usuario
// }

// type Direccion struct {
// 	Calle  string
// 	Ciudad string
// 	CP     string
// }

// func estructuras() {
// 	// Inicialización con valores cero
// 	var u1 Usuario
// 	fmt.Println(u1)

// 	// Inicialización con valores específicos
// 	u2 := Usuario{
// 		ID:     1,
// 		Nombre: "Jack",
// 		Email:  "jackcorreo@email.com",
// 		Activo: true,
// 	}
// 	fmt.Println(u2)

// 	// Inicialización parcial (otros campos toman valor cero)
// 	u3 := Usuario{
// 		Nombre: "stalin",
// 		Email:  "otrocorreo@email.com",
// 	}
// 	fmt.Println(u3)

// 	// Usando punteros (más eficiente para structs grandes)
// 	u4 := &Usuario{
// 		ID:     2,
// 		Nombre: "USS",
// 	}
// 	fmt.Println(*u4) // Para imprimir el valor del puntero, se debe desreferenciar

// 	var person struct {
// 		Name string
// 		Age  int
// 		Pet  string
// 	}

// 	person.Name = "bob"
// 	person.Age = 50
// 	person.Pet = "dog"

// 	fmt.Println("Person:", person)

// 	// Otra estructura anónima para mascota
// 	pet := struct {
// 		Name string
// 		Kind string
// 	}{
// 		Name: "Fido",
// 		Kind: "dog",
// 	}

// 	fmt.Println("Pet:", pet)
// }
