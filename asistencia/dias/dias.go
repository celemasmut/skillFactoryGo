package dias

import "fmt"

//Misdias struct
type Misdias struct {
	Dias [6]string
}

//ElegirDias es una funcion para elegir el dia de asistencia
func ElegirDias() int {
	Dias = [6]string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}
	for n := 0; n < 6; n++ {
		fmt.Printf("\nDia:%d %s", n+1, this.Dias[n])
	}
	fmt.Println("Elige dia de asistencia: ")
	var diaElegido int
	fmt.Scan(&diaElegido)
	return diaElegido - 1
}
