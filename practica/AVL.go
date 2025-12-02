package practica

/*Metodo Insert*/

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
