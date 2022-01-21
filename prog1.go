package main

import (
	"fmt"
)

type NombreCompleto struct {
	nombre   string
	apellido string
}

func (e NombreCompleto) String() string {
	return e.nombre + " " + e.apellido
}
func main() {
	fmt.Println("Deseas averiguar cual es tu nombre completo? Solo responde las siguientes preguntas:")
	fmt.Print("Cual es tu nombre? ")
	var data NombreCompleto = NombreCompleto{}
	fmt.Scanln(&data.nombre)
	fmt.Print("Y cual es tu apellido? ")
	fmt.Scanln(&data.apellido)
	fmt.Println("Bien, tu nombre completo es: " + data.String())
}
