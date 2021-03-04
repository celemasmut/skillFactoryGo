package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	leoArchivo()
	leoArchivo2()
	grabarArchivo()
	grabarArchivo2()
}

func leoArchivo() {
	//read file de io/ioutil. devuelve dos variables . en archivo el contenido del archivo que lee y en err el error.
	archivo, err := ioutil.ReadFile("./arch.txt")
	//con readfile gestiona el open y el close del archivo
	if err != nil { // nil = null
		fmt.Println("Hubo error")
	} else {
		//string transforma el archivo en string ya que el compilador lo toma como un arreglo.
		fmt.Println(string(archivo))
	}

}

func leoArchivo2() {
	archivo, err := os.Open("./arch.txt")
	//con readfile gestiona el open y el close del archivo
	if err != nil { // nil = null
		fmt.Println("Hubo error")
	} else { //en scanner lee el contenido del archivo
		scanner := bufio.NewScanner(archivo)
		//recorrer archivo. scanea linea por linea
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Printf("Linea > " + registro + "\n")
		}
	}
	archivo.Close()
}

func grabarArchivo() {
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("Hubo error")
	} else {
		//fprintln es agregar una linea en el archivo
		fmt.Fprintln(archivo, "Esto es la nueva linea")
		archivo.Close()
	}
}

func grabarArchivo2() {
	nombreArchivo := "./nuevoArchivo.txt"
	if Append(nombreArchivo, "\nSegunda Linea agregada") == false {
		fmt.Println("Error al agregar")
	}
}

func Append(archivo string, texto string) bool {
	//bandera: os.O_WRonly quiere decir que se puede tanto escribir como leer.
	//O_Apend es para agregar al final del archivo
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0664)
	var key bool
	if err != nil {
		fmt.Println("Hubo error en el OpenFile")
		key = false
	} else {
		//writestring retorna dos variables. si no uso uno de las variables usa _
		_, err2 := arch.WriteString(texto)
		if err2 != nil {
			fmt.Println("Hubo error en WriteString")
			key = false
		} else {
			key = true
		}
	}
	arch.Close()
	return key
}
