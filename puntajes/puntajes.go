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
	fmt.Println("El contador de las opciones es:")
	for opcion, cantidad := range myMap {
		color.Red("Opción %d => %d votos\n", opcion, cantidad)
	}
}

func mostrarMensaje(myMap map[int]int) {
	contadorpositivo := 0
	contadornegativo := 0

	for opcion, cantidad := range myMap {
		if opcion >= 4 {
			contadorpositivo += cantidad
		} else if opcion <= 2 {
			contadornegativo += cantidad
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
