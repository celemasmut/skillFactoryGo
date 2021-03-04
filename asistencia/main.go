package main

import (
	"fmt"

	misAlum "./asisalumnos"
	myDias "./dias"
)

type myDay struct {
	myDias.Misdias
}
type alumni struct {
	misAlum.Alumnos
}

var presente int

func main() {

	/*asistencia := map[string]string{
		"alumno1": "si",
		"alumno2": "no",
		"alumno3": "si",
		"alumno4": "si",
		"alumno5": "no",
		"alumno6": "si",
	}
	fmt.Println("\n\n Asistencia : ")

	for alumno, presente := range asistencia {

		fmt.Printf("\nEl alumno %s estuvo presente: %s", alumno, presente)
	}
	fmt.Println("\n\nLos Presentes fueron : ")
	for alumno, presente := range asistencia {
		if presente == "si" {
			fmt.Printf("El alumno %s\n", alumno)
		}
	}*/
	listAlum := make([]alumni, 5)
	day := new(myDay)
	diaElegido := day.ElegirDias() //interfaz
	switch diaElegido {
	case 0:
		fmt.Printf("Usted ha elegido al dia : %s", day.Misdias.Dias[diaElegido])
		listAlum[0].Alumnos.Alumno = "Juarez"
		listAlum[1].Alumnos.Alumno = "Correa"
		listAlum[2].Alumnos.Alumno = "Astrada"
		listAlum[3].Alumnos.Alumno = "Luna"
		listAlum[4].Alumnos.Alumno = "Pesantes"
		fmt.Println("\nAsistencia : 1=Presente  0=Ausente\n ")

		for i := 0; i < 5; i++ {
			presente = listAlum[i].AlumnAsist(listAlum[i].Alumnos.Alumno)
			listAlum[i].Presente(presente)
		}
		cont := 0
		for j := 0; j < 5; j++ {
			if listAlum[j].Asistencia == true {
				cont++
			}
		}
		fmt.Printf("\nLos Presentes son :%d", cont)
		break
	case 1:
		fmt.Printf("Usted ha elegido al dia : %s", day.Misdias.Dias[diaElegido])
		listAlum[0].Alumnos.Alumno = "Juarez"
		listAlum[1].Alumnos.Alumno = "Correa"
		listAlum[2].Alumnos.Alumno = "Astrada"
		listAlum[3].Alumnos.Alumno = "Luna"
		listAlum[4].Alumnos.Alumno = "Pesantes"
		fmt.Println("\nAsistencia : 1=Presente  0=Ausente\n ")

		for i := 0; i < 5; i++ {
			presente = listAlum[i].AlumnAsist(listAlum[i].Alumnos.Alumno)
			listAlum[i].Presente(presente)
		}
		cont := 0
		for j := 0; j < 5; j++ {
			if listAlum[j].Asistencia == true {
				cont++
			}
		}
		fmt.Printf("\nLos Presentes son :%d", cont)
		break
	case 2:
		fmt.Printf("Usted ha elegido al dia : %s", day.Misdias.Dias[diaElegido])
		listAlum[0].Alumnos.Alumno = "Juarez"
		listAlum[1].Alumnos.Alumno = "Correa"
		listAlum[2].Alumnos.Alumno = "Astrada"
		listAlum[3].Alumnos.Alumno = "Luna"
		listAlum[4].Alumnos.Alumno = "Pesantes"
		fmt.Println("\nAsistencia : 1=Presente  0=Ausente\n ")

		for i := 0; i < 5; i++ {
			presente = listAlum[i].AlumnAsist(listAlum[i].Alumnos.Alumno)
			listAlum[i].Presente(presente)
		}
		cont := 0
		for j := 0; j < 5; j++ {
			if listAlum[j].Asistencia == true {
				cont++
			}
		}
		fmt.Printf("\nLos Presentes son :%d", cont)
		break
	case 3:
		fmt.Printf("Usted ha elegido al dia : %s", day.Misdias.Dias[diaElegido])
		listAlum[0].Alumnos.Alumno = "Juarez"
		listAlum[1].Alumnos.Alumno = "Correa"
		listAlum[2].Alumnos.Alumno = "Astrada"
		listAlum[3].Alumnos.Alumno = "Luna"
		listAlum[4].Alumnos.Alumno = "Pesantes"
		fmt.Println("\nAsistencia : 1=Presente  0=Ausente\n ")

		for i := 0; i < 5; i++ {
			presente = listAlum[i].AlumnAsist(listAlum[i].Alumnos.Alumno)
			listAlum[i].Presente(presente)
		}
		cont := 0
		for j := 0; j < 5; j++ {
			if listAlum[j].Asistencia == true {
				cont++
			}
		}
		fmt.Printf("\nLos Presentes son :%d", cont)
		break
	case 4:
		fmt.Printf("Usted ha elegido al dia : %s", day.Misdias.Dias[diaElegido])
		listAlum[0].Alumnos.Alumno = "Juarez"
		listAlum[1].Alumnos.Alumno = "Correa"
		listAlum[2].Alumnos.Alumno = "Astrada"
		listAlum[3].Alumnos.Alumno = "Luna"
		listAlum[4].Alumnos.Alumno = "Pesantes"
		fmt.Println("\nAsistencia : 1=Presente  0=Ausente\n ")

		for i := 0; i < 5; i++ {
			presente = listAlum[i].AlumnAsist(listAlum[i].Alumnos.Alumno)
			listAlum[i].Presente(presente)
		}
		cont := 0
		for j := 0; j < 5; j++ {
			if listAlum[j].Asistencia == true {
				cont++
			}
		}
		fmt.Printf("\nLos Presentes son :%d", cont)
		break
	case 5:
		fmt.Printf("Usted ha elegido al dia : %s", day.Misdias.Dias[diaElegido])
		listAlum[0].Alumnos.Alumno = "Juarez"
		listAlum[1].Alumnos.Alumno = "Correa"
		listAlum[2].Alumnos.Alumno = "Astrada"
		listAlum[3].Alumnos.Alumno = "Luna"
		listAlum[4].Alumnos.Alumno = "Pesantes"
		fmt.Println("\nAsistencia : 1=Presente  0=Ausente\n ")

		for i := 0; i < 5; i++ {
			presente = listAlum[i].AlumnAsist(listAlum[i].Alumnos.Alumno)
			listAlum[i].Presente(presente)
		}
		cont := 0
		for j := 0; j < 5; j++ {
			if listAlum[j].Asistencia == true {
				cont++
			}
		}
		fmt.Printf("\nLos Presentes son :%d", cont)
		break
	case 6:
		fmt.Printf("Usted ha elegido al dia : %s", day.Misdias.Dias[diaElegido])
		listAlum[0].Alumnos.Alumno = "Juarez"
		listAlum[1].Alumnos.Alumno = "Correa"
		listAlum[2].Alumnos.Alumno = "Astrada"
		listAlum[3].Alumnos.Alumno = "Luna"
		listAlum[4].Alumnos.Alumno = "Pesantes"
		fmt.Println("\nAsistencia : 1=Presente  0=Ausente\n ")

		for i := 0; i < 5; i++ {
			presente = listAlum[i].AlumnAsist(listAlum[i].Alumnos.Alumno)
			listAlum[i].Presente(presente)
		}
		cont := 0
		for j := 0; j < 5; j++ {
			if listAlum[j].Asistencia == true {
				cont++
			}
		}
		fmt.Printf("\nLos Presentes son :%d", cont)
		break
	default:
		fmt.Println("Error no existe ese dia")
		break
	}

}
