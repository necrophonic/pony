package db_test

import (
	"testing"

	"github.com/necrophonic/pony/internal/db"
	"github.com/necrophonic/pony/internal/pony"
)

func TestGetPony(t *testing.T) {
	var pony *pony.Pony
	var err error

	// Fetch existing pony
	pony, err = db.GetPony("fluttershy")
	if err != nil {
		t.Errorf("Got error '%v', expected nil", err)
	}
	if pony == nil {
		t.Error("Failed to get known pony")
	}
	if pony.Element != "kindness" {
		t.Errorf("Got element '%s', expected '%s'", pony.Element, "kindness")
	}

	// Fetch fake pony
	pony, err = db.GetPony("nopony")
	if err == nil {
		t.Errorf("Got nil, but expected error '%v'", err)
	}
	if pony != nil {
		t.Errorf("Got '%s', expected no pony!", pony.Name)
	}
}

func TestSetPony(t *testing.T) {
	// Check pony doesn't exist yet
	var pony *pony.Pony
	var err error

	pony, err = db.GetPony("tempest")
	if pony != nil {
		t.Error("Expected pony not to exist yet!")
	}

	// Attempt to set then check
	err = db.SetPony("tempest", "evil")
	if err != nil {
		t.Errorf("Got error '%v', expected nil", err)
	}

	pony, err = db.GetPony("tempest")
	if err != nil {
		t.Errorf("Got error '%v', expected nil", err)
	}
	if pony == nil {
		t.Error("Couldn't fetch new pony")
	}

	if pony.Name != "tempest" {
		t.Errorf("Got pony name '%s', expected '%s'", pony.Name, "temptest")
	}
	if pony.Element != "evil" {
		t.Errorf("Got pony element '%s', expected '%s'", pony.Element, "evil")
	}

}
