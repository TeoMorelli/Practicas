package practica

import "sort"

//Un algoritmo Avido o gredy, se basa en buscar la mejor solucion local para un problema. Evidentemente esto solo funciona cuando se cumplen las propiedades de eleccion avida y subestructura optima.

/*
¿Cómo se piensa un algoritmo ávido?

Siempre este razonamiento:

¿Qué significa “mejor” en este problema?

¿Puedo ordenar los datos según ese criterio?

¿Si tomo siempre el mejor actual, nunca me arrepiento?

¿Puedo demostrar (o al menos justificar) que no hay contradicción?
*/

/*
Ejercicio 1 — Fácil
Selección de actividades

Tenés una lista de actividades, cada una con:

Hora de inicio

Hora de fin

Solo podés realizar una actividad a la vez.

🎯 Objetivo:

Seleccionar el máximo número de actividades sin superposición.

💡 Pistas:

📌 Este es el greedy más clásico de todos.
*/

type actividad struct {
	inicio int
	fin    int
}

type ListaActividades struct {
	actividades []*actividad
}

func (l *ListaActividades) SeleccionarActividades() []actividad {
	sort.Slice(l.actividades, func(i, j int) bool {
		return l.actividades[i].fin < l.actividades[j].fin
	})

	var resultado []actividad

	if len(l.actividades) == 0 {
		return resultado
	}

	resultado = append(resultado, *l.actividades[0])

	ultima := l.actividades[0]
	for i := 1; i < len(l.actividades); i++ {
		actual := l.actividades[i]
		if actual.inicio >= ultima.fin {
			resultado = append(resultado, *actual)
			ultima = actual
		}
	}
	return resultado
}

/*
Ejercicio 2 — Normal
Cambio de monedas

Tenés un monto N y un conjunto de monedas con valores fijos
(ej: 1, 5, 10, 25).

🎯 Objetivo:

Dar el cambio usando la menor cantidad de monedas posible.

Restricción:

Podés usar monedas ilimitadas

No todas las combinaciones funcionan con greedy (ojo 👀)

💡 Pistas:

¿Qué moneda elegís primero?

¿Por qué podría fallar con ciertos sistemas de monedas?
*/

var cambio = []int{1, 2, 5, 10, 25, 50, 100}

func mejorCambio(monto int) []int {
	var result []int
	i := len(cambio) - 1
	for monto > 0 {
		if monto >= cambio[i] {
			result = append(result, cambio[i])
			monto -= cambio[i]
		} else {
			i--
		}
	}
	return result
}

/*
Ejercicio A — Máximo número de tareas dentro de un tiempo total

Tenés tareas con:

duración

NO tienen ganancia

Tenés un tiempo total T

🎯 Objetivo:
Hacer la mayor cantidad de tareas posibles.

💡 Idea greedy:

Hacé primero las tareas más cortas

📌 Esto sí es greedy simple y funciona.
*/

type tareaTiempo struct {
	nombre string
	tiempo int
}

type ListaTareasTiempo struct {
	tareas []*tareaTiempo
}

func (l *ListaTareasTiempo) mayorCantidadDeTareas(tiempoMax int) []*tareaTiempo {
	sort.Slice(l.tareas, func(i, j int) bool {
		return l.tareas[i].tiempo < l.tareas[j].tiempo
	})

	var result []*tareaTiempo

	result = append(result, l.tareas[0])
	total := l.tareas[0].tiempo
	for i := 1; i < len(l.tareas); i++ {
		if total+l.tareas[i].tiempo <= tiempoMax {
			result = append(result, l.tareas[i])
			total += l.tareas[i].tiempo
		}
	}
	return result
}

/*
🟠 Ejercicio 3 — Difícil
Planificación de tareas con deadlines

Cada tarea tiene:

Un tiempo de ejecución (1 unidad)

Un deadline

Una ganancia

Solo podés hacer una tarea por unidad de tiempo.

🎯 Objetivo:

Maximizar la ganancia total, respetando los deadlines.

💡 Pistas:

No todas las tareas entran

¿Conviene elegir primero por deadline? ¿por ganancia?

¿Qué pasa si una tarea “cara” ocupa un lugar clave?

📌 Acá el greedy no es obvio.
*/

/*type tarea3 struct {
	ganancia int
	deadline int
}

type ListaTareas3 struct {
	tareas []*tarea3
}

func (l *ListaTareas3) masGanancia() []int {
	sort.Slice(l.tareas, func(i, j int) bool {
		return l.tareas[i].deadline > l.tareas[j].deadline
	})
	MaxDeadline := l.tareas[0].deadline
	for
}*/

/*
🔴 Ejercicio 4 — Muy Difícil
Cobertura mínima de intervalos

Te dan:

Un intervalo objetivo [L, R]

Un conjunto de sub-intervalos [li, ri]

🎯 Objetivo:

Cubrir completamente [L, R] usando la menor cantidad de intervalos.

Reglas:

Los intervalos pueden solaparse

No podés cortar intervalos

💡 Pistas:

¿Cuál intervalo conviene elegir primero?

¿Elegir el más largo?

¿El que empieza antes?

¿El que termina más lejos?

📌 Este ejercicio rompe cabezas y es muy de entrevistas fuertes.
*/

// muy dificil
