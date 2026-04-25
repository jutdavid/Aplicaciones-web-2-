package main

import (
	"errors"
	"fmt"
	"semana03-taller-relaciones/internal/cafeteria"
)

func main() {
	var repo cafeteria.Repository = cafeteria.NewRepoMemoria()
	cliente := cafeteria.Cliente{ID: 1, Nombre: "Ana"}
	producto := cafeteria.Producto{ID: 1, Nombre: "Cafe", Precio: 1.50}

	repo.GuardarCliente(cliente)
	repo.GuardarProducto(producto)

	t1, _ := repo.ObtenerCliente(1)
	fmt.Println("Cliente:", t1.Nombre)

	_, err := repo.ObtenerCliente(99)
	if err != nil {
		if errors.Is(err, cafeteria.ErrClienteNoEncontrado) {
			fmt.Println("El cliente no existe.")
		} else {
			fmt.Println("Error inesperado:", err)
		}
	}
}