package user

import "time"

//Usuario...
type Usuario struct {
	ID        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

//() significa que voy a estar trabajando en una struct usuario
func (this *Usuario) AltaUsuario(id int, nomb string, fechaAlta time.Time, status bool) {
	this.ID = id
	this.Nombre = nomb
	this.FechaAlta = fechaAlta
	this.Status = status
}
