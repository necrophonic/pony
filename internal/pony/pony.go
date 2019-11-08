// Package pony gives structs representing characters for MLP:FiM
package pony

type (
	// Pony is a single pony character
	Pony struct {
		Name    string
		Element string
	}

	// List is a list of existing ponies
	List struct {
		Ponies []Pony
	}
)
