package strm


// parseArgs deals with parsing the args of the Begin func, and the
// Strmer.Strm and Strmer.End methods.
func parseArgs(args []interface{}) (string, []interface{}, error) {

	// Confirm that there is at least 1 arg.
	//
	// The 1st arg (at index 0) is the name of
	// the strmdriver.Beginner, strmdriver.Strmer
	// or strmdriver.Ender.
	if 1 > len(args) {
//@TODO: Should include depth, so a better error message can be included.
		return "", nil, newMissingDriverNameComplainer()
	}

	// Confirm that the 1st arg is a string.
	//
	// Names are strings.
	arg0 := args[0]
	arg0String, ok := arg0.(string)
	if !ok {
//@TODO: Should include depth, so a better error message can be included.
		return "", nil, newNonStringDriverNameComplainer()
	}

	// Get the subArgs.
	subArgs := args[1:]

	// Return.
	return arg0String, subArgs, nil
}
