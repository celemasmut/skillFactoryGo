package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// Abbreviate transform a string in its acronym.
func abreviar(s string) string {
	s = cleanString(s)
	palabras := strings.Split(s, " ")
	sigla := ""

	for _, cadaPalabra := range palabras {
		letraInicial := string(cadaPalabra[0])
		sigla += strings.ToUpper(letraInicial)
	}

	return sigla
}

/*Funcion para limpiar la cadena de entrada
eliminando signos que no formen parte de las palabras

Dividir la entrada en palabras por cada palabra tomas la
primera letra para armar la sigla o acronimo*/

func cleanString(s string) string {
	s = strings.ReplaceAll(s, " - ", " ")
	s = strings.ReplaceAll(s, "-", " ")
	//lo que esta entre corchetes es lo que se indica que es lo que puede tener la palabra. el + indica que puede ser 0 o mas .
	//expresion regular toma solo letras espacios. si quiero numeros :("[^a-zA-Z0-9 ]+")
	reg, err := regexp.Compile("[^a-zA-Z ]+")
	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(s, "")
}

/*Entrada: "Organizacion Naciones Unidas"
Almacenamiento intermedio: ["Organizacion","Naciones","Unidas"]
["O""N""U"]
Salida :"ONU"
*/

type acronymTest struct {
	input    string
	expected string
}

var stringTestCases = []acronymTest{
	{
		input:    "Portable *Network Graphics",
		expected: "PNG",
	},
	{
		input:    "Ruby on 7Rails",
		expected: "ROR",
	},
	{
		input:    "First In, First Out",
		expected: "FIFO",
	},
	{
		input:    "GNU Image Manipulation Program",
		expected: "GIMP",
	},
	{
		input:    "Complementary metal-oxide semiconductor",
		expected: "CMOS",
	},
}

/*Programa va a trabajar sobre una cadena de caracteres , lo va a dividir
en palabras y tomar la primera letra de cada palabra para devolver el
acronimo

Entrada: cad de caract
salida: sigla
*/
func main() {

	//for _, test := range stringTestCases {
	actual := abreviar("Hola-Mundo-Go")
	fmt.Printf("Frase original [Hola-Mundo-Go], siglas [%s]\n", actual)
}
