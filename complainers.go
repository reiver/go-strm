package strm


// MissingDriverNameComplainer is an interface that wraps the Error and
// MissingDriverNameComplainer methods.
//
// A call to the Begin func, Strmer.Strm method and Strmer.End method
// must include a string as their 1st argument. (And thus must be
// passed at least 1 argument.)
//
// This 1st string argument is a name
//
// If this complainer is returned as an error it indicates that
// a call to the Begin func, Strmer.Strm method or Strmer.End method
// did not have at least 1 argument (i.e., it had 0 arguments) and
// thus did not have a name .
type MissingDriverNameComplainer interface {
	error
	MissingDriverNameComplainer()
}


// NonStringDriverNameComplainer is an interface that wraps the Error and
// NonStringDriverNameComplainer methods.
//
// A call to the Begin func, Strmer.Strm method and Strmer.End method
// must include a string as their 1st argument. (And thus must be
// passed at least 1 argument.)
//
// This 1st string argument is a name
//
// If this complainer is returned as an error it indicates that
// a call to the Begin func, Strmer.Strm method or Strmer.End method
// did indeed have at least 1 argument but that the 1st argument was
// not a string.
type NonStringDriverNameComplainer interface {
	error
	NonStringDriverNameComplainer()
}


type internalMissingDriverNameComplainer   struct{}
type internalNonStringDriverNameComplainer struct{}


func (*internalMissingDriverNameComplainer) MissingDriverNameComplainer() {
	// Nothing here.
}
func (*internalNonStringDriverNameComplainer) NonStringDriverNameComplainer() {
	// Nothing here.
}


func (*internalMissingDriverNameComplainer) Error() string {
	return "Missing driver name."
}
func (*internalNonStringDriverNameComplainer) Error() string {
	return "Non-string driver name."
}


func newMissingDriverNameComplainer() MissingDriverNameComplainer {
	complainer := internalMissingDriverNameComplainer{}

	return &complainer
}
func newNonStringDriverNameComplainer() NonStringDriverNameComplainer {
	complainer := internalNonStringDriverNameComplainer{}

	return &complainer
}
