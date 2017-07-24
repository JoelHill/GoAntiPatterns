package makecall

import (
	"github.com/JoelHill/GoAntiPatterns/packageUtil/antipattern"
	"github.com/JoelHill/GoAntiPatterns/packageUtil/betterway"
)

// makeAntiPatternUtilCalls - Test calls
func makeAntiPatternUtilCalls() {
	util.RandomBytes(20)

	util.RandomString(20)

	util.Cert("mycert")
}

// makeGenerationCalls - Test calls
func makeGenerationCalls() {
	generate.RandomBytes(20)

	generate.RandomString(20)

	generate.Cert("mycert")
}
