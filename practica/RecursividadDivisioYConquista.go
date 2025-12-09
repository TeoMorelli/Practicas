package practica

/*Contar dígitos de un número

Implementá una función recursiva que reciba un número entero positivo y devuelva cuántos dígitos tiene.
*/

func digitos( num int ) int {
  if num < 10 {
    return 1
  }
  return 1 + digitos(num / 10)
}

/*
Verificar si un arreglo es palíndromo

Dado un arreglo de enteros, escribir una función recursiva que determine si es palíndromo 
(se lee igual de izquierda a derecha que de derecha a izquierda).
*/

func esPalindromo( arr[] int ) bool {
  
}




