package practica

import "fmt"

/*
Los patrones de diseño son basicamente formas de reutilizacion de codigo, estos ofrecen una respuesta pra problema recurrentes
*/

/*Patron adapter, se utiliza para reutilizar codigo de una interfaz que no coincida con lo que espera el sistema*/

/*
“Adaptar un enchufe europeo a uno argentino”

Tenés dos interfaces incompatibles:

type EnchufeEuropeo interface {
    ConectarEuro() string
}
type EnchufeArgentino interface {
    ConectarArg() string
}

Objetivo:
Crear un Adapter que permita usar un EnchufeEuropeo como si fuera EnchufeArgentino
*/

type EnchufeEuropeo interface {
	ConectarEuro() string
}

type EnchufeArgentino interface {
	ConectarArg() string
}

type EnchufeEuro struct{}

func (e EnchufeEuro) ConectarEuro() {}

type Adapter struct {
	Eur EnchufeEuro
}

func (a *Adapter) ConectarArg() {
	a.Eur.ConectarEuro()
}

/*
Ejercicio 5 — Nivel Medio
“Adaptar un logger externo a tu sistema”

Supongamos que tu proyecto usa:

type Logger interface {
    LogInfo(msg string)
}


Pero la librería externa usa:

type ExternalLogger struct { ... }
func (e *ExternalLogger) WriteLog(level int, text string)


Objetivo:
Crear un Adapter que permita usar ExternalLogger como un Logger.

Restricción:
No podés modificar el código de ExternalLogger.
*/

type Logger interface {
	LogInfo(msg string)
}

type ExternalLogger struct{}

func (e *ExternalLogger) WriteLog(level int, text string)

type Adapter2 struct {
	ext *ExternalLogger
}

func (e *Adapter2) LogInfo(info string) {
	e.ext.WriteLog(0, info)
}

/*
Ejercicio 6 — Nivel Difícil
“Adaptar un procesador de pagos internacional”

Tu sistema tiene:

type PagoLocal interface {
    Pagar(monto float64) error
}

Pero la API externa funciona con:

type ForeignPayment interface {
    MakeTransaction(cents int) bool
}

Objetivo:
Crear un Adapter que:

convierta pesos → centavos

convierta el resultado booleano en un error Go

oculte la API externa completamente

Restricción difícil:
El Adapter debe permitir inyectar distintas APIs externas (usa composición, no herencia).
*/

//Pensar. Debo adaptar la funcion local para que funcione con la externa. debe permitir inyectar muchas apis.

type PagoLocal interface {
	Pagar(monto float64) error
}

type ForeignPayment interface {
	MakeTransaction(cents int) bool
}

type Adapter3 struct {
	fo ForeignPayment
}

func newAdapter3(api ForeignPayment) *Adapter3 {
	return &Adapter3{fo: api}
}

func (a *Adapter3) Pagar(monto float64) error {
	centavos := int(monto * 100)
	ok := a.fo.MakeTransaction(centavos)
	if ok {
		return nil
	}
	return fmt.Errorf("fallo la transaccion de %d centavos (%.2f pesos)", centavos, monto)
}

/*1. COMPOSITE

El patron composite se utiliza para tratar muchos elementos como uno solo, evidentemente sin distinciones particulares en estos.

Componente: define una interfaz comun

Simple: Representa los elementos individuales de la estructura

Compuesto: Elemento que contiene mas elementos, Simples como compuestos (una matriz). Se debe proveer un metodo que pueda agregar nuevos elementos


Ejercicio 1 — Nivel Básico (Fácil)
“Sistema de archivos simplificado”

Implementar un Composite que represente:

Archivo (hoja)

Carpeta (nodo compuesto)

Cada uno debe implementar la interfaz:

type Component interface {
    Nombre() string
    Tamaño() int
}

Objetivo:
Calcular el tamaño total de una carpeta (la suma recursiva de todos sus elementos).

Restricción:
No usar slices globales; solo composición recursiva.
*/

type NodoP interface {
	Nombre() string
	Tamaño() int
}

//tengo archivo y carpeta que guarda archivo.

type Archivo struct {
	nombre string
	tamaño int
}

func (a *Archivo) Nombre() string {
	return a.nombre
}

func (a *Archivo) Tamaño() int {
	return a.tamaño
}

type Carpeta struct {
	nombre string
	hijos  []NodoP
}

func (c *Carpeta) Nombre() string {
	return c.nombre
}

func (c *Carpeta) Tamaño() int {
	total := 0
	for _, t := range c.hijos {
		total += t.Tamaño()
	}
	return total
}

func (c *Carpeta) Add(n *NodoP) {
	c.hijos = append(c.hijos, *n)
}

/*Ejercicio 2 — Nivel Medio
“Estructura de menú de restaurante”

Modelar un Composite con:

Plato (leaf)

Sección (composite)

Cada Sección debe poder:

listar todos los platos que contiene (recursivamente)

obtener el precio más caro dentro de toda la sección

Objetivo:
Implementar un método:

PrecioMaximo() float64


que funcione recursivamente*/

type NodoR interface {
	Nombre() string
	Valor() float64
}

type Plato struct {
	nombre string
	valor float64
}

func (p *Plato) Nombre() string {
	return p.nombre
}

func (p *Plato) Valor() float64 {
	return p.valor
}

type Seccion struct {
	nombres string
	hijos []NodoR
}

func (s *Seccion) Nombre() string {
	return s.nombre
}

func (s *Seccion) Valor() float64 {
	total := 0
	for _, t := s.hijos {
		total += t.Valor()
	}
}

func (s *Seccion) listarPlatos() string[] {
	name := make(s.hijos, " ", len(s.hijos) - 1)
	for i := 0; i < len(s.hijos); i++ {
		name[i] = s.hijos[i].Nombre() 
	}
	return name
}

func (s *Seccion) precioMaximo() float64{
	max := 0
	for _, m := s.hijos {
		switch v := s.hijo(type) {
			case *Plato:
				if v.Valor() > max {
				max = v.valor()	
			}
			case *Carpeta:
				precioHijo := v.precioMaximo()
			if precioMaximo > max {
				max = precioHijo
			}
		}
	}
	return max
}


/*
“Evaluador de expresiones matemáticas”

Representar expresiones como un Composite:

Nodo hoja: números

Nodo compuesto: operadores (+, -, *, /)

Ejemplo de árbol:

      (*)
     /   \
   (+)    5
  /   \
 2     3


Objetivo:
Implementar:

Eval() float64


que evalúe toda la expresión de forma recursiva.

Restricciones:

Prohibido convertir a string o usar go/ast.

Debe ser recursive traverse.
*/















