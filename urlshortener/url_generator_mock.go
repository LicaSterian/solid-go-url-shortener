package urlshortener

type MockGenerator struct {
	value string
}

func (m *MockGenerator) Generate() string {
	return m.value
}
