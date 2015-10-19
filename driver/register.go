package strmdriver


// RegisterBeginner registers a Beginner.
//
// If you are writer a driver that you want to work with
// the strm.Begin func, then you would use this.
func RegisterBeginner(name string, beginner Beginner) {
	globalBeginners[name] = beginner
}


// RegisterStrmer registers a Strmer.
//
// If you are writer a driver that you want to work with
// the strm.Strmer.Strm method, then you would use this.
func RegisterStrmer(name string, strmer Strmer) {
	globalStrmers[name] = strmer
}


// RegisterEnder registers an Ender.
//
// If you are writer a driver that you want to work with
// the strm.Strmer.End func, then you would use this.
func RegisterEnder(name string, ender Ender) {
	globalEnders[name] = ender
}
