package nakereturn

import (
	"fmt"
	"math/rand"
)

// SmallFunction -
func SmallFunction(b int) (n int, err error) {
	n = rand.Intn(b)
	err = fmt.Errorf("This is an error: %d", b)
	return
}
