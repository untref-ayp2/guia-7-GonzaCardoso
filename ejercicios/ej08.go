package ejercicios

var resultado int = 0

// Escribir un método recursivo que devuelva el
// cociente y el resto de la división entera mediante
// restas sucesivas
func DivisionEntera(dividendo, divisor int) (cociente, resto int) {
	if dividendo < divisor {
		return resultado, dividendo
	}
	dividendo = dividendo - divisor
	resultado++

	return DivisionEntera(dividendo, divisor)

}
