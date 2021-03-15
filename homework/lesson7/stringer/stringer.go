package stringer

type Pill int

type Hospital int


//go:generate stringer -type=Hospital -output=medicine_homes.go
const (
	Central Hospital = iota
	Main
)

//go:generate stringer -type=Pill
const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

type Artist int

//go:generate stringer -type=Artist -linecomment
const (
	ArtistATB Artist = iota // andre
	ArtistPVD               // paul
	ArtistASOT              // armin
)



