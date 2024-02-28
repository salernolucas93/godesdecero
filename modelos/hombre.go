package modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

func (hombre *Hombre) Respirar() {
	hombre.respirando = true
}

func (hombre *Hombre) Pensar() {
	hombre.pensando = true
}

func (hombre *Hombre) Comer() {
	hombre.comiendo = true
}

func (hombre *Hombre) Sexo() string {
	return "Hombre"
}
