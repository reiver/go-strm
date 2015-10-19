package strm


import (
	"github.com/reiver/go-strm/driver"
)


type internalStrmer struct {
	src <-chan []interface{}
}


func newStrmer(src <-chan []interface{}) Strmer {
	strmer := internalStrmer{
		src:src,
	}

	return &strmer
}


func (strmer *internalStrmer) Strm(args ...interface{}) Strmer {

	// Parse args.
	driverName, subArgs, err := parseArgs(args)
	if nil != err {
                return newErrorStrmer(err)
	}

	// Get driver.
	strmDriver, ok := strmdriver.GetStrmer(driverName)
	if !ok {
		return newErrorStrmer(newMissingDriverNameComplainer())
	}

	// Execute driver.
	dst := make(chan []interface{})
	go strmDriver.Strm(strmer.src, dst, subArgs...)

	// Return.
	return newStrmer(dst)
}


func (strmer *internalStrmer) End(args ...interface{}) {

	// Parse args.
	driverName, subArgs, err := parseArgs(args)
	if nil != err {
//@TODO: Is there a better way to deal with this error?
                panic(err)
	}

	// Get driver.
	endDriver, ok := strmdriver.GetEnder(driverName)
	if !ok {
//@TODO: Is there a better way to deal with this error?
		panic(newMissingDriverNameComplainer())
	}

	// Execute driver.
	//
	// Not that unlike Strm() and Begin() this does
	// not spin up the driver method into its own
	// goroutine, but instead runs it inline.
	//
	// The reason is we want this to block.
	endDriver.End(strmer.src, subArgs...)
}
