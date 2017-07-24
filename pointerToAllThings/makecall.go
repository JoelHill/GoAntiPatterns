package makecall

import "github.com/JoelHill/GoAntiPatterns/pointerToAllThings/antipattern"
import "github.com/JoelHill/GoAntiPatterns/pointerToAllThings/betterway"

// makeAntiPatternUtilCalls - Test calls
func makeAntiPatternPointerCalls() bool {

	incomingRequest := withpointers.SomeModel{}

	// if not nil checked, could throw null pointer error.
	if incomingRequest.ID1 != nil && *incomingRequest.ID1 > 1 {
		return true
	}
	return false
}

// makeGenerationCalls - Test calls
func makeBetterPointerCalls() bool {

	incomingRequest := withlesspointers.SomeModel{}

	// no need to check for nil because its defaulted to 0
	if incomingRequest.ID1 > 1 {
		return true
	}
	return false
}
