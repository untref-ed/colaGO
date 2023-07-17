package cola

import (
	"fmt"
	"testing"
)

func TestCola(t *testing.T) {
	c := NuevaCola[int]()
	c.Encolar(1)
	c.Encolar(2)
	c.Encolar(3)
	for !c.EstaVacia() {
		dato, _ := c.Desencolar()
		fmt.Println(dato)
	}
	// Output:
	// 1
	// 2
	// 3
}

func TestColaVacia(t *testing.T) {
	c := NuevaCola[int]()
	_, err := c.Desencolar()
	fmt.Println(err)
	// Output:
	// Cola vacía
}
