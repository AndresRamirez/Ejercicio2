package main

import "fmt"

var numero1, numero2, total, ruta int

func main() {
MENU:
	fmt.Println("Digite:\n 1 -> para Sumar\n 2 -> para Restar\n 3 -> para Multiplicar\n 4 -> para Dividir\n 5 -> para Salir\n")
	fmt.Scanln(&ruta)
	if ruta == 1 || ruta == 2 || ruta == 3 || ruta == 4 || ruta == 5 {
		fmt.Println("número valido")
		switch ruta {
		case 1:
			fmt.Println("voy a sumar")
			fmt.Println("\nIngrese numero 1: ")
			fmt.Scanln(&numero1)

			fmt.Println("\nIngrese numero 2: ")
			fmt.Scanln(&numero2)

			total = numero1 + numero2

			fmt.Println("\nel total es: ", total)

			fmt.Println("\nDigite:\n1 -> para volver al menu\n2 -> para salir")
			fmt.Scanln(&ruta)

			if ruta == 1 {
				goto MENU
			} else {

				break
			}
		case 2:
			fmt.Println("voy a restar")
			fmt.Println("\nIngrese numero 1: ")
			fmt.Scanln(&numero1)

			fmt.Println("\nIngrese numero 2: ")
			fmt.Scanln(&numero2)

			total = numero1 - numero2

			fmt.Println("\nel total es: ", total)

			fmt.Println("\nDigite:\n1 -> para volver al menu\n2 -> para salir")
			fmt.Scanln(&ruta)

			if ruta == 1 {
				goto MENU
			} else {

				break
			}
		case 3:
			fmt.Println("voy a Multiplicar")
			fmt.Println("\nIngrese numero 1: ")
			fmt.Scanln(&numero1)

			fmt.Println("\nIngrese numero 2: ")
			fmt.Scanln(&numero2)

			total = numero1 * numero2

			fmt.Println("\nel total es: ", total)

			fmt.Println("\nDigite:\n1 -> para volver al menu\n2 -> para salir")
			fmt.Scanln(&ruta)

			if ruta == 1 {
				goto MENU
			} else {

				break
			}
		case 4:
			fmt.Println("voy a Dividir")
			fmt.Println("\nIngrese numero 1: ")
			fmt.Scanln(&numero1)

			fmt.Println("\nIngrese numero 2: ")
			fmt.Scanln(&numero2)

			total2 := float64(numero1) / float64(numero2)

			fmt.Println("\nel total es: ", total2)

			fmt.Println("\nDigite:\n1 -> para volver al menu\n2 -> para salir")
			fmt.Scanln(&ruta)

			if ruta == 1 {
				goto MENU
			} else {

				break
			}
		case 5:
			break
		}
	} else {
		fmt.Println("digite un numero valido")
		goto MENU
	}
}
