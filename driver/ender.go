package strmdriver


// Ender is an interface that wraps the End method.
//
// An Ender is the 2nd argument to the RegisterEnder
// func.
type Ender interface {
	End(<-chan []interface{}, ...interface{})
}

