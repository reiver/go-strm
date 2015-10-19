package strmdriver


// Strmer is an interface that wraps the Strm method.
//
// A Strmer is the 2nd argument to the RegisterStrmer
// func.
type Strmer interface {
	Strm(<-chan []interface{}, chan<- []interface{}, ...interface{})
}
