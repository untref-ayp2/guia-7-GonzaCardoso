package ejercicios

// Escriba un método recursivo que calcule el máximo
// común divisor entre dos números enteros.
// Nota: Se puede usar el algoritmo de Euclides para
// resolver este problema.
func MCD(a, b int) int {
	resto := 0
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	} else {
		resto = a % b
		return MCD(b, resto)
	}
}
