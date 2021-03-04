package main

import (
	"fmt"
	"time"

	miUs "./User"
)

type alumno struct {
	miUs.Usuario
}

func main() {
	user := new(alumno)
	user.AltaUsuario(3, "Luciano", time.Now(), true)
	fmt.Println(user.Usuario)

}
