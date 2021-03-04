package tarea2

import (
	"fmt"
)

//Curso es una estructura
type Curso struct {
	Lista     []Alumno
	NombCurso string
}

//AddCurso agrega un curso verifica si existe o no
func AddCurso(listaCurso []Curso, nombre string) Curso {
	materia := materiaExiste(listaCurso, nombre)
	if materia == true {
		fmt.Println("La materia ya existe")
	} else {
		return Curso{
			Lista:     []Alumno{},
			NombCurso: nombre,
		}
	}
	return Curso{}
}

func materiaExiste(listaMaterias []Curso, nombre string) bool {
	existe := false
	pos := len(listaMaterias)
	for i := 0; i < pos; i++ {
		if listaMaterias[i].NombCurso == nombre {
			existe = true
		}
	}
	return existe
}

//AddAlumnosACurso agrega la lista de alumnos al curso
func AddAlumnosACurso(listaCurso []Curso, alumnos []Alumno, nombreCur string) []Curso {
	switch nombreCur {
	case "Matematica":
		listaCurso[0].Lista = alumnos[:3]
		break
	case "Historia":
		listaCurso[1].Lista = alumnos[1:4]
		break
	case "TICs":
		listaCurso[2].Lista = alumnos[2:4]
		break
	default:
		fmt.Println("Tal materia no existe para agregar alumnos")
	}
	return listaCurso
}
