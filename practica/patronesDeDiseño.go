package practica

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

type EnchufeEuro struct {}

func (e EnchufeEuro) ConectarEuro() {}

type Adapter struct {
  Eur EnchufeEuro
}

func ( a *Adapter) ConectarArg() {
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

struct Adapter {
  ext *ExternalLogger
}

func (e *Adapter) LogInfo( info string ) {
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















