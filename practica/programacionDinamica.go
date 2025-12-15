package practica

/*
La programacion dinamica surge para contrarrestar los calculos inecesarios echos por algoritmos recursivos. Guaradando valores para no repetir calculos y encontrar la solucuion optima.
*/

func Fibonacci(n int) int {
	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

/*
🟢 Ejercicio 2 — Fácil
Suma mínima para llegar al final

Tenés un arreglo de números positivos.

Desde la posición i podés:

Avanzar 1 paso

Avanzar 2 pasos

Objetivo:
Encontrar la suma mínima para llegar al final.

Ejemplo:

arr := []int{1, 100, 1, 1, 1}
Resultado: 3


Explicación:
Camino óptimo: 1 → 1 → 1

Pista:

dp[i] = arr[i] + min(dp[i-1], dp[i-2])


👉 Idea nueva:
Elegir mínimo entre opciones.
*/

//dp[i] cantidad de formas de llegar al scalon i

//No comprendi consigna. !

/*
Ejercicio 3 — Normal
Cantidad de formas de subir escaleras

Tenés una escalera con n escalones.

Podés subir:

1 escalón

2 escalones

Pregunta:
¿Cuántas formas distintas hay de llegar arriba?

Ejemplo:

n = 4
Resultado: 5


Explicación:

(1+1+1+1)
(1+1+2)
(1+2+1)
(2+1+1)
(2+2)


Pista:

dp[i] = dp[i-1] + dp[i-2]


👉 Idea nueva:
DP de conteo, no de mínimo.
*/

func cantidadDeFormas(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

/*
Ejercicio 4 — Normal+
Máxima suma sin tomar elementos adyacentes

Dado un arreglo de enteros positivos.

Regla:
No podés tomar dos elementos consecutivos.

Objetivo:
Obtener la máxima suma posible.

Ejemplo:

arr := []int{2, 7, 9, 3, 1}
Resultado: 12


Explicación:
Tomar: 2 + 9 + 1 = 12

Pista:

dp[i] = max(
  dp[i-1],
  dp[i-2] + arr[i]
)


👉 Idea nueva:
Elegir entre tomar o no tomar.
*/

//Solo puedo tomar elemento separados, no uno al lado del otro.

func maximaSuma(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[len(nums)-1]
}

/*
Ejercicio 2 — Normal
Caminos en una grilla

Tenés una grilla de tamaño m x n.

Solo podés moverte:

Derecha

Abajo

Pregunta:
¿Cuántos caminos distintos existen desde (0,0) hasta (m-1,n-1)?

Ejemplo:

m = 3, n = 3
Resultado: 6


Restricciones:

No usar backtracking

Usar programación dinámica

Pista:

dp[i][j] = dp[i-1][j] + dp[i][j-1]
*/

func caminosDP(mat [][]int, n, m int) int {
	dp := make([][]int, m)
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return mat[m-1][n-1]
}

/*
🟠 Ejercicio 3 — Difícil
Mochila 0/1 (Knapsack)

Tenés:

n objetos

cada objeto tiene:

peso w[i]

valor v[i]

una mochila con capacidad máxima W

Objetivo:
Maximizar el valor total sin exceder W.

Restricción:

Cada objeto se puede tomar una sola vez

Ejemplo:

pesos  = [2, 3, 4]
valores = [4, 5, 6]
W = 5
Resultado: 9


Pista:

dp[i][w] = max(
  dp[i-1][w],
  dp[i-1][w - peso[i]] + valor[i]
)
*/

/*
🔴 Ejercicio 4 — Muy Difícil
Subsequencia común más larga (LCS)

Dadas dos cadenas A y B, encontrar la longitud de la subsecuencia común más larga.

Ejemplo:

A = "AGGTAB"
B = "GXTXAYB"
Resultado: 4   // "GTAB"


Restricciones:

No usar fuerza bruta

Complejidad O(n × m)

Pista:

dp[i][j]:
si A[i-1] == B[j-1]:
    dp[i][j] = dp[i-1][j-1] + 1
sino:
    dp[i][j] = max(dp[i-1][j], dp[i][j-1])
*/
