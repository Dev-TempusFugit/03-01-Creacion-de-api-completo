package model

// Community estructura de una comunidad
type Community struct {
	// Name es el nombre de una comunidad. Ej: EDteam
	Name string
}

// Communities slice de comunidades
type Communities []Community

// Person estructura de una persona
type Person struct {
	// Name nombre de la persona. Ej: Maxim
	Name string
	// Edad de la persona. Ej: 26
	Age uint8
	// Communities comunidades a las que pertenece
	Communities Communities
}

// Persons slice de Person
type Persons []Person
