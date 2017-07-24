package nonnakereturn

import (
	"errors"
	"fmt"
	"math/rand"
)

func something() error {
	return fmt.Errorf("something error")
}

// SmallThing -
func SmallThing(b int) (n int, err error) {
	n = rand.Intn(b)

	if somethingErr := something(); somethingErr != nil {
		return n, somethingErr
	}

	err = errors.New("explain what went wrong here")
	return n, err
}

func main() {
	_, err := SmallThing(10)
	fmt.Println("main err:", err)
}
