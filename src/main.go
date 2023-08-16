package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

func normalFunction(message string) {
	fmt.Println(message)
}

func fourArguments(a, b int, c string, d bool) {
	fmt.Println(a, b, c, d)
}

func returnValue(a int) int {
	return a * 2
}

func returnMultiValue(a int) (c, d int) {
	return a, a * 2
}

// funcion para verficar pares
func isEven(num int) bool {
	return num%2 == 0
}

// funcion para verficar user y password
func userValidator(user, password string) bool {
	return user == "Alejo" && password == "hola123"
}

// Defer
func connection() {
	defer fmt.Println("closing connection")
	defer fmt.Println("working with connection")
	defer fmt.Println("opening connection")
	fmt.Println("looking for connection")
}

// Continue
func findNumber(num int) {
	rango := rand.Intn(num)
	fmt.Println("Rango:", rango)
	for i := 0; i < rango; i++ {
		fmt.Println(i)
		if i == rango/2 {
			fmt.Println("Encontrado:", i)
			continue
		}
	}
}

// Break
func findNumberBreak(num int) {
	rango := rand.Intn(num)
	fmt.Println("Rango:", rango)
	for i := 0; i < rango; i++ {
		fmt.Println(i)
		if i == rango/2 {
			fmt.Println("Encontrado:", i)
			break
		}
	}
}

func main() {
	// Declaración de constantes
	const PI float64 = 3.14
	const PI2 = 3.1416

	fmt.Println("PI:", PI)
	fmt.Println("PI2:", PI2)

	//Declaración de variables enteras
	base := 12
	var altura int = 9
	// base = 5
	area := base * altura

	fmt.Println("area:", area)

	//zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("a:", a, "b:", b, "c:", c, "d:", d)

	//Operaciones
	x := 4
	y := 8

	result := x + y
	fmt.Println("La suma es:", result)

	result = y - x
	fmt.Println("La resta es:", result)

	result = x * y
	fmt.Println("El producto es:", result)

	result = y / x
	fmt.Println("El cociente es:", result)

	result = y % x
	fmt.Println("El residuo es:", result)

	//funciones

	normalFunction("Hello World")

	fourArguments(1, -9, "Hola Mundo", true)

	value := returnValue(25)
	fmt.Println(value)

	value1, value2 := returnMultiValue(8)
	fmt.Printf("value1: %v\nvalue2: %v\n", value1, value2)

	//para quedarme con alguno de los dos valores:

	valueFirst, _ := returnMultiValue(9)
	fmt.Println("valueFirst:", valueFirst)

	_, valueSecond := returnMultiValue(11)
	fmt.Println("valueSecond:", valueSecond)

	//for condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n")

	//for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("\n")

	//for forever
	// counterForever := 0
	// for {
	// fmt.Println(counterForever)
	// counterForever++
	// }

	//condicional if/else - operadores lógicos
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Se cumple")
	}

	//Convertir texto a número
	valorNumero, err := strconv.Atoi("53")
	// valorNumero, err := strconv.Atoi("hola53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("valor número:", valorNumero)

	fmt.Println(isEven(9))

	fmt.Println(userValidator("Alejo", "holamundo"))

	//switch
	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//switch sin condicion
	edad := 26
	switch {
	case edad > 18:
		fmt.Println("Es mayor de edad")
	case edad > 0 && edad < 18:
		fmt.Println("Es menor de edad")
	default:
		fmt.Println("Edad no valida")
	}

	//defer
	connection()

	fmt.Println("\n")

	//continue
	findNumber(20)

	fmt.Println("\n")

	//break
	findNumberBreak(15)
}
