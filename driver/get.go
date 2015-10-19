package strmdriver


// GetBeginner returns a Beginner registered under the
// name 'name' if it exists in the registry.
//
// Likely you will not use this func.
func GetBeginner(name string) (Beginner, bool) {

	beginner, ok := globalBeginners[name]
	if !ok {
		return nil, false
	}

	return beginner, true
}


// GetStrmer returns a Strmer registered under the
// name 'name' if it exists in the registry.
//
// Likely you will not use this func.
func GetStrmer(name string) (Strmer, bool) {

	strmer, ok := globalStrmers[name] 
	if !ok {
		return nil, false
	}

	return strmer, true
}


// GetEnder returns a Ender registered under the
// name 'name' if it exists in the registry.
//
// Likely you will not use this func.
func GetEnder(name string) (Ender, bool) {

	ender, ok := globalEnders[name]
	if !ok {
		return nil, false
	}

	return ender, ok
}
