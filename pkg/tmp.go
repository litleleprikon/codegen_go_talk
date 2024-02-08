package pkg

//easyjson:json
type User struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
