// Package db simulates a simple database access layer
package db

import (
	"errors"

	"github.com/necrophonic/pony/internal/pony"
)

// ponies pretends to be a set of database records
var ponies = map[string]string{
	"applejack":       "honesty",
	"pinkiepie":       "laughter",
	"rarity":          "generosity",
	"rainbowdash":     "loyalty",
	"twilightsparkle": "magic",
	"fluttershy":      "kindness",
}

// GetPony fetches a pony from the "database" by name
func GetPony(name string) (*pony.Pony, error) {
	if element, ok := ponies[name]; ok {
		return &pony.Pony{Name: name, Element: element}, nil
	}
	return nil, errors.New("failed to find pony")
}

// SetPony stores a pony to the "database"
func SetPony(name, element string) error {
	if _, ok := ponies[name]; ok {
		return errors.New("pony already exists")
	}
	ponies[name] = element
	return nil
}

// AllPonies returns all the ponies in the "database"
func AllPonies() *pony.List {

	list := make([]pony.Pony, 0, len(ponies))

	for name, element := range ponies {
		list = append(list, pony.Pony{Name: name, Element: element})
	}

	return &pony.List{Ponies: list}
}
