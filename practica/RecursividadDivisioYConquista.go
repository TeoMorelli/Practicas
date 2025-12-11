package practica

/*Contar dígitos de un número

Implementá una función recursiva que reciba un número entero positivo y devuelva cuántos dígitos tiene.
*/

func digitos(num int) int {
	if num < 10 {
		return 1
	}
	return 1 + digitos(num/10)
}

/*
Verificar si un arreglo es palíndromo

Dado un arreglo de enteros, escribir una función recursiva que determine si es palíndromo
(se lee igual de izquierda a derecha que de derecha a izquierda).
*/

func esPalindromo(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}
	if arr[0] == arr[len(arr)-1] {
		return true && esPalindromo(arr[1:len(arr)-1])
	}
	return false
}

/*
Cantidad de caminos en una grilla

Dada una grilla de tamaño n x m, contar recursivamente cuántas formas existen de ir desde la esquina superior izquierda hasta la esquina inferior derecha, moviéndose solo:
 - hacia la derecha
 - hacia abajo
*/

func caminos(mat [][]int) int {
	return caminosHelper(mat, 0, 0)
}

func caminosHelper(mat [][]int, i int, j int) int {
	//Primero condiciones: 1 - si se va del arreglo y 2 - Si llega al final
	if i >= len(mat) || j >= len(mat[0]) {
		return 0
	}
	if i == len(mat)-1 && j == len(mat[0])-1 {
		return 1
	}
	return caminosHelper(mat, i+1, j) + caminosHelper(mat, i, j+1)
}

/*Búsqueda binaria recursiva

Implementar una búsqueda binaria de forma recursiva que devuelva la posición de un elemento en un arreglo ordenado.
*/

func busquedaBinaria(arr []int, x int, inicio int, fin int) int {
	if inicio > fin {
		return -1
	}
	medio := len(arr) / 2

	if arr[medio] == x {
		return medio
	}
	if arr[medio] > x {
		return busquedaBinaria(arr, x, inicio, medio-1)
	} else {
		return busquedaBinaria(arr, x, medio+1, fin)
	}
}
