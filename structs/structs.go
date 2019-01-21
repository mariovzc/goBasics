package structs

import "fmt"

// PlatziCourse  Struct for courses
type PlatziCourse struct {
	Name   string
	Slug   string
	Skills []string
}

// Subscribe a user in the app
func (p PlatziCourse) Subscribe(name string) {
	fmt.Printf("La persona %s se ha registrado al curso %s", name, p.Name)
}

// PlatziCareer Struck for carrer
type PlatziCareer struct {
	PlatziCourse
}

// Subscribe a using in a career
func (p PlatziCareer) Subscribe(name string) {
	fmt.Printf("La persona %s se ha registrado al curso %s", name, p.Name)
}
