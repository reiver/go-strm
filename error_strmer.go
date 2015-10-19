package strm


type internalErrorStrmer struct {
	err error
}


func newErrorStrmer(err error) Strmer {
	strmer := internalErrorStrmer{
		err:err,
	}

	return &strmer
}

func (strmer *internalErrorStrmer) Strm(...interface{}) Strmer {
//@TODO: Is there a better way of dealing with this?
	panic(strmer.err)
}


func (strmer *internalErrorStrmer) End(...interface{}) {
//@TODO: Is there a better way of dealing with this?
	panic(strmer.err)
}
