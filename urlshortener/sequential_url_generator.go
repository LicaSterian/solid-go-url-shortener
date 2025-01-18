package urlshortener

import "fmt"

type SequentialURLGenerator struct {
	counter int
}

func (g *SequentialURLGenerator) Generate() string {
	g.counter++
	return fmt.Sprintf("%06d", g.counter)
}
