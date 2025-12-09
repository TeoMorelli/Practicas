package practica

//el codigo se realiza pensando en un Heap de maximos.

type NodoHeap struct {
	datos []int
}

func (h *NodoHeap) Top() int {
	return h.datos[0]
}

func (h *NodoHeap) Size() int {
	return len(h.datos)
}

func (h *NodoHeap) Insert(val int) {
	pos := len(h.datos) - 1
	h.datos[pos] = val
	h.upHeap(pos)
}

func (h *NodoHeap) upHeap(i int) {
	for i > 0 {
		padre := (i - 1) / 2
		if h.datos[i] >= h.datos[padre] {
			return
		}
		h.datos[i], h.datos[padre] = h.datos[padre], h.datos[i]
		i = padre
	}
}

//La funcion delete debe eliminar la raiz. Colocarl el primer elemento y con el downHeap acomodarlo.

func (h *NodoHeap) Delete() int {
	ElementoBorrado := h.datos[0]
	ultimo := len(h.datos) - 1
	h.datos[0], h.datos[ultimo] = h.datos[ultimo], h.datos[0]
	h.datos = h.datos[:ultimo]
	if len(h.datos) > 0 {
		h.downHeap(0)
	}
	return ElementoBorrado
}

// DownHeap, al estar primero, comparo con sus hijos. Si es menor se acomoda
func (h *NodoHeap) downHeap(pos int) {
	hijoDer := 2*pos + 1
	hijoIzq := 2*pos + 2
	for pos < len(h.datos) {
		if hijoIzq <= len(h.datos)-1 {
			if h.datos[pos] < h.datos[hijoIzq] {
				h.datos[pos], h.datos[hijoIzq] = h.datos[hijoIzq], h.datos[pos]
			}
		}
		if hijoDer <= len(h.datos)-1 {
			if h.datos[pos] < h.datos[hijoDer] {
				h.datos[pos], h.datos[hijoDer] = h.datos[hijoDer], h.datos[pos]
			}
		}
	}
}
