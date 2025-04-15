package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// Estructuras similares a clases
type Controlable interface {
	encender() bool
	apagar() bool
	Estadoactual() string
}

type Dispositivo struct {
	nombre string
	estado bool
}

func (d *Dispositivo) encender() error { //resever
	//Minuscula privado; Mayuscula publico
	//estoy usando el puntero de dispositivo
	if !d.estado {
		d.estado = true
		color.Green("Encendiendo %s...", d.nombre)
		return nil
	} else {
		//color.Yellow("%s ya está encendido.", d.nombre)
		//return fmt.Errorf("el dispositivo %s ya está encendido.", d.nombre)
		return errors.New("el dispositivo ya está encendido")
	}
}

func (d *Dispositivo) apagar() error {
	if d.estado {
		d.estado = false
		color.Green("Apagando %s...", d.nombre)
		return nil
	} else {
		//color.Yellow("%s ya está apagado.", d.nombre)
		return errors.New("el dispositivo ya está encendido")
	}
}

func (d *Dispositivo) Estadoactual() string {
	if d.estado {
		return "encendido"
	}
	return "apagado"
}

func ingresarDispositivo() (*Dispositivo, error) {
	reader := bufio.NewReader(os.Stdin)
	dispositivo := &Dispositivo{}

	fmt.Print("Introduce el nombre del dispositivo: ")
	nombre, err := reader.ReadString('\n')
	if err != nil {
		return dispositivo, fmt.Errorf("error al leer nombre: %w", err)
	}
	dispositivo.nombre = strings.TrimSpace(nombre)
	dispositivo.estado = false
	return dispositivo, nil
}

func visualizarDispositivo(d Dispositivo) {
	fmt.Println("-----------------------------")
	fmt.Println("Nombre:", d.nombre)
	fmt.Println("Estado:", d.Estadoactual())
	fmt.Println("-----------------------------")
}

func listarDispositivos(dispositivos []*Dispositivo) {
	if len(dispositivos) == 0 {
		color.Red("No hay dispositivos cargados.")
		return
	}
	for i, d := range dispositivos {
		fmt.Printf("%d. %s (%s)\n", i+1, d.nombre, d.Estadoactual())
	}
}

func seleccionarDispositivo(dispositivos []*Dispositivo) int {
	listarDispositivos(dispositivos)
	if len(dispositivos) == 0 {
		return -1
	}

	fmt.Print("Selecciona el número del dispositivo: ")
	reader := bufio.NewReader(os.Stdin)
	entrada, _ := reader.ReadString('\n')
	entrada = strings.TrimSpace(entrada)
	indice, err := strconv.Atoi(entrada)
	if err != nil || indice < 1 || indice > len(dispositivos) {
		color.Red("Selección no válida.")
		return -1
	}
	return indice - 1
}

func main() {
	var dispositivos []*Dispositivo
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nMenú:")
		fmt.Println("1. Ingresar nuevo dispositivo")
		fmt.Println("2. Listar dispositivos")
		fmt.Println("3. Encender dispositivo")
		fmt.Println("4. Apagar dispositivo")
		fmt.Println("5. Salir")
		fmt.Print("Opción: ")

		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "1":
			disp, err := ingresarDispositivo()
			if err != nil {
				color.Red("Error: %v", err)
			} else {
				dispositivos = append(dispositivos, disp)
				color.Green("Dispositivo agregado con éxito.")
			}
		case "2":
			listarDispositivos(dispositivos)
		case "3":
			indice := seleccionarDispositivo(dispositivos)
			if indice != -1 {
				if err := dispositivos[indice].encender(); err != nil {
					color.Yellow("Aviso: %v", err)
				}
			}
		case "4":
			indice := seleccionarDispositivo(dispositivos)
			if indice != -1 {
				if err := dispositivos[indice].apagar(); err != nil {
					color.Yellow("Aviso: %v", err)
				}
			}
		case "5":
			color.Blue("¡Hasta luego!")
			return
		default:
			color.Red("Opción no válida.")
		}
	}
}
