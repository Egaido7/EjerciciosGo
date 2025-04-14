package main

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	_votesQuantity = 10
	_voteMin       = 1
	_voteMax       = 5
)

func main() {
	votosCantidad := 0
	votos := map[int]int{}

	for {
		var voto int
		fmt.Println("votos realizados hasta ahora: ", votosCantidad)
		fmt.Println("Ingrese su voto: ")
		_, _ = fmt.Scanf("%d", &voto)

		if voto < _voteMin || voto > _voteMax {
			fmt.Println("...el voto debe ser un número entre el 1 y el 5...")
			continue
		}

		v, ok := votos[voto]
		if !ok {
			votos[voto] = 1
		} else {
			votos[voto] = v + 1
		}

		votosCantidad++
		if votosCantidad == _votesQuantity {
			break
		}
	}

	contar(votos)
	mostrarMensaje(votos)
}

func contar(myMap map[int]int) {
	contador1 := 0
	contador2 := 0
	contador3 := 0
	contador4 := 0
	contador5 := 0
	for _, v := range myMap {
		if v == 1 {
			contador1++
		} else if v == 2 {
			contador2++
		} else if v == 3 {
			contador3++
		} else if v == 4 {
			contador4++
		} else if v == 5 {
			contador5++
		} else {
			color.Red(fmt.Sprintf("Valor ingresado incorrecto: %v", v))
		}

	}
	fmt.Println("El contador de las opciones es:")
	fmt.Println("1==>", contador1)
	fmt.Println("2==>", contador2)
	fmt.Println("3==>", contador3)
	fmt.Println("4==>", contador4)
	fmt.Println("5==>", contador5)
}

func mostrarMensaje(myMap map[int]int) {
	contadorpositivo := 0
	contadornegativo := 0
	for _, v := range myMap {
		if v >= 4 {
			contadorpositivo++
		} else if v <= 2 {
			contadornegativo++
		}
	}
	fmt.Println("---------------------------------------")
	if contadorpositivo > contadornegativo {
		color.Green("¡Buen resultado!")
	} else if contadorpositivo < contadornegativo {
		color.Red("Resultado mejorable")

	} else {
		color.Yellow("Resultado neutro")
	}
}
