package name

import "fmt"

// GetName obtiene y retorna el nombre de la persona
// mayusulas como publicas - Minusculas como privadas
func GetName() string {
	var name string
	name = "sin nombre"

	fmt.Println("Ingresa tu nombre")
	fmt.Scanf("%s", &name)

	return name
}

// GetUnicode retorna un string
func GetUnicode() string {
	return "もしもし!"
}
