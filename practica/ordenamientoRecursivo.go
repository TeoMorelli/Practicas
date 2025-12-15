package practica

//Los oredenamientos recursivos son 3:

//MergeSort
//Se utliza para ordenar arreglos de forma recursiva y se basa en divide y conquistaras.
// Mergesort es un algoritmo Estable, que requiere espacio adicional proporcional al tamaño del arreglo, es decir no es In Place, basado en el paradigma “divide y vencerás”. Fue creado en 1945 por John von Neumann y es uno de los algoritmos de ordenamiento más eficientes y ampliamente utilizados.
// La complejidad o es O(n log n) ya que subdivide recursivamente el arreglo en mitades hasta llegar al caso base y luego combina los resultados de manera lineal.

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	medio := len(arr) / 2
	izquierda := MergeSort(arr[medio:])
	derecha := MergeSort(arr[:medio])
	return merge(izquierda, derecha)
}

func merge(izq []int, der []int) []int {
	var result []int
	i := 0
	j := 0
	for i < len(izq) && j < len(der) {
		if izq[i] <= der[j] {
			result = append(result, izq[i])
			i++
		} else {
			result = append(result, der[j])
			j++
		}
	}
	for i < len(izq) {
		result = append(result, izq[i])
		i++
	}
	for j < len(der) {
		result = append(result, der[j])
		j++
	}
	return result
}

//QuickSort
//Quicksort, como Mergesort, es un algoritmo de ordenamiento basado en el paradigma “divide y vencerás”. Sin embargo, su enfoque es diferente y su eficiencia en la práctica lo ha convertido en uno de los algoritmos de ordenamiento más utilizados. Quicksort no es un algoritmo de ordenamiento estable ni Online pero si es In-Place y estrictamente hablando tampoco se puede asegurar que sea de división y conquista

func QuickSort(arr []int, inicio int, fin int) {
	if inicio >= fin {
		return
	}
	mitad := Particionar(arr, inicio, fin)
	QuickSort(arr, inicio, mitad-1)
	QuickSort(arr, mitad+1, fin)
}

func Particionar(arr []int, inicio int, fin int) int {
	medio := (inicio + fin) / 2
	pivot_valor := arr[medio]
	Intercambiar(arr, medio, fin)

	i := inicio
	j := fin - 1

	for i <= j {
		for i <= j && arr[i] <= pivot_valor {
			i++
		}
		for i <= j && arr[j] > pivot_valor {
			j--
		}
		if i <= j {
			Intercambiar(arr, i, j)
			i++
			j--
		}
	}
	Intercambiar(arr, i, fin)
	return i
}

func Intercambiar(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// HeapSort
// Es un algoritmo de ordenamiento in-place y no estable
// Basado en la estructura de datos Heap (montículo)
// Complejidad: O(n log n) en todos los casos

func HeapSort(arr []int) {
	n := len(arr)

	// Construir el max-heap (reorganizar el arreglo)
	// Comenzamos desde el último nodo no hoja
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extraer elementos del heap uno por uno
	for i := n - 1; i > 0; i-- {
		// Mover la raíz (máximo) al final
		Intercambiar2(arr, 0, i)

		// Llamar heapify en el heap reducido
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n int, i int) {
	// Inicializar el más grande como raíz
	mayor := i
	izquierda := 2*i + 1
	derecha := 2*i + 2

	// Si el hijo izquierdo es mayor que la raíz
	if izquierda < n && arr[izquierda] > arr[mayor] {
		mayor = izquierda
	}

	// Si el hijo derecho es mayor que el mayor hasta ahora
	if derecha < n && arr[derecha] > arr[mayor] {
		mayor = derecha
	}

	// Si el mayor no es la raíz
	if mayor != i {
		Intercambiar2(arr, i, mayor)

		// Recursivamente heapify el subárbol afectado
		heapify(arr, n, mayor)
	}
}

func Intercambiar2(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
