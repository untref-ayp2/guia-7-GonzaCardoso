package ejercicios

// Escriba un método recursivo que tome un entero n
// devuelva su factorial
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
