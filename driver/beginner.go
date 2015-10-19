package strmdriver


// Beginner is an interface that wraps the Begin method.
//
// A Beginner is the 2nd argument to the RegisterBeginner
// func.
type Beginner interface {

	Begin(chan<- []interface{}, ...interface{})
}
