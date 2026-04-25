package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Cliente struct {
	ID      int
	Nombre  string
	Carrera string
	Saldo   float64
}

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}
type Pedido struct {
	ID         int
	ClienteID  int
	ProductoID int
	Cantidad   int
	Total      float64
	Fecha      string
}

func leerLinea(lector *bufio.Reader) string {
	linea, _ := lector.ReadString('\n')
	return strings.TrimSpace(linea)
}
func leerEntero(lector *bufio.Reader, prompt string) int {
	fmt.Print(prompt)
	texto := leerLinea(lector)
	n, err := strconv.Atoi(texto)
	if err != nil {
		return -1
	}
	return n
}
func leerFloat(lector *bufio.Reader, prompt string) float64 {
	fmt.Print(prompt)
	texto := leerLinea(lector)
	f, err := strconv.ParseFloat(texto, 64)
	if err != nil {
		return -1
	}
	return f
}
func mostrarMenu() {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  CAFETERIA ULEAM                         ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  1. Listar clientes                      ║")
	fmt.Println("║  2. Listar productos                     ║")
	fmt.Println("║  3. Buscar cliente                       ║")
	fmt.Println("║  4. Buscar producto 	                   ║")
	fmt.Println("║  5. Agregar cliente                      ║")
	fmt.Println("║  6. Agregar producto                     ║")
	fmt.Println("║  7. Eliminar cliente                     ║")
	fmt.Println("║  8. Eliminar producto                    ║")
	fmt.Println("║  0. Salir                                ║")
	fmt.Println("╚══════════════════════════════════════════╝")
}
func AgregarCliente(clientes []Cliente, nuevo Cliente) []Cliente {
	return append(clientes, nuevo)
}
func BuscarClientePorID(clientes []Cliente, id int) int {
	for i, c := range clientes {
		if c.ID == id {
			return i
		}
	}
	return -1
}
func ListarClientes(clientes []Cliente) {
	fmt.Println("\n=== CLIENTES REGISTRADOS ===")
	if len(clientes) == 0 {
		fmt.Println("(no hay clientes registrados)")
		return
	}
	for _, n := range clientes {
		fmt.Printf("  [%d] %s | %s | Saldo: $%.2f\n",
			n.ID, n.Nombre, n.Carrera, n.Saldo)
	}
}
func EliminarCliente(clientes []Cliente, id int) []Cliente {
	idx := BuscarClientePorID(clientes, id)
	if idx == -1 {
		fmt.Printf("⚠ Cliente con ID %d no existe.\n", id)
		return clientes
	}
	fmt.Printf("⚠ Cliente con ID %d eliminado.\n", id)
	return append(clientes[:idx], clientes[idx+1:]...)
}

func AgregarProducto(productos []Producto, nuevo Producto) []Producto {
	return append(productos, nuevo)
}
func BuscarProductoPorID(productos []Producto, id int) int {
	for i, p := range productos {
		if p.ID == id {
			return i
		}
	}
	return -1
}
func ListarProductos(productos []Producto) {
	fmt.Println("\n=== PRODUCTOS REGISTRADOS ===")
	if len(productos) == 0 {
		fmt.Println("(no hay productos registrados)")
		return
	}
	for _, p := range productos {
		fmt.Printf("  [%d] %s | $%.2f | Stock: %d | %s\n",
			p.ID, p.Nombre, p.Precio, p.Stock, p.Categoria)
	}
}
func EliminarProducto(productos []Producto, id int) []Producto {
	idx := BuscarProductoPorID(productos, id)
	if idx == -1 {
		fmt.Printf("⚠ Producto con ID %d no existe.\n", id)
		return productos
	}
	fmt.Printf("⚠ Producto con ID %d eliminado.\n", id)
	return append(productos[:idx], productos[idx+1:]...)
}
func DescontarSaldo(cliente *Cliente, saldo float64) error {
	if saldo <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}
	if cliente.Saldo < saldo {
		return fmt.Errorf("saldo insuficiente en %s (hay $%.2f, solicita $%.2f)",
			cliente.Nombre, cliente.Saldo, saldo)
	}
	cliente.Saldo -= saldo
	return nil
}
func DescontarStock(producto *Producto, stock int) error {
	if stock <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}
	if producto.Stock < stock {
		return fmt.Errorf("stock insuficiente en %s (hay %d, solicita %d)",
			producto.Nombre, producto.Stock, stock)
	}
	producto.Stock -= stock
	return nil
}
func RegistrarPedido(
	clientes []Cliente,
	productos []Producto,
	pedidos []Pedido,
	clienteID int,
	productoID int,
	cantidad int,
	fecha string,
) ([]Pedido, error) {
	PosiCliente := BuscarClientePorID(clientes, clienteID)
	if PosiCliente == -1 {
		return pedidos, errors.New("cliente no encontrado")
	}
	PosiProducto := BuscarProductoPorID(productos, productoID)
	if PosiProducto == -1 {
		return pedidos, errors.New("producto no encontrado")
	}
	if productos[PosiProducto].Stock < cantidad {
		return pedidos, fmt.Errorf("stock insuficiente en %s (hay %d, solicita %d)",
			productos[PosiProducto].Nombre, productos[PosiProducto].Stock, cantidad)
	}
	if cantidad < 1 || cantidad > 5 {
		return pedidos, errors.New("la cantidad debe estar entre 1 y 5")
	}
	err := DescontarStock(&productos[PosiProducto], cantidad)
	if err != nil {
		return pedidos, err
	}
	costo := float64(cantidad) * productos[PosiProducto].Precio
	err = DescontarSaldo(&clientes[PosiCliente], costo)
	if err != nil {
		// Si el saldo falla, devolvemos el stock que habíamos descontado
		productos[PosiProducto].Stock += cantidad
		return pedidos, err
	}
	nuevoID := len(pedidos) + 1
	nueva := Pedido{
		Id:         nuevoID,
		ClienteID:  clienteID,
		ProductoID: productoID,
		Cantidad:   cantidad,
		Total:      float64(cantidad) * productos[PosiProducto].Precio,
		Fecha:      fecha,
	}
	pedidos = append(pedidos, nueva)
	return pedidos, nil
}
func main() {
	clientes := []Cliente{
		{ID: 1, Nombre: "Juan Pérez", Carrera: "Ingeniería", Saldo: 150.00},
		{ID: 2, Nombre: "María Gómez", Carrera: "Derecho", Saldo: 200.00},
		{ID: 3, Nombre: "Carlos López", Carrera: "Medicina", Saldo: 100.00},
	}
	productos := []Producto{
		{ID: 1, Nombre: "Coca-Cola", Precio: 1.50, Stock: 50, Categoria: "Bebidas"},
		{ID: 2, Nombre: "Papas Fritas", Precio: 2.00, Stock: 30, Categoria: "Snacks"},
		{ID: 3, Nombre: "Chocolate", Precio: 1.00, Stock: 20, Categoria: "Dulces"},
		{ID: 4, Nombre: "Agua", Precio: 1.00, Stock: 100, Categoria: "Bebidas"},
	}

	lector := bufio.NewReader(os.Stdin)
	for {
		mostrarMenu()
		opcion := leerEntero(lector, "Elige una opción: ")

		switch opcion {
		case 1:
			ListarClientes(clientes)

		case 2:
			ListarProductos(productos)

		case 3:
			idx := BuscarClientePorID(clientes, 3)
			if idx == -1 {
				fmt.Println("Cliente no existe")
				return
			}
			fmt.Println("Encontrado:", clientes[idx].Nombre)

		case 4:
			idx := BuscarProductoPorID(productos, 3)
			if idx == -1 {
				fmt.Println("Producto no existe")
				return
			}
			fmt.Println("Encontrado:", productos[idx].Nombre)
		case 5:
			fmt.Println("\n--- Agregar Cliente ---")
			ID := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			Nombre := leerLinea(lector)
			fmt.Print("Carrera: ")
			Carrera := leerLinea(lector)
			fmt.Print("Saldo: ")
			Saldo := leerFloat(lector, "Saldo: ")
			clientes = AgregarCliente(clientes, Cliente{
				ID, Nombre, Carrera, Saldo,
			})
			fmt.Println("✓ Cliente agregado.")
		case 6:
			fmt.Println("\n--- Agregar Producto ---")
			ID := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			Nombre := leerLinea(lector)
			fmt.Print("Precio: ")
			Precio := leerFloat(lector, "Precio: ")
			fmt.Print("Stock: ")
			Stock := leerEntero(lector, "Stock: ")
			fmt.Print("Categoría: ")
			Categoria := leerLinea(lector)
			productos = AgregarProducto(productos, Producto{
				ID, Nombre, Precio, Stock, Categoria,
			})
			fmt.Println("✓ Producto agregado.")
		case 7:
			clientes = EliminarCliente(clientes, 3)
		case 8:
			productos = EliminarProducto(productos, 3)
		case 9:
			fmt.Println("\n--- Registrar Pedido ---")
			tid := leerEntero(lector, "ID del Cliente: ")
			nid := leerEntero(lector, "ID del Producto: ")
			cantidad := leerEntero(lector, "Cantidad: ")
			fecha := time.Now()

			cal := leerEntero(lector, "Calificación (1-5): ")
			fmt.Print("Comentario: ")
			comentario := leerLinea(lector)

			experiencias, err = RegistrarExperiencia(negocios, turistas,
				experiencias, tid, nid, idioma, cal, comentario)
			if err != nil {
				fmt.Println("⚠ No se pudo registrar:", err)
			} else {
				fmt.Println("✓ Experiencia registrada con éxito.")
			}
		case 0:
			fmt.Println("¡Hasta luego!")
			return

		default:
			fmt.Println("Opción no válida.")
		}
	}
}
