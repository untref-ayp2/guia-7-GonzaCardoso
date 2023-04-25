package ejercicios

import "fmt"

// Escribir un método recursivo que dado un número
// entero decimal devuelva su equivalente en
// binario en forma de string
func DecimalABinario(n int) string {
	bin := decimalABinario(n)
	return Invertir(bin)
}

func decimalABinario(n int) string {
	if n == 0 || n == 1 {
		return fmt.Sprint(n)
	}
	resto := n % 2
	resultado := n / 2
	return fmt.Sprint(resto) + decimalABinario(resultado)

}
