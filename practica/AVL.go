package practica

/*  Los Árboles de Búsqueda Binaria Balanceados o Árboles AVL (Adelson-Velsky y Landis) surgen para corregir una deficiencia crítica de los
Árboles de Búsqueda Binaria (BST) estándar.En el peor caso (un árbol degenerado o sesgado), la complejidad computacional para recorrer un BST 
asciende a O(n), lo que equivale a recorrer una lista enlazada. 
El objetivo de un AVL es asegurar que el tiempo de búsqueda se mantenga siempre en $O(\log n)$, lo que ofrece un 
rendimiento mucho más rápido y predecible.

¿Qué Significa que un AVL Esté Balanceado?
Un árbol AVL mantiene su eficiencia mediante un estricto criterio de balanceo:
Factor de Balanceo: Para cada nodo del árbol, la diferencia absoluta entre la altura de su subárbol izquierdo (Hizq) y la altura 
de su subárbol derecho (Hder) nunca debe ser mayor a 1.  Factor de Balanceo = |Hizq - Hder| <= 1 ,l Si este criterio se viola después de una 
inserción o eliminación, el árbol se considera desbalanceado y debe ser reestructurado.

Rotaciones para el BalanceoPara restaurar el balanceo tras una operación, los árboles AVL utilizan un sistema de rotaciones que reorganizan los nodos.
Esto asegura que la propiedad AVL se mantenga y la complejidad de las operaciones siga siendo O(log n).
Las rotaciones se clasifican en cuatro tipos, dependiendo de la configuración del desbalance:

- Rotación Simple a la Derecha (LL): Ocurre cuando se inserta un nodo en el subárbol izquierdo del subárbol izquierdo del nodo desbalanceado.

- Rotación Simple a la Izquierda (RR): Ocurre cuando se inserta un nodo en el subárbol derecho del subárbol derecho del nodo desbalanceado.

- Rotación Doble Izquierda-Derecha (LR): Ocurre cuando se inserta un nodo en el subárbol derecho del subárbol izquierdo del nodo desbalanceado. 
Esto requiere primero una rotación a la izquierda y luego una a la derecha.

- Rotación Doble Derecha-Izquierda (RL): Ocurre cuando se inserta un nodo en el subárbol izquierdo del subárbol derecho del nodo desbalanceado. 
Esto requiere primero una rotación a la derecha y luego una a la izquierda.
*/

/*Metodo Insert
Este metodo debe cumplir las propidades del AVL ademas de las rotaciones.
*/

func (bst *BST) insertAVL( int valor ) {
	bst.insertAVLrec(bst.root , valor)
}

func (bst *BST) insertAVLrec(n *Nodo, int val) *Nodo {
	
}


/*Metodo Eliminar*/

/*Metodo Search*/

/*7️⃣ Verificar si el árbol es completo*/

func (bst *BST) estaCompleto(n *Nodo) bool {
	if bst.Root == nil {
		return true
	}

	cola := []*Nodo{bst.Root}
	hijoVacio := false

	for len(cola) > 0 {
		actual := cola[0]
		cola = cola[1:]

		if actual.left != nil {
			if hijoVacio {
				return false
			}
			cola = append(cola, actual.left)
		} else {
			hijoVacio = true
		}

		if actual.right != nil {
			if hijoVacio {
				return false
			}
			cola = append(cola, actual.right)
		} else {
			hijoVacio = true
		}
	}
	return true
}
