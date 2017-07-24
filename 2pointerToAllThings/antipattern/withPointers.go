package withpointers

import "time"

// SomeModel -
type SomeModel struct {
	ID1           *int
	AnotherID1    *int
	Astring1      *string
	Abool1        *bool
	ID2           *int
	AnotherID2    *int
	Astring2      *string
	Abool2        *bool
	ID3           *int
	AnotherID3    *int
	FilterString  *string
	FilterBool    *bool
	SomeTime      *time.Time
	AnotherStruct *AnoterStruct
}

// AnoterStruct -
type AnoterStruct struct {
	id *int
}
