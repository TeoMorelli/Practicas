package practica

//Backtraking es una tecnica de que se utiliza cunado el codigo no tiene una forma viable de implementarse, como aquellos algoritmo

/*
🟢 Ejercicio 1 — Fácil
Generar todas las combinaciones de un conjunto

Dado un conjunto de números únicos, por ejemplo:

[1, 2, 3]

🎯 Objetivo:
Generar todas las combinaciones posibles (subconjuntos).

💡 Pistas:

En cada posición decidís:

incluir el elemento

no incluirlo

El orden no importa

📌 Esto enseña:

árbol binario de decisiones

backtracking puro
*/

//Subconjuntos de un conjunto, La cantidad posible es de 2 elevado a la totalidad de elementos.

func SubConjuntos(nums []int) [][]int {
  var result [][]int
  var actual []int

  var back func(i int)
  back = func (i int) {
      //Si i llegas a recorrer todo. Lo guardamos en el resultado
      if i == len(nums)  {
        copia := make([]int, len(actual)
        copy(copia, actual)
        result = append(resultado, copia)
        return
      }
      //aumentamos el indice recursivamente
      back(i+1)
      //agregar elementos.
      actual = append(actual, nums[i])
      back(i+1)
      actual = actual[:len(actual)-1]
  }
  back(0)
  return result
}

/*
🟡 Ejercicio 2 — Normal
Permutaciones de una lista

Dada una lista de números distintos:

[1, 2, 3]

🎯 Objetivo:
Generar todas las permutaciones posibles.

💡 Pistas:

El orden sí importa

Cada nivel del árbol fija una posición

No podés reutilizar un número ya usado

📌 Esto enseña:

control de usados

profundidad exacta
*/


                      
