package structs

// Platzi es ina interfaz para course and career
type Platzi interface {
	Subscribe(name string)
}

// InterfaceTest validate if interface works
func InterfaceTest() {
	platziCourse := PlatziCourse{
		Name:   "Go",
		Slug:   "Go",
		Skills: []string{"backend"}}

	platziCareer := new(PlatziCareer)
	platziCareer.Name = "Go Career"
	platziCareer.Slug = "Go Career"

	callSubscribe(platziCourse, "Mario")
	callSubscribe(platziCareer, "Mario2")
}

func callSubscribe(p Platzi, name string) {
	p.Subscribe(name)
}
