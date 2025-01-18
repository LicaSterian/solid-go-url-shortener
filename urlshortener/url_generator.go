package urlshortener

import (
	"fmt"
	"math/rand"
)

type URLGenerator interface {
	Generate() string
}

type RandomURLGenerator struct{}

func (g *RandomURLGenerator) Generate() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
