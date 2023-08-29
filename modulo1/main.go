package main

import "fmt"

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	aux := *x
	*x = *y
	*y = aux
}

func Suma(vector []int) int {
	largo := len(vector)
	if largo == 0 {
		return 0
	}
	return vector[largo-1] + Suma(vector[0:largo-1])
}

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

func Comparar(vector1 []int, vector2 []int) int {
	largo1 := len(vector1)
	largo2 := len(vector2)
	if largo1 == 0 && largo2 == 0 {
		return 0
	}
	if largo1 == 0 || vector1[0] < vector2[0] {
		return -1
	}
	if largo2 == 0 || vector2[0] < vector1[0] {
		return 1
	}
	return Comparar(vector1[1:largo1], vector2[1:largo2])
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	largo := len(vector)
	posicionMayor := 0
	if largo == 0 {
		posicionMayor = -1
	} else {
		for i := 0; i < largo-1; i++ {
			posicionMayor := i
			for j := i + 1; j < largo; j++ {
				if vector[j] > vector[posicionMayor] {
					posicionMayor = j
				}
			}
			vector[i], vector[posicionMayor] = vector[posicionMayor], vector[i]
		}
	}
	return posicionMayor
}

func main() {
	//x := 3
	//y := 6
	//Swap(&x, &y)
	//fmt.Println(x)
	//fmt.Println(y)
	//fmt.Println(EsCadenaCapicua("jibkj"))
	//Slice := []int{}
	Slice1 := []int{10, 20, 30, 40, 50}
	Slice2 := []int{10, 20, 30, 40}

	fmt.Println(Comparar(Slice2, Slice1))
}
