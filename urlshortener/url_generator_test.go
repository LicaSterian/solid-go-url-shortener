package urlshortener

import "testing"

func TestRandomURLGenerator(t *testing.T) {
	generator := RandomURLGenerator{}
	id := generator.Generate()
	if len(id) != 6 {
		t.Error("expected RandomURLGenerator to Generate an string id of length 6")
	}
}
