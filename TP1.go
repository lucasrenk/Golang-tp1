package tp1

import (
	"strconv"
)

// Producto contiene metodos que nos permiten acceder
// a atributos (ID, Precios) de un Producto.
type Producto interface {
	ID() int
	Precio() int
}

// Productos es una lista donde para cada producto
// se sabe el nombre del supermercado, el id y su precio.
// Esta estructura se puede cargar usando la funcion LeerProductos
// que carga informacion guardada en `productos.json`.
type Productos [][]string

// Carrito contiene el nombre del supermercado y el precio final luego
// de sumar todos los productos.
type Carrito struct {
	Tienda string
	Precio int
}

type Alfajor struct {
	Id    int
	Super string
	Monto int
}

func (a Alfajor) ID() int {
	return a.Id
}

func (a Alfajor) Precio() int {
	return a.Monto
}

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	return nil
	var precios []Carrito
	preciosPorSuper := make(map[string]int)
	for _, id := range ids {
		for _, product := range p {
			idProduct, err := strconv.Atoi(product[1])
			if err == nil {
				if id == idProduct {
					precio, err := strconv.Atoi(product[2])
					if err == nil {
						preciosPorSuper[product[0]] += precio
					}

				}
			}

		}
	}

	for super, precio := range preciosPorSuper {
		result := Carrito{Tienda: super, Precio: precio}
		precios = append(precios, result)
	}

	return precios
}

// Promedio recibe el ID del producto y retorna el precio promedio
// de ese producto tomando los precios de cada uno de los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	return 0
	suma := 0.0
	count := 0.0

	for _, producto := range p {
		idProduct, err := strconv.Atoi(producto[1])
		if err == nil && idProduct == idProducto {
			precio, err := strconv.Atoi(producto[2])
			if err == nil {
				count++
				suma += float64(precio)
			}
		}
	}
	if count == 0.0 {
		return 0
	}

	return suma / count
}

// BuscarMasBarato recibe ID de producto a buscar y devuelve la informacion 
// del producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	return nil, false
	producto := Alfajor{Id: idProducto}
	encontrado := false

	for _, prod := range p {
		idProduct, err := strconv.Atoi(prod[1])
		if err == nil && idProduct == idProducto {
			precio, err := strconv.Atoi(prod[2])
			if err == nil && (producto.Monto == 0 || precio < producto.Monto) {
				producto = Alfajor{Id: idProduct, Monto: precio}
				encontrado = true
			}
		}
	}
	return producto, encontrado
	return 0
	var sum_precio int
	var cant int
	for _, prod := range p {
		p_id, _ := strconv.Atoi(prod[1])
		if idProducto == p_id {
			precio, _ := strconv.Atoi(prod[2])
			sum_precio += precio
			cant++
		}
	}
	var precio_prom float64
	if cant > 0 {
		precio_prom = float64(sum_precio) / float64(cant)
	}
	return precio_prom
}

func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	return nil, false
	var min_price int
	var existe bool
	var tienda string
	for _, prods := range p {
		p_id, _ := strconv.Atoi(prods[1])
		if idProducto == p_id {
			existe = true
			precio, _ := strconv.Atoi(prods[2])
			if precio < min_price || min_price == 0 {
				min_price = precio
				tienda = prods[0]
			}
		}
	}
	pr := producto{tienda, idProducto, min_price}
	return pr, hallado
}
