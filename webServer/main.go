package main

import "net/http" //net es un paquete de redes luego el sub paquete http el cual tiene todo lo necesario para convertir el desarrollo en un servidor web

func main() {
	//handlefunc significa que se maneja el ingreso en el local host y eso ejecurar una funcion
	//donde tiene dos parametros
	//parametro "/" - la ruta que esta detectando
	//parametro home la funcion que tiene que ejecutar.
	http.HandleFunc("/", home)
	//http.HandleFunc("/login", login)

	//listenandserve lo que hace es escuchar y servir.
	//escucha un puerto y si hay info, sirve una peticion y ejecuta una instruccion.
	http.ListenAndServe(":3000", nil)
}

//la funcion tiene dos parametros. flujos de datos de un servidor - enviar y recibir respuestas.
//el puerto estara escuchando las dos peticiones
func home(w http.ResponseWriter, r *http.Request) {
	//w de write y r de recuest
	http.ServeFile(w, r, "./index.html")
}
