package main //Ejecutable

import "fmt" // Print

func main() { // Ejecutable
	fmt.Println("Hello world")
	//var x int //Default: 0
	//var cadena string //Default: ""
	//var booleano bool //Default: False
	//var arreglo []string
	var x int
	if x > 0 {
		x = 23
	}

	nombre := "Coco" //No hace falta ponerle que tipo de datos es
	if nombre != "Jesus" {
		fmt.Println("Cococ")
		nombre = "Jesus"
	}
}


// go run hello_world.go