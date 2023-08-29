package tp0

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	aux := *x
	*x = *y
	*y = aux
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	largo := len(vector)
	posicionMayor := 0
	if largo == 0 {
		posicionMayor = -1
	} else {
		for i := 0; i < largo; i++ {
			for j := 1; j < largo; j++ {
				if vector[j] > vector[posicionMayor] {
					posicionMayor = j
				}
			}
		}
	}
	return posicionMayor
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.

func Comparar(vector1 []int, vector2 []int) int {
	largo1 := len(vector1)
	largo2 := len(vector2)
	if largo1 == 0 && largo2 == 0 {
		return 0
	}
	if largo1 == 0 {
		return -1
	}
	if largo2 == 0 {
		return 1
	}
	if vector1[0] < vector2[0] {
		return -1
	}
	if vector2[0] < vector1[0] {
		return 1
	}
	return Comparar(vector1[1:largo1], vector2[1:largo2])
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	largo := len(vector)
	for i := 0; i < largo-1; i++ {
		posicionMenor := i
		for j := i + 1; j < largo; j++ {
			if vector[j] < vector[posicionMenor] {
				posicionMenor = j
			}
		}
		vector[i], vector[posicionMenor] = vector[posicionMenor], vector[i]
	}
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	largo := len(vector)
	if largo == 0 {
		return 0
	}
	return vector[largo-1] + Suma(vector[0:largo-1])
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).

func EsCadenaCapicua(cadena string) bool {
	largo := len(cadena)
	if largo <= 1 {
		return true
	}
	if cadena[0] != cadena[largo-1] {
		return false
	}
	return EsCadenaCapicua(cadena[1 : largo-1])
}
