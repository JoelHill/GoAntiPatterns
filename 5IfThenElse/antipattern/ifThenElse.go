package ifthenelse

type someType int

var (
	foo someType = 100
	bar someType = 200
)

func figureThingsOut(x int) someType {
	if x > 2 {
		return foo
	} else {
		return bar
	}
}
