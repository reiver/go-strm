package strm


// Strmer is an interface that wraps the Strm and End methods.
//
// Typically, one would call the Begin method to get a Strmer.
type Strmer interface {
        Strm(...interface{}) Strmer
        End(...interface{})
}
