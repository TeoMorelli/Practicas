package practica

type Nodo struct {
	valor  int
	left   *Nodo
	right  *Nodo
	altura int
}

type BST struct {
	Root *Nodo
}

func New(valor int) *Nodo {
	return &Nodo{
		left:  nil,
		right: nil,
		valor: valor,
	}
}

func (bst *BST) insert(valor int) {
	bst.insertRec(bst.Root, valor)
}

func (bst *BST) insertRec(nodo *Nodo, valor int) *Nodo {
	if bst.Root == nil {
		bst.Root = &Nodo{valor: valor, left: nil, right: nil}
		return bst.Root
	}
	if valor <= nodo.valor {
		nodo.left = bst.insertRec(nodo.left, valor)
	}
	if valor >= nodo.valor {
		nodo.right = bst.insertRec(nodo.right, valor)
	}
	return nodo
}

func (bst *BST) eliminar(valor int) {
	bst.Root = bst.eliminarRec(bst.Root, valor)
}

func (bst *BST) eliminarRec(nodo *Nodo, valor int) *Nodo {
	if nodo == nil {
		return nil
	}
	if valor > nodo.valor {
		nodo.right = bst.eliminarRec(nodo.right, valor)
	} else if valor < nodo.valor {
		nodo.left = bst.eliminarRec(nodo.left, valor)
	} else {
		if nodo.left == nil && nodo.right == nil {
			return nil
		}
		if nodo.left == nil {
			return nodo.left
		}
		if nodo.right == nil {
			return nodo.right
		}
		cambio := minNode(nodo.right)
		nodo.valor = cambio.valor
		nodo.right = bst.eliminarRec(nodo.right, cambio.valor)
	}
	return nodo
}

func minNode(n *Nodo) *Nodo {
	for n.left != nil {
		n = n.left
	}
	return n
}

// Contar todos los nodos del árbol

func (bst *BST) CountNodes(n *Nodo) int {
	if n == nil {
		return 0
	}
	return 1 + bst.CountNodes(n.left) + bst.CountNodes(n.right)
}

// Contar cuántas hojas tiene el árbol

func (bst *BST) CountLeaves(n *Nodo) int {
	if n == nil {
		return 0
	}
	if n.left == nil && n.right == nil {
		return 1
	}
	return bst.CountLeaves(n.left) + bst.CountLeaves(n.right)
}

// Calcular la altura del árbol

func (bst *BST) Height(n *Nodo) int {
	if n == nil {
		return 0
	}
	leftHeigth := bst.Height(n.left)
	rightHeigth := bst.Height(n.right)

	if leftHeigth > rightHeigth {
		return leftHeigth + 1
	}
	return rightHeigth + 1

}

/*4️⃣ Calcular el promedio de los valores del árbol */

func (bst *BST) promedio(n *Nodo) int {
	if n == nil {
		return 0
	}
	return bst.suma(n) / bst.CountNodes(n)
}

func (bst *BST) suma(n *Nodo) int {
	if n == nil {
		return 0
	}
	suma := n.valor
	return suma + bst.suma(n.left) + bst.suma(n.right)
}

/*5️⃣ Contar cuántos nodos tienen valor mayor al promedio*/

func (bst *BST) nodosSuperanPromedio(n *Nodo) int {
	if n == nil {
		return 0
	}
	contador := 0
	if n.valor > int(bst.promedio(n)) {
		contador++
	}
	return contador + bst.nodosSuperanPromedio(n.left) + bst.nodosSuperanPromedio(n.right)
}

/*6️⃣ Contar cuántos nodos tienen exactamente 1 hijo*/

func (bst *BST) nodosConUnHijo(n *Nodo) int {
	if n == nil {
		return 0
	}
	contador := 0
	if tieneUnHijo(n) {
		contador++
	}
	return contador + bst.nodosConUnHijo(n.left) + bst.nodosConUnHijo(n.right)
}

func tieneUnHijo(n *Nodo) bool {
	if n == nil {
		return false
	}
	if n.left == nil && n.right != nil {
		return true
	}
	if n.left != nil && n.right == nil {
		return true
	}
	return false
}
