package ejercicios

// Sean as y bs dos listas de enteros de tamaño n.
// Escribir una función que reciba como parámetros
// as, bs y un entero x y decida si x se puede
// escribir como suma de un elemento de as más un
// elemento de bs.
func SumaElementos(as, bs []int, x int) bool {
	if len(as) == 0 || len(bs) == 0 {
		return false
	}
	if as[0]+bs[0] == x {
		return true
	}
	return SumaElementos(as[1:], bs, x) || SumaElementos(as, bs[1:], x)
}

/*func sumaElementos(as, bs []int, x int, fin1, fin2 int) bool {
	if fin2 < 0 {
		return sumaElementos(as, bs, x, fin1-1, len(bs)-1)
	}
	if fin1 < 0 {
		return false
	}
	if as[fin1]+bs[fin2] == x {
		return true
	}
	return sumaElementos(as, bs, x, fin1, fin2-1)
}
*/
