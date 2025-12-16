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
  
  cantConjuntos := maht.Pow(2 , len(nums))
  result := make([][]int, cantConjuntos)
  index := 0
  for i := 0 ; i < len(nums) ; i++ {
    for j := 0; j < len(nums) ; j++ {
    result[j][] = nums[i]
    index ++
    }
  }
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
