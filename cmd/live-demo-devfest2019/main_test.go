package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/necrophonic/pony/internal/pony"
)

func TestPony(t *testing.T) {

	// Create a new router and plug in the defined routes
	r := httprouter.New()
	setupRoutes(r)

	ts := httptest.NewServer(r)
	defer ts.Close()

	cases := map[string][]string{
		"fluttershy":      []string{"fluttershy", "kindness"},
		"pinkiepie":       []string{"pinkiepie", "laughter"},
		"rainbowdash":     []string{"rainbowdash", "loyalty"},
		"rarity":          []string{"rarity", "generosity"},
		"twilightsparkle": []string{"twilightsparkle", "magic"},
		"applejack":       []string{"applejack", "honesty"},
	}

	for name, expected := range cases {
		// Again ensure the url is built using the base url from the test server
		// so that parameters are correctly captured.
		url := ts.URL + "/pony/" + name
		resp, err := http.Get(url)
		if err != nil {
			t.Fatal(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		// TODO Test returned json - can't just string match as can't guarantee
		// marshalled order of elements.
		// Unmarshal the result and check the data
		var p pony.Pony
		if err := json.Unmarshal(body, &p); err != nil {
			t.Errorf("bad pony unmarshal, error: %v", err)
		}

		if p.Name != expected[0] {
			t.Errorf("Wrong name, expected '%s', got '%s'", expected[0], p.Name)
		}
		if p.Element != expected[1] {
			t.Errorf("Wrong name, expected '%s', got '%s'", expected[1], p.Element)
		}

		resp.Body.Close()
	}
}
