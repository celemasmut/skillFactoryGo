package asisalumnos

import "fmt"

//"Alumnos..." es una estructura para la asistencia del Alumno
type Alumnos struct {
	Alumno     string
	Asistencia bool
}

//AlumnAsist es una funcion ....
func (this *Alumnos) AlumnAsist(nombre string) int {
	fmt.Printf("Esta presente %s ?", nombre)
	var presen int
	fmt.Scan(&presen)
	return presen
}

//Presente es una funcion.....
func (this *Alumnos) Presente(pre int) {
	if pre == 1 {
		this.Asistencia = true
	} else {
		this.Asistencia = false
	}
}
