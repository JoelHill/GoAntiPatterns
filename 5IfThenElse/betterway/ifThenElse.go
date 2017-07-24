package ifthenelse

type someType int

var (
	foo someType = 100
	bar someType = 200
)

func figureThingsOut(x int) someType {
	if x > 2 {
		return foo
	}

	// happy path is indented to the left
	return bar
}
