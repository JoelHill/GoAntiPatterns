package nakereturn

import (
	"fmt"
	"math/rand"
)

func something() error {
	return fmt.Errorf("something error")
}

// SmallThing -
func SmallThing(b int) (n int, err error) {
	n = rand.Intn(b)

	if err := something(); err != nil {
		fmt.Println("something() Err:", err)
		// return n, err
	}
	return
	// return n, err
}

func main() {
	_, err := SmallThing(10)
	// If you expect to get the error from "something" function returned to main, you would be mistaken.
	fmt.Println("main err:", err)
}
