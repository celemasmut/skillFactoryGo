package main

import "fmt"

func main() {
	// entre corch va el tipo de clave y luego indico el tipo de contenido.
	paises := make(map[string]string)
	paises["Mexico"] = "D.F"
	paises["Brazil"] = "Brazilia"
	fmt.Println(paises)
	campeonato := map[string]int{
		"BJ": 35,
		"RM": 29,
		"CH": 48,
		"BR": 56,
	}
	fmt.Println(campeonato) // me muestra la clave y valor ordenando en orden alfabetico

	delete(campeonato, "RM")
	fmt.Println(campeonato)

	//itera map
	for equipo, puntaje := range campeonato {
		fmt.Printf("El puntaje del equipo %s es de : %d\n", equipo, puntaje)
	}

	puntaje, existe := campeonato["BR"]
	fmt.Printf("El puntaje es %d y el equipo existe:%t", puntaje, existe)

}
