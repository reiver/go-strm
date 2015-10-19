package strm


import (
	"github.com/reiver/go-strm/driver"
)


// Begin returns a new Strmer.
func Begin(args ...interface{}) Strmer {

	// Parse args.
	driverName, subArgs, err := parseArgs(args)
	if nil != err {
		return newErrorStrmer(err)
	}

	// Get driver.
	beginDriver, ok := strmdriver.GetBeginner(driverName)
	if !ok {
		return newErrorStrmer(newMissingDriverNameComplainer())
	}

	// Execute driver.
	dst := make(chan []interface{})
	go beginDriver.Begin(dst, subArgs...)

	// Return.
	return newStrmer(dst)
}
