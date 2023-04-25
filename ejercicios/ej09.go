package ejercicios

// Escribir un método recursivo que tome un
// array de números enteros y devuelva la suma
// de todos sus elementos
func SumaArray(v []int) int {
	if len(v) == 0 {
		return 0
	}

	return sumaArray(v, len(v)-1)
}

func sumaArray(j []int, fin int) int {
	if fin == 0 {
		return j[fin]
	}
	return j[fin] + sumaArray(j, fin-1)
}
