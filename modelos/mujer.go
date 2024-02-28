package modelos

type Mujer struct {
	Hombre
}

func (mujer *Mujer) Respirar() {
	mujer.respirando = true
}

func (mujer *Mujer) Pensar() {
	mujer.pensando = true
}

func (mujer *Mujer) Comiendo() {
	mujer.comiendo = true
}

func (mujer *Mujer) Sexo() string {
	return "Mujer"
}

func (mujer *Mujer) EstaVivo() bool {
	return true
}
