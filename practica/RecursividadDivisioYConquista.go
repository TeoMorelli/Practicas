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
}

/*Búsqueda binaria recursiva

Implementar una búsqueda binaria de forma recursiva que devuelva la posición de un elemento en un arreglo ordenado.
*/

func busquedaBinaria(arr []int, x int) int {
	medio := len(arr) / 2
	if arr[medio] == x {
		return medio
	}
	if medio < x {
		busquedaBinaria(arr[:medio], x)
	}
	if medio > x {
		busquedaBinaria(arr[medio:], x)
	}
	return medio
}
