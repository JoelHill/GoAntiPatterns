package withlesspointers

import "time"

// SomeModel -
type SomeModel struct {
	ID1        int
	AnotherID1 int
	Astring1   string
	Abool1     bool
	ID2        int
	AnotherID2 int
	Astring2   string
	Abool2     bool
	ID3        int
	AnotherID3 int

	// May not always be part of the model
	FilterString *string
	FilterBool   *bool

	// time is a good use for a pointer
	SomeTime *time.Time

	// May not always be part of the model
	AnotherStruct *AnoterStruct
}

// AnoterStruct -
type AnoterStruct struct {
	id int
}
