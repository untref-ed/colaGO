package cola

import (
	"fmt"
)
// Cola genérica, soporta cualquier tipo de dato
// Se implementa con un slice
type Cola[T any] struct {
	datos []T
}

// Crea una nueva cola vacía O(1)
func NuevaCola[T any]() *Cola[T] {
	c := new(Cola[T])
	c.datos = make([]T, 0)
	return c
}

// Encola un dato. O(1)
func (c *Cola[T]) Encolar(dato T) {
	c.datos = append(c.datos, dato)
}

// Desencola un dato. Si la cola está vacía devuelve un error. O(1)
func (c *Cola[T]) Desencolar() (T, error) {
	var nulo T
	if len(c.datos) == 0 {
		return nulo, fmt.Errorf("Cola vacía")
	}
	dato := c.datos[0]
	c.datos = c.datos[1:]
	return dato, nil
}

// Devuelve el dato del frente de la cola. Si la cola está vacía devuelve un error. O(1)
func (c *Cola[T]) Frente() (T, error) {
	var nulo T
	if len(c.datos) == 0 {
		return nulo, fmt.Errorf("Cola vacía")
	}
	return c.datos[0], nil
}

// Devuelve true si la cola está vacía o false en caso contrario. O(1)
func (c *Cola[T]) EstaVacia() bool {
	return len(c.datos) == 0
}
