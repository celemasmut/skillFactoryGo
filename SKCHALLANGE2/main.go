package main

import (
	// para agregar el paquete tarea2 se crea el modulo de la carpeta donde contiene tanto tarea2 y el main
	//en este caso SKCHALLANGE2
	miAlumno "SKCHALLANGE2/tarea2"
	"fmt"
)

func main() {

	alumno1 := miAlumno.NewAlumno("Sabrina", "juarez", 34765976)
	alumno2 := miAlumno.NewAlumno("Jose", "Fernandez", 37356246)
	alumno3 := miAlumno.NewAlumno("Ayelen", "Burzoni", 43886357)
	alumno4 := miAlumno.NewAlumno("Tamara", "Davin", 34768640)
	alumno5 := miAlumno.NewAlumno("Pedro", "Lubon", 43765498)

	miAlumno.AddAlumno(alumno1)
	miAlumno.AddAlumno(alumno2)
	miAlumno.AddAlumno(alumno3)
	miAlumno.AddAlumno(alumno4)
	miAlumno.AddAlumno(alumno5)
	for index := range miAlumno.ListAlum {
		fmt.Println(miAlumno.ListAlum[index])
	}
	/*miAlumno.DeleteAlumno(alumno2.DNI)
	fmt.Println("\nSlice con eliminacion")
	for index := range miAlumno.ListAlum {
		fmt.Println(miAlumno.ListAlum[index])
	}*/
	listaDeCursos := make([]miAlumno.Curso, 3)
	listaDeCursos[0] = miAlumno.AddCurso(listaDeCursos, "Matematica")
	listaDeCursos[1] = miAlumno.AddCurso(listaDeCursos, "Historia")
	listaDeCursos[2] = miAlumno.AddCurso(listaDeCursos, "TICs")
	fmt.Println(listaDeCursos)
	for index := range listaDeCursos {
		listaDeCursos = miAlumno.AddAlumnosACurso(listaDeCursos, miAlumno.ListAlum, listaDeCursos[index].NombCurso)
	}
	for i := range listaDeCursos {
		fmt.Println(listaDeCursos[i])
	}
	alumno7 := miAlumno.NewAlumno("Luciano", "Kunsolo", 35864346)
	listaDeCursos[2].Lista = miAlumno.AddAlumno2(listaDeCursos[2].Lista, alumno7)
	fmt.Println(listaDeCursos[2])
	alumno6 := miAlumno.NewAlumno("Karla", "Panipea", 45346975)
	listaDeCursos[0].Lista = miAlumno.AddAlumno2(listaDeCursos[0].Lista, alumno6)
	fmt.Println(listaDeCursos[0])
	pos := mayorLista(listaDeCursos)
	fmt.Printf("El Curso con mas alumnos es el %s ", listaDeCursos[pos].NombCurso)

}

func mayorLista(listaDeCursos []miAlumno.Curso) int {
	pos := 0
	cant := 0
	for i := 0; i < len(listaDeCursos); i++ {
		if len(listaDeCursos[i].Lista) > cant {
			pos = i
			cant = len(listaDeCursos[i].Lista)
		}
	}
	return pos
}
