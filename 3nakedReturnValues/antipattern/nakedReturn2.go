package nakereturn

import (
	"fmt"
	"math/rand"
)

func something() error {
	return fmt.Errorf("a bad error")
}

// SmallThing -
func SmallThing(b int) (n int, err error) {
	n = rand.Intn(b)

	if err := something(); err != nil {
		// this line you will see "a bad error"
		fmt.Println("something() Err:", err)
		// return n, err
	}
	return
	// return n, err
}

func main() {
	_, err := SmallThing(10)
	// You expect -> main err: a bad error. But you would be mistaken.
	// The output -> main err: <nil>
	fmt.Println("main err:", err)
}
