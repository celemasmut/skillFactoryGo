package tarea2

import "fmt"

//Alumno es una estructura
type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int32
}

//ListAlum es un slice de alumnos
var ListAlum []Alumno

//NewAlumno funcion que retorna un Alumno
func NewAlumno(nombre string, apellido string, dni int32) Alumno {
	return Alumno{Nombre: nombre, Apellido: apellido, DNI: dni}
}

//AddAlumno agrega un alumno al slice preArmado
func AddAlumno(misAlumnos Alumno) {
	if ListAlum == nil {
		ListAlum = []Alumno{misAlumnos}
	} else {
		if existeAlumno(ListAlum, misAlumnos.DNI) == true {
			fmt.Println("El alumno ya existe")
		} else {
			ListAlum = append(ListAlum, misAlumnos)
		}
	}
}

//AddAlumno2 es una segunda funcion para agregar alumnos de a uno a las listas de los cursos.
func AddAlumno2(listaAlumno []Alumno, miAlumno Alumno) []Alumno {
	if existeAlumno(listaAlumno, miAlumno.DNI) == true {
		fmt.Println("El alumno ya existe")
	} else {
		listaAlumno = append(listaAlumno, miAlumno)
	}
	return listaAlumno
}

func existeAlumno(listaAlumno []Alumno, dni int32) bool {
	existe := false
	for indx := range ListAlum {
		if listaAlumno[indx].DNI == dni {
			existe = true
		}
	}
	return existe
}

//DeleteAlumno es una funcion que elimina un alumno del slice
func DeleteAlumno(dni int32) {
	for i := 0; i < len(ListAlum); i++ {
		if ListAlum[i].DNI == dni {
			ListAlum = append(ListAlum[:i], ListAlum[i+1:]...)
		}
	}
}
